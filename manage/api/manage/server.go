package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type ServerReq struct {
	g.Meta `path:"/server" tags:"Server" method:"get" summary:"获取设备服务器列表"`
}

type ServerRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.ServerTableRow `json:"items"  dc:"server list" `
}

type IndexServerReq struct {
	g.Meta `path:"/server/index" tags:"Server" method:"get" summary:"获取设备服务器列表" x-permission:"manage:server:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageServerSearch
}

type IndexServerRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.ServerTableRow `json:"items"  dc:"server list" `
}

type SaveServerReq struct {
	g.Meta `path:"/server/save" tags:"Server" method:"post" summary:"保存设备服务器" x-permission:"manage:server:save"`
	model.AuthorHeader
	req.ManageServerSave
}

type SaveServerRes struct {
	g.Meta `mime:"application/json" description:"保存设备服务器结果"`
	Id     int64 `json:"id" description:"设备服务器ID"` // 设备服务器ID
}

type DeleteServerReq struct {
	g.Meta `path:"/server/delete" tags:"Server" method:"delete" summary:"删除设备服务器" x-permission:"manage:server:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"设备服务器ID列表" v:"required#设备服务器ID列表不能为空"`
}

type DeleteServerRes struct {
	g.Meta `mime:"application/json"`
}

type ServerTypesReq struct {
	g.Meta `path:"/server/types" tags:"Server" method:"get" summary:"获取设备服务器类型列表"`
	model.AuthorHeader
}

type ServerTypesRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.ServerType `json:"data"  dc:"server type list" `
}

type TreeServerReq struct {
	g.Meta `path:"/server/tree" tags:"Server" method:"get" summary:"获取设备服务器树"`
	model.AuthorHeader
	req.ManageServerSearch
	model.PageListReq
}

type TreeServerRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"server tree list" `
}

type ReadServerReq struct {
	g.Meta `path:"/server/read/{Id}" method:"get" tags:"设备" summary:"获取服务信息." x-permission:"system:server:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"服务 id" v:"required|min:1#设备Id不能为空"`
}

type ReadServerRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.ServerInfo `json:"data" dc:"服务信息"`
}

type UpdateServerReq struct {
	g.Meta `path:"/server/update/{Id}" method:"put" tags:"服务管理" summary:"更新服务信息." x-permission:"manage:server:update"`
	model.AuthorHeader
	req.ManageServerUpdateInfo
}

type UpdateServerRes struct {
	g.Meta `mime:"application/json"`
}
