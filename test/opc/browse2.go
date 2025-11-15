package main

import (
	"context"
	"devinggo/manage/pkg/gateway"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

func join(a, b string) string {
	if a == "" {
		return b
	}
	return a + "." + b
}

const maxDepth = 10

func browse2(ctx context.Context, n *opcua.Node, path string, level int) (*gateway.NodeDef, error) {
	if level > maxDepth {
		return nil, nil
	}

	attrs, err := n.Attributes(ctx,
		ua.AttributeIDNodeClass,
		ua.AttributeIDBrowseName,
		ua.AttributeIDDescription,
		ua.AttributeIDAccessLevel,
		ua.AttributeIDDataType,
		ua.AttributeIDDisplayName,
	)
	if err != nil {
		return nil, err
	}

	def := &gateway.NodeDef{
		NodeID: n.ID,
	}

	switch err := attrs[0].Status; err {
	case ua.StatusOK:
		def.NodeClass = ua.NodeClass(attrs[0].Value.Int())
	default:
		return nil, err
	}

	// BrowseName
	switch err := attrs[1].Status; err {
	case ua.StatusOK:
		def.BrowseName = attrs[1].Value.String()
	default:
		return nil, err
	}

	switch err := attrs[2].Status; err {
	case ua.StatusOK:
		def.Description = attrs[2].Value.String()
	case ua.StatusBadAttributeIDInvalid:
		// ignore
	default:
		return nil, err
	}

	switch err := attrs[3].Status; err {
	case ua.StatusOK:
		def.AccessLevel = ua.AccessLevelType(attrs[3].Value.Int())
		def.Writable = def.AccessLevel&ua.AccessLevelTypeCurrentWrite == ua.AccessLevelTypeCurrentWrite
	case ua.StatusBadAttributeIDInvalid:
		// ignore
	default:
		return nil, err
	}

	switch err := attrs[4].Status; err {
	case ua.StatusOK:
		switch v := attrs[4].Value.NodeID().IntID(); v {
		case id.DateTime:
			def.DataType = "time.Time"
		case id.Boolean:
			def.DataType = "bool"
		case id.SByte:
			def.DataType = "int8"
		case id.Int16:
			def.DataType = "int16"
		case id.Int32:
			def.DataType = "int32"
		case id.Byte:
			def.DataType = "byte"
		case id.UInt16:
			def.DataType = "uint16"
		case id.UInt32:
			def.DataType = "uint32"
		case id.UtcTime:
			def.DataType = "time.Time"
		case id.String:
			def.DataType = "string"
		case id.Float:
			def.DataType = "float32"
		case id.Double:
			def.DataType = "float64"
		default:
			def.DataType = attrs[4].Value.NodeID().String()
		}
	case ua.StatusBadAttributeIDInvalid:
		// ignore
	default:
		return nil, err
	}

	switch err := attrs[5].Status; err {
	case ua.StatusOK:
		def.DisplayName = attrs[5].Value.String()
	default:
		return nil, err
	}

	def.Path = join(path, def.BrowseName)

	// 遍历子节点
	def.Children = []*gateway.NodeDef{}
	browseChildren := func(refType uint32) error {
		refs, err := n.ReferencedNodes(ctx, refType, ua.BrowseDirectionForward, ua.NodeClassAll, true)
		if err != nil {
			return fmt.Errorf("References: %d: %s", refType, err)
		}
		for _, rn := range refs {
			childNode, err := browse2(ctx, rn, def.Path, level+1)
			if err != nil {
				return fmt.Errorf("browse children: %s", err)
			}
			if childNode != nil {
				def.Children = append(def.Children, childNode)
			}
		}
		return nil
	}

	if err := browseChildren(id.HasComponent); err != nil {
		return nil, err
	}
	if err := browseChildren(id.Organizes); err != nil {
		return nil, err
	}
	if err := browseChildren(id.HasProperty); err != nil {
		return nil, err
	}

	return def, nil
}

func main() {
	endpoint := flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
	nodeID := flag.String("node", "i=84", "node id for the root node") // i=84 is the standard root node
	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
	flag.Parse()
	log.SetFlags(0)

	ctx := context.Background()

	c, err := opcua.NewClient(*endpoint)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close(ctx)

	id, err := ua.ParseNodeID(*nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %s", err)
	}

	nodeList, err := browse2(ctx, c.Node(id), "", 0)
	if err != nil {
		log.Fatal(err)
	}
	d, _ := json.Marshal(nodeList)
	f, _ := os.OpenFile("test.json", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	f.Write(d)
	f.Close()
}
