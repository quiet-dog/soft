package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func browseRecursive(ctx context.Context, c *opcua.Client, nodeID *ua.NodeID, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	// Browse 请求
	req := &ua.BrowseRequest{
		View: &ua.ViewDescription{
			ViewID: ua.NodeIDNull,
		},
		NodesToBrowse: []*ua.BrowseDescription{
			{
				NodeID:          nodeID,
				BrowseDirection: ua.BrowseDirectionForward,
				ReferenceTypeID: ua.NewNumericNodeID(0, ua.ReferenceTypeIDHierarchicalReferences),
				IncludeSubtypes: true,
				NodeClassMask:   ua.NodeClassAll,
				ResultMask:      ua.BrowseResultMaskAll,
			},
		},
	}

	resp, err := c.Browse(ctx, req)
	if err != nil {
		log.Printf("Browse failed for %v: %v", nodeID, err)
		return
	}

	if len(resp.Results) == 0 {
		return
	}

	for _, ref := range resp.Results[0].References {
		fmt.Printf("%s- NodeID: %s | BrowseName: %s | DisplayName: %s | NodeClass: %v\n",
			indent,
			ref.NodeID,
			ref.BrowseName.Name,
			ref.DisplayName.Text,
			ref.NodeClass,
		)

		// 继续递归子节点
		browseRecursive(ctx, c, ref.NodeID.NodeID(), level+1)
	}
}

func main() {
	ctx := context.Background()

	endpoint := "opc.tcp://127.0.0.1:4840"

	c, err := opcua.NewClient(endpoint)
	if err != nil {
		log.Fatal("create client:", err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatal("connect:", err)
	}
	defer c.Close(ctx)

	// 打印 NamespaceArray
	nsNode := c.Node(ua.NewNumericNodeID(0, 2255))
	val, err := nsNode.Value(ctx)
	if err != nil {
		log.Fatal("read namespace array:", err)
	}
	fmt.Println("Namespaces:", val.Value())

	fmt.Println("\n===== Browsing OPC UA Address Space =====")

	// 从 RootFolder 开始（ns=0;i=84）
	root := ua.NewNumericNodeID(0, 84)

	browseRecursive(ctx, c, root, 0)
}
