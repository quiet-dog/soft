package config

import (
	"context"
	"devinggo/manage/model/common"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/cache"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
)

const sensorTemplateCacheKey = "sensor-template"
const sensorTemplateCacheDuration = 10 * time.Second

type sSensorTemplateCache struct {
	base.BaseService
}

func init() {
	manage.RegisterManageSensorTemplateCache(NewManageSensorTemplateCache())
}

func NewManageSensorTemplateCache() *sSensorTemplateCache {
	return &sSensorTemplateCache{}
}

func (s *sSensorTemplateCache) Model(ctx context.Context) *gredis.Redis {
	return cache.GetRedisClient()
}
func (s *sSensorTemplateCache) Get(ctx context.Context, sensorId int64) (template string, err error) {
	key := fmt.Sprintf("%s-%d", sensorTemplateCacheKey, sensorId)

	v, err := s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}

	if v.IsEmpty() || v.IsNil() {
		template, err = s.Store(ctx, sensorId)
		return
	}

	template = v.String()
	aTemplate := common.AlarmTemplate{}
	aTemplate.Template = template
	return
}

func (s *sSensorTemplateCache) Store(ctx context.Context, sensorId int64) (template string, err error) {
	key := fmt.Sprintf("%s-%d", sensorTemplateCacheKey, sensorId)
	ex := int64(sensorTemplateCacheDuration.Seconds()) // time.Duration -> 秒 -> int6

	info, err := manage.ManageSensor().Read(ctx, sensorId)
	if err != nil {
		return
	}

	template = info.Template

	_, err = s.Model(ctx).Set(ctx, key, template, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &ex,
		},
	})

	return
}
