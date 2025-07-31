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

type AfterSelectInterface interface {
	AfterSelectHook(ctx context.Context, in *gdb.HookSelectInput, result *gdb.Result) (err error)
}
type BeforeSelectInterface interface {
	BeforeSelectHook(ctx context.Context, in *gdb.HookSelectInput) (err error)
}

type AfterInsertInterface interface {
	AfterInsertHook(ctx context.Context, in *gdb.HookInsertInput, result *sql.Result) (err error)
}
type BeforeInsertInterface interface {
	BeforeInsertHook(ctx context.Context, in *gdb.HookInsertInput) (err error)
}
type BeforeUpdateInterface interface {
	BeforeUpdateHook(ctx context.Context, in *gdb.HookUpdateInput) (err error)
}

type AfterUpdateInterface interface {
	AfterUpdateHook(ctx context.Context, in *gdb.HookUpdateInput, result *sql.Result) (err error)
}

type AfterDeleteInterface interface {
	AfterDeleteHook(ctx context.Context, in *gdb.HookDeleteInput, result *sql.Result) (err error)
}

type BeforeDeleteInterface interface {
	BeforeDeleteHook(ctx context.Context, in *gdb.HookDeleteInput) (err error)
}

type OrmHook interface {
	AfterSelectInterface
	BeforeSelectInterface

	AfterUpdateInterface
	BeforeUpdateInterface

	AfterInsertInterface
	BeforeInsertInterface

	AfterDeleteInterface
	BeforeDeleteInterface
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

			if v, ok := h.Inner.(BeforeSelectInterface); ok {
				err = v.BeforeSelectHook(ctx, in)
				if err != nil {
					return
				}
			}

			result, err = in.Next(ctx)
			if err != nil {
				return result, err
			}

			if v, ok := h.Inner.(AfterSelectInterface); ok {
				err = v.AfterSelectHook(ctx, in, &result)
				if err != nil {
					return
				}
			}

			// for _, v := range h.Inner {
			// 	if inn, ok := v.(HookSelect); ok {
			// 		result, err = inn.SelectHook(ctx, in)
			// 		if err != nil {
			// 			return result, err
			// 		}
			// 	}
			// }

			if *options.UserRelate {
				return hook.UserRelate(ctx, result, options.Params.([]string))
			}
			return
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			if v, ok := h.Inner.(BeforeInsertInterface); ok {
				err = v.BeforeInsertHook(ctx, in)
				if err != nil {
					return
				}
			}

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
			if err != nil {
				return
			}

			if v, ok := h.Inner.(AfterInsertInterface); ok {
				err = v.AfterInsertHook(ctx, in, &result)
				if err != nil {
					return
				}
			}
			g.Log().Debug(ctx, "Insert:", err)
			return
		},
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {

			if v, ok := h.Inner.(BeforeUpdateInterface); ok {
				err = v.BeforeUpdateHook(ctx, in)
				if err != nil {
					return
				}
			}

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
			if err != nil {
				return
			}

			if v, ok := h.Inner.(AfterUpdateInterface); ok {
				err = v.AfterUpdateHook(ctx, in, &result)
				if err != nil {
					return
				}
			}
			return
		},
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {

			if v, ok := h.Inner.(BeforeDeleteInterface); ok {
				err = v.BeforeDeleteHook(ctx, in)
				if err != nil {
					return
				}
			}
			if *options.CacheEvict {
				err = hook.CleanCache[gdb.HookDeleteInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)

			if err != nil {
				return
			}
			if v, ok := h.Inner.(AfterDeleteInterface); ok {
				err = v.AfterDeleteHook(ctx, in, &result)
				if err != nil {
					return
				}
			}
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
