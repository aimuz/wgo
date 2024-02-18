package wgo

import (
	"net/http"
	"net/url"

	"github.com/aimuz/wgo/rest"
)

const defaultBaseURL = "https://api.weixin.qq.com"

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

	hc *http.Client
}

// NewClient ...
func NewClient(appid string, opts ...ClientOption) *Client {
	base, _ := url.Parse(defaultBaseURL)
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

// NewClientWithSecret ...
func NewClientWithSecret(appid string, secret string, opts ...ClientOption) *Client {
	return NewClient(appid, append(opts, WithSecret(secret))...)
}

func WithSecret(secret string) ClientOption {
	return func(c *clientConfig) {
		c.secret = secret
	}
}

// WithToken sets the token for the WGo client
func WithToken(token string) ClientOption {
	return func(c *clientConfig) {
		transport := c.hc.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		c.hc.Transport = &RoundTripper{
			next: transport,
			TokenSource: tokenSourceFunc(func() (*Token, error) {
				return &Token{AccessToken: token}, nil
			}),
		}
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
		transport := c.hc.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		c.hc.Transport = &RoundTripper{
			next:        transport,
			TokenSource: tokenSource,
		}
	}
}

// NewRequest creates a new request
func (c *Client) NewRequest() *rest.Request { return rest.NewRequest(c.base, c.config.hc) }
