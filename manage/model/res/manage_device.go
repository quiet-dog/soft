package res

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type DeviceTableRow struct {
	base.BaseTable
	Name                 string      `json:"name" description:"设备名称"`                 // 设备名称
	Manufacturer         string      `json:"manufacturer" description:"制造商"`          // 制造商
	Model                string      `json:"model" description:"设备型号"`                // 设备型号
	InstallationLocation string      `json:"installationLocation" description:"安装位置"` // 安装位置
	AreaId               int64       `json:"areaId" description:"区域ID"`               // 区域ID
	ServerId             int64       `json:"serverId" description:"服务器ID"`            // 服务器ID
	Extend               *gjson.Json `json:"extend" description:"扩展信息"`               // 扩展信息
	Remark               string      `json:"remark" description:"备注"`                 // 备注

	AreaName   string `json:"areaName" description:"区域名称"`    // 区域名称
	ServerName string `json:"serverName" description:"服务器名称"` // 服务器名称
	IsOnline   bool   `json:"isOnline" description:"是否在线"`    // 是否在线
	OnlineTime string `json:"onlineTime" description:"在线时间"`  // 在线时间
}

type DeviceInfo struct {
	DeviceTableRow
	Server    *ServerTableRow `json:"server" description:"服务器信息"`   // 服务器信息
	ModelPath string          `json:"modelPath" description:"模型路径"` // 模型路径
}

type DeviceSensorInfo struct {
	DeviceInfo
	Sensors []*SensorTableRow `json:"sensors" description:"传感器列表"` // 模型路径
}

type DeviceSensorInfoTableRow struct {
	Rows  []*DeviceSensorInfo `json:"rows" description:"设备传感器信息列表"`
	Total int                 `json:"total" description:"总条数"`
}
