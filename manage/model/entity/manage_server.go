// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageServer is the golang structure for table manage_server.
type ManageServer struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`     // 主键
	Name      string      `json:"name"      orm:"name"       description:"服务器名称"`  // 服务器名称
	Ip        string      `json:"ip"        orm:"ip"         description:"服务器ip"`  // 服务器ip
	Port      string      `json:"port"      orm:"port"       description:"端口"`     // 端口
	Type      string      `json:"type"      orm:"type"       description:"服务器类型"`  // 服务器类型
	Username  string      `json:"username"  orm:"username"   description:"服务器用户名"` // 服务器用户名
	Password  string      `json:"password"  orm:"password"   description:"服务器密码"`  // 服务器密码
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`    // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`    // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`   // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`   // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`     // 备注
	Interval  int64       `json:"interval"  orm:"interval"   description:""`       //
}
