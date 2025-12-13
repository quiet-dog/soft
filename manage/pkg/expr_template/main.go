package expr_template

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/expr-lang/expr"
)

type ExprTemplate string

type ExprEnv struct {
	Value interface{} `json:"value" expr:"value"`
}

// 不经转换的值转换为float64
func (t *ExprTemplate) ToValueFloat64(value any) (result float64, err error) {
	vType := reflect.TypeOf(value)
	switch vType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.(int)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.(uint)), nil
	case reflect.Float32, reflect.Float64:
		return value.(float64), nil
	// case reflect.String:
	// 	return value.(string), nil
	// case reflect.Bool:
	// return value.(bool), nil
	default:
		return 0, fmt.Errorf("result is not a numeric type, got %T", value)
	}

}

// 根据expr模板转换为float64
func (t *ExprTemplate) ToExprValueFloat64(value any) (result any, err error) {
	if *t == "" {
		// 如果value是数值类型的话直接返回
		vType := reflect.TypeOf(value)
		switch vType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(value.(int)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return float64(value.(uint)), nil
		case reflect.Float32, reflect.Float64:
			return value.(float64), nil
		case reflect.String:
			return value.(string), nil
		case reflect.Bool:
			return value.(bool), nil
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			jsonStr, err := json.Marshal(value)
			if err != nil {
				return result, err
			}
			return string(jsonStr), nil
		default:
			return 0, fmt.Errorf("result is not a numeric type, got %T", value)
		}
	}

	env := ExprEnv{
		Value: value,
	}
	d, err := json.Marshal(env)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(d, &env)
	if err != nil {
		return result, err
	}
	program, err := expr.Compile(string(*t)) // Pass the struct as an environment.
	if err != nil {
		return result, err
	}

	data, err := expr.Run(program, env)
	if err != nil {
		fmt.Println("============expr run error============", err, value, string(*t), reflect.TypeOf(value))
		return result, err
	}

	vType := reflect.TypeOf(data)
	switch vType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(data.(int)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(data.(uint)), nil
	case reflect.Float32, reflect.Float64:
		return data.(float64), nil
	case reflect.String:
		return data.(string), nil
	case reflect.Bool:
		return data.(bool), nil
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return result, err
		}
		return string(jsonStr), nil
	default:
		return 0, fmt.Errorf("result is not a numeric type, got %T", data)
	}

}

// 根据expr转换为需要向influxdb插入的数据
func (t *ExprTemplate) ToExprValueInfluxdbFloat64(value any) string {
	v, err := t.ToExprValueFloat64(value)
	if err != nil {
		return fmt.Sprintf("%v", 0.0)
	}
	// if err != nil {
	// 	return fmt.Sprintf("%f", 0.0)
	// }
	// return fmt.Sprintf("%f", v)
	// 都转换为字符串
	return fmt.Sprintf("%v", v)
}

func (t *ExprTemplate) ToExprValue(value any) (result any, err error) {
	env := ExprEnv{
		Value: value,
	}
	d, err := json.Marshal(env)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(d, &env)
	if err != nil {
		return result, err
	}
	program, err := expr.Compile(string(*t)) // Pass the struct as an environment.
	if err != nil {
		return result, err
	}
	data, err := expr.Run(program, env)
	if err != nil {
		return result, err
	}
	return data, nil
}

// 不经expr转换直接转换为influxdb的float64
func (t *ExprTemplate) ToValueInfluxdbFloat64(value any) string {
	v, err := t.ToValueFloat64(value)
	if err != nil {
		return fmt.Sprintf("%f", 0.0)
	}
	return fmt.Sprintf("%f", v)
}

func (t *ExprTemplate) ToExprValueInfluxdb(value any) string {
	v, err := t.ToExprValue(value)
	if err != nil {
		return fmt.Sprintf("%fi", 0.0)
	}
	vType := reflect.TypeOf(v)
	switch vType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%di", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%di", v)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v)
	case reflect.String:
		return fmt.Sprintf("\"%s\"", v)
	case reflect.Bool:
		return fmt.Sprintf("%t", v)
	default:
		json, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("%f", v)
		}
		return fmt.Sprintf("\"%s\"", string(json))
	}
}
