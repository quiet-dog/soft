// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package manage

import (
	"context"
	"database/sql"
	"devinggo/manage/model/common"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/model/res/device"
	"devinggo/manage/pkg/expr_template"
	"devinggo/manage/pkg/gateway"
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
)

type (
	IManageArea interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAreaSearch) (rs []*res.AreaTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageAreaSave) (id int64, err error)
		IsExitAreaById(ctx context.Context, id int64) (bool, error)
		Tree(ctx context.Context, in *req.ManageAreaSearch) (rs []*res.AreaTree, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		UpdateInfo(ctx context.Context, in *req.ManageAreaUpdateInfo) (out sql.Result, err error)
		AllTreeById(ctx context.Context, in *req.ManageAreaSearch) (items []*res.AreaTree, err error)
	}

	IManageAlarmLabel interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAlarmLabelSearch) (rs []*res.AlarmLabelTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageAlarmLabelSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Read(ctx context.Context, alarmLabelId int64) (alarmLabelInfo *res.AlarmLabelInfo, err error)
	}

	IManageServer interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (rs []*res.ServerTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageServerSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Types(ctx context.Context) (rs []*res.ServerType, err error)
		Tree(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (rs []*res.AreaTree, err error)
		Read(ctx context.Context, serverId int64) (DeviceInfo *res.ServerInfo, err error)
		UpdateInfo(ctx context.Context, in *req.ManageServerUpdateInfo) (out sql.Result, err error)
	}

	IManageDevice interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceSearch) (rs []*res.DeviceTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageDeviceSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Tree(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceSearch) (rs []*res.AreaTree, err error)
		GetInfoById(ctx context.Context, deviceId int64) (DeviceInfo *res.DeviceInfo, err error)
		GetInfoByIds(ctx context.Context, deviceIds []int64) (DeviceInfos []*res.DeviceInfo, err error)
		Read(ctx context.Context, deviceId int64) (DeviceInfo *res.DeviceInfo, err error)
		TestConnect(ctx context.Context, in *req.DeviceTestConnectReq) (err error)
		ImportModel(ctx context.Context, in *req.DeviceImportModelReq) (err error)
		// GetOpc(ctx context.Context, deviceId int64) (opc *res.OpcInfo, err error)
	}

	IManageSensorType interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageSensorTypeSearch) (rs []*res.SensorTypeTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageSensorTypeSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Tree(ctx context.Context, req *model.PageListReq, in *req.ManageSensorTypeSearch) (rs []*res.AreaTree, err error)
	}

	IManageSensor interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageSensorSearch) (rs []*res.SensorTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageSensorSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Tree(ctx context.Context, req *model.PageListReq, in *req.ManageSensorSearch) (rs []*res.AreaTree, err error)
		ReadData(ctx context.Context, req *req.ManageSensorReadData) (rs *common.TemplateEnv, err error)
		TranslateData(ctx context.Context, req *req.ManageSensorTranslate) (rs common.Value, err error)
		ReadInfluxdbFormat(ctx context.Context, sensorId int64) (out *common.SensorToInfluxdb, err error)
		Read(ctx context.Context, sensorId int64) (sensorInfo *res.SensorInfo, err error)
		ReadEchart(ctx context.Context, re *model.PageListReq, in *req.ManageInfluxdbOneSensorSearch) (out *res.SensorEchart, err error)
	}

	IManageOpc interface {
		InitOpc(ctx context.Context, serverId int64) (result []*device.OpcTree, err error)
		Tree(ctx context.Context, in *req.OpcTreeReq) (rs []*res.OpcTree, err error)
		ReadData(ctx context.Context, opcId int64) (rs *common.TemplateEnv, err error)
		Read(ctx context.Context, opcId int64) (opcInfo *res.OpcInfo, err error)
	}

	IManageInfluxdb interface {
		StoreDataChannel(ctx context.Context, msg gateway.Msg) (err error)
		Store(ctx context.Context, data gateway.Value, sensorId int64) (err error)
		SearchSensorDataList(ctx context.Context, req *model.PageListReq, in *req.ManageInfluxdbSearch) (out *res.SensorDataList, err error)
		SearchSensorEchart(ctx context.Context, re *model.PageListReq, in *req.ManageInfluxdbOneSensorSearch) (out *res.SensorDataList, err error)
	}

	IManageModbus interface {
		TestDataByDeviceId(ctx context.Context, deviceId int64, in *req.ManageSensorReadData) (rs *common.TemplateEnv, err error)
	}

	IManageSensorDataCache interface {
		Model(ctx context.Context) *gredis.Redis
		Store(ctx context.Context, key int64, value gateway.Value) (v *gvar.Var, err error)
		Get(ctx context.Context, key int64) (t gateway.Value, err error)
		Delete(ctx context.Context, key int64) (n int64, err error)
		StoreDevice(ctx context.Context, sensorId int64) (v *gvar.Var, err error)
		GetDevice(ctx context.Context, deviceId int64) (data []int64, err error)
		DeleteDevice(ctx context.Context, sensorId int64) (v *gvar.Var, err error)
	}

	IManageSensorTemplateCache interface {
		Get(ctx context.Context, sensorId int64) (template expr_template.ExprTemplate, err error)
		Store(ctx context.Context, sensorId int64) (template expr_template.ExprTemplate, err error)
	}

	IManageThreshold interface {
		Model(ctx context.Context) *gdb.Model
		Save(ctx context.Context, in *req.ThresholdRow) (id int64, err error)
		Delete(ctx context.Context, in *req.ThresholdRow) (err error)
		CleanDiff(ctx context.Context, in *req.ThresholdRow) (err error)
		AddThreshold(ctx context.Context, in *req.ManageThresholdAddReq) (err error)
		GetSensorThresholds(ctx context.Context, sensorId int64) (out []*req.ThresholdRow, err error)
	}

	IManageThresholdCache interface {
		Model(ctx context.Context) *gredis.Redis
		Get(ctx context.Context, sensorId int64) (out []*req.ThresholdRow, err error)
		Store(ctx context.Context, sensorId int64) (out []*req.ThresholdRow, err error)
	}

	IManageEvent interface {
		Save(ctx context.Context, in *req.ManageEventReq) (id int64, err error)
		CheckEvent(ctx context.Context, sensorId int64, value gateway.Value) (id int64, err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageEventSearch) (res []*res.EventTableRow, total int, err error)
	}

	IManageAlarm interface {
		Save(ctx context.Context, in req.ManageAlarmSave) (id int64, err error)
		Read(ctx context.Context, alarmId int64) (alarmInfo *res.AlarmInfo, err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAlarmSearch) (res []*res.AlarmTableRow, total int, err error)
		Delete(ctx context.Context, ids []int64) (err error)
	}

	IManageDeviceControl interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageDeviceControlSearch) (res []*res.DeviceControlTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageDeviceControlSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Read(ctx context.Context, sensorId int64) (deviceControlInfo *res.DeviceControlInfo, err error)
		UpdateInfo(ctx context.Context, in *req.ManageDeviceControlInfo) (out sql.Result, err error)
		AddControl(ctx context.Context, in *req.ManageAddDeviceControlInfo) (err error)
		Control(ctx context.Context, controlId int64) (err error)
	}

	IManageThird interface {
		SendSensorDataByWs(ctx context.Context, value gateway.Msg)
	}
)

