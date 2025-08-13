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
	DeviceControlController = deviceControlController{}
)

type deviceControlController struct {
	base.BaseController
}

func (c *deviceControlController) IndexDeviceControlCon(ctx context.Context, in *manage.IndexDeviceControlReq) (out *manage.IndexDeviceControlRes, err error) {
	out = &manage.IndexDeviceControlRes{}
	items, totalCount, err := sManage.ManageDeviceControl().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageDeviceControlSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.DeviceControlTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *deviceControlController) SaveDeviceControl(ctx context.Context, in *manage.SaveDeviceControlReq) (out *manage.SaveDeviceControlRes, err error) {
	out = &manage.SaveDeviceControlRes{}
	id, err := sManage.ManageDeviceControl().Save(ctx, &in.ManageDeviceControlSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *deviceControlController) DeleteDeviceControl(ctx context.Context, in *manage.DeleteDeviceControlReq) (out *manage.DeleteDeviceControlRes, err error) {
	out = &manage.DeleteDeviceControlRes{}
	err = sManage.ManageDeviceControl().Delete(ctx, in.Ids)
	return
}

func (c *deviceControlController) AddControl(ctx context.Context, in *manage.SaveAddDeviceControlReq) (out *manage.SaveAddDeviceControlRes, err error) {
	out = &manage.SaveAddDeviceControlRes{}
	err = sManage.ManageDeviceControl().AddControl(ctx, &in.ManageAddDeviceControlInfo)
	return
}

func (c *deviceControlController) Control(ctx context.Context, in *manage.DeviceControlSendReq) (out *manage.DeleteDeviceControlRes, err error) {
	out = &manage.DeleteDeviceControlRes{}
	err = sManage.ManageDeviceControl().Control(ctx, in.Id)
	return
}
