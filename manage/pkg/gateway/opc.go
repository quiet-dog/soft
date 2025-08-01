package gateway

import (
	"context"
	"fmt"
	"log"
	"slices"
	"sync"
	"time"

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
		opcua.SecurityPolicy(""),
		opcua.SecurityModeString(""),
		opcua.CertificateFile(""),
		opcua.PrivateKeyFile(""),
		opcua.AuthAnonymous(),
		opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
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
	wg := &sync.WaitGroup{}

	c.isOnline = true

	// // start callback-based subscription
	// wg.Add(1)
	// go startCallbackSub(ctx, m, 2*time.Second, 0, wg, "ns=3;i=3")

	// // start channel-based subscription
	wg.Add(1)
	go c.startChanSub(ctx, m, c.conf.SubTime, 0, wg)

	wg.Wait()
	c.isOnline = false
	return
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

	defer cleanup(ctx, sub, wg)

	<-ctx.Done()
}

func (c *OpcClient) startChanSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, wg *sync.WaitGroup, nodes ...string) {
	ch := make(chan *monitor.DataChangeMessage, 16)
	for _, v := range c.nodes {
		nodes = append(nodes, v.NodeId)
	}
	sub, err := m.ChanSubscribe(ctx, &opcua.SubscriptionParameters{Interval: interval}, ch, nodes...)
	if err != nil {
		log.Fatal(err)
	}
	c.sub = sub
	defer cleanup(ctx, sub, wg)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(5 * time.Second):
			if err = c.TestPing(); err != nil {
				c.isOnline = false
				c.client.Close(ctx)
				c.cancel()
				return
			}
		case msg := <-ch:
			fmt.Println("===============================来自通道信息")
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
	}
}

func cleanup(ctx context.Context, sub *monitor.Subscription, wg *sync.WaitGroup) {
	log.Printf("stats: sub=%d delivered=%d dropped=%d", sub.SubscriptionID(), sub.Delivered(), sub.Dropped())
	sub.Unsubscribe(ctx)
	wg.Done()
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
