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
func (s *sEvent) CheckEvent(ctx context.Context, sensorId int64, value gateway.Value) (id int64, err error) {
	thresholds, err := manage.ManageThresholdCache().Get(ctx, sensorId)
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
			s.floatInsertOrUpdateEvent(ctx, v.AlarmLabelId, sensorId, eV, value.CreateTime.UnixNano())
		} else {
			// 是否需要解除
			s.floatLiftAlarm(ctx, v.AlarmLabelId, sensorId, eV)
		}
	}

	return
}

// 判断浮点是是否插入或更新报警事件
func (s *sEvent) floatInsertOrUpdateEvent(ctx context.Context, alarmLabelId int64, sensorId int64, value float64, timeSamp int64) (err error) {

	var labelInfo *res.AlarmLabelInfo
	labelInfo, err = manage.ManageAlarmLabel().Read(ctx, alarmLabelId)
	if err != nil {
		return
	}

	page := model.PageListReq{}
	page.Page = 1
	page.PageSize = 1
	var alarms []*res.AlarmTableRow
	alarms, _, err = manage.ManageAlarm().GetPageListForSearch(ctx, &page, &req.ManageAlarmSearch{
		SensorId: sensorId,
		IsLift:   "0",
		Level:    labelInfo.Level,
	})
	if err != nil {
		return
	}
	// var alarmId int64
	if len(alarms) > 0 {
		// alarmId = alarms[0].Id
	} else {
		_, err = manage.ManageAlarm().Save(ctx, req.ManageAlarmSave{
			IsLift:   false,
			Level:    labelInfo.Level,
			SensorId: sensorId,
			Color:    labelInfo.Color,
			SendTime: timeSamp,
		})
		if err != nil {
			return
		}
	}

	// // 存储事件
	// _, err = s.Save(ctx, &req.ManageEventReq{
	// 	SensorId:    sensorId,
	// 	Value:       value,
	// 	Level:       labelInfo.Level,
	// 	Color:       labelInfo.Color,
	// 	Description: "",
	// 	AlarmId:     alarmId,
	// })
	return
}

// 判断浮点是是否需要解除报警
func (s *sEvent) floatLiftAlarm(ctx context.Context, alarmLabelId int64, sensorId int64, value float64) (err error) {

	var labelInfo *res.AlarmLabelInfo
	labelInfo, err = manage.ManageAlarmLabel().Read(ctx, alarmLabelId)
	if err != nil {
		return
	}

	page := model.PageListReq{}
	page.Page = 1
	page.PageSize = 1
	var alarms []*res.AlarmTableRow
	alarms, _, err = manage.ManageAlarm().GetPageListForSearch(ctx, &page, &req.ManageAlarmSearch{
		SensorId: sensorId,
		IsLift:   "0",
		Level:    labelInfo.Level,
	})
	if err != nil {
		return
	}
	var alarmId int64
	if len(alarms) > 0 {
		alarmId = alarms[0].Id
	}
	if alarmId > 0 {
		NewManageAlarm().LiftAlarm(ctx, alarmId)
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
