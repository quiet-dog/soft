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
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"fmt"
	"net/url"
	"time"

	"github.com/goburrow/modbus"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
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

// type deviceHook struct {
// }

// func (h *deviceHook) AfterSelectHook(ctx context.Context, in *gdb.HookSelectInput, result *gdb.Result) (err error) {
// 	for _, item := range *result {
// 		if !item["area_id"].IsEmpty() {
// 			item["area_name"], _ = dao.ManageArea.Ctx(ctx).WherePri(item["area_id"].Int64()).Value("name")
// 		}
// 		if !item["server_id"].IsEmpty() {
// 			item["server_name"], _ = dao.ManageServer.Ctx(ctx).WherePri(item["server_id"].Int64()).Value("name")
// 		}
// 		if !item[dao.ManageDevice.Columns().Id].IsEmpty() {
// 			data, _ := manage.ManageSensorDataCache().GetDevice(ctx, item[dao.ManageDevice.Columns().Id].Int64())
// 			if len(data) > 0 {
// 				item["is_online"] = g.NewVar(true)
// 			} else {
// 				item["is_online"] = g.NewVar(false)
// 			}
// 		}
// 	}
// 	return
// }

func (s *sDevice) Model(ctx context.Context) *gdb.Model {
	// dHook := &deviceHook{}
	return dao.ManageDevice.Ctx(ctx).Handler().Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sDevice) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceSearch) (res []*res.DeviceTableRow, total int, err error) {
	m := s.handleDeviceSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

// 获取设备列表,更加详细,带传感器数据和当前数据
func (s *sDevice) GetPageListForSearchHaveSensors(ctx context.Context, r *model.PageListReq, in *req.ManageDeviceSearch) (res []*res.DeviceSensorInfo, total int, err error) {
	m := s.handleDeviceSearch(ctx, in)
	err = orm.GetPageList(m, r).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}

	for _, item := range res {
		item.Sensors, _, err = NewManageSensor().GetPageListForSearch(ctx, &model.PageListReq{}, &req.ManageSensorSearch{
			DeviceIds: []int64{item.Id},
		})
		if utils.IsError(err) {
			return nil, 0, err
		}
		for _, sensor := range item.Sensors {
			val, err := NewManageSensorDataCache().Get(ctx, sensor.Id)
			if err != nil {
				continue
			}
			sensor.Value = val.Value
			sensor.DataTime = val.CreateTime
		}
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

func (s *sDevice) ReadSensorInfo(ctx context.Context, deviceId int64) (info *res.DeviceSensorInfo, err error) {
	deviceInfo, err := s.GetInfoById(ctx, deviceId)
	if err != nil {
		return
	}

	if err = gconv.Struct(deviceInfo, &info); err != nil {
		return
	}

	// 获取对应的传感器
	res, _, err := manage.ManageSensor().GetPageListForSearch(ctx, &model.PageListReq{}, &req.ManageSensorSearch{
		DeviceId: deviceId,
	})
	if err != nil {
		return
	}

	info.Sensors = append(info.Sensors, res...)

	return
}

func (s sDevice) SaveSensorInfo(ctx context.Context, r *req.DeviceSensorInfoSaveReq) (err error) {
	if r == nil {
		return nil
	}

	if len(r.Sensors) == 0 {
		NewManageSensor().Model(ctx).Where("device_id", r.DeviceId).Delete()
	}

	for _, v := range r.Sensors {
		fmt.Println(v.Extend.Get("id"))
		_, err = NewManageSensor().Save(ctx, v)
		if utils.IsError(err) {
			return
		}
	}
	return
}

func (s *sDevice) TestConnect(ctx context.Context, req *req.DeviceTestConnectReq) (err error) {

	serverInfo, err := manage.ManageServer().Read(ctx, req.ServerId)
	if err != nil {
		return
	}
	return s.handleDeviceConnect(serverInfo, req.Extend)
}

func (s *sDevice) ImportModel(ctx context.Context, req *req.DeviceImportModelReq) (err error) {
	s.Model(ctx).WherePri(req.DeviceId).Data("model_path", req.Path).Update()
	return
}

// 获取设备传感器报警列表
func (s *sDevice) GetSensorAlarmList(ctx context.Context, deviceId int64) (sensors []*res.SensorAlarmRow, err error) {
	err = NewManageSensor().Model(ctx).Where("device_id", deviceId).Scan(&sensors)
	if err != nil {
		return
	}

	for _, sensor := range sensors {
		var thresholds []*req.ThresholdRow
		thresholds, err = NewManageThreshold().GetSensorThresholds(ctx, sensor.Id)
		if err != nil {
			return
		}
		sensor.Thresholds = thresholds
	}

	return
}

func (s *sDevice) SaveSensorAlarmList(ctx context.Context, deviceId int64, sensors []*res.SensorAlarmRow) (err error) {

	thresholdService := NewManageThreshold()
	// if len(sensors) == 0 {
	// NewManageSensor().Model(ctx).Where("device_id", deviceId).Delete()
	// }
	//  goframe 子查询删除
	// thresholdService.Model(ctx).Where("sensor_id in (?)", ).Delete()
	// 创建子sql语句
	subSql := g.DB().Model("manage_sensor").Where("device_id", deviceId).Fields("id")
	thresholdService.Model(ctx).Where("sensor_id in ?", subSql).Delete()
	for _, sensor := range sensors {
		for _, threshold := range sensor.Thresholds {
			thresholdService.Save(ctx, threshold)
		}
	}
	return
}

func (s *sDevice) GetSensorNow(ctx context.Context, deviceId int64) (out []*res.SensorInfo, err error) {

	NewManageSensor().Model(ctx).Where("device_id", deviceId).Scan(&out)
	for _, sensor := range out {
		val, err := NewManageSensorDataCache().Get(ctx, sensor.Id)

		// 没有数据的话从influxdb获取
		if err != nil {

			list, total, err := NewManageInfluxdb().SearchTable(ctx, &model.PageListReq{
				PageReq: page.PageReq{
					Page:     1,
					PageSize: 1,
				},
			}, &req.ManageInfluxdbSearch{
				// Precision: 1,
				SensorIds: []int64{sensor.Id},
				DeviceId:  deviceId,
			})

			if err != nil || total == 0 {
				continue
			}

			sensor.Value = list[0][fmt.Sprintf("c_%d", sensor.Id)]
			sensor.DataTime = list[0]["time"].(time.Time)
			continue
		}
		sensor.Value = val.Value
		sensor.DataTime = val.CreateTime
	}
	return
}

func (s *sDevice) getAllChildrenIds(ctx context.Context, parentId int64) ([]int64, error) {
	result := []int64{parentId}
	queue := []int64{parentId}

	for len(queue) > 0 {
		// var next []int64
		var rows []*do.ManageArea

		err := NewManageArea().Model(ctx).WhereIn("parent_id", queue).Fields("id").Scan(&rows)
		// fmt.Println("next====================", next)
		next := make([]int64, 0, len(rows))
		for _, r := range rows {
			next = append(next, gconv.Int64(r.Id))
		}
		if err != nil {
			return nil, err
		}

		if len(next) == 0 {
			break
		}

		result = append(result, next...)
		queue = next
	}

	return result, nil
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

	if len(in.AreaIds) > 0 {
		ids := []int64{}
		for _, id := range in.AreaIds {
			childrenIds, err := s.getAllChildrenIds(ctx, id)
			if err != nil {
				continue
			}
			ids = append(ids, childrenIds...)
		}
		query = query.WhereIn("area_id", ids)
	}

	return
}

func (s *sDevice) handleDeviceConnect(serverInfo *res.ServerInfo, extend *gjson.Json) (err error) {
	switch serverInfo.Type {
	case gateway.SERVER_MODBUS_TCP:
		{
			err = s.handleModbusTcp(serverInfo, extend.Get("slaveId").Int())
		}
	case gateway.SERVER_MODBUS_RTU:
		{
			err = s.handleModbusRtu(serverInfo, extend.Get("slaveId").Int())
		}
	case gateway.SERVER_MODBUS_RTU_OVER_TCP:
		{
			err = s.handleModbusRtuOverTcp(serverInfo, extend.Get("slaveId").Int())
		}
	}
	return
}

// modbus tcp
func (s *sDevice) handleModbusTcp(serverInfo *res.ServerInfo, slave int) (err error) {
	url := fmt.Sprintf("%s:%s", serverInfo.Ip, serverInfo.Port)
	handler := modbus.NewTCPClientHandler(url)
	handler.SlaveId = byte(slave)
	handler.Timeout = 5 * time.Second

	err = handler.Connect()
	if err != nil {
		return
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	rs, err := client.ReadHoldingRegisters(0, 1)
	if err != nil {
		return
	}
	fmt.Println("rs====================", rs)
	return
}

// modbus rtu
func (s *sDevice) handleModbusRtu(serverInfo *res.ServerInfo, slave int) (err error) {
	// handle := modbus.NewRTUClientHandler(sensor)
	// handle.BaudRate = baudRate

	return
}

// modbus rtu_over_tcp
func (s *sDevice) handleModbusRtuOverTcp(serverInfo *res.ServerInfo, slave int) (err error) {
	u, err := url.Parse(fmt.Sprintf("%s:%s", serverInfo.Ip, serverInfo.Port))
	if err != nil {
		return
	}

	handler := modbus.NewTCPClientHandler(u.String())
	handler.SlaveId = byte(slave)
	handler.Timeout = 5 * time.Second

	err = handler.Connect()
	if err != nil {
		return
	}
	defer handler.Close()
	// r := modbus.NewRTUClientHandler("")

	// c := modbus.NewClient2(r, handler)
	return
}
