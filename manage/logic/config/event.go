package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/common"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
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
func (s *sEvent) CheckEvent(ctx context.Context, sensorId int64, value common.TemplateEnv) (id int64, err error) {
	thresholds, err := manage.ManageThresholdCache().Get(ctx, sensorId)
	if err != nil {
		fmt.Println("=============查询阈值失败===========")
		return
	}

	var template string
	template, err = manage.ManageSensorTemplateCache().Get(ctx, sensorId)
	if err != nil {
		fmt.Println("=============查询模板失败===========")
		return
	}

	eV := value.Value.ToValueExprFloat64(template)
	fmt.Println("=============数值转换成功===========", len(thresholds))
	// 是否有阈值设置
	for _, v := range thresholds {
		aAlarmTemplate := common.AlarmTemplate{
			Template: v.Template,
		}

		// 判断是否报警
		if aAlarmTemplate.IsAlarm(eV) {
			var labelInfo *res.AlarmLabelInfo
			labelInfo, err = manage.ManageAlarmLabel().Read(ctx, v.AlarmLabelId)
			if err != nil {
				return
			}

			page := model.PageListReq{}
			page.Page = 1
			page.PageSize = 1
			var alarms []*res.AlarmTableRow
			alarms, _, err = manage.ManageAlarm().GetPageListForSearch(ctx, &page, &req.ManageAlarmSearch{
				SensorId: sensorId,
				IsLift:   false,
				Level:    labelInfo.Level,
			})
			if err != nil {
				return
			}

			var alarmId int64
			if len(alarms) > 0 {
				alarmId = alarms[0].Id
			} else {
				alarmId, err = manage.ManageAlarm().Save(ctx, req.ManageAlarmSave{
					IsLift:   false,
					Level:    labelInfo.Level,
					SensorId: sensorId,
					Color:    labelInfo.Color,
				})
				if err != nil {
					return
				}
			}

			_, err = s.Save(ctx, &req.ManageEventReq{
				SensorId:    sensorId,
				Value:       eV,
				Level:       labelInfo.Level,
				Color:       labelInfo.Color,
				Desctiption: "",
				AlarmId:     alarmId,
			})
			break
		}

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
