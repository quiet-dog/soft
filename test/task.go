package main

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

func main() {
	var (
		err error
		ctx = context.Background()
	)

	gg := gcron.New()
	_, err = gg.Add(ctx, "* * * * * *", func(ctx context.Context) {
		g.Log().Print(ctx, "Every second")
	}, "MySecondCronJob")
	if err != nil {
		panic(err)
	}

	_, err = gg.Add(ctx, "0 30 * * * *", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour on the half hour")
	})
	if err != nil {
		panic(err)
	}

	_, err = gg.Add(ctx, "@hourly", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour")
	})
	if err != nil {
		panic(err)
	}

	_, err = gg.Add(ctx, "@every 1h30m", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour thirty")
	})
	if err != nil {
		panic(err)
	}

	g.Dump(gcron.Entries())

	time.Sleep(3 * time.Second)

	g.Log().Print(ctx, `stop cronjob "MySecondCronJob"`)
	gg.Stop("MySecondCronJob")

	time.Sleep(3 * time.Second)

	g.Log().Print(ctx, `start cronjob "MySecondCronJob"`)
	gg.Start("MySecondCronJob")

	time.Sleep(3 * time.Second)
}
