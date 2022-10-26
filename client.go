package wgo

import (
	"net/http"
	"net/url"

	"github.com/aimuz/wgo/rest"
)

const apiBase = "https://api.weixin.qq.com"

// Client is a WGo client.
type Client struct {
	config clientConfig
	base   *url.URL
}

// ClientOption is a client option for the WGo Client config.
type ClientOption func(*clientConfig)

type clientConfig struct {
	appid  string
	secret string
	token  string

	hc *http.Client
}

// NewClient ...
func NewClient(opts ...ClientOption) *Client {
	base, _ := url.Parse(apiBase)
	c := &Client{
		base: base,
		config: clientConfig{
			hc: http.DefaultClient,
		},
	}
	for _, opt := range opts {
		opt(&c.config)
	}
	return c
}

// WithAPPIDAndSecret sets the appid and secret for the WGo client
func WithAPPIDAndSecret(appid string, secret string) ClientOption {
	return func(c *clientConfig) {
		c.appid = appid
		c.secret = secret
	}
}

// WithToken sets the token for the WGo client
func WithToken(token string) ClientOption {
	return func(c *clientConfig) {
		c.token = token
	}
}

// WithHTTPClient sets the http client for the WGo client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.hc = client
	}
}

// WithTokenSource sets the token source for the WGo client
func WithTokenSource(tokenSource TokenSource) ClientOption {
	return func(c *clientConfig) {
		c.hc = &http.Client{
			Transport: &RoundTripper{
				next:        http.DefaultTransport,
				TokenSource: tokenSource,
			},
		}
	}
}

// NewRequest creates a new request
func (c *Client) NewRequest() *rest.Request {
	return rest.NewRequest(c.base, c.config.hc)
}
