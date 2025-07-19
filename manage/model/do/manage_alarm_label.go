// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageAlarmLabel is the golang structure of table manage_alarm_label for DAO operations like Where/Data.
type ManageAlarmLabel struct {
	g.Meta    `orm:"table:manage_alarm_label, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 标签名称
	Color     interface{} // 标签颜色
	Level     interface{} // 标签等级
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
