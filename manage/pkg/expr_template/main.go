package expr_template

import (
	"fmt"

	"github.com/expr-lang/expr"
)

type ExprTemplate string

type ExprEnv struct {
	Value interface{} `json:"value" expr:"value"`
}

func (t *ExprTemplate) ToValueFloat64(value any) (result float64, err error) {
	switch v := value.(type) {
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
		return 0, fmt.Errorf("result is not a numeric type, got %T", value)
	}

}

// 根据expr模板转换为float64
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
		return result, err
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

// 根据expr转换为需要向influxdb插入的数据
func (t *ExprTemplate) ToExprValueInfluxdbFloat64(value any) string {
	v, err := t.ToExprValueFloat64(value)
	if err != nil {
		return fmt.Sprintf("%fi", 0.0)
	}
	return fmt.Sprintf("%f", v)
}

// 直接转换为influxdb的float64
func (t *ExprTemplate) ToValueInfluxdbFloat64(value any) string {
	v, err := t.ToValueFloat64(value)
	if err != nil {
		return fmt.Sprintf("%fi", 0.0)
	}
	return fmt.Sprintf("%fi", v)
}
