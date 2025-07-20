package gateway

import (
	"fmt"
	"log"
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
		if err = opcClient.TestPing(); err != nil {
			return
		}
		go client.StoreOpcServer()
	}

	// 等待服务器连接完成
	time.Sleep(2 * time.Second)
	return
}

func (c *Client) StoreOpcServer() {
	for {
		err := c.client.(*OpcClient).connectAndSubscribeOnce(c.channel)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		log.Println("[WARN] 连接中断，5 秒后重试...")
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) AddNodes(nodes ...string) {
	if v, ok := c.client.(*OpcClient); ok {
		n := []OpcNode{}
		for _, node := range nodes {
			g := gjson.New(node)
			fmt.Println(g.Get("nodeId").String())
			n = append(n, OpcNode{
				NodeId: g.Get("nodeId").String(),
				ID:     g.Get("id").Int64(),
			})
		}
		v.AddNodes(n...)
	}
}

func (c *Client) IsOnline() bool {
	if v, ok := c.client.(*OpcClient); ok {
		return v.isOnline
	}
	return false
}

func (c *Client) Channel() chan Value {
	return c.channel
}

func (c *Client) Close() {
	if v, ok := c.client.(*OpcClient); ok {
		v.client.Close(v.ctx)
	}
	close(c.channel)
	log.Println("[INFO] 客户端已关闭")
}
