// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageAlarmLabel is the golang structure for table manage_alarm_label.
type ManageAlarmLabel struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`   // 主键
	Name      string      `json:"name"      orm:"name"       description:"标签名称"` // 标签名称
	Color     string      `json:"color"     orm:"color"      description:"标签颜色"` // 标签颜色
	Level     string      `json:"level"     orm:"level"      description:"标签等级"` // 标签等级
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`  // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`  // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`   // 备注
}
