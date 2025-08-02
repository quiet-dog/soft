package req

type ManageAlarmSave struct {
	IsLift   bool   `json:"isLift" description:"是否解除报警"`  // 是否解除报警
	Level    string `json:"level" description:"报警等级"`     // 报警等级
	SensorId int64  `json:"sensorId" description:"传感器Id"` // 传感器id
	Color    string `json:"color" description:"颜色"`       // 颜色
}

type ManageAlarmSearch struct {
	SensorId int64  `json:"sensorId" description:"传感器id"` // 传感器id
	IsLift   bool   `json:"isLift" description:"是否解除报警"`  // 是否解除报警
	Level    string `json:"level" description:"报警等级"`     // 报警等级
}
