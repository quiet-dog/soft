package main

import (
	"devinggo/manage/pkg/gateway"
	"time"
)

func main() {
	gateway.NewClinet(gateway.Config{
		Host: "127.0.0.1",
		Port: "4840",
		Type: "opc",
	})
	time.Sleep(3 * time.Second)
}
