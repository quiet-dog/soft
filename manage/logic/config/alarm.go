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
	websocket2 "devinggo/modules/system/pkg/websocket"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAlarm struct {
	base.BaseController
}

func init() {
	manage.RegisterManageAlarm(NewManageAlarm())
}

func NewManageAlarm() *sAlarm {
	return &sAlarm{}
}

func (s *sAlarm) Model(ctx context.Context) *gdb.Model {
	return dao.ManageAlarm.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sAlarm) Save(ctx context.Context, in req.ManageAlarmSave) (id int64, err error) {
	var alarm *do.ManageAlarm
	if err = gconv.Struct(in, &alarm); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(alarm).Insert()
	if utils.IsError(err) {
		return 0, err
	}

	id, err = rs.LastInsertId()
	s.sendMsg(ctx, id)
	return
}

func (s *sAlarm) Read(ctx context.Context, alarmId int64) (alarmInfo *res.AlarmInfo, err error) {
	alarmInfo = &res.AlarmInfo{}
	err = s.Model(ctx).Where(dao.ManageAlarm.Columns().Id, alarmId).Scan(&alarmInfo)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sAlarm) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAlarmSearch) (res []*res.AlarmTableRow, total int, err error) {
	m := s.handleAlarmSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sAlarm) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sAlarm) sendMsg(ctx context.Context, alarmId int64) {
	alarmInfo, err := s.Read(ctx, alarmId)
	if err != nil {
		return
	}
	toId := "1"
	clientIdWResponse := &websocket2.ClientIdWResponse{
		ID: toId,
		WResponse: &websocket2.WResponse{
			BindEvent: "alarm",
			Event:     websocket2.IdMessage,
			Data:      alarmInfo,
			Code:      200,
			RequestId: "0",
		},
	}
	websocket2.PublishIdMessage(ctx, toId, clientIdWResponse)
}

// 报警解除
func (s *sAlarm) LiftAlarm(ctx context.Context, alarmId int64) (err error) {
	_, err = s.Model(ctx).Where(dao.ManageAlarm.Columns().Id, alarmId).Data(g.Map{
		dao.ManageAlarm.Columns().IsLift:  true,
		dao.ManageAlarm.Columns().EndTime: time.Now().UnixMilli(),
	}).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sAlarm) handleAlarmSearch(ctx context.Context, in *req.ManageAlarmSearch) *gdb.Model {
	m := s.Model(ctx)

	if in.SensorId > 0 {
		m = m.Where(dao.ManageAlarm.Columns().SensorId, in.SensorId)
	}

	if in.IsLift != "" {
		if in.IsLift == "1" {
			m = m.Where(dao.ManageAlarm.Columns().IsLift, 1)
		} else {
			m = m.Where(dao.ManageAlarm.Columns().IsLift, 0)
		}
	}

	if in.Level != "" {
		m = m.Where(dao.ManageAlarm.Columns().Level, in.Level)
	}

	return m
}
