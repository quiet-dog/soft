package common

import (
	"encoding/json"
	"strconv"
	"time"
)

type TemplateEnv struct {
	Value      interface{} `json:"value" expr:"value"`
	CreateTime time.Time   `json:"createTime" expr:"createTime"`
	Type       string      `json:"type" expr:"type"`
}

func (temp *TemplateEnv) PrepareExprEnv() map[string]interface{} {
	var value interface{}

	switch v := temp.Value.(type) {
	case json.Number:
		f, err := v.Float64()
		if err == nil {
			value = f
		} else {
			value = 0.0
		}
	case string:
		// 可能是可解析的 float 字符串，也可能是普通字符串
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			value = f
		} else {
			value = v
		}
	case float64, int, int64:
		// 保留数值类型
		value = v
	case map[string]interface{}:
		// 对象保持原样，表达式中可访问其字段
		value = v
	case nil:
		value = nil
	default:
		// fallback: 如果是 struct，也允许透传
		value = v
	}

	return map[string]interface{}{
		"value":      value,
		"type":       temp.Type,
		"createTime": temp.CreateTime,
	}
}
