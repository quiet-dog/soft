package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/expr-lang/expr"
)

type Value struct {
	Value interface{} `json:"value"`
}
type TemplateEnv struct {
	Value      Value     `json:"value" expr:"value"`
	CreateTime time.Time `json:"createTime" expr:"createTime"`
	Type       string    `json:"type" expr:"type"`
}

func (temp *TemplateEnv) PrepareExprEnv() map[string]interface{} {
	value := temp.Value.ToValue()

	return map[string]interface{}{
		"value":      value,
		"type":       temp.Type,
		"createTime": temp.CreateTime,
	}
}

func (v *Value) ToValue() interface{} {
	switch val := v.Value.(type) {
	case json.Number:
		f, err := val.Float64()
		if err == nil {
			return f
		}
		return 0.0
	case string:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f
		}
		return val
	case float64, int, int64:
		return val
	case map[string]interface{}:
		return val
	case nil:
		return nil
	default:
		return val // fallback，防止 panic
	}
}

func (v *Value) ToValueInfluxdb() string {
	switch val := v.Value.(type) {
	case int, int64, int32:
		return fmt.Sprintf("%di", val) // InfluxDB整数需要加 i
	case float64, float32:
		return fmt.Sprintf("%f", val)
	case string:
		return fmt.Sprintf("\"%s\"", val) // 字符串需要用引号包起来
	case bool:
		return fmt.Sprintf("%t", val)
	default:
		return fmt.Sprintf("\"%v\"", val) // fallback，防止 panic
	}
}

func (v *Value) UnmarshalJSON(data []byte) error {
	// 先用 interface{} 接收
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// 这里你可以根据 raw 类型做额外判断
	switch val := raw.(type) {
	case map[string]interface{}:
		// 如果是 map，直接赋值
		v.Value = val
	case string:
		// 如果是字符串，尝试转换为 float64
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			v.Value = f
		} else {
			v.Value = val // 如果转换失败，保留原字符串
		}
	case float64, int, int64:
		// 如果是数字类型，直接赋值
		v.Value = val
	case bool:
		// 如果是布尔类型，直接赋值
		v.Value = val
	case nil:
		// 如果是 nil，直接赋值
		v.Value = nil
	default:
		v.Value = val
	}
	return nil
}

func (v *Value) ToValueExpr(template string) any {
	pre := map[string]interface{}{
		"value": v.ToValue(),
	}
	program, err := expr.Compile(template, expr.Env(pre))
	if err != nil {
		return v.Value
	}
	result, err := expr.Run(program, pre)
	if err != nil {
		return v.Value
	}
	return result
}

func (v *Value) ToValueExprInfluxdb(template string) string {
	switch val := v.ToValueExpr(template).(type) {
	case int, int64, int32:
		return fmt.Sprintf("%di", val) // InfluxDB整数需要加 i
	case float64, float32:
		return fmt.Sprintf("%f", val)
	case string:
		return fmt.Sprintf("\"%s\"", val) // 字符串需要用引号包起来
	case bool:
		return fmt.Sprintf("%t", val)
	default:
		return fmt.Sprintf("\"%v\"", val) // fallback，防止 panic
	}
}
