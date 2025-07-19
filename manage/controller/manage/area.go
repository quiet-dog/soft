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
	AreaController = areaController{}
)

type areaController struct {
	base.BaseController
}

func (c *areaController) IndexArea(ctx context.Context, in *manage.IndexAreaReq) (out *manage.IndexAreaRes, err error) {
	out = &manage.IndexAreaRes{}
	items, totalCount, err := sManage.ManageArea().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageAreaSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.AreaTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *areaController) SaveArea(ctx context.Context, in *manage.SaveAreaReq) (out *manage.SaveAreaRes, err error) {
	out = &manage.SaveAreaRes{}
	id, err := sManage.ManageArea().Save(ctx, &in.ManageAreaSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *areaController) DeleteArea(ctx context.Context, in *manage.DeleteAreaReq) (out *manage.DeleteAreaRes, err error) {
	out = &manage.DeleteAreaRes{}
	err = sManage.ManageArea().Delete(ctx, in.Ids)
	return
}

func (c *areaController) Tree(ctx context.Context, in *manage.AreaTreeReq) (out *manage.AreaTreeRes, err error) {
	out = &manage.AreaTreeRes{}
	items, err := sManage.ManageArea().Tree(ctx, &in.ManageAreaSearch)
	if err != nil {
		return
	}
	out.Data = items
	return
}
