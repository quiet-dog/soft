package req

import "devinggo/manage/model/base"

type ManageAreaSearch struct {
	Name     string `json:"name" description:"区域名称"`     // 区域名称
	Remark   string `json:"remark" description:"备注"`     // 备注
	ParentId int64  `json:"parentId" description:"父级ID"` // 父级ID
	base.BaseIds
}

type ManageAreaSave struct {
	Name     string `json:"name" v:"required|max-length:50" description:"区域名称"` // 区域名称
	Remark   string `json:"remark" description:"备注"`                            // 备注
	ParentId int64  `json:"parentId" description:"父级ID"`                        // 父级ID
	Sort     int    `json:"sort" description:"排序"`                              // 排序
}

type ManageAreaUpdateInfo struct {
	ManageAreaSave
	base.BaseId
}

type ManageAreaDelete struct {
	base.BaseIds
}

type ManageAreaTreeById struct {
	base.BaseId
}
