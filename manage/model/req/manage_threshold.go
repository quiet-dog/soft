package req

type ThresholdRow struct {
	SensorId     int64  `json:"sensorId" description:"传感器Id"`      // 传感器Id
	AlarmLabelId int64  `json:"alarmLabelId" description:"报警标签Id"` // 报警标签id
	Template     string `json:"template" description:"expr模板"`     // expr模板
	Sort         int64  `json:"sort" description:"优先级"`            // 报警优先级
	Color        string `json:"color" description:"颜色"`            // 颜色
	Level        string `json:"level" description:"等级"`            // 等级
}

type ManageThresholdAddReq struct {
	SensorId   int64
	Thresholds []*ThresholdRow
}
