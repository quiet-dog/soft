package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceReq struct {
	g.Meta `path:"/device" tags:"Device" method:"get" summary:"获取设备列表"`
}

type DeviceRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.DeviceTableRow `json:"items"  dc:"device list" `
}

type IndexDeviceReq struct {
	g.Meta `path:"/device/index" tags:"Device" method:"get" summary:"获取设备列表" x-permission:"manage:device:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageDeviceSearch
}

type IndexDeviceRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.DeviceTableRow `json:"items"  dc:"device list" `
}

type SaveDeviceReq struct {
	g.Meta `path:"/device/save" tags:"Device" method:"post" summary:"保存设备" x-permission:"manage:device:save"`
	model.AuthorHeader
	req.ManageDeviceSave
}

type SaveDeviceRes struct {
	g.Meta `mime:"application/json" description:"保存设备服务器结果"`
	Id     int64 `json:"id" description:"设备服务器ID"` // 设备服务器ID
}

type DeleteDeviceReq struct {
	g.Meta `path:"/device/delete" tags:"Device" method:"delete" summary:"删除设备" x-permission:"manage:device:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"设备ID列表" v:"required#设备ID列表不能为空"`
}

type DeleteDeviceRes struct {
	g.Meta `mime:"application/json"`
}

type TreeDeviceReq struct {
	g.Meta `path:"/device/tree" tags:"Device" method:"get" summary:"获取设备树"`
	model.AuthorHeader
	req.ManageDeviceSearch
	model.PageListReq
}

type TreeDeviceRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"device tree list" `
}

type ReadDeviceReq struct {
	g.Meta `path:"/device/read/{Id}" method:"get" tags:"设备" summary:"获取设备信息." x-permission:"system:device:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"岗位 id" v:"required|min:1#设备Id不能为空"`
}

type ReadDeviceRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.DeviceInfo `json:"data" dc:"设备信息"`
}
