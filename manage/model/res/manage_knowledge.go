package res

import "github.com/gogf/gf/v2/os/gtime"

type ManageKnowledge struct {
	Name string `json:"name"  description:"文件名称" `

	Code string `json:"code"  description:"文件编号" `

	KnowledgeType string `json:"knowledge_type"  description:"知识库类型ID" `

	Remark string `json:"remark"  description:"备注" `

	Id uint64 `json:"id"  description:"主键" `

	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间" `

	Path string `json:"path"  description:"文件路径" `
}

type ManageKnowledgeExcel struct {
	Name string `json:"name"  v:"required"  description:"文件名称"  excelName:"文件名称" excelIndex:"1"  `

	Code string `json:"code"  v:"required"  description:"文件编号"  excelName:"文件编号" excelIndex:"2"  `

	KnowledgeType string `json:"knowledge_type"  v:"required"  description:"知识库类型ID"  excelName:"知识库类型ID" excelIndex:"3"  `

	Remark string `json:"remark"  description:"备注"  excelName:"备注" excelIndex:"9"  `

	Path string `json:"path"  v:"required"  description:"上传文件"  excelName:"上传文件" excelIndex:"10"  `
}
