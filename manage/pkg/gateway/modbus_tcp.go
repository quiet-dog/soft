package gateway

import (
	"context"
	"fmt"
	"maps"
	"time"

	"github.com/goburrow/modbus"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcron"
)

type ModbusTcpClient struct {
	client   *modbus.TCPClientHandler
	conf     Config
	isOnline bool
	ctx      context.Context
	cancel   context.CancelFunc
	channel  chan Value
	nodes    []*ModbusDevice
	timer    *gcron.Entry
}

type ModbusSensor struct {
	StartAddress uint16 `json:"startAddress"`
	Quantity     uint16 `json:"quantity"`
	SensorId     int64  `json:"sensorId"`
	ReadType     int64  `json:"readType"`
}

type ModbusDevice struct {
	SlaveId  uint16
	DeviceId int64
	Sensors  map[int64]ModbusSensor
}

func (m *ModbusTcpClient) TestPing() (err error) {
	handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", m.conf.Host, m.conf.Port))
	if err = handler.Connect(); err != nil {
		return
	}

	defer handler.Close()
	return
}

func durationToCron(duration time.Duration) string {
	seconds := int(duration.Seconds())

	if seconds == 0 {
		return fmt.Sprintln("*/5 * * * * *")
	}

	// 如果秒数小于 60，意味着每隔多少秒执行一次
	if seconds < 60 {
		return fmt.Sprintf("*/%d * * * * *", seconds) // 每 `seconds` 秒执行一次
	}

	// 如果大于等于 60 秒，可以进一步转换成分钟、小时等
	minutes := seconds / 60
	if minutes < 60 {
		return fmt.Sprintf("%d */%d * * * *", seconds%60, minutes) // 每 `minutes` 分钟执行一次
	}

	hours := minutes / 60
	if hours < 24 {
		return fmt.Sprintf("* %d * * *", hours) // 每 `hours` 小时执行一次
	}

	days := hours / 24
	return fmt.Sprintf("* %d * * *", days) // 每 `days` 天执行一次
}

// 初始化
func (c *ModbusTcpClient) connectAndSubscribeOnce(channel chan Value) (err error) {
	c.channel = channel
	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancel
	url := fmt.Sprintf("%s:%s", c.conf.Host, c.conf.Port)

	g := gcron.New()

	handler := modbus.NewTCPClientHandler(url)
	handler.Timeout = 2 * time.Second
	if err = handler.Connect(); err != nil {
		return
	}
	c.client = handler

	defer func() {
		c.isOnline = false
		handler.Close()
	}()

	c.isOnline = true
	//  durationToCron(c.conf.SubTime)
	c.timer, err = g.Add(ctx, durationToCron(c.conf.SubTime), func(ctx context.Context) {
		c.client.Close()
		if err := c.client.Connect(); err != nil {

			fmt.Println("===========================测试111")

			c.isOnline = false
			return
		}
		for _, device := range c.nodes {
			for _, sensor := range device.Sensors {
				client := modbus.NewClient(handler)
				var rs []byte
				// 读寄存器
				if sensor.ReadType == 1 {
					rs, err = client.ReadHoldingRegisters(sensor.StartAddress, sensor.Quantity)
					if err != nil {
						continue
					}
				}

				// 读写寄存器
				if sensor.ReadType == 2 {
					rs, err = client.ReadInputRegisters(sensor.StartAddress, sensor.Quantity)
					if err != nil {
						continue
					}
				}

				if len(rs) < 2 || len(rs)%2 != 0 {
					continue
				}

				registerValues := []uint16{}

				for i := 0; i < len(rs)-1; i += 2 {
					registerValues = append(registerValues, uint16(rs[0])<<8|uint16(rs[1]))
				}

				msg := Value{
					ID:         sensor.SensorId,
					Value:      registerValues,
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

func (c *ModbusTcpClient) AddNodes(devices ...ModbusDevice) {
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

func (c *ModbusTcpClient) Control(commands ...gjson.Json) (err error) {
	for _, command := range commands {
		startAddr := command.Get("startAddr").Uint16()
		values := command.Get("value").Int64s()
		handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", c.conf.Host, c.conf.Port))
		handler.Timeout = 3 * time.Second
		handler.SlaveId = c.client.SlaveId
		err = handler.Connect()
		if err != nil {
			return
		}
		defer handler.Close()

		// 把 values 转成字节切片，Modbus寄存器每个2字节
		byteValues := make([]byte, 0, len(values)*2)
		for _, v := range values {
			val := uint16(v) // 转uint16，注意超出范围需处理
			high := byte(val >> 8)
			low := byte(val & 0xFF)
			byteValues = append(byteValues, high, low)
		}

		client := modbus.NewClient(handler)
		_, err = client.WriteMultipleRegisters(startAddr, uint16(len(values)), byteValues)
		if err != nil {
			return
		}
	}

	return
}
