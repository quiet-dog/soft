package thirdmodule

import (
	"context"
	"devinggo/manage/global"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/pkg/websocket"
)

type ThirdSendMsg struct {
	SensorId int64 `json:"sensorId"`
	Data     any   `json:"data"`
	ExprData any   `json:"exprData"`
}

func Start() {

	channels := global.DeviceGateway.RegisterChannel(100)
	for channel := range channels {
		t, err := manage.ManageSensorTemplateCache().Get(context.Background(), channel.Value.ID)
		if err != nil {
			continue
		}

		manage.ManageSensorDataCache().GetDevice(context.Background(), channel.Value.DeviceId)

		thirdSendMsg := ThirdSendMsg{
			SensorId: channel.Value.ID,
			Data:     channel,
		}

		exprData, err := t.ToExprValue(t)
		if err == nil {
			thirdSendMsg.ExprData = exprData
		}

		websocket.SendToTopic("sensor_data", &websocket.WResponse{
			Event: websocket.Publish,
			Data:  thirdSendMsg,
			Code:  200,
		})
	}
}
