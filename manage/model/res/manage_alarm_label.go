package res

import "devinggo/manage/model/base"

type AlarmLabelTableRow struct {
	base.BaseTable
	Remark string `json:"remark" description:"备注"`  // 备注
	Name   string `json:"name" description:"标签名称"`  // 区域名称
	Level  string `json:"level" description:"标签等级"` // 标签等级
	Color  string `json:"color" description:"标签颜色"` // 标签颜色
}
