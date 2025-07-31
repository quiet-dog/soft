package base

import "github.com/gogf/gf/v2/os/gtime"

type BaseTable struct {
	Id        int64       `json:"id"                 description:"主键"` // 主键
	CreatedBy int64       `json:"createdBy"  description:"创建者"`        // 创建者
	UpdatedBy int64       `json:"updatedBy"  description:"更新者"`        // 更新者
	CreatedAt *gtime.Time `json:"createdAt"  description:""`           //
	UpdatedAt *gtime.Time `json:"updatedAt"  description:""`           //
}

type BaseIds struct {
	Ids []int64 `json:"ids" description:"主键集合"` // 主键集合
}

type BaseId struct {
	Id int64 `json:"id"`
}
