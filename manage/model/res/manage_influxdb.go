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

type SensorEchart struct {
	CSeriesData    []interface{} `json:"cSeiresData" description:"真实数据"`       // 真实数据
	ESeriesData    []interface{} `json:"eSeiresData" description:"转换后数据"`      // 转换后数据
	XData          []interface{} `json:"xData" description:"x轴数据"`             // x轴数据
	SensorId       int64         `json:"sensorId" description:"传感器ID"`         // 传感器ID
	SensorName     string        `json:"sensorName" description:"传感器名称"`       // 传感器名称
	SensorTypeName string        `json:"sensorTypeName" description:"传感器类型名称"` // 传感器类型名称
	SensorUnit     string        `json:"sensorUnit" description:"传感器单位"`       // 传感器单位
	SensorTypeId   int64         `json:"sensorTypeId" description:"传感器类型ID"`   // 传感器类型ID
	DeviceId       int64         `json:"deviceId" description:"设备ID"`
}
