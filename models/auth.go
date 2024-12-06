package models

import (
	"time"
)

type Credentials struct {
	Username  string  `json:"username,omitempty"`
	Password  string  `json:"password,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	Cookie    bool    `json:"cookie,omitempty"`
	CsrfToken bool    `json:"csrfToken,omitempty"`
}

type AuthorizationToken struct {
	// the date and time when the response was generated
	ResponseTime *time.Time `json:"responseTime,omitempty"`
	// the result of the request
	Status string `json:"status"`
	// the major and minor version of the API
	ApiVersion string `json:"apiVersion"`
	// whether the requested API is deprecated, default false
	Deprecated *bool `json:"deprecated,omitempty"`
	// authorization bearer token
	Data string `json:"data"`
}
