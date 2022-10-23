package wgo

import "context"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GrantType string

const (
	GrantTypeClientCredential  GrantType = "client_credential"
	GrantTypeAuthorizationCode GrantType = "authorization_code"
)

// IsValid checks GetAccessToken grant_type
func (gt GrantType) IsValid() bool {
	return gt == GrantTypeClientCredential
}

func (c *Client) GetAccessToken(ctx context.Context, gt GrantType) (*AccessToken, error) {
	if gt.IsValid() {
		// TODO
	}

	ret := AccessToken{}
	err := c.NewRequest().Get().
		RequestURI("/cgi-bin/token").
		Param("appid", c.config.appid).
		Param("secret", c.config.secret).
		Param("grant_type", string(gt)).
		Do(ctx).
		Into(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

type CheckSessionKey struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

type SignType string

const (
	SignTypeHMACSHA1   SignType = "hmac_sha1"
	SignTypeHMACSHA256 SignType = "hmac_sha256"
)

// IsValid checks GetAccessToken grant_type
func (st SignType) IsValid() bool {
	return st == SignTypeHMACSHA1 || st == SignTypeHMACSHA256
}

// CheckSessionKey ...
func (c *Client) CheckSessionKey(ctx context.Context, openID, signature string, sigMethod SignType) error {
	if sigMethod.IsValid() {
		// TODO
	}

	ret := struct{}{}
	err := c.NewRequest().Get().
		RequestURI("/wxa/checksession").
		Param("openid", openID).
		Param("signature", signature).
		Param("sig_method", string(sigMethod)).
		Do(ctx).
		Into(&ret)
	return err
}

type Code2Session struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

// Code2Session ...
func (c *Client) Code2Session(ctx context.Context, code string) (*Code2Session, error) {
	ret := Code2Session{}
	err := c.NewRequest().Get().
		RequestURI("/sns/jscode2session").
		Param("appid", c.config.appid).
		Param("secret", c.config.secret).
		Param("js_code", code).
		Param("grant_type", string(GrantTypeAuthorizationCode)).
		Do(ctx).
		Into(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}
