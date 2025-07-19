package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sServer struct {
	base.BaseService
}

func init() {
	manage.RegisterManageServer(NewManageServer())
}

func NewManageServer() *sServer {
	return &sServer{}
}

func (s *sServer) Model(ctx context.Context) *gdb.Model {
	return dao.ManageServer.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sServer) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (res []*res.ServerTableRow, total int, err error) {
	m := s.handleServerSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sServer) Save(ctx context.Context, in *req.ManageServerSave) (id int64, err error) {
	var area *do.ManageServer
	if err = gconv.Struct(in, &area); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(area).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sServer) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sServer) Types(ctx context.Context) (rs []*res.ServerType, err error) {
	t := []string{
		gateway.SERVER_OPC,
		gateway.SERVER_MODBUS_TCP,
		gateway.SERVER_MODBUS_RTU,
		gateway.SERVER_MODBUS_RTU_OVER_TCP,
		gateway.SERVER_MQTT,
	}
	for _, v := range t {
		rs = append(rs, &res.ServerType{
			Label: v,
			Value: v,
		})
	}
	return
}

func (s *sServer) Tree(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (out []*res.AreaTree, err error) {
	out = []*res.AreaTree{}
	items, _, err := s.GetPageListForSearch(ctx, req, in)
	for _, item := range items {
		out = append(out, &res.AreaTree{
			Label:    item.Name,
			Value:    item.Id,
			Children: nil,
			IsLeaf:   true, // Assuming devices are leaf nodes
		})
	}
	return
}

func (s *sServer) Read(ctx context.Context, serverId int64) (serverInfo *res.ServerInfo, err error) {
	serverInfo = &res.ServerInfo{}
	err = s.Model(ctx).Where(dao.ManageServer.Columns().Id, serverId).Scan(&serverInfo)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sServer) handleServerSearch(ctx context.Context, in *req.ManageServerSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		query = query.WhereLike(dao.ManageServer.Table()+".name", "%"+in.Name+"%")
	}
	if !g.IsEmpty(in.Ip) {
		query = query.WhereLike(dao.ManageServer.Table()+".ip", "%"+in.Ip+"%")
	}
	if !g.IsEmpty(in.Type) {
		query = query.Where(dao.ManageServer.Table()+".type", in.Type)
	}

	return
}
