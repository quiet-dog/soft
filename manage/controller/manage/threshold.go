package manage

import (
	"context"
	"devinggo/manage/api/manage"
	sManage "devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"
)

var (
	ThresholdController = thresholdController{}
)

type thresholdController struct {
	base.BaseController
}

func (c *thresholdController) AddThreshold(ctx context.Context, in *manage.SaveThresholdReq) (out *manage.SaveThresholdRes, err error) {
	out = &manage.SaveThresholdRes{}
	err = sManage.ManageThreshold().AddThreshold(ctx, &in.ManageThresholdAddReq)
	return
}

func (c *thresholdController) GetThresholdInfo(ctx context.Context, in *manage.GetThresholdReq) (out *manage.GetThresholdRes, err error) {
	out = &manage.GetThresholdRes{}
	out.Data, err = sManage.ManageThreshold().GetSensorThresholds(ctx, in.Id)
	return
}
