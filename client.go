package wgo

import (
	"net/http"
	"net/url"

	"github.com/aimuz/wgo/rest"
)

const apiBase = "https://api.weixin.qq.com"

type Client struct {
	config clientConfig
	base   *url.URL
}

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

func WithAPPIDAndSecret(appid string, secret string) ClientOption {
	return func(c *clientConfig) {
		c.appid = appid
		c.secret = secret
	}
}

func WithToken(token string) ClientOption {
	return func(c *clientConfig) {
		c.token = token
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.hc = client
	}
}

func (c *Client) NewRequest() *rest.Request {
	return rest.NewRequest(c.base, c.config.hc)
}
