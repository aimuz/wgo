package wgo

import (
	"context"
	"errors"
)

// GrantType 授权类型
type GrantType string

const (
	// GrantTypeClientCredential 获取 AccessToken 使用该值
	GrantTypeClientCredential GrantType = "client_credential"

	// GrantTypeAuthorizationCode 小程序登录使用该类型
	GrantTypeAuthorizationCode GrantType = "authorization_code"
)

// SignType ...
type SignType string

const (
	// SignTypeHMACSHA1 ...
	// Deprecated: Suggested Use [SignTypeHMACSHA256]
	SignTypeHMACSHA1 SignType = "hmac_sha1"

	// SignTypeHMACSHA256 ...
	SignTypeHMACSHA256 SignType = "hmac_sha256"
)

// ErrInvalidSignType is returned when sign type is invalid
var ErrInvalidSignType = errors.New("wgo: invalid sign type")

// IsValid checks CheckSessionKey sig_method
func (st SignType) IsValid() error {
	ok := st == SignTypeHMACSHA1 || st == SignTypeHMACSHA256
	if ok {
		return nil
	}
	return ErrInvalidSignType
}

// CheckSessionKey 校验服务器所保存的登录态 session_key 是否合法。
// 为了保持 session_key 私密性，接口不明文传输 session_key，而是通过校验登录态签名完成。
// see: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.checkSessionKey.html
func (c *Client) CheckSessionKey(ctx context.Context, openID, signature string) error {
	return c.CheckSessionKeyWithSignType(ctx, openID, signature, SignTypeHMACSHA256)
}

// CheckSessionKeyWithSignType 校验服务器所保存的登录态 session_key 是否合法。
// 为了保持 session_key 私密性，接口不明文传输 session_key，而是通过校验登录态签名完成。
// see: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.checkSessionKey.html
func (c *Client) CheckSessionKeyWithSignType(ctx context.Context, openID, signature string, st SignType) error {
	if err := st.IsValid(); err != nil {
		return err
	}

	err := c.NewRequest().Get().
		RequestURI("/wxa/checksession").
		Param("openid", openID).
		Param("signature", signature).
		Param("sig_method", string(st)).
		Do(ctx).
		Into(NewJSONValidator(nil))
	return err
}

// Code2SessionResponse is the response to an Code2Session
type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

// Code2Session 登录凭证校验。
// 通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
// see: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.code2Session.html
func (c *Client) Code2Session(ctx context.Context, code string) (*Code2SessionResponse, error) {
	resp := Code2SessionResponse{}
	err := c.NewRequest().Get().
		RequestURI("/sns/jscode2session").
		Param("appid", c.config.appid).
		//Param("secret", c.config.secret).
		Param("js_code", code).
		Param("grant_type", string(GrantTypeAuthorizationCode)).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, err
}
