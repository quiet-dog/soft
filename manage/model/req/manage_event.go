package req

import "devinggo/manage/model/base"

type ManageEventReq struct {
	SensorId    int64   `json:"sensorId" description:"传感器id"`   // 传感器id
	Value       float64 `json:"value" description:"传感器数值"`      // 传感器数值
	Description string  `json:"description" description:"事件描述"` // 事件描述
	Level       string  `json:"level" description:"报警的等级"`      // 报警等级
	Color       string  `json:"color" description:"颜色标签"`       // 颜色标签
	AlarmId     int64   `json:"alarmId" description:"报警Id"`     // 报警id
}

type ManageEventSearch struct {
	base.BaseIds
	AlarmIds  []int64 `json:"alarmIds" description:"报警事件Ids"` // 报警事件Ids
	SensorIds []int64 `json:"sensorIds" description:"传感器Ids"` // 传感器Ids
}
