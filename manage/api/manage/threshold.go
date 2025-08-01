package manage

import (
	"devinggo/manage/model/base"
	"devinggo/manage/model/req"
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SaveThresholdReq struct {
	g.Meta `path:"/threshold/save" tags:"Threshold" method:"post" summary:"保存传感器阈值" x-permission:"manage:threshold:save"`
	model.AuthorHeader
	req.ManageThresholdAddReq
}

type SaveThresholdRes struct {
	g.Meta `mime:"application/json" description:"保存设备传感器阈值结果"`
	Id     int64 `json:"id" description:"设备服务器ID"` // 设备服务器ID
}

type GetThresholdReq struct {
	g.Meta `path:"/threshold/info" tags:"Threshold" method:"post" summary:"获取传感器阈值信息" x-permission:"manage:threshold:info"`
	model.AuthorHeader
	base.BaseId
}

type GetThresholdRes struct {
	g.Meta `mime:"application/json" description:"获取设备传感器阈值结果"` // 获取设备传感器阈值结果
	Data   []*req.ThresholdRow                                 `json:"data" description:"传感器阈值结果"` // 获取设备
}
