package config

import (
	"context"
	"devinggo/manage/model/req"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/cache"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
)

type sThresholdCache struct {
	base.BaseService
}

const thresholdCacheKey = "threshold"
const thresholdCacheDuration = 10 * time.Second

func init() {
	manage.RegisterManageThresholdCache(NewManageThresholdCache())
}

func NewManageThresholdCache() *sThresholdCache {
	return &sThresholdCache{}
}

func (s *sThresholdCache) Model(ctx context.Context) *gredis.Redis {
	return cache.GetRedisClient()
}

func (s *sThresholdCache) Get(ctx context.Context, sensorId int64) (thresholds []*req.ThresholdRow, err error) {

	key := fmt.Sprintf("%s-%d", thresholdCacheKey, sensorId)

	v, err := s.Model(ctx).Get(ctx, key)

	if err != nil {
		return
	}

	if v.IsEmpty() || v.IsNil() {
		thresholds, err = s.Store(ctx, sensorId)
		return
	}

	if err = v.Scan(&thresholds); err != nil {
		return
	}

	// 重新设置到期时间
	_, err = s.Model(ctx).Expire(ctx, key, int64(thresholdCacheDuration.Seconds()))
	if err != nil {
		return
	}

	return
}

func (s *sThresholdCache) Store(ctx context.Context, sensorId int64) (out []*req.ThresholdRow, err error) {
	key := fmt.Sprintf("%s-%d", thresholdCacheKey, sensorId)
	ex := int64(thresholdCacheDuration.Seconds()) // time.Duration -> 秒 -> int64

	out, err = manage.ManageThreshold().GetSensorThresholds(ctx, sensorId)
	if err != nil {
		return
	}

	_, err = s.Model(ctx).Set(ctx, key, out, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &ex,
		},
	})

	return
}
