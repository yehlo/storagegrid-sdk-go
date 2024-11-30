package services

import (
	"context"
	"net/http"
)

type HTTPClient interface {
	DoParsed(ctx context.Context, method, path string, body interface{}, output interface{}) error
	DoUnparsed(ctx context.Context, method, path string, body interface{}) (*http.Response, error)
}
