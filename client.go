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
	hc     *http.Client
}

type ClientOption func(*clientConfig)

type clientConfig struct {
	appid  string
	secret string
	token  string
}

func NewClient(opts ...ClientOption) *Client {
	// TODO implement NewClient
	return &Client{}
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

func (c *Client) NewRequest() *rest.Request {
	return rest.NewRequest(c.base, c.hc)
}
