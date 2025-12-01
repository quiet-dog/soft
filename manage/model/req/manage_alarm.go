package req

type ManageAlarmSave struct {
	IsLift      bool   `json:"isLift" description:"是否解除报警"`    // 是否解除报警
	Level       string `json:"level" description:"报警等级"`       // 报警等级
	SensorId    int64  `json:"sensorId" description:"传感器Id"`   // 传感器id
	Color       string `json:"color" description:"颜色"`         // 颜色
	SendTime    int64  `json:"sendTime" description:"发送时间"`    // 发送时间
	ThresholdId int64  `json:"thresholdId" description:"阈值id"` // 阈值id
}

type ManageAlarmSearch struct {
	SensorId int64  `json:"sensorId" description:"传感器id"` // 传感器id
	IsLift   string `json:"isLift" description:"是否解除报警"`  // 是否解除报警
	Level    string `json:"level" description:"报警等级"`     // 报警等级
}
