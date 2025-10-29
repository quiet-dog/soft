package cmd

import (
	"context"
	thirdmodule "devinggo/manage/modules/third_module"

	"github.com/gogf/gf/v2/os/gcmd"
)

// 是否开启推送给第三方
var (
	ThirdSensorData = &gcmd.Command{
		Name:  "third_sensor_data",
		Usage: "third_sensor_data",
		Brief: "是否开启推送给第三方",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化设备网关
			thirdmodule.Start()
			return
		},
	}
)
