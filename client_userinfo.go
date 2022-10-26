package wgo

import (
	"context"
	"errors"
)

// GetUserInfoRequest is the request for the GetUserInfo method
type GetUserInfoRequest struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang"`
}

// GetUserInfoResponse is the response to a GetUserInfo request
type GetUserInfoResponse struct {
	Subscribe      int    `json:"subscribe"`
	OpenID         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	UnionID        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagIDList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

// ErrOpenIDEmpty is returned when the OpenID is empty
var ErrOpenIDEmpty = errors.New("OpenID empty")

// GetUserInfo Users Basic Information UnionID
// see: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (c *Client) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	if len(req.OpenID) == 0 {
		return nil, ErrOpenIDEmpty
	}
	resp := GetUserInfoResponse{}
	err := c.NewRequest().Get().
		RequestURI("/cgi-bin/user/info").
		Param("openid", req.OpenID).
		Param("lang", req.Lang).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPluginOpenPId ...
func (c *Client) GetPluginOpenPId() {
	// TODO implement GetPluginOpenPId
	path := "/cgi-bin/token"
	_ = path
}

// CheckEncryptedData ...
func (c *Client) CheckEncryptedData() {
	// TODO implement CheckEncryptedData
	path := "/cgi-bin/token"
	_ = path
}

// GetPaidUnionID ...
func (c *Client) GetPaidUnionID() {
	// TODO implement GetPaidUnionID
	path := "/cgi-bin/token"
	_ = path
}
