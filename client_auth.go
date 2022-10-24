package wgo

import (
	"context"
	"errors"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GrantType string

const (
	GrantTypeClientCredential  GrantType = "client_credential"
	GrantTypeAuthorizationCode GrantType = "authorization_code"
)

var ErrInvalidGrantType = errors.New("invalid grant type")

// IsValid checks GetAccessToken grant_type
func (gt GrantType) IsValid() error {
	if gt == GrantTypeClientCredential {
		return nil
	}
	return ErrInvalidGrantType
}

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

func (c *Client) GetAccessToken(ctx context.Context) (*AccessTokenResponse, error) {
	return c.GetAccessTokenWithGrantType(ctx, GrantTypeClientCredential)
}

type SignType string

const (
	SignTypeHMACSHA1   SignType = "hmac_sha1"
	SignTypeHMACSHA256 SignType = "hmac_sha256"
)

var ErrInvalidSignType = errors.New("invalid sign type")

// IsValid checks GetAccessToken grant_type
func (st SignType) IsValid() error {
	ok := st == SignTypeHMACSHA1 || st == SignTypeHMACSHA256
	if ok {
		return nil
	}
	return ErrInvalidSignType
}

// CheckSessionKey ...
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

func (c *Client) CheckSessionKeyWithHMACSHA256(ctx context.Context, openID, signature string) error {
	return c.CheckSessionKey(ctx, openID, signature, SignTypeHMACSHA256)
}

type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

// Code2Session ...
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
