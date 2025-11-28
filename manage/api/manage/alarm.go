package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type AlarmReq struct {
	g.Meta `path:"/alarm" tags:"Alarm" method:"get" summary:"获取设备报警列表"`
}

type AlarmRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AlarmTableRow `json:"items"  dc:"alarm list" `
}

type IndexAlarmReq struct {
	g.Meta `path:"/alarm/index" tags:"Alarm" method:"get" summary:"获取设备报警列表" x-permission:"manage:alarm:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageAlarmSearch
}

type IndexAlarmRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AlarmTableRow `json:"items"  dc:"alarm list" `
}

type SaveAlarmReq struct {
	g.Meta `path:"/alarm/save" tags:"Alarm" method:"post" summary:"保存报警" x-permission:"manage:alarm:save"`
	model.AuthorHeader
	req.ManageAlarmSave
}

type SaveAlarmRes struct {
	g.Meta `mime:"application/json" description:"保存报警结果"`
	Id     int64 `json:"id" description:"报警标签ID"` // 报警
}

type DeleteAlarmReq struct {
	g.Meta `path:"/alarm/delete" tags:"Alarm" method:"delete" summary:"删除报警" x-permission:"manage:alarm:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"报警标签ID列表" v:"required#报警标签ID列表不能为空"`
}

type DeleteAlarmRes struct {
	g.Meta `mime:"application/json"`
}

type LiftAlarmReq struct {
	g.Meta `path:"/alarm/lift" tags:"Alarm" method:"post" summary:"解除报警" x-permission:"manage:alarm:lift"`
	model.AuthorHeader
	Id int64 `json:"id" description:"报警ID" v:"required#报警ID不能为空"`
}

type LiftAlarmRes struct {
	g.Meta `mime:"application/json"`
}

type ReadAlarmReq struct {
	g.Meta `path:"/alarm/read/{Id}" method:"get" tags:"Alarm" summary:"获取报警信息" x-permission:"manage:alarm:read"`
	model.AuthorHeader
	Id int64 `json:"id" description:"报警ID" v:"required#报警ID不能为空"`
}

type ReadAlarmRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.AlarmInfo `json:"data" dc:"报警信息"`
}
