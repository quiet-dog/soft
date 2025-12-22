// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ManageKnowledge is the golang structure of table manage_knowledge for DAO operations like Where/Data.
type ManageKnowledge struct {
	g.Meta        `orm:"table:manage_knowledge, do:true"`
	Id            interface{} // 主键
	Name          interface{} // 文件名称
	Code          interface{} // 文件编号
	KnowledgeType interface{} // 知识库类型ID
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	Remark        interface{} // 备注
	Path          interface{} // 文件路径
}
