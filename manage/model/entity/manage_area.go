// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageArea is the golang structure for table manage_area.
type ManageArea struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`   // 主键
	ParentId  int64       `json:"parentId"  orm:"parent_id"  description:"父ID"`  // 父ID
	Name      string      `json:"name"      orm:"name"       description:"区域名称"` // 区域名称
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`   // 排序
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`  // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`  // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`   // 备注
	Level     string      `json:"level"     orm:"level"      description:""`     //
}
