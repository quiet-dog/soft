package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/common"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sEvent struct {
	base.BaseService
}

func init() {
	manage.RegisterManageEvent(NewManageEvent())
}

func NewManageEvent() *sEvent {
	return &sEvent{}
}

func (s *sEvent) Model(ctx context.Context) *gdb.Model {
	return dao.ManageEvent.Ctx(ctx).Hook(hook.Bind()).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sEvent) Save(ctx context.Context, in *req.ManageEventReq) (id int64, err error) {
	var event *do.ManageEvent
	if err = gconv.Struct(in, &event); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(event).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()

	return
}

func (s *sEvent) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageEventSearch) (res []*res.EventTableRow, total int, err error) {
	m := s.handleEventSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

// 检查是否报警，创建事件
func (s *sEvent) CheckEvent(ctx context.Context, sensorId int64, value gateway.Value) (id int64, isAlarm bool, err error) {
	thresholds, err := manage.ManageThresholdCache().Get(ctx, sensorId)
	// 暂不处理的阈值设置
	if err != nil {
		fmt.Println("=============查询阈值失败===========")
		return
	}

	template, err := manage.ManageSensorTemplateCache().Get(ctx, sensorId)
	if err != nil {
		fmt.Println("=============查询模板失败===========", err)
		return
	}

	eV, err := template.ToExprValueFloat64(value.Value)
	if err != nil {
		return
	}

	// 是否有阈值设置
	for _, v := range thresholds {
		aAlarmTemplate := common.AlarmTemplate{
			Template: v.Template,
		}

		// 判断是否报警
		if aAlarmTemplate.IsAlarmFloat64(eV) {
			// 是否更新事件还是插入事件
			id, _ = s.floatInsertOrUpdateEvent(ctx, v, sensorId, eV, value.CreateTime.UnixNano())
			return id, true, nil
		} else {
			// 是否需要解除
			id, _ = s.floatLiftAlarm(ctx, v, sensorId, eV)
			return id, false, nil
		}
	}
	return
}

// 判断浮点是是否插入或更新报警事件
func (s *sEvent) floatInsertOrUpdateEvent(ctx context.Context, thresholdInfo *req.ThresholdRow, sensorId int64, value any, timeSamp int64) (alarmId int64, err error) {

	alarmId, err = manage.ManageAlarmSensorCache().Get(ctx, sensorId, thresholdInfo.AlarmLabelId)
	if err != nil {
		return
	}

	if alarmId == 0 {
		alarmId, err = manage.ManageAlarm().Save(ctx, req.ManageAlarmSave{
			IsLift:   false,
			Level:    thresholdInfo.Level,
			SensorId: sensorId,
			Color:    thresholdInfo.Color,
			SendTime: timeSamp,
		})
	}

	return
}

// 判断浮点是是否需要解除报警
func (s *sEvent) floatLiftAlarm(ctx context.Context, thresholdInfo *req.ThresholdRow, sensorId int64, value any) (alarmId int64, err error) {

	alarmId, err = manage.ManageAlarmSensorCache().Get(ctx, sensorId, thresholdInfo.AlarmLabelId)

	if alarmId > 0 {
		NewManageAlarm().LiftAlarm(ctx, alarmId)
		manage.ManageAlarmSensorCache().Delete(ctx, sensorId, thresholdInfo.AlarmLabelId)
	}
	return
}

func (s *sEvent) handleEventSearch(ctx context.Context, in *req.ManageEventSearch) *gdb.Model {
	m := s.Model(ctx)

	if len(in.AlarmIds) > 0 {
		m.WhereIn(dao.ManageEvent.Columns().AlarmId, in.AlarmIds)
	}

	if len(in.Ids) > 0 {
		m.WhereIn(dao.ManageEvent.Columns().Id, in.Ids)
	}

	if len(in.SensorIds) > 0 {
		m.WhereIn(dao.ManageEvent.Columns().SensorId, in.SensorIds)
	}

	return m
}
