package manage

import (
	"context"
	"devinggo/manage/api/manage"
	"devinggo/manage/model/res"
	sManage "devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	ServerController = serverController{}
)

type serverController struct {
	base.BaseController
}

func (c *serverController) IndexServer(ctx context.Context, in *manage.IndexServerReq) (out *manage.IndexServerRes, err error) {
	out = &manage.IndexServerRes{}
	items, totalCount, err := sManage.ManageServer().GetPageListForSearch(ctx, &in.PageListReq, &in.ManageServerSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.ServerTableRow, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *serverController) SaveServer(ctx context.Context, in *manage.SaveServerReq) (out *manage.SaveServerRes, err error) {
	out = &manage.SaveServerRes{}
	id, err := sManage.ManageServer().Save(ctx, &in.ManageServerSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *areaController) DeleteServer(ctx context.Context, in *manage.DeleteServerReq) (out *manage.DeleteServerRes, err error) {
	out = &manage.DeleteServerRes{}
	err = sManage.ManageServer().Delete(ctx, in.Ids)
	return
}

func (c *serverController) ServerTypes(ctx context.Context, in *manage.ServerTypesReq) (out *manage.ServerTypesRes, err error) {
	out = &manage.ServerTypesRes{}
	out.Data, err = sManage.ManageServer().Types(ctx)
	return
}

func (c *serverController) TreeServer(ctx context.Context, in *manage.TreeServerReq) (out *manage.TreeServerRes, err error) {
	out = &manage.TreeServerRes{}
	out.Data, err = sManage.ManageServer().Tree(ctx, &in.PageListReq, &in.ManageServerSearch)
	return
}

func (c *serverController) Read(ctx context.Context, in *manage.ReadServerReq) (out *manage.ReadServerRes, err error) {
	out = &manage.ReadServerRes{}
	item, err := sManage.ManageServer().Read(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = item
	return
}

func (c *serverController) UpdateInfo(ctx context.Context, in *manage.UpdateServerReq) (out *manage.UpdateServerRes, err error) {
	out = &manage.UpdateServerRes{}
	_, err = sManage.ManageServer().UpdateInfo(ctx, &in.ManageServerUpdateInfo)
	return
}
