package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

func main() {
	ctx := context.Background()

	endpoint := "opc.tcp://172.20.10.2:17778"

	c, err := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone), opcua.SecurityPolicy(ua.SecurityPolicyURINone))
	if err != nil {
		log.Fatal("create client:", err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatal("connect:", err)
	}
	defer c.Close(ctx)

	rootNode := ua.NewNumericNodeID(0, id.ObjectsFolder)

	// 打印 NamespaceArray
	nsNode := c.Node(ua.NewNumericNodeID(0, 2255))
	val, err := nsNode.Value(ctx)
	if err != nil {
		log.Fatal("read namespace array:", err)
	}
	fmt.Println("Namespaces:", val.Value())

	fmt.Println("\n===== Browsing OPC UA Address Space =====")

	// 从 RootFolder 开始（ns=0;i=84）
	// root := ua.NewNumericNodeID(0, 84)
	// 重新，上面的都是错误的

}
