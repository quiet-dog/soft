package res

import (
	"devinggo/manage/model/base"
	"devinggo/manage/model/req"
)

type AlarmTableRow struct {
	req.ManageAlarmSave
	base.BaseTable
}
type AlarmInfo struct {
	AlarmTableRow
}
