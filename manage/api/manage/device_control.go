package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceControlReq struct {
	g.Meta `path:"/deviceControl" tags:"DeviceControl" method:"get" summary:"获取设备控制列表"`
}

type DeviceControlRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.DeviceControlTableRow `json:"items"  dc:"deviceControl list" `
}

type IndexDeviceControlReq struct {
	g.Meta `path:"/deviceControl/index" tags:"DeviceControl" method:"get" summary:"获取设备控制列表" x-permission:"manage:deviceControl:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageDeviceControlSearch
}

type IndexDeviceControlRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.DeviceControlTableRow `json:"items"  dc:"deviceControl list" `
}

type SaveDeviceControlReq struct {
	g.Meta `path:"/deviceControl/save" tags:"DeviceControl" method:"post" summary:"保存设备控制" x-permission:"manage:deviceControl:save"`
	model.AuthorHeader
	req.ManageDeviceControlSave
}

type SaveDeviceControlRes struct {
	g.Meta `mime:"application/json" description:"保存设备控制服务器结果"`
	Id     int64 `json:"id" description:"设备控制服务器ID"` // 设备控制服务器ID
}

type DeleteDeviceControlReq struct {
	g.Meta `path:"/deviceControl/delete" tags:"DeviceControl" method:"delete" summary:"删除设备控制" x-permission:"manage:deviceControl:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"设备控制ID列表" v:"required#设备控制ID列表不能为空"`
}

type DeleteDeviceControlRes struct {
	g.Meta `mime:"application/json"`
}

type SaveAddDeviceControlReq struct {
	g.Meta `path:"/deviceControl/addControl" tags:"DeviceControl" method:"post" summary:"保存设备控制" x-permission:"manage:deviceControl:addControl"`
	model.AuthorHeader
	req.ManageAddDeviceControlInfo
}

type SaveAddDeviceControlRes struct {
	g.Meta `mime:"application/json" description:"保存设备控制服务器结果"`
	Id     int64 `json:"id" description:"设备控制服务器ID"` // 设备控制服务器ID
}

type DeviceControlSendReq struct {
	g.Meta `path:"/deviceControl/control" tags:"DeviceControl" method:"get" summary:"保存设备控制" x-permission:"manage:deviceControl:control"`
	model.AuthorHeader
	Id int64 `json:"id" description:"设备控制ID" v:"required#设备控制ID不能为空"`
}

type DeviceControlSendRes struct {
	g.Meta `mime:"application/json" description:"保存设备控制服务器结果"`
}
