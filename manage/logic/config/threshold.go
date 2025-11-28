package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"slices"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sThreshold struct {
	base.BaseService
}

func init() {
	manage.RegisterManageThreshold(NewManageThreshold())
}

func NewManageThreshold() *sThreshold {
	return &sThreshold{}
}

func (s *sThreshold) Model(ctx context.Context) *gdb.Model {
	return dao.ManageThreshold.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sThreshold) Save(ctx context.Context, in *req.ThresholdRow) (id int64, err error) {
	var threshold *do.ManageThreshold
	if err = gconv.Struct(in, &threshold); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(threshold).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()

	return
}

func (s *sThreshold) Delete(ctx context.Context, in *req.ThresholdRow) (err error) {
	_, err = s.Model(ctx).
		Where(dao.ManageThreshold.Columns().SensorId, in.SensorId).
		Where(dao.ManageThreshold.Columns().AlarmLabelId, in.AlarmLabelId).
		Where(dao.ManageThreshold.Columns().Template, in.Template).
		Delete()
	return
}

func (s *sThreshold) CleanDiff(ctx context.Context, in *req.ThresholdRow) (err error) {
	_, err = s.Model(ctx).
		Where(dao.ManageThreshold.Columns().SensorId, in.SensorId).
		Where(dao.ManageThreshold.Columns().AlarmLabelId, in.AlarmLabelId).
		Where(dao.ManageThreshold.Columns().Template, in.Template).
		Delete()
	return
}

// 添加阈值
func (s *sThreshold) AddThreshold(ctx context.Context, in *req.ManageThresholdAddReq) (err error) {
	sensorInfo, err := manage.ManageSensor().Read(ctx, in.SensorId)
	if err != nil {
		return
	}

	// 查询阈值
	var thresholds []*req.ThresholdRow
	if err = s.Model(ctx).
		Where(dao.ManageThreshold.Columns().SensorId, sensorInfo.Id).
		Scan(&thresholds); err != nil {
		return
	}

	diffThresholds := []*req.ThresholdRow{}
	for _, v := range thresholds {
		if !slices.ContainsFunc(in.Thresholds, func(value *req.ThresholdRow) bool {
			return value.Template == v.Template && value.AlarmLabelId == v.AlarmLabelId
		}) {
			diffThresholds = append(diffThresholds, v)
		}
	}

	// 清除不要的
	for _, v := range diffThresholds {
		if err = s.CleanDiff(ctx, v); err != nil {
			return
		}
	}

	// 清除需要到，不进行特殊处理
	for _, v := range in.Thresholds {
		if err = s.Delete(ctx, v); err != nil {
			return
		}
	}

	// 按数组顺序排序保存
	for i, v := range in.Thresholds {
		i++
		_, err = s.Save(ctx, &req.ThresholdRow{
			SensorId:     in.SensorId,
			AlarmLabelId: v.AlarmLabelId,
			Sort:         int64(i),
			Template:     v.Template,
		})
		if err != nil {
			return
		}
	}
	manage.ManageThresholdCache().Store(ctx, in.SensorId)
	return
}

// 获取阈值
func (s *sThreshold) GetSensorThresholds(ctx context.Context, sensorId int64) (out []*req.ThresholdRow, err error) {
	err = s.Model(ctx).
		Where(dao.ManageThreshold.Columns().SensorId, sensorId).Scan(&out)

	// 获取对应的阈值标签信息
	for _, v := range out {
		var labelInfo *res.AlarmLabelInfo
		labelInfo, err = manage.ManageAlarmLabel().Read(ctx, v.AlarmLabelId)
		if err != nil {
			return
		}
		v.Color = labelInfo.Color
		v.Level = labelInfo.Level
	}
	return
}
