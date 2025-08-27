package common

import "devinggo/manage/pkg/expr_template"

type SensorToInfluxdb struct {
	SensorId     int64                      `json:"sensorId"`
	DeviceId     int64                      `json:"deviceId"`
	SensorTypeId int64                      `json:"sensorTypeId"`
	Template     expr_template.ExprTemplate `json:"template"`
}
