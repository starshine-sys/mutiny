// Package httputil implements a wrapper around the standard library's HTTP client.
package httputil

import (
	"encoding/json"
	"io"
	"net/http"

	"emperror.dev/errors"
)

// Client is a http client with convenience methods for making requests.
type Client struct {
	Client *http.Client

	// optional: if set, this is added to any requests as the "User-Agent" header.
	UserAgent string
	// optional: if set, this is added to any requests as the "Authorization" header.
	Token string

	opts []RequestOption
}

// New returns a new Client.
func New() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

// WrapClient wraps an existing *http.Client.
func WrapClient(c *http.Client) *Client {
	return &Client{Client: c}
}

// AddOptions adds the given RequestOptions to the
func (c *Client) AddOptions(opts ...RequestOption) {
	c.opts = append(c.opts, opts...)
}

// Request makes a request. It is the caller's responsibility to close the response body.
func (c *Client) Request(method, url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	err = c.applyOpts(req, opts)
	if err != nil {
		return nil, errors.Wrap(err, "applying options")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}

	return c.Client.Do(req)
}

func (c *Client) applyOpts(req *http.Request, extra []RequestOption) (err error) {
	for _, opt := range c.opts {
		err = opt(req)
		if err != nil {
			return err
		}
	}
	for _, opt := range extra {
		err = opt(req)
		if err != nil {
			return err
		}
	}
	return nil
}

// RequestJSON makes a request returning a JSON body.
func (c *Client) RequestJSON(out interface{}, method, url string, opts ...RequestOption) (code int, err error) {
	resp, err := c.Request(method, url, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "request")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "reading body")
	}

	if len(b) == 0 || out == nil {
		return resp.StatusCode, nil
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return resp.StatusCode, errors.Wrap(err, "unmarshaling JSON")
	}
	return resp.StatusCode, nil
}
