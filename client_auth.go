package wgo

import (
	"context"
	"errors"
)

// GrantType ...
type GrantType string

const (
	// GrantTypeClientCredential ...
	GrantTypeClientCredential GrantType = "client_credential"
	// GrantTypeAuthorizationCode ...
	GrantTypeAuthorizationCode GrantType = "authorization_code"
)

// ErrInvalidGrantType is returned when an grant type is invalid
var ErrInvalidGrantType = errors.New("invalid grant type")

// IsValid checks GetAccessToken grant_type
func (gt GrantType) IsValid() error {
	if gt == GrantTypeClientCredential {
		return nil
	}
	return ErrInvalidGrantType
}

// AccessTokenResponse is the response to an GetAccessToken
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessTokenWithGrantType Obtains the globally unique credential (`access_token`) for calling backend APIs on Mini Programs.
//
//	The access_token is required for calls to most backend APIs. Developers need to save it properly.
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) GetAccessTokenWithGrantType(ctx context.Context, gt GrantType) (*AccessTokenResponse, error) {
	if err := gt.IsValid(); err != nil {
		return nil, err
	}

	resp := AccessTokenResponse{}
	err := c.NewRequest().Get().
		RequestURI("/cgi-bin/token").
		Param("appid", c.config.appid).
		Param("secret", c.config.secret).
		Param("grant_type", string(gt)).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// GetAccessToken Obtains the globally unique credential (`access_token`) for calling backend APIs on Mini Programs.
//
//	The access_token is required for calls to most backend APIs. Developers need to save it properly.
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) GetAccessToken(ctx context.Context) (*AccessTokenResponse, error) {
	return c.GetAccessTokenWithGrantType(ctx, GrantTypeClientCredential)
}

// SignType ...
type SignType string

const (
	// SignTypeHMACSHA1 ...
	SignTypeHMACSHA1 SignType = "hmac_sha1"
	// SignTypeHMACSHA256 ...
	SignTypeHMACSHA256 SignType = "hmac_sha256"
)

// ErrInvalidSignType is returned when sign type is invalid
var ErrInvalidSignType = errors.New("invalid sign type")

// IsValid checks CheckSessionKey sig_method
func (st SignType) IsValid() error {
	ok := st == SignTypeHMACSHA1 || st == SignTypeHMACSHA256
	if ok {
		return nil
	}
	return ErrInvalidSignType
}

// CheckSessionKey Verifies whether the `session_key` of the login status saved on the server is legal.
//
//	To ensure the confidentiality of the `session_key`, the API does not transmit the session_key in clear text.
//	Rather, it checks the login status signature.
//	see: https://developers.weixin.qq.com/minigame/en/dev/api-backend/open-api/login/auth.checkSessionKey.html
func (c *Client) CheckSessionKey(ctx context.Context, openID, signature string, sigMethod SignType) error {
	if err := sigMethod.IsValid(); err != nil {
		return err
	}

	resp := struct{}{}
	err := c.NewRequest().Get().
		RequestURI("/wxa/checksession").
		Param("openid", openID).
		Param("signature", signature).
		Param("sig_method", string(sigMethod)).
		Do(ctx).
		Into(&resp)
	return err
}

// CheckSessionKeyWithHMACSHA256 Verifies whether the `session_key` of the login status saved on the server is legal.
//
//	To ensure the confidentiality of the `session_key`, the API does not transmit the session_key in clear text.
//	Rather, it checks the login status signature.
//	see: https://developers.weixin.qq.com/minigame/en/dev/api-backend/open-api/login/auth.checkSessionKey.html
func (c *Client) CheckSessionKeyWithHMACSHA256(ctx context.Context, openID, signature string) error {
	return c.CheckSessionKey(ctx, openID, signature, SignTypeHMACSHA256)
}

// Code2SessionResponse is the response to an Code2Session
type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

// Code2Session Verifies the login credential.
//
//	The temporary login credential code is obtained via the wx.login API and passed into WeChat DevTools,
//	which calls this API to complete the login procedure. For more usages, see Mini Program Login.
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/login/auth.code2Session.html
func (c *Client) Code2Session(ctx context.Context, code string) (*Code2SessionResponse, error) {
	resp := Code2SessionResponse{}
	err := c.NewRequest().Get().
		RequestURI("/sns/jscode2session").
		Param("appid", c.config.appid).
		Param("secret", c.config.secret).
		Param("js_code", code).
		Param("grant_type", string(GrantTypeAuthorizationCode)).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
