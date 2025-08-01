// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ManageThreshold is the golang structure for table manage_threshold.
type ManageThreshold struct {
	SensorId     int64  `json:"sensorId"     orm:"sensor_id"      description:"传感器id"`   // 传感器id
	AlarmLabelId int64  `json:"alarmLabelId" orm:"alarm_label_id" description:"报警标签id"`  // 报警标签id
	Template     string `json:"template"     orm:"template"       description:"expr表达式"` // expr表达式
	Sort         int    `json:"sort"         orm:"sort"           description:"优先级"`     // 优先级
}
