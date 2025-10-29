package config

import (
	"context"
	"database/sql"
	"devinggo/manage/dao"
	"devinggo/manage/global"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDeviceControl struct {
	base.BaseService
}

func init() {
	manage.RegisterManageDeviceControl(NewManageDeviceControl())
}

func NewManageDeviceControl() *sDeviceControl {
	return &sDeviceControl{}
}

func (s *sDeviceControl) Model(ctx context.Context) *gdb.Model {
	return dao.ManageDeviceControl.Ctx(ctx).Hook(hook.Bind()).Handler().Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sDeviceControl) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceControlSearch) (res []*res.DeviceControlTableRow, total int, err error) {
	m := s.handleSersorControlSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sDeviceControl) Save(ctx context.Context, in *req.ManageDeviceControlSave) (id int64, err error) {
	var control *do.ManageDeviceControl
	if err = gconv.Struct(in, &control); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(control).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sDeviceControl) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sDeviceControl) Read(ctx context.Context, sensorId int64) (deviceControlInfo *res.DeviceControlInfo, err error) {
	deviceControlInfo = &res.DeviceControlInfo{}
	err = s.Model(ctx).Where(dao.ManageSensor.Columns().Id, sensorId).Scan(&deviceControlInfo)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sDeviceControl) UpdateInfo(ctx context.Context, in *req.ManageDeviceControlInfo) (out sql.Result, err error) {
	if g.IsEmpty(in.Id) {
		err = myerror.MissingParameter(ctx, "控制id为空")
		return
	}

	var control *do.ManageDeviceControl
	if err = gconv.Struct(in, &control); err != nil {
		return
	}

	out, err = s.Model(ctx).OmitEmptyData().Data(control).Where(dao.ManageArea.Columns().Id, in.Id).Update()
	if utils.IsError(err) {
		return
	}

	return
}

func (s *sDeviceControl) AddControl(ctx context.Context, in *req.ManageAddDeviceControlInfo) (err error) {
	commandIds := []int64{}
	deviceId := 0
	for _, v := range in.Command {
		deviceId = int(v.DeviceId)
		if v.Id != 0 {
			s.UpdateInfo(ctx, v)
			commandIds = append(commandIds, v.Id)
		} else {
			var id int64
			id, err = s.Save(ctx, &req.ManageDeviceControlSave{
				DeviceId: v.DeviceId,
				Extend:   v.Extend,
				Name:     v.Name,
			})
			if err != nil {
				return
			}
			commandIds = append(commandIds, id)
		}
	}

	s.Model(ctx).
		WhereNot(dao.ManageDeviceControl.Columns().Id, commandIds).
		Where(dao.ManageDeviceControl.Columns().DeviceId, deviceId).
		Delete()

	return
}

func (s *sDeviceControl) Control(ctx context.Context, controlId int64) (err error) {
	info, err := s.Read(ctx, controlId)
	if err != nil {
		return
	}

	deviceInfo, err := manage.ManageDevice().Read(ctx, info.DeviceId)
	if err != nil {
		return
	}

	info.Extend.Set("slaveId", deviceInfo.Extend.Get("slaveId").Uint8())
	err = global.DeviceGateway.Control(deviceInfo.ServerId, info.Extend)

	return
}

func (s *sDeviceControl) handleSersorControlSearch(ctx context.Context, in *req.ManageDeviceControlSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if len(in.DeviceIds) > 0 {
		query = query.WhereIn(dao.ManageDeviceControl.Columns().DeviceId, in.DeviceIds)
	}
	return
}
