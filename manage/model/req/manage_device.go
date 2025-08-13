package req

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type ManageDeviceSearch struct {
	Name                 string `json:"name" description:"设备名称"`                 // 设备名称
	Manufacturer         string `json:"manufacturer" description:"制造商"`          // 制造商
	Model                string `json:"model" description:"设备型号"`                // 设备型号
	InstallationLocation string `json:"installationLocation" description:"安装位置"` // 安装位置
	ServerId             int64  `json:"serverId" description:"服务器ID"`            // 服务器ID
}

type ManageDeviceSave struct {
	Name                 string      `json:"name" v:"required|max-length:255" description:"设备名称"`                 // 设备名称
	Manufacturer         string      `json:"manufacturer" v:"required|max-length:255" description:"制造商"`          // 制造商
	Model                string      `json:"model" v:"required|max-length:255" description:"设备型号"`                // 设备型号
	InstallationLocation string      `json:"installationLocation" v:"required|max-length:255" description:"安装位置"` // 安装位置
	Extend               *gjson.Json `json:"extend" description:"扩展信息"`                                           // 扩展信息
	AreaId               int64       `json:"areaId" v:"required" description:"区域ID"`                              // 区域ID
	ServerId             int64       `json:"serverId" v:"required" description:"服务器ID"`                           // 服务器ID
	Remark               string      `json:"remark" description:"备注"`                                             // 备注
}

type ManageDeviceDelete struct {
	base.BaseIds
}

type DeviceTestConnectReq struct {
	ServerId int64       `json:"serverId" v:"required|max-length:255" description:"服务id"` // 设备id
	Extend   *gjson.Json `json:"extend" description:"扩展信息"`                               // 扩展信息
}

type DeviceImportModelReq struct {
	Path     string `json:"path"`
	DeviceId int64  `json:"deviceId" v:"required" description:"服务id"` // 设备id
}
