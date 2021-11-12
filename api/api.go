// Package api contains types and methods for interacting with Revolt's REST API.
package api

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/starshine-sys/mutiny/revolt"
	"github.com/starshine-sys/mutiny/util/httputil"
)

// Client is an API client.
type Client struct {
	*httputil.Client
}

// New returns a new API client.
func New(token string) *Client {
	client := httputil.New()
	client.AddOptions(httputil.WithExactHeader("x-bot-token", token))
	return &Client{Client: client}
}

// RequestJSON ...
func (c *Client) RequestJSON(out interface{}, method, url string, opts ...httputil.RequestOption) (err error) {
	code, err := c.Client.RequestJSON(out, method, revolt.APIBaseURL+url, opts...)
	if err != nil {
		return err
	}

	switch code {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusForbidden:
		return ErrForbidden
	}
	return nil
}

// HTTP error codes
const (
	ErrNotFound  = errors.Sentinel("not found")
	ErrForbidden = errors.Sentinel("forbidden")
)
