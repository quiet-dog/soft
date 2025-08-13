package gateway

import (
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
)

const (
	SERVER_OPC                 = "opc"
	SERVER_MODBUS_TCP          = "modbus_tcp"
	SERVER_MODBUS_RTU          = "modbus_rtu"
	SERVER_MODBUS_RTU_OVER_TCP = "modbus_rtu_over_tcp"
	SERVER_MQTT                = "mqtt"
)

type Config struct {
	ServiceId int64         `json:"serviceId"`
	Host      string        `json:"host"`
	Port      string        `json:"port"`
	Type      string        `json:"type"`
	Extend    *gjson.Json   `json:"extend"`
	SubTime   time.Duration `json:"subTime"` // 订阅时间间隔
}

type Msg struct {
	Value     Value `json:"value"`
	ServiceId int64 `json:"serviceId"`
}

type Value struct {
	ID         int64       `json:"id"`
	Value      interface{} `json:"value"`
	CreateTime time.Time   `json:"createTime"`
	Type       string      `json:"type"`
	DeviceId   int64       `json:"deviceId"`
}

func (v *OpcNodes) FindByNodeId(nodeId string) *OpcNode {
	for _, node := range *v {
		if node.NodeId == nodeId {
			return &node
		}
	}
	return nil
}
