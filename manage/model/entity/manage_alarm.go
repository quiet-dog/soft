// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageAlarm is the golang structure for table manage_alarm.
type ManageAlarm struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`     // 主键
	IsLift    float64     `json:"isLift"    orm:"is_lift"    description:"是否解除报警"` // 是否解除报警
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`    // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`    // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`   // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`   // 删除时间
	Level     string      `json:"level"     orm:"level"      description:""`       //
	SensorId  int64       `json:"sensorId"  orm:"sensor_id"  description:"传感器Id"`  // 传感器Id
	Color     string      `json:"color"     orm:"color"      description:"颜色"`     // 颜色
	EndTime   int64       `json:"endTime"   orm:"end_time"   description:"解除报警时间"` // 解除报警时间
	SendTime  int64       `json:"sendTime"  orm:"send_time"  description:"触发时间"`   // 触发时间
}
