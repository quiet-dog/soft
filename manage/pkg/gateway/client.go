package gateway

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type Client struct {
	client  interface{}
	channel chan Value
}

func NewClinet(cfg Config) (client Client, err error) {
	client = Client{}
	client.channel = make(chan Value, ChannelSize)
	if cfg.Type == SERVER_OPC {
		opcClient := OpcClient{
			conf: cfg,
		}
		client.client = &opcClient
		// if err = opcClient.TestPing(); err != nil {
		// 	return
		// }
		go client.StoreOpcServer()
	}

	if cfg.Type == SERVER_MODBUS_TCP {
		tcpClient := ModbusTcpClient{
			conf: cfg,
		}
		client.client = &tcpClient
		// if tcpClient.TestPing(); err != nil {
		// 	return
		// }
		go client.StoreModebusTcpServer()
	}

	// 等待服务器连接完成
	time.Sleep(2 * time.Second)
	return
}

// opc存储并且开始长连接服务
func (c *Client) StoreOpcServer() {
	for {
		err := c.client.(*OpcClient).connectAndSubscribeOnce(c.channel)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		log.Println("[WARN] opc 连接中断，5 秒后重试...")
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) StoreModebusTcpServer() {
	for {
		err := c.client.(*ModbusTcpClient).connectAndSubscribeOnce(c.channel)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		log.Println("[WARN] modbus 连接中断，5 秒后重试...")
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) AddNodes(nodes ...*gjson.Json) {
	// opc 设备的节点添加
	if v, ok := c.client.(*OpcClient); ok {
		n := []OpcNode{}
		for _, node := range nodes {
			fmt.Println(node.Get("nodeId").String())
			n = append(n, OpcNode{
				NodeId:   node.Get("nodeId").String(),
				ID:       node.Get("id").Int64(),
				DeviceId: node.Get("deviceId").Int64(),
			})
		}
		v.AddNodes(n...)
	}

	if v, ok := c.client.(*ModbusTcpClient); ok {
		n := []ModbusDevice{}
		for _, node := range nodes {
			device := ModbusDevice{}
			sensors := make(map[int64]ModbusSensor)
			if node.Get("sensors").IsMap() {
				for k, value := range node.Get("sensors").Map() {
					if kInt, err := strconv.Atoi(k); err != nil {
						continue
					} else {
						if sV, ok := value.(ModbusSensor); ok {
							sensors[int64(kInt)] = sV
						}
					}
				}
			}
			device.Sensors = sensors
			n = append(n, device)
		}

		v.AddNodes(n...)
	}

}

func (c *Client) IsOnline() bool {
	if v, ok := c.client.(*OpcClient); ok {
		return v.isOnline
	}
	if v, ok := c.client.(*ModbusTcpClient); ok {
		return v.isOnline
	}
	return false
}

func (c *Client) Control(commands ...gjson.Json) (err error) {
	// opc 设备的节点添加
	if v, ok := c.client.(*ModbusTcpClient); ok {
		return v.Control(commands...)
	}

	if v, ok := c.client.(*OpcClient); ok {
		return v.Control(commands...)
	}

	return
}

func (c *Client) Channel() chan Value {
	return c.channel
}

func (c *Client) Close() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in client.Close(): %v", r)
		}
	}()
	if v, ok := c.client.(*OpcClient); ok {
		if c.client != nil && v.ctx != nil {
			v.client.Close(v.ctx)
		}
	}

	if v, ok := c.client.(*ModbusTcpClient); ok {
		if c.client != nil && v.cancel != nil {
			v.cancel()
		}
	}

	close(c.channel)
	log.Println("[INFO] 客户端已关闭")
}
