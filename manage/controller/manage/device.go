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
	DeviceController = deviceController{}
)

type deviceController struct {
	base.BaseController
}

func (c *deviceController) IndexDevice(ctx context.Context, in *manage.IndexDeviceReq) (out *manage.IndexDeviceRes, err error) {
	out = &manage.IndexDeviceRes{}
	items, totalCount, err := sManage.ManageDevice().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageDeviceSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.DeviceTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *deviceController) SaveDevice(ctx context.Context, in *manage.SaveDeviceReq) (out *manage.SaveDeviceRes, err error) {
	out = &manage.SaveDeviceRes{}
	id, err := sManage.ManageDevice().Save(ctx, &in.ManageDeviceSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *areaController) DeleteDevice(ctx context.Context, in *manage.DeleteDeviceReq) (out *manage.DeleteDeviceRes, err error) {
	out = &manage.DeleteDeviceRes{}
	err = sManage.ManageDevice().Delete(ctx, in.Ids)
	return
}

func (c *deviceController) TreeDevice(ctx context.Context, in *manage.TreeDeviceReq) (out *manage.TreeDeviceRes, err error) {
	out = &manage.TreeDeviceRes{}
	out.Data, err = sManage.ManageDevice().Tree(ctx, &in.PageListReq, &in.ManageDeviceSearch)
	return
}

func (c *deviceController) Read(ctx context.Context, in *manage.ReadDeviceReq) (out *manage.ReadDeviceRes, err error) {
	out = &manage.ReadDeviceRes{}
	item, err := sManage.ManageDevice().Read(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = item
	return
}

func (c *deviceController) TestConnect(ctx context.Context, in *manage.DeviceTestConnectReq) (out *manage.DeviceTestConnectRes, err error) {
	out = &manage.DeviceTestConnectRes{}
	err = sManage.ManageDevice().TestConnect(ctx, &in.DeviceTestConnectReq)
	return
}
