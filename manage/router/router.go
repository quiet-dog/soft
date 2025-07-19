package router

import (
	"devinggo/manage/controller/manage"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/manage", func(group *ghttp.RouterGroup) {
		group.Bind(
			manage.AreaController,
			manage.AlarmLabelController,
			manage.ServerController,
			manage.DeviceController,
			manage.SensorTypeController,
			manage.SensorController,
			manage.OpcController,
		)
	})

}
