package cmd

import (
	"context"
	"devinggo/manage/global"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	DeviceGateway = &gcmd.Command{
		Name:  "device_gateway",
		Usage: "device_gateway",
		Brief: "设备网关服务，处理设备数据采集和通信",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化设备网关
			global.InitDeviceGateway()
			return
		},
	}
)
