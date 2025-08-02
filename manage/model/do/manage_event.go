// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageEvent is the golang structure of table manage_event for DAO operations like Where/Data.
type ManageEvent struct {
	g.Meta      `orm:"table:manage_event, do:true"`
	Id          interface{} // 主键
	SensorId    interface{} // 传感器Id
	Value       interface{} // 报警时候的数值
	Description interface{} // 报警描述
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Level       interface{} // 报警等级
	Color       interface{} // 报警颜色
	AlarmId     interface{} //
}
