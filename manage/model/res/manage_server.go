package res

import "devinggo/manage/model/base"

type ServerTableRow struct {
	base.BaseTable
	Name     string `json:"name" description:"服务器名称"`   // 服务器名称
	Ip       string `json:"ip" description:"IP地址"`      // IP地址
	Port     string `json:"port" description:"端口"`      // 端口
	Type     string `json:"type" description:"服务器类型"`   // 服务器类型
	Username string `json:"username" description:"用户名"` // 用户名
	Password string `json:"password" description:"密码"`  // 密码
	Remark   string `json:"remark" description:"备注"`    // 备注
}

type ServerType struct {
	Label string `json:"label" description:"服务器类型名称"` // 服务器类型名称
	Value string `json:"value" description:"服务器类型值"`  // 服务器类型值
}

type ServerInfo struct {
	ServerTableRow
}
