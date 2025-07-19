// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageOpc is the golang structure of table manage_opc for DAO operations like Where/Data.
type ManageOpc struct {
	g.Meta          `orm:"table:manage_opc, do:true"`
	Id              interface{} // 主键
	NodeId          interface{} // 标签名称
	ServerId        interface{} // 服务id
	NodeClass       interface{} // 变量名称
	Type            interface{} // 变量类型
	CreatedBy       interface{} // 创建者
	UpdatedBy       interface{} // 更新者
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
	Remark          interface{} // 备注
	ParentId        interface{} //
	NamespacesIndex interface{} //
	BrowseName      interface{} //
}
