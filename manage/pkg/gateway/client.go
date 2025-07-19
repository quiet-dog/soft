package gateway

import (
	"log"
	"time"
)

type Client struct {
	client interface{}
}

func NewClinet(cfg Config) (client Client, err error) {
	client = Client{}
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
	return
}

func (c *Client) StoreOpcServer() {
	for {
		err := c.client.(*OpcClient).connectAndSubscribeOnce()
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		log.Println("[WARN] 连接中断，5 秒后重试...")
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) IsOnline() bool {
	if v, ok := c.client.(*OpcClient); ok {
		return v.isOnline
	}

	return false
}
