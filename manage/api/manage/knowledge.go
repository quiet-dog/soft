package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IndexManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/index" method:"get" tags:"知识库信息表" summary:"分页列表" x-permission:":manageKnowledge:index" `
	model.AuthorHeader
	model.PageListReq
	req.ManageKnowledgeSearch
}

type IndexManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.ManageKnowledge `json:"items"  dc:"list" `
}

type ListManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/list" method:"get" tags:"知识库信息表" summary:"列表" x-permission:":manageKnowledge:list" `
	model.AuthorHeader
	model.ListReq
	req.ManageKnowledgeSearch
}

type ListManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.ManageKnowledge `json:"data"  dc:"list" `
}

type SaveManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/save" method:"post" tags:"知识库信息表" summary:"新增" x-permission:":manageKnowledge:save"`
	model.AuthorHeader
	req.ManageKnowledgeSave
}

type SaveManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id" dc:"id"`
}

type ReadManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/read/{Id}" method:"get" tags:"知识库信息表" summary:"获取单个信息" x-permission:":manageKnowledge:read"`
	model.AuthorHeader
	Id uint64 `json:"id" dc:"知识库信息表 id" v:"required|min:1#Id不能为空"`
}

type ReadManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	Data   res.ManageKnowledge `json:"data" dc:"信息数据"`
}

type UpdateManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/update/{Id}" method:"put" tags:"知识库信息表" summary:"更新" x-permission:":manageKnowledge:update"`
	model.AuthorHeader
	req.ManageKnowledgeUpdate
}

type UpdateManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/delete" method:"delete" tags:"知识库信息表" summary:"删除" x-permission:":manageKnowledge:delete"`
	model.AuthorHeader
	Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/recycle" method:"get" tags:"知识库信息表" summary:"回收站列表" x-permission:":manageKnowledge:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.ManageKnowledgeSearch
}

type RecycleManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.ManageKnowledge `json:"items"  dc:"list" `
}

type RealDeleteManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/realDelete" method:"delete" tags:"知识库信息表" summary:"单个或批量真实删除 （清空回收站）" x-permission:":manageKnowledge:realDelete"`
	model.AuthorHeader
	Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/recovery" method:"put" tags:"知识库信息表" summary:"单个或批量恢复在回收站的" x-permission:":manageKnowledge:recovery"`
	model.AuthorHeader
	Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type ExportManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/export" method:"post" tags:"知识库信息表" summary:"导出" x-permission:":manageKnowledge:export"`
	model.AuthorHeader
	model.ListReq
	req.ManageKnowledgeSearch
}

type ExportManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type ImportManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/import" method:"post" mime:"multipart/form-data" tags:"知识库信息表" summary:"导入" x-permission:":manageKnowledge:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/downloadTemplate" method:"post,get" tags:"知识库信息表" summary:"下载导入模板." x-exceptAuth:"true" x-permission:":manageKnowledge:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/changeStatus" method:"put" tags:"知识库信息表" summary:"更改状态" x-permission:":manageKnowledge:changeStatus"`
	model.AuthorHeader
	Id     uint64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int    `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/numberOperation" method:"put" tags:"知识库信息表" summary:"数字运算操作" x-permission:":manageKnowledge:update"`
	model.AuthorHeader
	Id          uint64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteManageKnowledgeReq struct {
	g.Meta `path:"/manageKnowledge/remote" method:"post" tags:"知识库信息表" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:":manageKnowledge:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteManageKnowledgeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.ManageKnowledge `json:"items"  dc:"list" `
	Data  []res.ManageKnowledge `json:"data"  dc:"list" `
}
