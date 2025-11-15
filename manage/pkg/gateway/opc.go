package gateway

import (
	"context"
	"fmt"
	"log"
	"slices"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

type OpcNodes []OpcNode

type OpcNode struct {
	ID       int64
	NodeId   string
	DeviceId int64
}

type OpcClient struct {
	conf     Config
	client   *opcua.Client
	isOnline bool
	sub      *monitor.Subscription
	nodes    OpcNodes
	ctx      context.Context
	cancel   context.CancelFunc
	channel  chan Value
}

func (c *OpcClient) TestPing() (err error) {
	// 测试是否可以脸上
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	endpointURL := fmt.Sprintf("opc.tcp://%s:%s", c.conf.Host, c.conf.Port)

	endpoints, err := opcua.GetEndpoints(ctx, endpointURL)
	if err != nil {
		log.Printf("[ERROR] 获取 OPC UA 服务器端点失败: %v", err)
		return
	}

	ep, err := opcua.SelectEndpoint(endpoints, "", ua.MessageSecurityModeNone)
	if err != nil {
		log.Printf("[ERROR] 选择端点失败: %v", err)
		return
	}

	opts := []opcua.Option{
		opcua.SecurityPolicy("None"),
		opcua.SecurityMode(ua.MessageSecurityModeNone),
		opcua.AuthAnonymous(),
		opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
	}

	client, err := opcua.NewClient(ep.EndpointURL, opts...)

	if err != nil {
		return
	}

	if err = client.Connect(ctx); err != nil {
		log.Printf("[ERROR] 连接 OPC UA 服务器失败: %v", err)
		return
	}
	defer client.Close(ctx)

	log.Println("[SUCCESS] 成功连接到 OPC UA 服务器", len(c.nodes))
	return
}

func (c *OpcClient) connectAndSubscribeOnce(channel chan Value) (err error) {
	c.channel = channel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.ctx = ctx
	c.cancel = cancel

	endpoints, err := opcua.GetEndpoints(ctx, fmt.Sprintf("opc.tcp://%s:%s", c.conf.Host, c.conf.Port))
	if err != nil {
		return
	}
	ep, err := opcua.SelectEndpoint(endpoints, "", ua.MessageSecurityModeFromString(""))
	if err != nil {
		return
	}

	opts := []opcua.Option{
		opcua.SecurityPolicy(c.conf.Extend.Get("policy").String()),
		opcua.SecurityModeString(c.conf.Extend.Get("mode").String()),
		opcua.CertificateFile(c.conf.Extend.Get("certPath").String()),
		opcua.PrivateKeyFile(c.conf.Extend.Get("keyPath").String()),
		// opcua.AuthAnonymous(),
		// opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
	}

	username := c.conf.Extend.Get("username").String()
	password := c.conf.Extend.Get("password").String()

	fmt.Println(c.conf.Extend.Get("policy").String(), c.conf.Extend.Get("mode").String(), c.conf.Extend.Get("certPath").String(), c.conf.Extend.Get("keyPath").String())
	if username != "" && password != "" {
		opts = append(opts, opcua.AuthUsername(username, password), opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeUserName))
	} else {
		opts = append(opts, opcua.AuthAnonymous(), opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous))
	}

	client, err := opcua.NewClient(ep.EndpointURL, opts...)
	if err != nil {
		return
	}

	c.client = client
	if err = client.Connect(ctx); err != nil {
		return
	}

	defer client.Close(ctx)

	m, err := monitor.NewNodeMonitor(client)
	if err != nil {
		return
	}

	m.SetErrorHandler(func(_ *opcua.Client, sub *monitor.Subscription, err error) {
		log.Printf("error: sub=%d err=%s", sub.SubscriptionID(), err.Error())
	})

	go c.watchOnline(ctx)
	c.startChanSub(ctx, m, c.conf.SubTime, 0)

	return
}

func (c *OpcClient) watchOnline(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done(): // 使用传入的 ctx
			c.isOnline = false
			return // 退出 goroutine
		case <-ticker.C:
			fmt.Println("=====", c.client.State().String())
			if c.client != nil && c.client.State().String() == "Connected" {
				c.isOnline = true
			} else {
				c.isOnline = false
			}
		}
	}
}

func (c *OpcClient) startCallbackSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, wg *sync.WaitGroup, nodes ...string) {
	sub, err := m.Subscribe(
		ctx,
		&opcua.SubscriptionParameters{
			Interval: interval,
		},
		func(s *monitor.Subscription, msg *monitor.DataChangeMessage) {
			if msg.Error != nil {
				log.Printf("[callback] sub=%d error=%s", s.SubscriptionID(), msg.Error)
			} else {
				log.Printf("[callback] sub=%d ts=%s node=%s value=%v", s.SubscriptionID(), msg.SourceTimestamp.UTC().Format(time.RFC3339), msg.NodeID, msg.Value.Value())
			}
			time.Sleep(lag)
		},
		nodes...)

	if err != nil {
		log.Fatal(err)
	}

	defer cleanup(ctx, sub)

	<-ctx.Done()
}

