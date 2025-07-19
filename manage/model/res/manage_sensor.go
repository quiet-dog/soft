package res

import (
	"devinggo/manage/model/base"
)

type SensorTableRow struct {
	base.BaseTable
	Name         string `json:"name" description:"传感器名称"`           // 传感器名称
	SensorTypeId int64  `json:"sensorTypeId" description:"传感器类型ID"` // 传感器类型ID
	DeviceId     int64  `json:"deviceId" description:"设备ID"`        // 设备ID
	Extend       string `json:"extend" description:"扩展信息"`          // 扩展信息
	Template     string `json:"template" description:"模板内容"`
}
