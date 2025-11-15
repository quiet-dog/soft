package manage

import (
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type InitOpcReq struct {
	g.Meta   `path:"/opc/init" tags:"Device" method:"get" summary:""初始化OPC设备树" x-permission:"manage:opc:init"`
	ServerId int64 `json:"serverId" description:"OPC服务器ID" v:"required#OPC服务器ID不能为空"`
}

type InitOpcRes struct {
	g.Meta `mime:"application/json" description:"初始化OPC设备树结果"`
	Data   *gateway.NodeDef `json:"data" description:"OPC设备树"`
}

type OpcTreeReq struct {
	g.Meta `path:"/opc/tree" tags:"Device" method:"get" summary:"获取OPC设备树" x-permission:"manage:opc:tree"`
	model.AuthorHeader
	req.OpcTreeReq
}

type OpcTreeRes struct {
	g.Meta `mime:"application/json" description:"OPC设备树结果"`
	Data   []*res.OpcTree `json:"data" description:"OPC设备树"`
}
