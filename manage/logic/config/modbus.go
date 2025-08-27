package config

import (
	"context"
	"devinggo/manage/model/common"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"fmt"
	"net/url"
	"time"

	"github.com/goburrow/modbus"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type sModbus struct {
	base.BaseService
}

func init() {
	manage.RegisterManageModbus(NewManageModbus())
}

func NewManageModbus() *sModbus {
	return &sModbus{}
}

func (s *sModbus) TestDataByDeviceId(ctx context.Context, deviceId int64, in *req.ManageSensorReadData) (rs *common.TemplateEnv, err error) {
	rs = &common.TemplateEnv{}
	deviceInfo, err := manage.ManageDevice().Read(ctx, deviceId)
	if err != nil {
		return
	}

	serverInfo, err := manage.ManageServer().Read(ctx, deviceInfo.ServerId)
	if err != nil {
		return
	}

	in.Extend.Set("slaveId", deviceInfo.Extend.Get("slaveId").Int())

	rs, err = s.handleSensorConnect(serverInfo, in.Extend)

	return
}

func (s *sModbus) ReadData(ctx context.Context, serverType string, host string, slave int, start int64, quantity int64) (rs *common.TemplateEnv, err error) {

	return
}

func (s *sModbus) handleSensorConnect(serverInfo *res.ServerInfo, extend *gjson.Json) (rs *common.TemplateEnv, err error) {
	switch serverInfo.Type {
	case gateway.SERVER_MODBUS_TCP:
		{
			rs, err = s.handleModbusTcpSensor(serverInfo, extend.Get("slaveId").Int(), extend.Get("start").Int64(), extend.Get("quantity").Int64())
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
func (s *sModbus) handleModbusTcpSensor(serverInfo *res.ServerInfo, slave int, start int64, quantity int64) (rs *common.TemplateEnv, err error) {
	rs = &common.TemplateEnv{}
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
	r, err := client.ReadHoldingRegisters(uint16(start), uint16(quantity))
	if err != nil {
		return
	}

	if len(r) < 2 {
		return nil, fmt.Errorf("数据读取失败")
	}

	if len(r)%2 > 0 {
		return nil, fmt.Errorf("数据长度格式失败")
	}

	registerValues := []uint16{}

	for i := 0; i < len(r)-1; i += 2 {
		registerValues = append(registerValues, uint16(r[0])<<8|uint16(r[1]))
	}

	value := common.Value{
		Value: registerValues,
	}
	rs.Value = value
	rs.CreateTime = time.Now()
	rs.Type = "ArrayUint16"
	return
}

// modbus rtu
func (s *sModbus) handleModbusRtu(serverInfo *res.ServerInfo, slave int) (err error) {
	// handle := modbus.NewRTUClientHandler(sensor)
	// handle.BaudRate = baudRate

	return
}

// modbus rtu_over_tcp
func (s *sModbus) handleModbusRtuOverTcp(serverInfo *res.ServerInfo, slave int) (err error) {
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
