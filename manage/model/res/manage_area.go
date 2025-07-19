package res

import "devinggo/manage/model/base"

type AreaTableRow struct {
	base.BaseTable
	// Name   string `json:"name"   description:"区域名称"` // 区域名称
	Remark   string `json:"remark" description:"备注"`     // 备注
	ParentId int64  `json:"parentId" description:"父级ID"` // 父级ID
	Name     string `json:"name" description:"区域名称"`     // 区域名称
	Sort     int64  `json:"sort" description:"排序"`       // 排序
}

type AreaTree struct {
	Label    string     `json:"label" description:"名称"`      // 名称
	Value    int64      `json:"value" description:"ID"`      // ID
	Children []AreaTree `json:"children" description:"子节点"`  // 子节点
	IsLeaf   bool       `json:"isLeaf" description:"是否叶子节点"` // 是否叶子节点
}
