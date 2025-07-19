package req

import "devinggo/manage/model/base"

type ManageSensorTypeSearch struct {
	Name string `json:"name" description:"设备名称"` // 设备名称
}

type ManageSensorTypeSave struct {
	Name   string `json:"name" v:"required|max-length:255" description:"设备名称"` // 设备名称
	Unit   string `json:"unit" v:"required|max-length:50" description:"单位"`    // 单位
	Remark string `json:"remark" description:"备注"`                             // 备注
}

type ManageSensorTypeDelete struct {
	base.BaseIds
}
