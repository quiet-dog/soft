package req

type ManageKnowledgeSearch struct {
	Name string `json:"name" description:"文件名称" `

	Code string `json:"code" description:"文件编号" `

	KnowledgeType string `json:"knowledge_type" description:"知识库类型ID" `
}

type ManageKnowledgeSave struct {
	Name string `json:"name"  v:"required"  description:"文件名称" `

	Code string `json:"code"  v:"required"  description:"文件编号" `

	KnowledgeType string `json:"knowledge_type"  v:"required"  description:"知识库类型ID" `

	Remark string `json:"remark"  description:"备注" `

	Path string `json:"path"  v:"required"  description:"上传文件" `
}

type ManageKnowledgeUpdate struct {
	Id int64 `json:"id"  description:"主键" `

	Name string `json:"name"  v:"required"  description:"文件名称" `

	Code string `json:"code"  v:"required"  description:"文件编号" `

	KnowledgeType string `json:"knowledge_type"  v:"required"  description:"知识库类型ID" `

	Remark string `json:"remark"  description:"备注" `

	Path string `json:"path"  v:"required"  description:"上传文件" `
}
