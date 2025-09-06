package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

var (
	// append this list after implementing saml auth for sdk
	implementedAuthorizeEndpoints = []string{"/authorize"}
)

type Client struct {
	baseURL      *url.URL
	httpClient   *http.Client
	credentials  *models.Credentials
	skipSSL      bool
	token        string
	tokenExpires time.Time
	mu           sync.Mutex
}

type ClientOption func(*Client)

func WithCredentials(creds *models.Credentials) ClientOption {
	return func(c *Client) {
		c.credentials = creds
	}
}

func WithSkipSSL() ClientOption {
	return func(c *Client) {
		c.skipSSL = true
	}
}

func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		// Required because url.Parse returns an empty string for the hostname if there was no schema
		if !strings.HasPrefix(endpoint, "https://") && !strings.HasPrefix(endpoint, "http://") {
			endpoint = "https://" + endpoint
		}

		if !strings.HasSuffix(endpoint, "/") {
			endpoint = endpoint + "/"
		}

		c.baseURL, _ = url.Parse(endpoint)
	}
}

func newClient(options ...ClientOption) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	// err if no endpoint is set
	if c.baseURL == nil {
		return nil, fmt.Errorf("no endpoint set")
	}

	var transCfg = &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.skipSSL}, // #nosec G402
	}

	c.httpClient.Transport = transCfg

	return c, nil
}

func (c *Client) authorize(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != "" && time.Now().Before(c.tokenExpires) {
		return nil
	}

	resp, err := c.DoUnparsed(ctx, "POST", "/authorize", c.credentials)
	if err != nil {
		return err
	}

	token := &models.AuthorizationToken{}
	err = c.parseResponse(resp, token)
	if err != nil {
		return err
	}

	c.token = token.Data
	timeFormat := "Mon, 02 Jan 2006 15:04:05 GMT"
	c.tokenExpires, err = time.Parse(timeFormat, resp.Header.Get("Expires"))
	if err != nil {
		return fmt.Errorf("failed to parse token expiration: %w", err)
	}

	return nil
}

func (c *Client) newRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	var err error
	var reqBody []byte
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	// Create a new request
	surl := c.baseURL.String() + path
	req, err := http.NewRequest(method, surl, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Ensure Content-Type is set for requests with a body.
	if body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// if the path is not an authorize endpoint, set the token in the header
	if !slices.Contains(implementedAuthorizeEndpoints, path) {
		// If there's no token yet or it's expired, authorize
		if c.token == "" || time.Now().After(c.tokenExpires) {
			err = c.authorize(ctx)
			if err != nil {
				return nil, err
			}
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}

func (c *Client) parseResponse(resp *http.Response, resourceType interface{}) error {
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	err = json.Unmarshal(respBody, &resourceType)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

func (c *Client) DoUnparsed(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	resp, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, (fmt.Errorf("API error: %s (code: %d)", resp.Status, resp.StatusCode))
	}

	return resp, nil
}

func (c *Client) DoParsed(ctx context.Context, method string, path string, body interface{}, output interface{}) error {
	resp, err := c.DoUnparsed(ctx, method, path, body)
	if err != nil {
		return err
	}

	// If there's no output, we're done
	if output != nil {
		err = c.parseResponse(resp, output)
		if err != nil {
			return err
		}
	}

	return nil
}