var (
	localManageArea                IManageArea
	localManageAlarmLabel          IManageAlarmLabel
	localManageServer              IManageServer
	localManageDevice              IManageDevice
	localManageSensorType          IManageSensorType
	localManageSensor              IManageSensor
	localManageOpc                 IManageOpc
	localInfluxdb                  IManageInfluxdb
	localModbus                    IManageModbus
	localManageSensorDataCache     IManageSensorDataCache
	localManageThreshold           IManageThreshold
	localManageThresholdCache      IManageThresholdCache
	localManageEvent               IManageEvent
	localManageSensorTemplateCache IManageSensorTemplateCache
	localManageAlarm               IManageAlarm
	localManageDeviceControl       IManageDeviceControl
	localManageThird               IManageThird
)

func ManageArea() IManageArea {
	if localManageArea == nil {
		panic("implement not found for interface IManageArea, forgot register?")
	}
	return localManageArea
}

func RegisterManageArea(i IManageArea) {
	localManageArea = i
}

func ManageAlarmLabel() IManageAlarmLabel {
	if localManageAlarmLabel == nil {
		panic("implement not found for interface IManageAlarmLabel, forgot register?")
	}
	return localManageAlarmLabel
}

func RegisterManageAlarmLabel(i IManageAlarmLabel) {
	localManageAlarmLabel = i
}

