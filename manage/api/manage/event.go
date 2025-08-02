package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type IndexEventReq struct {
	g.Meta `path:"/event/index" tags:"Event" method:"get" summary:"获取设备事件列表" x-permission:"manage:event:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageEventSearch
}

type IndexEventRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.EventTableRow `json:"items"  dc:"event list" `
}
