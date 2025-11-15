package manage

import (
	"context"
	"devinggo/manage/api/manage"
	sManage "devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"
)

var (
	OpcController = opcController{}
)

type opcController struct {
	base.BaseController
}

func (c *opcController) InitOpc(ctx context.Context, in *manage.InitOpcReq) (out *manage.InitOpcRes, err error) {
	out = &manage.InitOpcRes{}
	// Initialize OPC connection and fetch namespace tree
	out.Data, err = sManage.ManageOpc().InitOpc(ctx, in.ServerId, "")
	return
}

func (c *opcController) OpcTree(ctx context.Context, in *manage.OpcTreeReq) (out *manage.OpcTreeRes, err error) {
	out = &manage.OpcTreeRes{}
	// Fetch the OPC device tree based on the provided request
	out.Data, err = sManage.ManageOpc().Tree(ctx, &in.OpcTreeReq)
	if err != nil {
		return nil, err
	}
	return
}

func (c *opcController) OpcNodeIsExit(ctx context.Context, in *manage.OpcNodeIdIsExitReq) (out *manage.OpcNodeIdIsExitRes, err error) {
	out = &manage.OpcNodeIdIsExitRes{}
	out.Data, err = sManage.ManageOpc().OpcNodeIsExit(ctx, &in.OpcReadByServer)
	return
}
