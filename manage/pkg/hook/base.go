package hook

import (
	"context"
	"database/sql"
	"devinggo/modules/system/pkg/hook"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type HookHand interface {
	SelectHook(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error)
	InsertHook(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error)
	UpdateHook(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error)
	DeleteHook(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error)
}

type HookSelect interface {
	SelectHook(ctx context.Context, in *gdb.HookSelectInput) (gdb.Result, error)
}

type HookInsert interface {
	InsertHook(ctx context.Context, in *gdb.HookInsertInput) (sql.Result, error)
}

type HookUpdate interface {
	UpdateHook(ctx context.Context, in *gdb.HookUpdateInput) (sql.Result, error)
}

type HookDelete interface {
	DeleteHook(ctx context.Context, in *gdb.HookDeleteInput) (sql.Result, error)
}

type Hook struct {
	Inner interface{}
}

func (h *Hook) Bind() gdb.HookHandler {
	defaultAutoCreatedUpdatedBy := true
	defaultCacheEvict := true
	defaultUserRelate := false
	var options = hook.HookOptions{
		AutoCreatedUpdatedBy: &defaultAutoCreatedUpdatedBy,
		CacheEvict:           &defaultCacheEvict,
		UserRelate:           &defaultUserRelate,
	}
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return result, err
			}
			if h.Inner != nil {
				result, err = h.Inner.(HookSelect).SelectHook(ctx, in)
				if err != nil {
					return result, err
				}
			}
			if *options.UserRelate {
				return hook.UserRelate(ctx, result, options.Params.([]string))
			}
			return
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			if *options.AutoCreatedUpdatedBy {
				err = hook.AutoCreatedUpdatedByInsert(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if *options.CacheEvict {
				err = hook.CleanCache[gdb.HookInsertInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}
			//g.Log().Debug(ctx, "in:", in)
			result, err = in.Next(ctx)
			g.Log().Debug(ctx, "Insert:", err)
			return
		},

		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if *options.AutoCreatedUpdatedBy {
				err = hook.AutoCreatedUpdatedByUpdatefunc(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if *options.CacheEvict {
				err = hook.CleanCache[gdb.HookUpdateInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			return
		},

		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {

			if *options.CacheEvict {
				err = hook.CleanCache[gdb.HookDeleteInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			return
		},
	}
}

func Bind(Inner ...interface{}) gdb.HookHandler {
	h := &Hook{}
	if len(Inner) > 0 {
		h.Inner = Inner[0]
	}
	return h.Bind()
}
