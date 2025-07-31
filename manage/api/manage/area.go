package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type AreaReq struct {
	g.Meta `path:"/area" tags:"Area" method:"get" summary:"获取区域列表"`
}

type AreaRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AreaTableRow `json:"items"  dc:"area list" `
}

type IndexAreaReq struct {
	g.Meta `path:"/area/index" tags:"Area" method:"get" summary:"获取区域列表" x-permission:"manage:area:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageAreaSearch
}

type IndexAreaRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AreaTableRow `json:"items"  dc:"area list" `
}

type SaveAreaReq struct {
	g.Meta `path:"/area/save" tags:"Area" method:"post" summary:"保存区域" x-permission:"manage:area:save"`
	model.AuthorHeader
	req.ManageAreaSave
}

type SaveAreaRes struct {
	g.Meta `mime:"application/json" description:"保存区域结果"`
	Id     int64 `json:"id" description:"区域ID"` // 区域ID
}

type AreaTreeReq struct {
	g.Meta `path:"/area/tree" tags:"Area" method:"get" summary:"获取区域树"`
	model.AuthorHeader
	req.ManageAreaSearch
}

type AreaTreeRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"area tree list" `
}

type DeleteAreaReq struct {
	g.Meta `path:"/area/delete" tags:"Area" method:"delete" summary:"删除区域" x-permission:"manage:area:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"区域ID列表" v:"required#区域ID列表不能为空"`
}

type DeleteAreaRes struct {
	g.Meta `mime:"application/json" description:"删除区域结果"`
}

type UpdateAreaReq struct {
	g.Meta `path:"/area/update/{Id}" method:"put" tags:"区域管理" summary:"更新区域信息." x-permission:"manage:area:update"`
	model.AuthorHeader
	req.ManageAreaUpdateInfo
}

type UpdateAreaRes struct {
	g.Meta `mime:"application/json"`
}

type AllTreeByIdReq struct {
	g.Meta `path:"/area/treeTop" method:"get" tags:"区域管理" summary:"更新区域信息." x-permission:"manage:area:treeTop"`
	model.AuthorHeader
	req.ManageAreaSearch
}

type AllTreeByIdRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"area tree list" `
}
