package manage

import (
	"context"
	"devinggo/manage/api/manage"
	"devinggo/manage/model/res"
	sManage "devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	EventController = eventController{}
)

type eventController struct {
	base.BaseController
}

func (c *eventController) IndexEvent(ctx context.Context, in *manage.IndexEventReq) (out *manage.IndexEventRes, err error) {
	out = &manage.IndexEventRes{}
	items, totalCount, err := sManage.ManageEvent().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageEventSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.EventTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}
