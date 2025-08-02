package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	config = gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	group = "cache"
	ctx   = gctx.New()
)

func main() {
	gredis.SetConfig(&config, group)

	// _, err := g.Redis(group).Set(ctx, "key", "value")
	// if err != nil {
	// 	g.Log().Fatal(ctx, err)
	// }
	// value, err := g.Redis(group).Get(ctx, "key")
	// if err != nil {
	// 	g.Log().Fatal(ctx, err)
	// }

	// scan to struct
	type User struct {
		Id    uint64
		Name  string
		Score []User
	}

	_, err := g.Redis(group).Set(ctx, "user", "1111")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	va, err := g.Redis(group).Get(ctx, "user")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	if va.IsEmpty() || va.IsNil() {
		panic("22")
	}

}
