package config

import (
	"context"
	"devinggo/manage/pkg/expr_template"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/cache"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
)

const sensorTemplateCacheKey = "sensor-template"
const sensorTemplateCacheDuration = 10 * time.Second

var exSensorTemplate = int64(sensorTemplateCacheDuration.Seconds()) // time.Duration -> 秒 -> int6

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

func (s *sSensorTemplateCache) Get(ctx context.Context, sensorId int64) (template expr_template.ExprTemplate, err error) {
	key := fmt.Sprintf("%s-%d", sensorTemplateCacheKey, sensorId)

	v, err := s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}

	if v.IsEmpty() || v.IsNil() {
		return
	}

	template = expr_template.ExprTemplate(v.String())
	return
}

func (s *sSensorTemplateCache) Store(ctx context.Context, sensorId int64) (template expr_template.ExprTemplate, err error) {
	key := fmt.Sprintf("%s-%d", sensorTemplateCacheKey, sensorId)

	info, err := manage.ManageSensor().Read(ctx, sensorId)
	if err != nil {
		return
	}

	template = expr_template.ExprTemplate(info.Template)

	_, err = s.Model(ctx).Set(ctx, key, string(template), gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &exSensorTemplate,
		},
	})

	return
}
