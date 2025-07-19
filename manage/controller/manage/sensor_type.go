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
	SensorTypeController = sensorTypeController{}
)

type sensorTypeController struct {
	base.BaseController
}

func (c *sensorTypeController) IndexSensorType(ctx context.Context, in *manage.IndexSensorTypeReq) (out *manage.IndexSensorTypeRes, err error) {
	out = &manage.IndexSensorTypeRes{}
	items, totalCount, err := sManage.ManageSensorType().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageSensorTypeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SensorTypeTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *sensorTypeController) SaveSensorType(ctx context.Context, in *manage.SaveSensorTypeReq) (out *manage.SaveSensorTypeRes, err error) {
	out = &manage.SaveSensorTypeRes{}
	id, err := sManage.ManageSensorType().Save(ctx, &in.ManageSensorTypeSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *areaController) DeleteSensorType(ctx context.Context, in *manage.DeleteSensorTypeReq) (out *manage.DeleteSensorTypeRes, err error) {
	out = &manage.DeleteSensorTypeRes{}
	err = sManage.ManageSensorType().Delete(ctx, in.Ids)
	return
}

func (c *sensorTypeController) TreeSensorType(ctx context.Context, in *manage.TreeSensorTypeReq) (out *manage.TreeSensorTypeRes, err error) {
	out = &manage.TreeSensorTypeRes{}
	out.Data, err = sManage.ManageSensorType().Tree(ctx, &in.PageListReq, &in.ManageSensorTypeSearch)
	return
}
