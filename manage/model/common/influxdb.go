package common

type SensorToInfluxdb struct {
	SensorId     int64  `json:"sensorId"`
	DeviceId     int64  `json:"deviceId"`
	SensorTypeId int64  `json:"sensorTypeId"`
	Template     string `json:"template"`
}
