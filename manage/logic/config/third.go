package config

import (
	"context"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
)

type sThird struct {
	base.BaseService
}

func init() {
	manage.RegisterManageThird(NewManageThird())
}

func NewManageThird() *sThird {
	return &sThird{}
}

func (s *sThird) SendSensorDataByWs(ctx context.Context, value gateway.Msg) {
	return
}
