package common

import (
	"time"
)

type Value struct {
	Value interface{} `json:"value"`
}
type TemplateEnv struct {
	Value      Value     `json:"value" expr:"value"`
	CreateTime time.Time `json:"createTime" expr:"createTime"`
	Type       string    `json:"type" expr:"type"`
}
