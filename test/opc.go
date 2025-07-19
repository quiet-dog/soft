package main

import (
	"context"
	"devinggo/manage/model/res/device"
	"fmt"
	"log"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func browseNode(c *opcua.Client, nodeID *ua.NodeID, indent string, visited map[string]bool) {
	// 避免重复访问
	key := nodeID.String()
	if visited[key] {
		return
	}
	visited[key] = true

	req := &ua.BrowseRequest{
		View: &ua.ViewDescription{
			ViewID: ua.NewTwoByteNodeID(0),
		},
		NodesToBrowse: []*ua.BrowseDescription{
			{
				NodeID:          nodeID,
				BrowseDirection: ua.BrowseDirectionForward,
				ReferenceTypeID: ua.NewNumericNodeID(0, 0), // all references
				IncludeSubtypes: true,
				NodeClassMask:   0, // all node classes
				ResultMask:      uint32(ua.BrowseResultMaskAll),
			},
		},
	}

	resp, err := c.Browse(context.Background(), req)
	if err != nil {
		log.Printf("Browse error at node %s: %v", nodeID, err)
		return
	}

	for _, ref := range resp.Results[0].References {
		fmt.Printf("%s- [%s] %s (ns=%d, id=%v)\n", indent, ref.NodeClass, ref.BrowseName.Name, ref.NodeID.NodeID.Namespace(), ref.NodeID.NodeID.String())
		browseNode(c, ref.NodeID.NodeID, indent+"  ", visited)
	}
}

func browseNamespace3Nodes(c *opcua.Client, nodeID *ua.NodeID, indent string, visited map[string]bool, targetNS uint16) {
	key := nodeID.String()
	if visited[key] {
		return
	}
	visited[key] = true

	req := &ua.BrowseRequest{
		View: &ua.ViewDescription{
			ViewID: ua.NewTwoByteNodeID(0),
		},
		NodesToBrowse: []*ua.BrowseDescription{
			{
				NodeID:          nodeID,
				BrowseDirection: ua.BrowseDirectionForward,
				ReferenceTypeID: ua.NewNumericNodeID(0, 0), // all references
				IncludeSubtypes: true,
				NodeClassMask:   0, // all
				ResultMask:      uint32(ua.BrowseResultMaskAll),
			},
		},
	}

	resp, err := c.Browse(context.Background(), req)
	if err != nil {
		log.Printf("Browse error at node %s: %v", nodeID, err)
		return
	}

	for _, ref := range resp.Results[0].References {
		ns := ref.NodeID.NodeID.Namespace()
		if ns == targetNS {
			// if ref.NodeClass == ua.NodeClassVariable {
			// 	node := c.Node(ref.NodeID.NodeID)
			// 	val, err := node.Value(context.Background())
			// 	if err != nil {
			// 		log.Printf("%s  读取值失败: %v", indent, err)
			// 	} else {
			// 		fmt.Printf("%s  Value: %v\n", indent, val.Value())
			// 		if val.DataValue() != nil {
			// 			fmt.Printf("%s  SourceTimestamp: %v\n", indent, val.DataValue().SourceTimestamp)
			// 		} else {
			// 			fmt.Printf("%s  SourceTimestamp: <nil>\n", indent)
			// 		}
			// 	}
			// }
			fmt.Printf("%s- [%s] %s (ns=%d, id=%v)  %s\n", indent, ref.NodeClass, ref.BrowseName.Name, ns, ref.NodeID.NodeID.String(), ref.NodeID.NodeID.Type())
			browseNamespace3Nodes(c, ref.NodeID.NodeID, indent+"  ", visited, targetNS)
		} else {
			// 也可以递归，但不打印，或者根据需求决定
			// browseNamespace3Nodes(c, ref.NodeID.NodeID, indent+"  ", visited, targetNS)
		}
	}
}

func browseNamespaceTree(c *opcua.Client, nodeID *ua.NodeID, visited map[string]bool, targetNS uint16, treeName string) *device.OpcTree {
	key := nodeID.String()
	if visited[key] {
		return nil
	}
	visited[key] = true

	root := &device.OpcTree{
		NodeClass:    0,
		VariableName: treeName,
		Namespace:    fmt.Sprintf("%d", targetNS),
		NodeId:       nodeID.String(),
		Type:         nodeID.Type().String(),
		Children:     []*device.OpcTree{},
	}

	req := &ua.BrowseRequest{
		View: &ua.ViewDescription{
			ViewID: ua.NewTwoByteNodeID(0),
		},
		NodesToBrowse: []*ua.BrowseDescription{
			{
				NodeID:          nodeID,
				BrowseDirection: ua.BrowseDirectionForward,
				ReferenceTypeID: ua.NewNumericNodeID(0, 0),
				IncludeSubtypes: true,
				NodeClassMask:   0,
				ResultMask:      uint32(ua.BrowseResultMaskAll),
			},
		},
	}

	resp, err := c.Browse(context.Background(), req)
	if err != nil || len(resp.Results) == 0 || resp.Results[0].StatusCode != ua.StatusOK {
		log.Printf("Browse error at node %s: %v", nodeID, err)
		return root // 返回空的 root 也可以
	}

	for _, ref := range resp.Results[0].References {
		childNode := ref.NodeID.NodeID
		childKey := childNode.String()
		if visited[childKey] {
			continue
		}

		ns := childNode.Namespace()
		if ns == targetNS {
			child := &device.OpcTree{
				NodeClass:    ref.NodeClass,
				VariableName: ref.BrowseName.Name,
				Namespace:    fmt.Sprintf("%d", ns),
				NodeId:       childNode.String(),
				Children:     []*device.OpcTree{},
			}
			fmt.Printf("Found node: %s (ns=%d, id=%v)\n", ref.BrowseName.Name, ns, childNode.String())
			if childNode != nil {
				child.Type = childNode.Type().String()
			}
			child.Children = browseNamespaceTree(c, childNode, visited, targetNS, ref.DisplayName.Text).Children
			root.Children = append(root.Children, child)
		}
	}

	return root
}

func main() {
	endpoint := "opc.tcp://127.0.0.1:4840"
	c, err := opcua.NewClient(endpoint)
	if err != nil {
		log.Fatalf("Failed to create OPC UA client: %v", err)
	}

	if err := c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer c.Close(context.Background())

	// 先打印 NamespaceArray，确认索引和URI
	nsNode := c.Node(ua.NewNumericNodeID(0, 2255))
	val, err := nsNode.Value(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	namespaces, ok := val.Value().([]string)
	if !ok {
		log.Fatalf("Expected []string but got %T", val)
	}
	fmt.Println("Namespaces:")
	for i, ns := range namespaces {
		fmt.Printf("  %d: %s\n", i, ns)
	}

	// 从标准根节点开始递归浏览，获取所有命名空间下的节点
	// rootNode := ua.NewNumericNodeID(0, 84)
	visited := make(map[string]bool)
	// browseNode(c, rootNode, "", visited)

	objectsNode := ua.NewNumericNodeID(0, 85)
	visited = make(map[string]bool)
	fmt.Println("Browsing nodes in namespace ns=3:")
	// browseNamespace3Nodes(c, objectsNode, "", visited, 2)
	nodes := browseNamespaceTree(c, objectsNode, visited, 2, objectsNode.String())
	// fmt.Println("Namespace 3 Tree:", len(nodes.Children))
	for _, child := range nodes.Children {
		fmt.Printf("Node: %s, ID: %s, Type: %s\n", child.VariableName, child.NodeId, child.Type)
		for _, grandChild := range child.Children {
			fmt.Printf("  Child Node: %s, ID: %s, Type: %s\n", grandChild.VariableName, grandChild.NodeId, grandChild.Type)
			for _, greatGrandChild := range grandChild.Children {
				fmt.Printf("    Grandchild Node: %s, ID: %s, Type: %s\n", greatGrandChild.VariableName, greatGrandChild.NodeId, greatGrandChild.Type)
			}
		}
	}
}
