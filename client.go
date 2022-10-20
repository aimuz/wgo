package wgo

const apiBase = "https://api.weixin.qq.com"

type Client struct {
	config clientConfig
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
