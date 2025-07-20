package main

import (
	"devinggo/manage/pkg/gateway"
	"fmt"
	"time"
)

func main() {

	g := gateway.NewGateway()

	cfg := gateway.Config{
		Host: "127.0.0.1",
		Port: "4840",
		Type: "opc",
	}
	// c, err := gateway.NewClinet(cfg)
	g.AddClient(1, cfg)
	c, ok := g.Client(1)

	if !ok {
		// panic(err)
	}
	c.AddNodes(`{"id":1,"nodeId":"ns=3;i=3"}`)

	go func() {
		// for v := range c.Channel() {
		// 	fmt.Println("========", v.Value)
		// }
		// fmt.Println("Channel closed")
		ch := g.RegisterChannel(100)
		for v := range ch {
			fmt.Println("Received value:", v.Value.Value, "from channel")
		}

	}()
	// c.Close()
	time.Sleep(30 * time.Second)
}
