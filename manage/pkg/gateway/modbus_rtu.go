package gateway

import (
	"context"
	"fmt"
	"maps"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/simonvetter/modbus"
)

type ModbusRtuClient struct {
	client   *modbus.ModbusClient
	conf     Config
	isOnline bool
	ctx      context.Context
	cancel   context.CancelFunc
	channel  chan Value
	nodes    []*ModbusDevice
	timer    *gcron.Entry
}

func (m *ModbusRtuClient) getConf() (conf modbus.ClientConfiguration) {
	var speed uint = 2
	var dataBits uint = 8
	var stopBits uint = 1
	var parity uint = modbus.PARITY_NONE

	if m.conf.Extend != nil {
		if !m.conf.Extend.Get("speed").IsEmpty() {
			speed = m.conf.Extend.Get("speed").Uint()
		}
		if !m.conf.Extend.Get("dataBits").IsEmpty() {
			dataBits = m.conf.Extend.Get("dataBits").Uint()
		}
		if !m.conf.Extend.Get("stopBits").IsEmpty() {
			stopBits = m.conf.Extend.Get("stopBits").Uint()
		}
		if !m.conf.Extend.Get("parity").IsEmpty() {
			parity = m.conf.Extend.Get("parity").Uint()
		}
	}
	conf.URL = m.conf.Extend.Get("url").String()
	conf.Speed = speed
	conf.Timeout = m.conf.SubTime * time.Second
	conf.DataBits = dataBits
	conf.StopBits = stopBits
	conf.Parity = parity
	return
}

func (m *ModbusRtuClient) TestPing() (err error) {

	conf := m.getConf()
	client, err := modbus.NewClient(&conf)
	if err != nil {
		return
	}
	err = client.Open()
	if err != nil {
		return
	}
	defer client.Close()

	return
}

// 初始化
func (c *ModbusRtuClient) connectAndSubscribeOnce(channel chan Value) (err error) {
	c.channel = channel
	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancel

	g := gcron.New()
	conf := c.getConf()
	client, err := modbus.NewClient(&conf)
	if err != nil {
		return
	}
	err = client.Open()
	if err != nil {
		return
	}
	defer client.Close()
	c.client = client

	c.isOnline = true
	//  durationToCron(c.conf.SubTime)
	c.timer, err = g.Add(ctx, durationToCron(c.conf.SubTime), func(ctx context.Context) {
		for _, device := range c.nodes {
			if c.client == nil {
				continue
			}
			c.client.SetUnitId(uint8(device.SlaveId))
			for _, sensor := range device.Sensors {
				var rs []uint16
				// 读寄存器
				if sensor.ReadType == 1 {
					rs, err = c.client.ReadRegisters(sensor.StartAddress, sensor.Quantity, modbus.HOLDING_REGISTER)
					if err != nil {
						continue
					}
				} else if sensor.ReadType == 2 {
					rs, err = c.client.ReadRegisters(sensor.StartAddress, sensor.Quantity, modbus.INPUT_REGISTER)
					if err != nil {
						continue
					}
				} else {
					continue
				}

				msg := Value{
					ID:         sensor.SensorId,
					Value:      rs,
					CreateTime: time.Now(),
					Type:       "ArrayUnit16",
					DeviceId:   device.DeviceId,
				}

				select {
				case <-c.ctx.Done():
					c.isOnline = false
					return
				case c.channel <- msg:
					c.isOnline = true
				default:
					c.isOnline = false
					return
				}

			}
		}
	})
	if err != nil {
		return
	}

	<-c.ctx.Done()
	c.isOnline = false
	c.client.Close()
	c.timer.Stop()

	return
}

// AddNodes 添加设备节点
func (c *ModbusRtuClient) AddNodes(devices ...ModbusDevice) {
	for _, v := range devices {
		var isDeviceExit bool
		for _, vv := range c.nodes {
			if v.DeviceId == vv.DeviceId {
				isDeviceExit = true
				maps.Copy(vv.Sensors, v.Sensors)
			}
		}
		if !isDeviceExit {
			c.nodes = append(c.nodes, &v)
		}
	}
}

// Control 写入寄存器
func (c *ModbusRtuClient) Control(commands ...gjson.Json) (err error) {
	for _, command := range commands {
		startAddr := command.Get("startAddr").Uint16()
		values := command.Get("value").Int64s()
		slaveId := command.Get("slaveId").Uint8()

		// 创建新客户端（与原始代码保持一致）
		cfg := c.getConf()
		client, err := modbus.NewClient(&cfg)
		if err != nil {
			return fmt.Errorf("创建客户端失败: %v", err)
		}

		// 设置从站 ID（与 c.client 保持一致）
		client.SetUnitId(slaveId)

		// 打开连接
		if err = client.Open(); err != nil {
			return fmt.Errorf("连接失败: %v", err)
		}
		defer client.Close()

		// 将 values 转换为 uint16 切片
		registerValues := make([]uint16, len(values))
		for i, v := range values {
			if v < 0 || v > 0xFFFF {
				return fmt.Errorf("值超出 uint16 范围: %d", v)
			}
			registerValues[i] = uint16(v)
		}

		// 写入多个寄存器
		err = client.WriteRegisters(startAddr, registerValues)
		if err != nil {
			return fmt.Errorf("写入寄存器失败: %v", err)
		}
	}

	return nil
}
