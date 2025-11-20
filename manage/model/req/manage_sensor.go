package req

import (
	"devinggo/manage/model/base"
	"devinggo/manage/model/common"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type ManageSensorSearch struct {
	Name         string  `json:"name" description:"传感器名称"`           // 传感器名称
	SensorTypeId int64   `json:"sensorTypeId" description:"传感器类型ID"` // 传感器类型ID
	DeviceId     int64   `json:"deviceId" description:"设备ID"`        // 设备ID
	DeviceIds    []int64 `json:"deviceIds" description:"设备ID列表"`     // 设备ID列表
}

type ManageSensorSave struct {
	Id           int64       `json:"id"`
	Name         string      `json:"name" v:"required|max-length:255" description:"传感器名称"` // 传感器名称
	SensorTypeId int64       `json:"sensorTypeId" v:"required" description:"传感器类型ID"`      // 传感器类型ID
	DeviceId     int64       `json:"deviceId" v:"required" description:"设备ID"`             // 设备ID
	Extend       *gjson.Json `json:"extend" v:"required" description:"扩展信息"`               // 扩展信息
	Template     string      `json:"template" description:"模板内容"`
	Remark       string      `json:"remark" description:"备注"` // 备注
	Unit         string      `json:"unit" description:"单位"`   // 单位
}

type ManageSensorDelete struct {
	base.BaseIds
}

type ManageSensorReadData struct {
	Extend *gjson.Json `json:"extend" v:"required" description:"扩展信息"`
	Type   string      `json:"type" v:"required" description:"设备类型"`
}

type ManageSensorTranslate struct {
	Env      common.TemplateEnv `json:"env" description:"扩展信息"`
	Template string             `json:"template" description:"expr模板"`
}
