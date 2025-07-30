package req

type ManageServerSearch struct {
	Name string `json:"name" description:"服务器名称"` // 服务器名称
	Ip   string `json:"ip" description:"服务器IP"`   // 服务器IP
	Type string `json:"type" description:"服务器类型"` // 服务器类型
}

type ManageServerSave struct {
	Name     string `json:"name" v:"required|max-length:255" description:"服务器名称"` // 服务器名称
	Ip       string `json:"ip" v:"required|ip" description:"服务器IP"`               // 服务器IP
	Type     string `json:"type" v:"required" description:"服务器类型"`                // 服务器类型
	Port     string `json:"port" description:"服务器端口"`                             // 服务器端口
	Username string `json:"username" description:"服务器用户名"`                        // 服务器用户名
	Password string `json:"password" description:"服务器密码"`                         // 服务器密码
	Interval int64  `json:"interval" description:"采集间隔"`                          // 采集间隔
	Remark   string `json:"remark" description:"备注"`                              // 备注
}
