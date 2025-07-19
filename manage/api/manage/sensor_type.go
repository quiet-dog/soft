package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type SensorTypeReq struct {
	g.Meta `path:"/sensorType" tags:"SensorType" method:"get" summary:"获取传感器类型列表"`
}

type SensorTypeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SensorTypeTableRow `json:"items"  dc:"sensorType list" `
}

type IndexSensorTypeReq struct {
	g.Meta `path:"/sensorType/index" tags:"SensorType" method:"get" summary:"获取传感器类型列表" x-permission:"manage:sensorType:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageSensorTypeSearch
}

type IndexSensorTypeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SensorTypeTableRow `json:"items"  dc:"sensorType list" `
}

type SaveSensorTypeReq struct {
	g.Meta `path:"/sensorType/save" tags:"SensorType" method:"post" summary:"保存传感器类型" x-permission:"manage:sensorType:save"`
	model.AuthorHeader
	req.ManageSensorTypeSave
}

type SaveSensorTypeRes struct {
	g.Meta `mime:"application/json" description:"保存传感器类型结果"`
	Id     int64 `json:"id" description:"传感器类型ID"` // 设备服务器ID
}

type DeleteSensorTypeReq struct {
	g.Meta `path:"/sensorType/delete" tags:"SensorType" method:"delete" summary:"删除设备" x-permission:"manage:sensorType:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"设备ID列表" v:"required#传感器类型ID列表不能为空"`
}

type DeleteSensorTypeRes struct {
	g.Meta `mime:"application/json"`
}

type TreeSensorTypeReq struct {
	g.Meta `path:"/sensorType/tree" tags:"SensorType" method:"get" summary:"获取传感器类型"`
	model.AuthorHeader
	req.ManageSensorTypeSearch
	model.PageListReq
}

type TreeSensorTypeRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"sensorType tree list" `
}
