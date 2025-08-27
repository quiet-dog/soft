package res

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type SensorTableRow struct {
	base.BaseTable
	Name           string      `json:"name" description:"传感器名称"`           // 传感器名称
	SensorTypeId   int64       `json:"sensorTypeId" description:"传感器类型ID"` // 传感器类型ID
	SensorTypeName string      `json:"sensorTypeName" description:"传感器名称"` // 传感器名称
	DeviceId       int64       `json:"deviceId" description:"设备ID"`        // 设备ID
	DeviceName     string      `json:"deviceName" description:"设备名称"`      // 设备名称
	Extend         *gjson.Json `json:"extend" description:"扩展信息"`          // 扩展信息
	Template       string      `json:"template" description:"模板内容"`        // 模板内容
	Value          any         `json:"value" description:"当前值"`            // 当前值
	IsOnline       bool        `json:"isOnline" description:"是否在线"`        // 是否在线
}

type SensorInfo struct {
	SensorTableRow
}
