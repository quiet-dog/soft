package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/hook"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDevice struct {
	base.BaseService
}

func init() {
	manage.RegisterManageDevice(NewManageDevice())
}

func NewManageDevice() *sDevice {
	return &sDevice{}
}

type deviceHook struct {
}

func (h *deviceHook) SelectHook(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
	result, err = in.Next(ctx)
	for _, item := range result {
		if !item["area_id"].IsEmpty() {
			item["area_name"], _ = dao.ManageArea.Ctx(ctx).WherePri(item["area_id"]).Value("name")
		}
	}
	return result, err
}

func (s *sDevice) Model(ctx context.Context) *gdb.Model {
	dHook := &deviceHook{}
	return dao.ManageDevice.Ctx(ctx).Hook(hook.Bind(dHook)).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sDevice) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceSearch) (res []*res.DeviceTableRow, total int, err error) {
	m := s.handleDeviceSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sDevice) Save(ctx context.Context, in *req.ManageDeviceSave) (id int64, err error) {
	var device *do.ManageDevice
	if err = gconv.Struct(in, &device); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(device).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sDevice) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sDevice) Tree(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceSearch) (out []*res.AreaTree, err error) {
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

func (s *sDevice) GetInfoById(ctx context.Context, deviceId int64) (deviceInfo *res.DeviceInfo, err error) {
	deviceInfo = &res.DeviceInfo{}
	err = s.Model(ctx).Where(dao.ManageDevice.Columns().Id, deviceId).Scan(&deviceInfo)
	if utils.IsError(err) {
		return
	}
	// 服务器信息
	dao.ManageServer.Ctx(ctx).WherePri(deviceInfo.ServerId).Scan(&deviceInfo.Server)
	return
}

func (s *sDevice) GetInfoByIds(ctx context.Context, deviceIds []int64) (deviceInfos []*res.DeviceInfo, err error) {
	err = s.Model(ctx).WhereIn(dao.ManageDevice.Columns().Id, deviceIds).Scan(&deviceInfos)
	if utils.IsError(err) {
		return
	}
	for _, deviceInfo := range deviceInfos {
		// 服务器信息
		if deviceInfo.ServerId > 0 {
			dao.ManageServer.Ctx(ctx).WherePri(deviceInfo.ServerId).Scan(&deviceInfo.Server)
		}
	}
	return
}

func (s *sDevice) Read(ctx context.Context, deviceId int64) (deviceInfo *res.DeviceInfo, err error) {
	deviceInfo, err = s.GetInfoById(ctx, deviceId)
	return
}

func (s *sDevice) handleDeviceSearch(ctx context.Context, in *req.ManageDeviceSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if in == nil {
		return query
	}

	if !g.IsEmpty(in.Name) {
		query = query.WhereLike("name", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.InstallationLocation) {
		query = query.WhereLike("installation_location", "%"+in.InstallationLocation+"%")
	}

	if !g.IsEmpty(in.Manufacturer) {
		query = query.WhereLike("manufacturer", "%"+in.Manufacturer+"%")
	}

	if !g.IsEmpty(in.Model) {
		query = query.WhereLike("model", "%"+in.Model+"%")
	}

	if in.ServerId > 0 {
		query = query.Where("server_id", in.ServerId)
	}

	return
}
