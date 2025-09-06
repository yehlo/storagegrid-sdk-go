package services

import (
	"context"
	"net/http"
)

// HTTPClient interface defines the contract for HTTP operations
// This allows for easy mocking and testing
type HTTPClient interface {
	DoParsed(ctx context.Context, method, path string, body interface{}, output interface{}) error
	DoUnparsed(ctx context.Context, method, path string, body interface{}) (*http.Response, error)
}
