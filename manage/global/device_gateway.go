package global

import (
	"context"
	"devinggo/manage/model/req"
	"devinggo/manage/pkg/gateway"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/model"
	"fmt"

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
		if server.Type == gateway.SERVER_OPC {
			cfg := gateway.Config{
				Host: server.Ip,
				Port: server.Port,
				Type: server.Type,
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
			for _, sensor := range sensors {
				opc, err := manage.ManageOpc().Read(context.Background(), sensor.Extend.Get("id").Int64())
				if err != nil {
					continue
				}
				jsonB, err := gjson.Marshal(map[string]interface{}{
					"id":     sensor.Id,
					"nodeId": opc.NodeId,
				})
				if err != nil {
					continue
				}
				c.AddNodes(string(jsonB))
			}

		} else {
			// panic("Unsupported gateway type: " + server.Type)
		}
	}

	channel := DeviceGateway.RegisterChannel(1000)
	for v := range channel {
		fmt.Println("注册开始监听============")
		fmt.Println("Received value============:", v.Value, "from channel")
		manage.ManageInfluxdb().StoreDataChannel(context.Background(), v)

	}
}
