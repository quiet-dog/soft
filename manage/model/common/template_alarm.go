package common

import (
	"github.com/expr-lang/expr"
)

type AlarmTemplate struct {
	Template string
}

func (a *AlarmTemplate) IsAlarmFloat64(value float64) bool {
	pre := map[string]interface{}{
		"value": value,
	}

	program, err := expr.Compile(a.Template, expr.Env(pre))
	if err != nil {
		return false
	}

	output, err := expr.Run(program, pre)
	if err != nil {
		panic(err)
	}

	if v, ok := output.(bool); ok {
		return v
	}
	return false
}
