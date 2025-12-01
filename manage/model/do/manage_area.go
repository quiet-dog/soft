// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageArea is the golang structure of table manage_area for DAO operations like Where/Data.
type ManageArea struct {
	g.Meta    `orm:"table:manage_area, do:true"`
	Id        interface{} // 主键
	ParentId  interface{} // 父ID
	Name      interface{} // 区域名称
	Sort      interface{} // 排序
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
	Level     interface{} //
}
