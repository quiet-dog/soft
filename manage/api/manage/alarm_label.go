package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type AlarmLabelReq struct {
	g.Meta `path:"/alarmLabel" tags:"AlarmLabel" method:"get" summary:"获取设备报警标签列表"`
}

type AlarmLabelRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AlarmLabelTableRow `json:"items"  dc:"alarmLabel list" `
}

type IndexAlarmLabelReq struct {
	g.Meta `path:"/alarmLabel/index" tags:"AlarmLabel" method:"get" summary:"获取设备报警标签列表" x-permission:"manage:alarmLabel:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageAlarmLabelSearch
}

type IndexAlarmLabelRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AlarmLabelTableRow `json:"items"  dc:"alarmLabel list" `
}

type SaveAlarmLabelReq struct {
	g.Meta `path:"/alarmLabel/save" tags:"AlarmLabel" method:"post" summary:"保存报警标签" x-permission:"manage:alarmLabel:save"`
	model.AuthorHeader
	req.ManageAlarmLabelSave
}

type SaveAlarmLabelRes struct {
	g.Meta `mime:"application/json" description:"保存报警标签结果"`
	Id     int64 `json:"id" description:"报警标签ID"` // 报警标签
}

type DeleteAlarmLableReq struct {
	g.Meta `path:"/alarmLabel/delete" tags:"AlarmLabel" method:"delete" summary:"删除报警标签" x-permission:"manage:alarmLabel:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"报警标签ID列表" v:"required#报警标签ID列表不能为空"`
}

type DeleteAlarmLabelRes struct {
	g.Meta `mime:"application/json"`
}
