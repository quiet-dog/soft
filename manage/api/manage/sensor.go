package manage

import (
	"devinggo/manage/model/common"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"

	"github.com/gogf/gf/v2/frame/g"
)

type SensorReq struct {
	g.Meta `path:"/sensor" tags:"Sensor" method:"get" summary:"获取传感器列表"`
}

type SensorRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SensorTableRow `json:"items"  dc:"sensor list" `
}

type IndexSensorReq struct {
	g.Meta `path:"/sensor/index" tags:"Sensor" method:"get" summary:"获取传感器列表" x-permission:"manage:sensor:index"`
	model.AuthorHeader
	model.PageListReq
	req.ManageSensorSearch
}

type IndexSensorRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SensorTableRow `json:"items"  dc:"sensor list" `
}

type SaveSensorReq struct {
	g.Meta `path:"/sensor/save" tags:"Sensor" method:"post" summary:"保存传感器" x-permission:"manage:sensor:save"`
	model.AuthorHeader
	req.ManageSensorSave
}

type SaveSensorRes struct {
	g.Meta `mime:"application/json" description:"保存传感器服务器结果"`
	Id     int64 `json:"id" description:"传感器服务器ID"` // 传感器服务器ID
}

type DeleteSensorReq struct {
	g.Meta `path:"/sensor/delete" tags:"Sensor" method:"delete" summary:"删除传感器" x-permission:"manage:sensor:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" description:"传感器ID列表" v:"required#传感器ID列表不能为空"`
}

type DeleteSensorRes struct {
	g.Meta `mime:"application/json"`
}

type TreeSensorReq struct {
	g.Meta `path:"/sensor/tree" tags:"Sensor" method:"get" summary:"获取传感器树"`
	model.AuthorHeader
	req.ManageSensorSearch
	model.PageListReq
}

type TreeSensorRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.AreaTree `json:"data"  dc:"sensor tree list" `
}

type SensorReadDataReq struct {
	g.Meta `path:"/sensor/readData" tags:"Sensor" method:"post" summary:"读取传感器数据" x-permission:"manage:sensor:readData"`
	model.AuthorHeader
	req.ManageSensorReadData
}

type SensorReadDataRes struct {
	g.Meta `mime:"application/json"`
	Data   *common.TemplateEnv `json:"data"  dc:"sensor tree list" `
}

type SensorTranslateReq struct {
	g.Meta `path:"/sensor/translate" tags:"Sensor" method:"post" summary:"转换传感器数据" x-permission:"manage:sensor:translate"`
	model.AuthorHeader
	req.ManageSensorTranslate
}

type SensorTranslateRes struct {
	g.Meta `mime:"application/json"`
	Data   any `json:"data"  dc:"sensor tree list" `
}

type SensorDataListReq struct {
	g.Meta `path:"/sensor/data" tags:"Sensor" method:"get" summary:"转换传感器数据" x-permission:"manage:sensor:data"`
	model.AuthorHeader
	req.ManageInfluxdbSearch
	model.PageListReq
}

type SensorDataListRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.SensorDataList `json:"data"  dc:"sensor data list" `
}

type ReadSensorReq struct {
	g.Meta `path:"/sensor/read/{Id}" method:"get" tags:"传感器" summary:"获取传感器." x-permission:"system:sensor:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"服务 id" v:"required|min:1#设备Id不能为空"`
}

type ReadSensorRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.SensorInfo `json:"data" dc:"服务信息"`
}

type ReadEchartSensorReq struct {
	g.Meta `path:"/sensor/readEchart" method:"post" tags:"传感器" summary:"获取传感器." x-permission:"system:readEchart:read"`
	model.AuthorHeader
	model.PageListReq
	req.ManageInfluxdbOneSensorSearch
}

type ReadEchartSensorRes struct {
	g.Meta `mime:"application/json"`
	Data   *res.SensorEchart `json:"data" dc:"服务信息"`
}
