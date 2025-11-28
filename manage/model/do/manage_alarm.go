// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageAlarm is the golang structure of table manage_alarm for DAO operations like Where/Data.
type ManageAlarm struct {
	g.Meta    `orm:"table:manage_alarm, do:true"`
	Id        interface{} // 主键
	IsLift    interface{} // 是否解除报警
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Level     interface{} //
	SensorId  interface{} // 传感器Id
	Color     interface{} // 颜色
	EndTime   interface{} // 解除报警时间
	SendTime  interface{} // 触发时间
}
