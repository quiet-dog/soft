package cmd

import (
	"context"
	thirdswmodule "devinggo/manage/modules/third_sw_module"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	ThirdSw = &gcmd.Command{
		Name:  "third_sw",
		Usage: "third_sw",
		Brief: "third_sw",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			thirdswmodule.Start(ctx)
			return
		},
	}
)
