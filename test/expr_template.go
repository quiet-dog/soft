package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

type ExprTemplate string

type ExprEnv struct {
	Value interface{} `json:"value" expr:"value"`
}

func (t *ExprTemplate) ToExprValueFloat64(value any) (result float64, err error) {
	env := ExprEnv{
		Value: value,
	}

	program, err := expr.Compile(string(*t), expr.Env(env)) // Pass the struct as an environment.
	if err != nil {
		return result, err
	}

	data, err := expr.Run(program, env)
	if err != nil {
		return
	}

	switch v := data.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("result is not a numeric type, got %T", data)
	}
}

func main() {
	var a interface{}
	a = 1

	exprTemplate := ExprTemplate("value + 1")
	result, err := exprTemplate.ToExprValueFloat64(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
