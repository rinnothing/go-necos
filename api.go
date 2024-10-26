package go_necos

import (
	"context"
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"net/url"
)

// Client is structure used to make calls to api easier
type Client struct {
	http.Client
	DefaultQuery url.Values
	Domain       string
}

func newClient() *Client {
	return &Client{Domain: DefaultDomain}
}

// Get is a wrapper for GET http method
func (c *Client) Get(path string, query url.Values, result interface{}) error {
	return c.CallAPI(http.MethodGet, path, query, result)
}

// GetWithContext is a wrapper for GET http method
func (c *Client) GetWithContext(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodGet, path, query, result)
}

// Post is a wrapper for POST http method
func (c *Client) Post(path string, query url.Values, result interface{}) error {
	return c.CallAPI(http.MethodPost, path, query, result)
}

// PostWithContext is a wrapper for POST http method
func (c *Client) PostWithContext(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPost, path, query, result)
}

// applyDefaultQuery adds parameters from DefaultQuery that aren't present in query to query
func (c *Client) applyDefaultQuery(query url.Values) {
	for k, v := range c.DefaultQuery {
		if _, ok := query[k]; ok {
			continue
		}

		query[k] = v
	}
}

// CallAPI is a plain api call, which uses context.Background()
func (c *Client) CallAPI(method, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(context.Background(), method, path, query, result)
}

// CallAPIWithContext is a plain api call
//
// At first it builds query suffix from provided url.Values and DefaultQuery, makes request, and marshals response data
func (c *Client) CallAPIWithContext(ctx context.Context, method, path string, query url.Values, result interface{}) error {
	var queryEnc string
	if query == nil {
		queryEnc = c.DefaultQuery.Encode()
	} else if c.DefaultQuery == nil {
		queryEnc = query.Encode()
	} else {
		cloneQuery := maps.Clone(query)
		c.applyDefaultQuery(cloneQuery)
		queryEnc = cloneQuery.Encode()
	}

	reqPath := c.Domain + path
	if queryEnc != "" {
		reqPath += "?" + queryEnc
	}

	req, err := http.NewRequestWithContext(ctx, method, reqPath, http.NoBody)
	if err != nil {
		return err
	}

	response, err := c.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err = response.Body.Close(); err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}
