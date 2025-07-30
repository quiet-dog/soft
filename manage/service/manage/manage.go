// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package manage

import (
	"context"
	"devinggo/manage/model/common"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/model/res/device"
	"devinggo/manage/pkg/gateway"
	"devinggo/modules/system/model"
)

type (
	IManageArea interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAreaSearch) (rs []*res.AreaTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageAreaSave) (id int64, err error)
		IsExitAreaById(ctx context.Context, id int64) (bool, error)
		Tree(ctx context.Context, in *req.ManageAreaSearch) (rs []*res.AreaTree, err error)
		Delete(ctx context.Context, ids []int64) (err error)
	}

	IManageAlarmLabel interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAlarmLabelSearch) (rs []*res.AlarmLabelTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageAlarmLabelSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
	}

	IManageServer interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (rs []*res.ServerTableRow, total int, err error)
		Save(ctx context.Context, in *req.ManageServerSave) (id int64, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		Types(ctx context.Context) (rs []*res.ServerType, err error)
		Tree(ctx context.Context, req *model.PageListReq, in *req.ManageServerSearch) (rs []*res.AreaTree, err error)
		Read(ctx context.Context, serverId int64) (DeviceInfo *res.ServerInfo, err error)
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
		Store(ctx context.Context, data common.TemplateEnv, sensorId int64) (err error)
		SearchSensorDataList(ctx context.Context, req *model.PageListReq, in *req.ManageInfluxdbSearch) (out *res.SensorDataList, err error)
		SearchSensorEchart(ctx context.Context, re *model.PageListReq, in *req.ManageInfluxdbOneSensorSearch) (out *res.SensorDataList, err error)
	}

	IManageModbus interface {
		TestDataByDeviceId(ctx context.Context, deviceId int64, in *req.ManageSensorReadData) (rs *common.TemplateEnv, err error)
	}
)

var (
	localManageArea       IManageArea
	localManageAlarm      IManageAlarmLabel
	localManageServer     IManageServer
	localManageDevice     IManageDevice
	localManageSensorType IManageSensorType
	localManageSensor     IManageSensor
	localManageOpc        IManageOpc
	localInfluxdb         IManageInfluxdb
	localModbus           IManageModbus
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
	if localManageAlarm == nil {
		panic("implement not found for interface IManageAlarm, forgot register?")
	}
	return localManageAlarm
}

func RegisterManageAlarmLabel(i IManageAlarmLabel) {
	localManageAlarm = i
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
