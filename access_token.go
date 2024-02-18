package wgo

//// ErrInvalidGrantType is returned when an grant type is invalid
//var ErrInvalidGrantType = NewError(40002, "invalid grant_type")
//
//// IsValid checks GetAccessToken grant_type
//func (gt GrantType) IsValid() error {
//	if gt == GrantTypeClientCredential {
//		return nil
//	}
//	return ErrInvalidGrantType
//}
//
//// AccessTokenResponse is the response to an GetAccessToken
//type AccessTokenResponse struct {
//	AccessToken string `json:"access_token"`
//	ExpiresIn   int    `json:"expires_in"`
//}
//
//// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）
//// see: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/access-token/auth.getAccessToken.html
//func (c *Client) GetAccessToken(ctx context.Context) (*AccessTokenResponse, error) {
//	resp := AccessTokenResponse{}
//	err := c.NewRequest().Get().
//		RequestURI("/cgi-bin/token").
//		Param("appid", c.config.appid).
//		Param("secret", c.config.secret).
//		Param("grant_type", string(GrantTypeClientCredential)).
//		Do(ctx).
//		Into(NewJSONValidator(&resp))
//	if err != nil {
//		return nil, err
//	}
//	return &resp, err
//}
//
//// GetStableAccessToken 获取稳定版接口调用凭据
//// forceRefresh 默认使用 false。
////  1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；
////  2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
////
//// see: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/access-token/auth.getStableAccessToken.html
//func (c *Client) GetStableAccessToken(ctx context.Context, forceRefresh bool) (*AccessTokenResponse, error) {
//	resp := AccessTokenResponse{}
//	err := c.NewRequest().Get().
//		RequestURI("/cgi-bin/stable_token").
//		Param("appid", c.config.appid).
//		Param("secret", c.config.secret).
//		Param("grant_type", string(GrantTypeClientCredential)).
//		Param("force_refresh", strconv.FormatBool(forceRefresh)).
//		Do(ctx).
//		Into(NewJSONValidator(&resp))
//	if err != nil {
//		return nil, err
//	}
//	return &resp, err
//}