func (c *OpcClient) startChanSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, nodes ...string) {
	ch := make(chan *monitor.DataChangeMessage, 16)
	for _, v := range c.nodes {
		nodes = append(nodes, v.NodeId)
	}
	sub, err := m.ChanSubscribe(ctx, &opcua.SubscriptionParameters{Interval: interval}, ch, nodes...)
	if err != nil {
		log.Fatal(err)
	}
	c.sub = sub
	defer cleanup(ctx, sub)

	for msg := range ch {
		if msg.Error != nil {
			log.Printf("[channel ] sub=%d error=%s", sub.SubscriptionID(), msg.Error)
		} else {
			if c.channel == nil {
				continue
			}
			nodes := []Value{}
			for _, node := range c.nodes {
				if node.NodeId == msg.NodeID.String() {
					nodes = append(nodes, Value{
						ID:         node.ID,
						Value:      msg.Value.Value(),
						CreateTime: msg.SourceTimestamp,
						Type:       msg.Value.Type().String(),
					})
					fmt.Println("Received nodes:=========", node.ID)
				}
			}
			for _, v := range nodes {
				select {
				case c.channel <- v:
				default:
					// log.Println("[WARN] Channel is full, dropping message")
					continue
				}
			}

			// log.Printf("[channel ] sub=%d ts=%s node=%s value=%v", sub.SubscriptionID(), msg.SourceTimestamp.UTC().Format(time.RFC3339), msg.NodeID, msg.Value.Value())
		}
		time.Sleep(lag)
	}
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	// case <-time.After(5 * time.Second):
	// 	// 	if err = c.TestPing(); err != nil {
	// 	// 		c.isOnline = false
	// 	// 		c.client.Close(ctx)
	// 	// 		c.cancel()
	// 	// 		return
	// 	// 	}
	// 	case msg := <-ch:
	// 		fmt.Println("===============================来自通道信息")
	// 		if msg.Error != nil {
	// 			log.Printf("[channel ] sub=%d error=%s", sub.SubscriptionID(), msg.Error)
	// 		} else {
	// 			if c.channel == nil {
	// 				continue
	// 			}
	// 			nodes := []Value{}
	// 			for _, node := range c.nodes {
	// 				if node.NodeId == msg.NodeID.String() {
	// 					nodes = append(nodes, Value{
	// 						ID:         node.ID,
	// 						Value:      msg.Value.Value(),
	// 						CreateTime: msg.SourceTimestamp,
	// 						Type:       msg.Value.Type().String(),
	// 					})
	// 					fmt.Println("Received nodes:=========", node.ID)
	// 				}
	// 			}
	// 			for _, v := range nodes {
	// 				select {
	// 				case c.channel <- v:
	// 				default:
	// 					// log.Println("[WARN] Channel is full, dropping message")
	// 					continue
	// 				}
	// 			}

	// 			// log.Printf("[channel ] sub=%d ts=%s node=%s value=%v", sub.SubscriptionID(), msg.SourceTimestamp.UTC().Format(time.RFC3339), msg.NodeID, msg.Value.Value())
	// 		}
	// 		time.Sleep(lag)

	// 	}
	// }
}

func cleanup(ctx context.Context, sub *monitor.Subscription) {
	log.Printf("stats: sub=%d delivered=%d dropped=%d", sub.SubscriptionID(), sub.Delivered(), sub.Dropped())
	sub.Unsubscribe(ctx)
}

func (c *OpcClient) AddNodes(nodes ...OpcNode) {
	for _, v := range nodes {
		if !slices.Contains(c.nodes, v) {
			c.nodes = append(c.nodes, v)
			if c.sub == nil {
				return
			}
			c.sub.AddNodes(c.ctx, v.NodeId)
		}
	}
}

func (c *OpcClient) Control(extends ...*gjson.Json) (err error) {
	for _, v := range extends {
		if c.client != nil && c.isOnline {
			controlType := v.Get("type").String()
			if controlType == "methond" {
				method := v.Get("methodId").String()
				object := v.Get("objectId").String()
				value := v.Get("value").Interface()
				methodID := ua.NewStringNodeID(2, method)
				objectID := ua.NewStringNodeID(2, object)

				// 方法参数
				inputArgs := []*ua.Variant{
					ua.MustVariant(value), // 比如传true表示启动
				}

				resp, err := c.client.Call(c.ctx, &ua.CallMethodRequest{
					ObjectID:       objectID,
					MethodID:       methodID,
					InputArguments: inputArgs,
				})
				if err != nil {
					return fmt.Errorf("call method error: %s", err)
				}

				if resp.StatusCode != ua.StatusOK {
					return fmt.Errorf("method call failed: %v", resp.StatusCode)
				}
			}
			if controlType == "value" {
				id := v.Get("nodeId").String()
				if id == "" {
					return fmt.Errorf("nodeId 为空")
				}
				nodeID, err := ua.ParseNodeID(id)
				if err != nil {
					return fmt.Errorf("节点ID错误")
				}
				value := v.Get("value").Interface()

				fmt.Println("nodeId", nodeID.String())
				// 假设写入值是布尔值true，代表启动
				val, err := variantFromValueByNodeID(c.client, nodeID, value)
				if err != nil {
					fmt.Println(err.Error())
					return fmt.Errorf("创建val失败" + err.Error())
				}

				// 写节点
				req := &ua.WriteRequest{
					NodesToWrite: []*ua.WriteValue{
						{
							NodeID:      nodeID,
							AttributeID: ua.AttributeIDValue,
							Value: &ua.DataValue{
								Value: val,
							},
						},
					},
				}

				resp, err := c.client.Write(c.ctx, req)
				if err != nil {
					// log.Fatalf("Write error: %s", err)
					return fmt.Errorf("write error: %s", err)
				}

				if resp.Results[0] != ua.StatusOK {
					return fmt.Errorf("write failed with status: %v", resp.Results[0])
				}
			}
			return fmt.Errorf("未有控制命令")
		}
	}
	return
}
