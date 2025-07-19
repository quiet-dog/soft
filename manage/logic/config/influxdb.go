package config

import (
	"context"
	"devinggo/manage/model/common"
	"devinggo/manage/model/req"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"fmt"

	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
)

type sInfluxdb struct {
	base.BaseService
}

func init() {
	manage.RegisterManageInfluxdb(NewManageInfluxdb())
}

func NewManageInfluxdb() *sInfluxdb {
	return &sInfluxdb{}
}

func (s *sInfluxdb) Model(ctx context.Context) (*influxdb3.Client, error) {
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     "http://localhost:8181",
		Token:    "apiv3_KDfgrll4Hg3VKFpOt5wLOtStjWSmZNeIcW-obG1SYJGc5W2OAZRrH-pXq_5Q-_E7LT0bhKwcMOglg-Ml2J3EJg",
		Database: "DATABASE_NAME",
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *sInfluxdb) Store(ctx context.Context, data common.TemplateEnv, sensorId int64) (err error) {
	c, err := s.Model(context.Background())
	if err != nil {
		return
	}
	defer c.Close()

	influxdbData, err := manage.ManageSensor().ReadInfluxdbFormat(ctx, sensorId)
	if err != nil {
		return
	}

	formatData, err := manage.ManageSensor().TranslateData(ctx, &req.ManageSensorTranslate{
		Env:      data,
		Template: influxdbData.Template,
	})
	if err != nil {
		return
	}

	currend := data.Value
	// line := "1,sensor=2 value=23.5,current=45i"
	line := fmt.Sprintf("%d,sensor=%d, %d_c=%d,%d=%s", influxdbData.DeviceId,
		influxdbData.SensorId,
		influxdbData.SensorTypeId,
		currend,
		influxdbData.SensorId,
		formatData,
	)
	err = c.Write(ctx, []byte(line))
	return
}
