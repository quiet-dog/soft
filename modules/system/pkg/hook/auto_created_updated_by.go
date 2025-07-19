// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/orm"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func AutoCreatedUpdatedByInsert(ctx context.Context, in *gdb.HookInsertInput) (err error) {
	hasCreatedBy := gstr.InArray(orm.GetTableFieds(in.Model), "created_by")
	hasUpdatedBy := gstr.InArray(orm.GetTableFieds(in.Model), "updated_by")
	if hasCreatedBy || hasUpdatedBy {
		userId := contexts.New().GetUserId(ctx)
		if !g.IsEmpty(in.Data) && !g.IsEmpty(userId) {
			for _, data := range in.Data {
				if hasCreatedBy {
					if _, ok := data["created_by"]; !ok {
						data["created_by"] = contexts.New().GetUserId(ctx)
					}
					if _, ok := data["createdBy"]; !ok {
						data["created_by"] = contexts.New().GetUserId(ctx)
					}
				}
				if hasUpdatedBy {
					if _, ok := data["updated_by"]; !ok {
						data["updated_by"] = contexts.New().GetUserId(ctx)
					}
					if _, ok := data["updatedBy"]; !ok {
						data["updated_by"] = contexts.New().GetUserId(ctx)
					}
				}
			}
		}
	}
	return
}

func AutoCreatedUpdatedByUpdatefunc(ctx context.Context, in *gdb.HookUpdateInput) (err error) {
	//g.Log().Debug(ctx, "in", in)
	if gstr.InArray(orm.GetTableFieds(in.Model), "updated_by") {
		userId := contexts.New().GetUserId(ctx)
		if !g.IsEmpty(in.Data) && !g.IsEmpty(userId) {
			switch in.Data.(type) {
			case map[string]interface{}:
				//g.Log().Info(ctx, "map")
				if _, ok := in.Data.(map[string]interface{})["updated_by"]; !ok {
					in.Data.(map[string]interface{})["updated_by"] = userId
				}
				if _, ok := in.Data.(map[string]interface{})["updatedBy"]; !ok {
					in.Data.(map[string]interface{})["updatedBy"] = userId
				}
			case string:
				in.Data = in.Data.(string) + ", `updated_by` = " + gconv.String(userId)
			}
		}
	}
	return
}
