package res

import (
	"devinggo/manage/model/base"
	"devinggo/manage/model/req"
)

type AlarmTableRow struct {
	req.ManageAlarmSave
	base.BaseTable
	EndTime  int64 `json:"endTime" description:"解除报警时间"` // 解除报警时间
	SendTime int64 `json:"sendTime" description:"发送时间"`  // 发送时间
}
type AlarmInfo struct {
	AlarmTableRow
}
