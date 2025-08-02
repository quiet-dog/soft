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
	AlarmController = alarmController{}
)

type alarmController struct {
	base.BaseController
}

func (c *alarmController) IndexAlarm(ctx context.Context, in *manage.IndexAlarmReq) (out *manage.IndexAlarmRes, err error) {
	out = &manage.IndexAlarmRes{}
	items, totalCount, err := sManage.ManageAlarm().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageAlarmSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.AlarmTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *alarmController) SaveAlarm(ctx context.Context, in *manage.SaveAlarmReq) (out *manage.SaveAlarmRes, err error) {
	out = &manage.SaveAlarmRes{}
	id, err := sManage.ManageAlarm().Save(ctx, in.ManageAlarmSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *alarmController) DeleteAlarm(ctx context.Context, in *manage.DeleteAlarmReq) (out *manage.DeleteAlarmRes, err error) {
	out = &manage.DeleteAlarmRes{}
	err = sManage.ManageAlarm().Delete(ctx, in.Ids)
	return
}
