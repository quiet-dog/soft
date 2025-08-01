// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ManageThreshold is the golang structure of table manage_threshold for DAO operations like Where/Data.
type ManageThreshold struct {
	g.Meta       `orm:"table:manage_threshold, do:true"`
	SensorId     interface{} // 传感器id
	AlarmLabelId interface{} // 报警标签id
	Template     interface{} // expr表达式
	Sort         interface{} // 优先级
}
