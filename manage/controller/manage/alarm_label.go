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
	AlarmLabelController = alarmLabelController{}
)

type alarmLabelController struct {
	base.BaseController
}

func (c *alarmLabelController) IndexAlarmLabel(ctx context.Context, in *manage.IndexAlarmLabelReq) (out *manage.IndexAlarmLabelRes, err error) {
	out = &manage.IndexAlarmLabelRes{}
	items, totalCount, err := sManage.ManageAlarmLabel().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageAlarmLabelSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.AlarmLabelTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *alarmLabelController) SaveAlarmLabel(ctx context.Context, in *manage.SaveAlarmLabelReq) (out *manage.SaveAlarmLabelRes, err error) {
	out = &manage.SaveAlarmLabelRes{}
	id, err := sManage.ManageAlarmLabel().Save(ctx, &in.ManageAlarmLabelSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *alarmLabelController) DeleteAlarmLabel(ctx context.Context, in *manage.DeleteAlarmLableReq) (out *manage.DeleteAlarmLabelRes, err error) {
	out = &manage.DeleteAlarmLabelRes{}
	err = sManage.ManageAlarmLabel().Delete(ctx, in.Ids)
	return
}
