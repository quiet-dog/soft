// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageKnowledge is the golang structure for table manage_knowledge.
type ManageKnowledge struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`      // 主键
	Name          string      `json:"name"          orm:"name"           description:"文件名称"`    // 文件名称
	Code          string      `json:"code"          orm:"code"           description:"文件编号"`    // 文件编号
	KnowledgeType string      `json:"knowledgeType" orm:"knowledge_type" description:"知识库类型ID"` // 知识库类型ID
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`     // 创建者
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`     // 更新者
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`    // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`    // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`    // 删除时间
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`      // 备注
	Path          string      `json:"path"          orm:"path"           description:"文件路径"`    // 文件路径
}
