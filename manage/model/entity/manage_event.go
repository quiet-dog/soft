// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageEvent is the golang structure for table manage_event.
type ManageEvent struct {
	Id          int64       `json:"id"          orm:"id"          description:"主键"`      // 主键
	SensorId    int64       `json:"sensorId"    orm:"sensor_id"   description:"传感器Id"`   // 传感器Id
	Value       float64     `json:"value"       orm:"value"       description:"报警时候的数值"` // 报警时候的数值
	Description string      `json:"description" orm:"description" description:"报警描述"`    // 报警描述
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`     // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`     // 更新者
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`    // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`    // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`    // 删除时间
	Level       string      `json:"level"       orm:"level"       description:"报警等级"`    // 报警等级
	Color       string      `json:"color"       orm:"color"       description:"报警颜色"`    // 报警颜色
	AlarmId     int64       `json:"alarmId"     orm:"alarm_id"    description:""`        //
}
