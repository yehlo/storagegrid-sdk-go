package models

import "time"

type Response[T any] struct {
	ResponseTime *time.Time `json:"responseTime,omitempty"`
	Status       string     `json:"status"`
	APIVersion   string     `json:"apiVersion"`
	Deprecated   *bool      `json:"deprecated,omitempty"`
	Data         T          `json:"data,omitempty"`
}
