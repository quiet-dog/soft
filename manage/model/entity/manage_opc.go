// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageOpc is the golang structure for table manage_opc.
type ManageOpc struct {
	Id              int64       `json:"id"              orm:"id"               description:"主键"`   // 主键
	NodeId          string      `json:"nodeId"          orm:"node_id"          description:"标签名称"` // 标签名称
	ServerId        int64       `json:"serverId"        orm:"server_id"        description:"服务id"` // 服务id
	NodeClass       string      `json:"nodeClass"       orm:"node_class"       description:"变量名称"` // 变量名称
	Type            string      `json:"type"            orm:"type"             description:"变量类型"` // 变量类型
	CreatedBy       int64       `json:"createdBy"       orm:"created_by"       description:"创建者"`  // 创建者
	UpdatedBy       int64       `json:"updatedBy"       orm:"updated_by"       description:"更新者"`  // 更新者
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"` // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:"删除时间"` // 删除时间
	Remark          string      `json:"remark"          orm:"remark"           description:"备注"`   // 备注
	ParentId        int64       `json:"parentId"        orm:"parent_id"        description:""`     //
	NamespacesIndex int         `json:"namespacesIndex" orm:"namespaces_index" description:""`     //
	BrowseName      string      `json:"browseName"      orm:"browse_name"      description:""`     //
}
