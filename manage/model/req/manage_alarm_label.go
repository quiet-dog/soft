package req

type ManageAlarmLabelSearch struct {
	Level  string `json:"level" description:"标签等级"` // 标签等级
	Name   string `json:"name" description:"标签名称"`  // 标签名称
	Remark string `json:"remark" description:"备注"`  // 备注
}

type ManageAlarmLabelSave struct {
	Name   string `json:"name" v:"required|max-length:255" description:"标签名称"` // 标签名称
	Level  string `json:"level" v:"required" description:"标签等级"`               // 标签等级
	Remark string `json:"remark" description:"备注"`                             // 备注
	Color  string `json:"color" description:"标签颜色"`                            // 标签颜色
}
