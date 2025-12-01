package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/simonvetter/modbus"
)

// ModbusTcpClient 结构体，适配 simonvetter/modbus 库
type ModbusTcpClient struct {
	client   *modbus.ModbusClient
	conf     Config
	isOnline bool
	ctx      context.Context
	cancel   context.CancelFunc
	channel  chan Value
	nodes    []*ModbusDevice
	timer    *gcron.Entry
}

// ModbusSensor 定义传感器配置
type ModbusSensor struct {
	StartAddress uint16 `json:"startAddress"`
	Quantity     uint16 `json:"quantity"`
	SensorId     int64  `json:"sensorId"`
	ReadType     int64  `json:"readType"`
	SlaveId      uint8  `json:"slaveId"`
}

// ModbusDevice 定义设备配置
type ModbusDevice struct {
	SlaveId  uint16
	DeviceId int64
	Sensors  map[int64]ModbusSensor
}

// TestPing 测试连接
func (m *ModbusTcpClient) TestPing() (err error) {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     fmt.Sprintf("tcp://%s:%s", m.conf.Host, m.conf.Port),
		Timeout: 2 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("创建客户端失败: %v", err)
	}

	err = client.Open()
	if err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}
	defer client.Close()
	return nil
}

// durationToCron 将时间间隔转换为 cron 表达式
func durationToCron(duration time.Duration) string {
	seconds := int(duration.Seconds())

	if seconds == 0 {
		return "*/5 * * * * *"
	}

	if seconds < 60 {
		return fmt.Sprintf("*/%d * * * * *", seconds)
	}

	minutes := seconds / 60
	if minutes < 60 {
		return fmt.Sprintf("%d */%d * * * *", seconds%60, minutes)
	}

	hours := minutes / 60
	if hours < 24 {
		return fmt.Sprintf("* %d */%d * *", minutes%60, hours)
	}

	days := hours / 24
	return fmt.Sprintf("* * %d * *", days)
}

// connectAndSubscribeOnce 初始化并订阅数据
func (c *ModbusTcpClient) connectAndSubscribeOnce(channel chan Value) (err error) {
	c.channel = channel
	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancel

	// 创建 Modbus 客户端
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     fmt.Sprintf("tcp://%s:%s", c.conf.Host, c.conf.Port),
		Timeout: 3 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("创建客户端失败: %v", err)
	}

	// 打开连接
	if err = client.Open(); err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}
	c.client = client

	// 确保在函数退出时清理资源
	defer func() {
		c.isOnline = false
		client.Close()
		if c.timer != nil {
			c.timer.Stop()
		}
	}()

	c.isOnline = true

	// 创建定时任务
	g := gcron.New()
	c.timer, err = g.Add(ctx, durationToCron(c.conf.SubTime), func(ctx context.Context) {

		for _, device := range c.nodes {

			// 设置从站 ID
			for _, sensor := range device.Sensors {
				c.client.SetUnitId(uint8(sensor.SlaveId))
				// 选择寄存器类型
				var regType modbus.RegType

				switch sensor.ReadType {
				case 1:
					regType = modbus.HOLDING_REGISTER
				case 2:
					regType = modbus.INPUT_REGISTER
				default:
					fmt.Printf("无效的 ReadType: %d (SensorID: %d)\n", sensor.ReadType, sensor.SensorId)
					continue
				}

				// 读取寄存器
				rs, err := c.client.ReadRegisters(sensor.StartAddress, sensor.Quantity, regType)
				if err != nil {

					fmt.Printf("读取寄存器失败 (SlaveId: %d) (SensorID: %d) (StartAddress: %d) (Quantity: %d) (RegType: %d) (URL: %s) (ReadType: %d): %v\n", device.SlaveId, sensor.SensorId, sensor.StartAddress, sensor.Quantity, regType, fmt.Sprintf("tcp://%s:%s", c.conf.Host, c.conf.Port), sensor.ReadType, err)
					continue
				}

				// 构造消息
				msg := Value{
					ID:         sensor.SensorId,
					Value:      rs,
					CreateTime: time.Now(),
					Type:       "ArrayUnit16",
					DeviceId:   device.DeviceId,
				}

				// 发送消息到通道
				select {
				case <-c.ctx.Done():
					c.isOnline = false
					return
				case c.channel <- msg:
					fmt.Println("===========读取寄存器============", msg)
					c.isOnline = true
				default:
					fmt.Println("===========读取寄存器============", 22)
					c.isOnline = false
					return
				}
			}
		}
	})
	if err != nil {
		return fmt.Errorf("添加定时任务失败: %v", err)
	}

	<-c.ctx.Done()
	c.isOnline = false
	return nil
}

// AddNodes 添加设备节点
func (c *ModbusTcpClient) AddNodes(devices ...ModbusDevice) {

	for _, v := range devices {

		var isDeviceExit bool
		for _, vv := range c.nodes {
			if v.DeviceId == vv.DeviceId {
				isDeviceExit = true
			}
		}
		if !isDeviceExit {
			c.nodes = append(c.nodes, &v)
		}
	}
}

// Control 写入寄存器
func (c *ModbusTcpClient) Control(commands ...*gjson.Json) (err error) {
	for _, command := range commands {
		startAddr := command.Get("startAddr").Uint16()
		values := command.Get("value").Int64s()
		slaveId := command.Get("slaveId").Uint8()

		// 创建新客户端（与原始代码保持一致）
		client, err := modbus.NewClient(&modbus.ClientConfiguration{
			URL:     fmt.Sprintf("tcp://%s:%s", c.conf.Host, c.conf.Port),
			Timeout: 3 * time.Second,
		})
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
