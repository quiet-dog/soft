package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAlarmLabel struct {
	base.BaseController
}

func init() {
	manage.RegisterManageAlarmLabel(NewManageAlarmLabel())
}

func NewManageAlarmLabel() *sAlarmLabel {
	return &sAlarmLabel{}
}

func (s *sAlarmLabel) Model(ctx context.Context) *gdb.Model {
	return dao.ManageAlarmLabel.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sAlarmLabel) Save(ctx context.Context, in *req.ManageAlarmLabelSave) (id int64, err error) {

	var alarmLabel *do.ManageAlarmLabel
	if err = gconv.Struct(in, &alarmLabel); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(alarmLabel).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sAlarmLabel) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAlarmLabelSearch) (res []*res.AlarmLabelTableRow, total int, err error) {
	m := s.handleAlarmLabelSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sAlarmLabel) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sAlarmLabel) handleAlarmLabelSearch(ctx context.Context, in *req.ManageAlarmLabelSearch) *gdb.Model {
	m := s.Model(ctx)
	if in.Name != "" {
		m = m.WhereLike("name", "%"+in.Name+"%")
	}

	if in.Level != "" {
		m = m.WhereLike("level", "%"+in.Level+"%")
	}

	if in.Remark != "" {
		m = m.WhereLike("remark", "%"+in.Remark+"%")
	}

	return m
}
