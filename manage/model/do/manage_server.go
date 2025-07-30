// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageServer is the golang structure of table manage_server for DAO operations like Where/Data.
type ManageServer struct {
	g.Meta    `orm:"table:manage_server, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 服务器名称
	Ip        interface{} // 服务器ip
	Port      interface{} // 端口
	Type      interface{} // 服务器类型
	Username  interface{} // 服务器用户名
	Password  interface{} // 服务器密码
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
	Interval  interface{} //
}
