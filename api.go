package go_necos

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Client is structure used to make calls to api easier
type Client struct {
	http.Client
	DefaultQuery url.Values
}

// Get is wrapper for GET http method
func (c *Client) Get(path string, query url.Values, result interface{}) error {
	return c.CallAPI(http.MethodGet, path, query, result)
}

// GetWithContext is wrapper for GET http method
func (c *Client) GetWithContext(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodGet, path, query, result)
}

// Post is wrapper for POST http method
func (c *Client) Post(path string, query url.Values, result interface{}) error {
	return c.CallAPI(http.MethodPost, path, query, result)
}

// PostWithContext is wrapper for POST http method
func (c *Client) PostWithContext(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPost, path, query, result)
}

// applyDefaultQuery adds parameters from DefaultQuery that aren't present in query to query
func (c *Client) applyDefaultQuery(query *url.Values) {
	for k, v := range c.DefaultQuery {
		if _, ok := (*query)[k]; ok {
			continue
		}

		(*query)[k] = v
	}
}

// CallAPI is plain api call, which uses context.Background()
func (c *Client) CallAPI(method, path string, query url.Values, result interface{}) error {
	return c.CallAPIWithContext(context.Background(), method, path, query, result)
}

// CallAPIWithContext is plain api call
//
// At first it builds query suffix from provided url.Values and DefaultQuery, makes request, and marshals response data
func (c *Client) CallAPIWithContext(ctx context.Context, method, path string, query url.Values, result interface{}) error {
	c.applyDefaultQuery(&query)
	queryEnc := query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, path+queryEnc, http.NoBody)
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
	return json.Unmarshal(body, result)
}
