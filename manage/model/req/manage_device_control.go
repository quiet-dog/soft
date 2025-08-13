package req

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type ManageDeviceControlSave struct {
	DeviceId int64       `json:"deviceId" v:"required" description:"设备ID"` // 设备ID
	Extend   *gjson.Json `json:"extend" v:"required" description:"扩展信息"`   // 扩展信息
	Name     string      `json:"name" description:"控制名称"`                  // 控制名称
}

type ManageDeviceControlInfo struct {
	ManageDeviceControlSave
	base.BaseId
}

type ManageAddDeviceControlInfo struct {
	Command []*ManageDeviceControlInfo `json:"command"`
}

type ManageDeviceControlSearch struct {
	DeviceIds []int64 `json:"deviceIds" description:"设备ID列表"` // 设备ID列表
}
