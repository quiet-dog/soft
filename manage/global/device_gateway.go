package global

import (
	"context"
	"devinggo/manage/model/req"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/model"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
)

var DeviceGateway *gateway.Gateway

func InitDeviceGateway() {
	DeviceGateway = gateway.NewGateway()
	l, _, err := manage.ManageServer().GetPageListForSearch(context.Background(), &model.PageListReq{}, &req.ManageServerSearch{})
	if err != nil {
		panic(err)
	}
	for _, server := range l {

		duration := server.Interval * int64(time.Second)
		cfg := gateway.Config{
			Host:    server.Ip,
			Port:    server.Port,
			Type:    server.Type,
			SubTime: time.Duration(duration),
			Extend:  server.Extend,
		}

		DeviceGateway.AddClient(server.Id, cfg)
		c, ok := DeviceGateway.Client(server.Id)
		if !ok {
			// panic("Failed to get client for server ID: " + string(server.Id))
			continue
		}
		devices, _, err := manage.ManageDevice().GetPageListForSearch(context.Background(), &model.PageListReq{}, &req.ManageDeviceSearch{
			ServerId: server.Id,
		})
		if err != nil {
			continue
		}
		deviceIds := make([]int64, 0, len(devices))
		for _, device := range devices {
			deviceIds = append(deviceIds, device.Id)
		}
		sensors, _, _ := manage.ManageSensor().GetPageListForSearch(context.Background(), &model.PageListReq{}, &req.ManageSensorSearch{})

		if server.Type == gateway.SERVER_OPC {

			for _, sensor := range sensors {
				opc, err := manage.ManageOpc().Read(context.Background(), sensor.Extend.Get("id").Int64())
				if err != nil {
					continue
				}

				g := gjson.New(nil, false)
				g.Set("id", sensor.Id)
				g.Set("nodeId", opc.NodeId)
				g.Set("deviceId", sensor.DeviceId)
				c.AddNodes(g)
			}

		} else if server.Type == gateway.SERVER_MODBUS_TCP {
			node := []*gjson.Json{}
			for _, device := range devices {
				deviceJson := gjson.New(nil)
				sensorMap := make(map[int64]gateway.ModbusSensor)
				for _, sensor := range sensors {
					if sensor.DeviceId == device.Id {
						deviceJson.Set("deviceId", device.Id)
						deviceJson.Set("slaveId", device.Extend.Get("slaveId").Uint16())
						sensorMap[sensor.Id] = gateway.ModbusSensor{
							StartAddress: sensor.Extend.Get("start").Uint16(),
							Quantity:     sensor.Extend.Get("quantity").Uint16(),
							SensorId:     sensor.Id,
							ReadType:     sensor.Extend.Get("readType").Int64(),
						}
					}
				}
				deviceJson.Set("sensors", sensorMap)
				deviceJson.Set("slaveId", device.Extend.Get("slaveId").Uint16())
				deviceJson.Set("deviceId", device.Id)
				node = append(node, deviceJson)
			}
			c.AddNodes(node...)

		} else {
			// panic("Unsupported gateway type: " + server.Type)
		}
	}

	channel := DeviceGateway.RegisterChannel(1000)
	for v := range channel {
		manage.ManageInfluxdb().StoreDataChannel(context.Background(), v)

	}
}
