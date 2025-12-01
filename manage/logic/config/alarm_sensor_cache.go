package config

import (
	"context"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/pkg/cache"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
)

const alarmSensorCacheKey = "alarm-sensor"

type sAlarmSensorCache struct {
	base.BaseController
}

func init() {
	manage.RegisterManageAlarmSensorCache(NewManageAlarmSensorCache())
}

func NewManageAlarmSensorCache() *sAlarmSensorCache {
	return &sAlarmSensorCache{}
}

func (s *sAlarmSensorCache) Model(ctx context.Context) *gredis.Redis {
	return cache.GetRedisClient()
}

func (s *sAlarmSensorCache) Get(ctx context.Context, sensorId int64, thresholdId int64) (alarmId int64, err error) {
	key := fmt.Sprintf("%s-%d-%d", alarmSensorCacheKey, sensorId, thresholdId)
	v, err := s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}
	return v.Int64(), nil
}

func (s *sAlarmSensorCache) Store(ctx context.Context, sensorId int64, thresholdId int64, alarmId int64) (err error) {
	key := fmt.Sprintf("%s-%d-%d", alarmSensorCacheKey, sensorId, thresholdId)
	_, err = s.Model(ctx).Set(ctx, key, alarmId, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			// 一直缓存
			EX: nil,
		},
	})
	return
}

func (s *sAlarmSensorCache) Delete(ctx context.Context, sensorId int64, thresholdId int64) (err error) {
	key := fmt.Sprintf("%s-%d-%d", alarmSensorCacheKey, sensorId, thresholdId)
	_, err = s.Model(ctx).Del(ctx, key)
	return
}
