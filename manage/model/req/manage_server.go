package req

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type ManageServerSearch struct {
	Name      string  `json:"name" description:"服务器名称"`       // 服务器名称
	Ip        string  `json:"ip" description:"服务器IP"`         // 服务器IP
	Type      string  `json:"type" description:"服务器类型"`       // 服务器类型
	DeviceIds []int64 `json:"deviceIds" description:"设备ID列表"` // 设备ID列表
}

type ManageServerSave struct {
	Name     string      `json:"name" v:"required|max-length:255" description:"服务器名称"` // 服务器名称
	Ip       string      `json:"ip" v:"required|ip" description:"服务器IP"`               // 服务器IP
	Type     string      `json:"type" v:"required" description:"服务器类型"`                // 服务器类型
	Port     string      `json:"port" description:"服务器端口"`                             // 服务器端口
	Extend   *gjson.Json `json:"extend" description:"扩展信息"`                            // 扩展信息
	Username string      `json:"username" description:"服务器用户名"`                        // 服务器用户名
	Password string      `json:"password" description:"服务器密码"`                         // 服务器密码
	Interval int64       `json:"interval" description:"采集间隔"`                          // 采集间隔
	Remark   string      `json:"remark" description:"备注"`                              // 备注
}

type ManageServerUpdateInfo struct {
	ManageServerSave
	base.BaseId
}
