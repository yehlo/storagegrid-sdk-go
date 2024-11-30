package models

import "time"

type Response struct {
	ResponseTime *time.Time  `json:"responseTime,omitempty"`
	Status       string      `json:"status"`
	ApiVersion   string      `json:"apiVersion"`
	Deprecated   *bool       `json:"deprecated,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
