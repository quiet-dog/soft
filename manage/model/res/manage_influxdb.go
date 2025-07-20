package res

type SensorDataList struct {
	SensorId       int64                    `json:"sensorId" description:"传感器ID"`         // 传感器ID
	SensorName     string                   `json:"sensorName" description:"传感器名称"`       // 传感器名称
	SensorTypeName string                   `json:"sensorTypeName" description:"传感器类型名称"` // 传感器类型名称
	SensorUnit     string                   `json:"sensorUnit" description:"传感器单位"`       // 传感器单位
	SensorTypeId   int64                    `json:"sensorTypeId" description:"传感器类型ID"`   // 传感器类型ID
	DeviceId       int64                    `json:"deviceId" description:"设备ID"`
	Total          int64                    `json:"total" description:"总记录数"` // 总记录数
	Rows           []map[string]interface{} `json:"rows" description:"数据行"`   // 数据行
}
