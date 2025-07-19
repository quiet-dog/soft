package res

import (
	"devinggo/manage/model/base"
)

type SensorTypeTableRow struct {
	base.BaseTable
	Name   string `json:"name" description:"设备名称"` // 设备名称
	Unit   string `json:"unit" description:"单位"`   // 单位
	Remark string `json:"remark" description:"备注"` // 备注
}
