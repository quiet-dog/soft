package config

import (
	"context"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/cache"
	"fmt"
	"slices"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
)

const sensorDataGroup = "sensorData"
const deviceDataGroup = "deviceData"
const duration = 10 * time.Second

var exSensorDataCache = int64(duration.Seconds()) // time.Duration -> 秒 -> int64

type sSensorDataCache struct {
	base.BaseService
}

func init() {
	manage.RegisterManageSensorDataCache(NewManageSensorDataCache())
}

func NewManageSensorDataCache() *sSensorDataCache {
	return &sSensorDataCache{}
}

func (s *sSensorDataCache) Model(ctx context.Context) *gredis.Redis {
	// redispubsub.New()

	return cache.GetRedisClient()
}

func (s *sSensorDataCache) Store(ctx context.Context, key int64, value gateway.Value) (v *gvar.Var, err error) {
	// 打印value.value的类型
	v, err = s.Model(ctx).Set(ctx, fmt.Sprintf("%s-%d", sensorDataGroup, key), value, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &exSensorDataCache,
		},
	})
	if err != nil {
		fmt.Println("=============存储传感器数据失败===========", err)
	}

	s.StoreDevice(ctx, key)
	return
}

func (s *sSensorDataCache) Get(ctx context.Context, key int64) (t gateway.Value, err error) {
	v, err := s.Model(ctx).Get(ctx, fmt.Sprintf("%s-%d", sensorDataGroup, key))
	if err != nil {
		return
	}

	if v.IsNil() || v.IsEmpty() {
		return t, fmt.Errorf("为空")
	}
	err = v.Scan(&t)
	return
}

func (s *sSensorDataCache) Delete(ctx context.Context, key int64) (n int64, err error) {
	n, err = s.Model(ctx).Del(ctx, fmt.Sprintf("%s-%d", sensorDataGroup, key))
	if err != nil {
		return
	}

	s.DeleteDevice(ctx, key)
	return
}

func (s *sSensorDataCache) StoreDevice(ctx context.Context, sensorId int64) (v *gvar.Var, err error) {
	sensorInfo, err := manage.ManageSensor().Read(ctx, sensorId)
	if err != nil {
		return
	}

	key := fmt.Sprintf("%s-%d", deviceDataGroup, sensorInfo.DeviceId)
	v, err = s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}

	data := []int64{}
	if v.IsEmpty() || v.IsNil() {
		data = append(data, sensorId)
	} else {
		if err = v.Scan(&data); err != nil {
			return
		}

		if !slices.Contains(data, sensorId) {
			data = append(data, sensorId)
		}
	}

	// ex := int64(duration.Seconds()) // time.Duration -> 秒 -> int64
	v, err = s.Model(ctx).Set(ctx, key, data, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &exSensorDataCache,
		},
	})

	return
}

func (s *sSensorDataCache) GetDevice(ctx context.Context, deviceId int64) (data []int64, err error) {
	key := fmt.Sprintf("%s-%d", deviceDataGroup, deviceId)
	v, err := s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}

	if !v.IsEmpty() && !v.IsNil() {
		if err = v.Scan(&data); err != nil {
			return
		}
	}
	return data, fmt.Errorf("缓存传感器为空")
}

func (s *sSensorDataCache) DeleteDevice(ctx context.Context, sensorId int64) (v *gvar.Var, err error) {
	sensorInfo, err := manage.ManageSensor().Read(ctx, sensorId)
	if err != nil {
		return
	}

	key := fmt.Sprintf("%s-%d", deviceDataGroup, sensorInfo.DeviceId)
	v, err = s.Model(ctx).Get(ctx, key)
	if err != nil {
		return
	}

	data := []int64{}
	if v.IsEmpty() || v.IsNil() {
		return
	} else {
		if err = v.Scan(&data); err != nil {
			return
		}
		if slices.Contains(data, sensorId) {
			data = slices.DeleteFunc(data, func(v int64) bool {
				return v == sensorId
			})
		}
	}

	v, err = s.Model(ctx).Set(ctx, key, data)
	return
}
