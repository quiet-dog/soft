package req

type ManageInfluxdbSearch struct {
	DeviceId  int64   `json:"deviceId" description:"设备ID"`   // InfluxDB服务器ID
	SensorIds []int64 `json:"sensorIds" description:"传感器ID"` // 传感器ID
	SensorId  int64   `json:"sensorId" description:"传感器ID"`  // 传感器ID
	BeginTime int64   `json:"beginTime" description:"开始时间"`  // 开始时间
	EndTime   int64   `json:"endTime" description:"结束时间"`    // 结束时间
	Precision int64   `json:"precision" description:"精度"`    // 精度
}

type ManageInfluxdbOneSensorSearch struct {
	// DeviceId  int64 `json:"deviceId" description:"设备ID"`  // InfluxDB服务器ID
	SensorId  int64 `json:"sensorId" description:"传感器ID"` // 传感器ID
	BeginTime int64 `json:"beginTime" description:"开始时间"` // 开始时间
	EndTime   int64 `json:"endTime" description:"结束时间"`   // 结束时间
	Precision int64 `json:"precision" description:"精度"`   // 精度
}
