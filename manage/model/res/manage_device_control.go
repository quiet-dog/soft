package res

import (
	"devinggo/manage/model/base"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type DeviceControlTableRow struct {
	base.BaseTable
	DeviceId int64       `json:"deviceId" description:"设备ID"` // 设备ID
	Extend   *gjson.Json `json:"extend" description:"扩展信息"`   // 扩展信息
	Name     string      `json:"name"`
}

type DeviceControlInfo struct {
	DeviceControlTableRow
}
