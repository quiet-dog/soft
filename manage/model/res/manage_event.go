package res

import (
	"devinggo/manage/model/base"
	"devinggo/manage/model/req"
)

type EventTableRow struct {
	base.BaseTable
	req.ManageEventReq
}
