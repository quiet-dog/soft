package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

func main() {
	// 创建服务器
	srv := server.New(
		server.EndPoint("0.0.0.0", 4843),
	)

	// 添加对象和变量
	ns := srv.Node(ua.NewStringNodeID(0, "MyDevice"))

	// 创建3个温度
	a := opcua.Node{
		ID: ua.NewStringNodeID(2, "Temp"),
	}
	humNode := ns.AddVariable(ua.NewStringNodeID(0, "Hum"), ua.NewDouble(0))
	voltNode := ns.AddVariable(ua.NewStringNodeID(0, "Volt"), ua.NewDouble(0))

	// 启动后台更新协程
	go func() {
		for {
			time.Sleep(time.Second)
			for i := 1; i <= 3; i++ {
				ns.SetVariable("Temp"+string('0'+i), ua.NewDouble(rand.Float64()*100))
				ns.SetVariable("Hum"+string('0'+i), ua.NewDouble(rand.Float64()*100))
				ns.SetVariable("Volt"+string('0'+i), ua.NewDouble(rand.Float64()*24))
			}
		}
	}()

	// 启动服务器
	log.Println("OPC UA Server running at opc.tcp://localhost:4840")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
