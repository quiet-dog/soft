package gateway

import (
	"fmt"
	"log"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type Gateway struct {
	Devices  GMap[int64, *Client]
	Channels GMap[<-chan Msg, chan Msg] // 向客户端发送数据的通道
}

func NewGateway() *Gateway {
	return &Gateway{
		Devices: GMap[int64, *Client]{},
	}
}

func (g *Gateway) AddClient(serviceId int64, conf Config) (*Client, error) {
	client, err := NewClinet(conf)
	if err != nil {
		return nil, err
	}
	g.Devices.Store(serviceId, &client)
	log.Printf("[INFO] 添加设备 %d 的客户端成功", serviceId)
	// 启动监听客户端通道
	go g.listenClient(serviceId, &client)
	return &client, nil
}

func (g *Gateway) DeleteClient(serviceId int64) {
	if client, ok := g.Devices.Load(serviceId); ok {
		client.Close()
		g.Devices.Delete(serviceId)
		log.Printf("[INFO] 删除设备 %d 的客户端成功", serviceId)
	} else {
		log.Printf("[WARN] 设备 %d 的客户端不存在", serviceId)
	}
}

func (g *Gateway) GetClient(serviceId int64) (*Client, bool) {
	if client, ok := g.Devices.Load(serviceId); ok {
		return client, true
	}
	log.Printf("[WARN] 设备 %d 的客户端不存在", serviceId)
	return nil, false
}

func (g *Gateway) Control(serviceId int64, commands ...*gjson.Json) (err error) {
	if client, ok := g.Devices.Load(serviceId); ok {
		return client.Control(commands...)
	}
	log.Printf("[WARN] 设备 %d 的客户端不存在", serviceId)
	return fmt.Errorf("设备不存在")
}

// 更新客户端配置
func (g *Gateway) UpdateClientConfig(serviceId int64, conf Config) error {
	if client, ok := g.Devices.Load(serviceId); ok {
		client.Close() // 先关闭旧的连接
		newClient, err := NewClinet(conf)
		if err != nil {
			return err
		}
		g.Devices.Store(serviceId, &newClient)
		log.Printf("[INFO] 更新设备 %d 的客户端配置成功", serviceId)
	}
	return nil
}

// 获取对应客户端
func (g *Gateway) Client(serviceId int64) (*Client, bool) {
	if client, ok := g.Devices.Load(serviceId); ok {
		return client, true
	}
	log.Printf("[WARN] 设备 %d 的客户端不存在", serviceId)
	return nil, false
}

func (g *Gateway) GetOnline(serviceId int64) bool {
	if client, ok := g.Devices.Load(serviceId); ok {
		if client != nil {
			return client.IsOnline()
		}
		return false
	}
	log.Printf("[WARN] 设备 %d 的客户端不存在", serviceId)
	return false
}

// 注册通道
func (g *Gateway) RegisterChannel(size int) <-chan Msg {
	ch := make(chan Msg, size)
	g.Channels.Store(ch, ch)
	return ch
}

// 取消注册通道
func (g *Gateway) UnregisterChannel(ch <-chan Msg) {
	if v, ok := g.Channels.Load(ch); ok {
		g.Channels.Delete(ch)
		// 关闭通道以通知所有监听者
		close(v)
		log.Println("[INFO] 通道已取消注册")
	} else {
		log.Println("[WARN] 通道不存在，无法取消注册")
	}
}

func (g *Gateway) listenClient(serviceId int64, client *Client) {
	for v := range client.Channel() {
		// 将值发送到所有注册的通道
		g.Channels.Range(func(key <-chan Msg, value chan Msg) bool {
			select {
			case value <- Msg{
				ServiceId: serviceId,
				Value:     v,
			}:
				fmt.Println("Sent value:", v.Value, "to channel for serviceId", serviceId)
			default:
				return true
			}
			return true
		})
	}
}
