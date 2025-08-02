package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"

	"github.com/gogf/gf/v2/database/gdb"
)

type sSensorControl struct {
	base.BaseService
}

func init() {
	manage.RegisterManageSensorControl(NewManageSensorControl())
}

func NewManageSensorControl() *sSensorControl {
	return &sSensorControl{}
}

func (s *sSensorControl) Model(ctx context.Context) *gdb.Model {
	return dao.ManageSensorControl.Ctx(ctx).Hook(hook.Bind()).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}
