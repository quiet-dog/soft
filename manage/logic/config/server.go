package config

import (
	"context"
	"database/sql"
	"devinggo/manage/dao"
	"devinggo/manage/global"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"os"
	"runtime"
	"strings"
	"time"

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

// type serverHook struct{}

// 创建前
// func (s *serverHook) BeforeInsertHook(ctx context.Context, in *gdb.HookInsertInput) (err error) {
// 	for _, m := range in.Data {
// 		// 是否存在相同的ip和端口
// 		if !g.IsEmpty(m[dao.ManageServer.Columns().Ip]) && !g.IsEmpty(m[dao.ManageServer.Columns().Port]) {
// 			var result gdb.Record
// 			result, err = dao.ManageServer.Ctx(ctx).
// 				Where(dao.ManageServer.Columns().Ip, m[dao.ManageServer.Columns().Ip]).
// 				Where(dao.ManageServer.Columns().Port, m[dao.ManageServer.Columns().Port]).
// 				One()
// 			if err != nil {
// 				return
// 			}
// 			if !result.IsEmpty() {
// 				return errors.New("服务器已经存在")
// 			}
// 		}
// 	}

// 	return
// }

// // 创建后
// func (s *serverHook) AfterInsertHook(ctx context.Context, in *gdb.HookInsertInput, result *sql.Result) (err error) {

// 	return
// }

// // 更新前
// func (s *serverHook) BeforeUpdateHook(ctx context.Context, in *gdb.HookUpdateInput) (err error) {

// 	switch m := in.Data.(type) {
// 	case map[string]interface{}:
// 		// 是否存在相同的ip和端口
// 		if !g.IsEmpty(m["ip"]) && !g.IsEmpty(m["port"]) {
// 			var result gdb.Record
// 			result, err = dao.ManageServer.Ctx(ctx).
// 				WhereNot(dao.ManageServer.Columns().Id, m[dao.ManageServer.Columns().Id]).
// 				Where(dao.ManageServer.Columns().Ip, m[dao.ManageServer.Columns().Ip]).
// 				Where(dao.ManageServer.Columns().Port, m[dao.ManageServer.Columns().Port]).
// 				One()
// 			if err != nil {
// 				return
// 			}
// 			if !result.IsEmpty() {
// 				return errors.New("服务器已经存在")
// 			}
// 		}
// 	}

// 	return
// }

// 查询后
// func (s *serverHook) AfterSelectHook(ctx context.Context, in *gdb.HookSelectInput, result *gdb.Result) (err error) {
// 	for _, item := range *result {
// 		if !item[dao.ManageServer.Columns().Id].IsEmpty() {
// 			item["is_online"] = g.NewVar(global.DeviceGateway.GetOnline(item[dao.ManageServer.Columns().Id].Int64()))
// 		}

// 	}
// 	return
// }

func (s *sServer) Model(ctx context.Context) *gdb.Model {
	// h := &serverHook{}
	return dao.ManageServer.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sServer) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (res []*res.ServerTableRow, total int, err error) {
	m := s.handleServerSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}

	// 获取在线状态
	s.getIsOnline(res...)

	return
}

// 创建服务
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
	if err != nil {
		return
	}

	if in.Type == gateway.SERVER_OPC {
		_, err = manage.ManageOpc().InitOpc(ctx, id)
		if err != nil {
			s.Delete(ctx, []int64{id})
			return
		}
	}

	// 测试是否连通
	_, err = global.DeviceGateway.AddClient(id, gateway.Config{
		Type:    in.Type,
		Port:    in.Port,
		Host:    in.Ip,
		SubTime: time.Duration(in.Interval * int64(time.Second)),
	})

	if err != nil {
		s.Delete(ctx, []int64{id})
	}

	return
}

func (s *sServer) Delete(ctx context.Context, ids []int64) (err error) {
	dao.ManageOpc.Ctx(ctx).Unscoped().Where(dao.ManageOpc.Columns().ServerId, ids).Delete()
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

func (s *sServer) UpdateInfo(ctx context.Context, in *req.ManageServerUpdateInfo) (out sql.Result, err error) {
	if g.IsEmpty(in.Id) {
		err = myerror.MissingParameter(ctx, "服务id为空")
		return
	}

	var server *do.ManageServer
	if err = gconv.Struct(in, &server); err != nil {
		return
	}

	out, err = s.Model(ctx).OmitEmptyData().Data(server).Where(dao.ManageServer.Columns().Id, in.Id).Update()
	if utils.IsError(err) {
		return
	}

	global.DeviceGateway.DeleteClient(in.Id)
	global.DeviceGateway.AddClient(in.Id, gateway.Config{
		Type:    in.Type,
		Host:    in.Ip,
		Port:    in.Port,
		SubTime: time.Duration(in.Interval * int64(time.Second)),
	})

	return
}

// 获取串口地址
func (s *sServer) GetSerialPort(ctx context.Context) (ports []string, err error) {
	// 如果是linux
	if runtime.GOOS == "linux" {
		// 获取485的串口设备
		files, err := os.ReadDir("/dev")
		if err != nil {
			return ports, err
		}
		for _, file := range files {
			name := file.Name()
			if strings.HasPrefix(name, "ttyUSB") || strings.HasPrefix(name, "ttyS") || strings.HasPrefix(name, "ttyAMA") {
				ports = append(ports, "/dev/"+file.Name())
			}

		}
	}

	if runtime.GOOS == "windows" {
		// 获取com口

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

func (s *sServer) getIsOnline(servers ...*res.ServerTableRow) {
	for _, v := range servers {
		v.IsOnline = global.DeviceGateway.GetOnline(v.Id)
	}
}
