package thirdswmodule

import (
	"context"

	"github.com/gorilla/websocket"
)

type Node struct {
	NodeId string `json:"nodeId"`
	Param  string `json:"param"`
}

type ServerJson struct {
	Url   string `json:"url"`
	Nodes []Node `json:"nodes"`
}

// DeviceDTO represents a device with various properties
type DeviceDTO struct {
	DeviceType           string                  `json:"deviceType"`           // 设备类型
	DeviceID             int64                   `json:"deviceId"`             // 设备ID
	EnvironmentAlarmInfo EnvironmentAlarmInfoDTO `json:"environmentAlarmInfo"` // 环境档案数据信息
	EquipmentInfo        EquipmentInfoDTO        `json:"equipmentInfo"`        // 设备信息
	DateSource           string                  `json:"dateSource"`           // 数据来源
}

// EnvironmentAlarmInfoDTO represents environment alarm information
type EnvironmentAlarmInfoDTO struct {
	EnvironmentID    int64   `json:"environmentId"`    // 设备ID
	Value            float64 `json:"value"`            // 数据
	Unit             string  `json:"unit"`             // 单位
	Power            float64 `json:"power"`            // 功耗
	WaterValue       float64 `json:"waterValue"`       // 用水量
	ElectricityValue float64 `json:"electricityValue"` // 用电量
}

// EquipmentInfoDTO represents equipment information

// EquipmentInfoDTO represents equipment information
type EquipmentInfoDTO struct {
	EquipmentID int64   `json:"equipmentId"` // 设备ID
	ThresholdID int64   `json:"thresholdId"` // 阈值传感器ID
	SensorName  string  `json:"sensorName"`  // 传感器名称
	Value       float64 `json:"value"`       // 传感器值
}

type WebSocketReadWriteCloser struct {
	Conn *websocket.Conn
}

// Read 实现 io.Reader 接口
func (w *WebSocketReadWriteCloser) Read(p []byte) (int, error) {
	_, message, err := w.Conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	n := copy(p, message)
	return n, nil
}

// Write 实现 io.Writer 接口
func (w *WebSocketReadWriteCloser) Write(p []byte) (int, error) {
	err := w.Conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close 实现 io.Closer 接口
func (w *WebSocketReadWriteCloser) Close() error {
	return w.Conn.Close()
}

var url = "ws://127.0.0.1:9020/ws"

// Start 启动 WebSocket 客户端并管理重连
func Start(ctx context.Context) {
	// for {
	// 	if global.DeviceGateway == nil {
	// 		log.Println("device gateway is nil")
	// 		continue
	// 	}
	// 	channels := global.DeviceGateway.RegisterChannel(100)

	// 	var data []ServerJson
	// 	// 读取文件./opc.json 解析
	// 	d, err := os.ReadFile("./opc.json")
	// 	if err == nil {
	// 		err = json.Unmarshal(d, &data)
	// 		if err != nil {
	// 			log.Println("parse opc.json failed:", err)
	// 		}
	// 	}

	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			log.Println("WebSocket client stopped by context")
	// 			return
	// 		default:
	// 		}

	// 		// 尝试连接
	// 		c, _, err := websocket.DefaultDialer.Dial(url, nil)
	// 		if err != nil {
	// 			log.Println("connect failed, retry in 3s:", err)
	// 			time.Sleep(3 * time.Second)
	// 			continue
	// 		}
	// 		log.Println("connected to server")

	// 		// 使用 STOMP 客户端连接
	// 		rwc := WebSocketReadWriteCloser{
	// 			Conn: c,
	// 		}
	// 		stompConn, err := stomp.Connect(&rwc)
	// 		if err != nil {
	// 			log.Println("connect to stomp failed:", err)
	// 			c.Close()
	// 			continue
	// 		}
	// 		defer stompConn.Disconnect()

	// 		// 用 done 通道通知读 goroutine退出
	// 		done := make(chan struct{})

	// 		// 读消息 goroutine
	// 		go func(conn *websocket.Conn) {
	// 			defer close(done)
	// 			for {
	// 				_, message, err := conn.ReadMessage()
	// 				if err != nil {
	// 					log.Println("read error:", err)
	// 					return
	// 				}
	// 				log.Printf("recv: %s", message)
	// 			}
	// 		}(c)

	// 		// 主循环处理发送和断线
	// 		run := true
	// 		for run {
	// 			select {
	// 			case <-ctx.Done():
	// 				log.Println("WebSocket client stopped by context")
	// 				c.Close()
	// 				return
	// 			case <-done:
	// 				log.Println("connection closed, will reconnect...")
	// 				run = false
	// 			case msg := <-channels:
	// 				serverInfo, err := manage.ManageServer().Read(ctx, msg.ServiceId)
	// 				if err != nil {
	// 					log.Println("read server info failed:", err)
	// 					continue
	// 				}

	// 				nodeInfo, err := manage.ManageSensor().Read(ctx, msg.Value.ID)
	// 				if err != nil {
	// 					log.Println("read node info failed:", err)
	// 					continue
	// 				}

	// 				nodeId, err := manage.ManageOpc().Read(ctx, nodeInfo.Extend.Get("id").Int64())
	// 				if err != nil {
	// 					log.Println("read node id failed:", err)
	// 					continue
	// 				}

	// 				t, err := manage.ManageSensorTemplateCache().Get(context.Background(), msg.Value.ID)
	// 				if err != nil {
	// 					continue
	// 				}

	// 				for _, s := range data {
	// 					if s.Url == serverInfo.Ip {
	// 						for _, node := range s.Nodes {
	// 							if node.NodeId == nodeId.NodeId {

	// 								result := DeviceDTO{}
	// 								r := strings.Split(node.Param, "-")

	// 								if len(r) >= 4 {
	// 									for i := 0; i < len(r); i++ {
	// 										if r[i] == "deviceType" {
	// 											result.DeviceType = r[i+1]
	// 										} else if r[i] == "environmentId" {
	// 											if id, err := strconv.Atoi(r[i+1]); err == nil {
	// 												result.EnvironmentAlarmInfo.EnvironmentID = int64(id)
	// 											}
	// 										} else if r[i] == "thresholdId" {
	// 											if id, err := strconv.Atoi(r[i+1]); err == nil {
	// 												result.EquipmentInfo.ThresholdID = int64(id)
	// 											}
	// 										} else if r[i] == "equipmentId" {
	// 											if id, err := strconv.Atoi(r[i+1]); err == nil {
	// 												result.EquipmentInfo.EquipmentID = int64(id)
	// 											}
	// 										}
	// 									}
	// 								}
	// 								// 发送数据处理
	// 								if result.DeviceType == "设备档案" {
	// 									result.EquipmentInfo.Value, _ = t.ToExprValueFloat64(msg.Value.Value)
	// 								} else if result.DeviceType == "环境档案" {
	// 									result.EnvironmentAlarmInfo.Value, _ = t.ToExprValueFloat64(msg.Value.Value)
	// 								}
	// 								result.DateSource = msg.Value.CreateTime.Format("2006-01-02 15:04:05")
	// 								sd, err := gjson.Marshal(result)
	// 								if err != nil {
	// 									log.Println("marshal result failed:", err)
	// 									continue
	// 								}
	// 								c.WriteMessage(websocket.TextMessage, sd)
	// 							}
	// 						}
	// 					}
	// 				}

	// 			}
	// 		}

	// 		// 等几秒再重连
	// 		c.Close()
	// 		time.Sleep(3 * time.Second)
	// 	}
	// }
}
