package testing

import (
	"context"
	"net/http"
)

// MockHTTPClient implements the services.HTTPClient interface for testing actual service implementations
// This allows testing the real service logic while mocking only the HTTP layer
type MockHTTPClient struct {
	DoParseFunc    func(ctx context.Context, method, path string, body interface{}, output interface{}) error
	DoUnparsedFunc func(ctx context.Context, method, path string, body interface{}) (*http.Response, error)
}

func (m *MockHTTPClient) DoParsed(ctx context.Context, method, path string, body interface{}, output interface{}) error {
	if m.DoParseFunc != nil {
		return m.DoParseFunc(ctx, method, path, body, output)
	}
	return nil
}

func (m *MockHTTPClient) DoUnparsed(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	if m.DoUnparsedFunc != nil {
		return m.DoUnparsedFunc(ctx, method, path, body)
	}
	return &http.Response{StatusCode: 200}, nil
}
