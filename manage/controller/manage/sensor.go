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
	SensorController = sernsorController{}
)

type sernsorController struct {
	base.BaseController
}

func (c *sernsorController) IndexSensor(ctx context.Context, in *manage.IndexSensorReq) (out *manage.IndexSensorRes, err error) {
	out = &manage.IndexSensorRes{}
	items, totalCount, err := sManage.ManageSensor().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageSensorSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SensorTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *sernsorController) SaveSensor(ctx context.Context, in *manage.SaveSensorReq) (out *manage.SaveSensorRes, err error) {
	out = &manage.SaveSensorRes{}
	id, err := sManage.ManageSensor().Save(ctx, &in.ManageSensorSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *areaController) DeleteSensor(ctx context.Context, in *manage.DeleteSensorReq) (out *manage.DeleteSensorRes, err error) {
	out = &manage.DeleteSensorRes{}
	err = sManage.ManageSensor().Delete(ctx, in.Ids)
	return
}

func (c *sernsorController) TreeSensor(ctx context.Context, in *manage.TreeSensorReq) (out *manage.TreeSensorRes, err error) {
	out = &manage.TreeSensorRes{}
	out.Data, err = sManage.ManageSensor().Tree(ctx, &in.PageListReq, &in.ManageSensorSearch)
	return
}

func (c *sernsorController) ReadData(ctx context.Context, in *manage.SensorReadDataReq) (out *manage.SensorReadDataRes, err error) {
	out = &manage.SensorReadDataRes{}
	out.Data, err = sManage.ManageSensor().ReadData(ctx, &in.ManageSensorReadData)
	return
}

func (c *sernsorController) TranslateData(ctx context.Context, in *manage.SensorTranslateReq) (out *manage.SensorTranslateRes, err error) {
	out = &manage.SensorTranslateRes{}
	out.Data, err = sManage.ManageSensor().TranslateData(ctx, &in.ManageSensorTranslate)
	return
}

func (c *sernsorController) SearchSensorDataList(ctx context.Context, in *manage.SensorDataListReq) (out *manage.SensorDataListRes, err error) {
	out = &manage.SensorDataListRes{}
	out.Data = &res.SensorDataList{}
	out.Data, err = sManage.ManageInfluxdb().SearchSensorDataList(ctx, &in.PageListReq, &in.ManageInfluxdbSearch)
	return
}

func (c *sernsorController) Read(ctx context.Context, in *manage.ReadSensorReq) (out *manage.ReadSensorRes, err error) {
	out = &manage.ReadSensorRes{}
	item, err := sManage.ManageSensor().Read(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = item
	return
}

func (c *sernsorController) ReadEchart(ctx context.Context, in *manage.ReadEchartSensorReq) (out *manage.ReadEchartSensorRes, err error) {
	out = &manage.ReadEchartSensorRes{}
	out.Data, err = sManage.ManageSensor().ReadEchart(ctx, &in.PageListReq, &in.ManageInfluxdbOneSensorSearch)
	return
}

func (c *sernsorController) ReadHistoryData(ctx context.Context, in *manage.ReadHistoryDataSensorReq) (out *manage.ReadHistoryDataSensorRes, err error) {

	out = &manage.ReadHistoryDataSensorRes{}
	out.Data, err = sManage.ManageSensor().ReadHistoryData(ctx, &in.PageListReq, &in.ManageInfluxdbOneSensorSearch)
	return
}