func ManageServer() IManageServer {
	if localManageServer == nil {
		panic("implement not found for interface IManageServer, forgot register?")
	}
	return localManageServer
}

func RegisterManageServer(i IManageServer) {
	localManageServer = i
}

func ManageDevice() IManageDevice {
	if localManageDevice == nil {
		panic("implement not found for interface IManageDevice, forgot register?")
	}
	return localManageDevice
}

func RegisterManageDevice(i IManageDevice) {
	localManageDevice = i
}

func ManageSensorType() IManageSensorType {
	if localManageSensorType == nil {
		panic("implement not found for interface IManageSensorType, forgot register?")
	}
	return localManageSensorType
}

func RegisterManageSensorType(i IManageSensorType) {
	localManageSensorType = i
}

func ManageSensor() IManageSensor {
	if localManageSensor == nil {
		panic("implement not found for interface IManageSensor, forgot register?")
	}
	return localManageSensor
}

func RegisterManageSensor(i IManageSensor) {
	localManageSensor = i
}

func ManageOpc() IManageOpc {
	if localManageOpc == nil {
		panic("implement not found for interface IManageOpc, forgot register?")
	}
	return localManageOpc
}

func RegisterManageOpc(i IManageOpc) {
	localManageOpc = i
}

func ManageInfluxdb() IManageInfluxdb {
	if localInfluxdb == nil {
		panic("implement not found for interface localInfluxdb, forgot register?")
	}
	return localInfluxdb
}

func RegisterManageInfluxdb(i IManageInfluxdb) {
	localInfluxdb = i
}

func ManageModbus() IManageModbus {
	if localModbus == nil {
		panic("implement not found for interface localModbus, forgot register?")
	}
	return localModbus
}

func RegisterManageModbus(i IManageModbus) {
	localModbus = i
}

func ManageSensorDataCache() IManageSensorDataCache {
	if localManageSensorDataCache == nil {
		panic("implement not found for interface localManageSensorDataCache, forgot register?")
	}
	return localManageSensorDataCache
}

func RegisterManageSensorDataCache(i IManageSensorDataCache) {
	localManageSensorDataCache = i
}

func ManageThreshold() IManageThreshold {
	if localManageThreshold == nil {
		panic("implement not found for interface localManageThreshold, forgot register?")
	}
	return localManageThreshold
}

func RegisterManageThreshold(i IManageThreshold) {
	localManageThreshold = i
}

func ManageThresholdCache() IManageThresholdCache {
	if localManageThresholdCache == nil {
		panic("implement not found for interface localManageThresholdCache, forgot register?")
	}
	return localManageThresholdCache
}

func RegisterManageThresholdCache(i IManageThresholdCache) {
	localManageThresholdCache = i
}

func ManageEvent() IManageEvent {
	if localManageEvent == nil {
		panic("implement not found for interface localManageEvent, forgot register?")
	}
	return localManageEvent
}

func RegisterManageEvent(i IManageEvent) {
	localManageEvent = i
}

// IManageSensorTemplateCache
func ManageSensorTemplateCache() IManageSensorTemplateCache {
	if localManageSensorTemplateCache == nil {
		panic("implement not found for interface localManageSensorTemplateCache, forgot register?")
	}
	return localManageSensorTemplateCache
}

func RegisterManageSensorTemplateCache(i IManageSensorTemplateCache) {
	localManageSensorTemplateCache = i
}

func ManageAlarm() IManageAlarm {
	if localManageAlarm == nil {
		panic("implement not found for interface localManageSensorTemplateCache, forgot register?")
	}
	return localManageAlarm
}

func RegisterManageAlarm(i IManageAlarm) {
	localManageAlarm = i
}

func ManageDeviceControl() IManageDeviceControl {
	if localManageDeviceControl == nil {
		panic("implement not found for interface localManageDeviceControl, forgot register?")
	}
	return localManageDeviceControl
}

func RegisterManageDeviceControl(i IManageDeviceControl) {
	localManageDeviceControl = i
}

func ManageThird() IManageThird {
	if localManageThird == nil {
		panic("implement not found for interface localManageThird, forgot register?")
	}
	return localManageThird
}

func RegisterManageThird(i IManageThird) {
	localManageThird = i
}
