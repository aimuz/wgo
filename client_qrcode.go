package wgo

import (
	"context"
	"strconv"
)

// GetQRCodeRequest is a request to get the GetQRCode
type GetQRCodeRequest struct {
	Path      string          `json:"path"`
	Width     int             `json:"width"`
	AutoColor bool            `json:"auto_color"`
	LineColor QRCodeLineColor `json:"line_color"`
	IsHyaline bool            `json:"is_hyaline"`
}

// QRCodeLineColor It takes effect when auto_color is "false".
//
//	The color is specified via RGB, for example, {"r":"xxx","g":"xxx","b":"xxx"}, and indicated by a decimal value.
type QRCodeLineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// GetQRCodeResponse is the response of GetQRCode
type GetQRCodeResponse struct {
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

// GetQRCode Obtains Mini Program QR code in scenarios where a small number of codes are required.
//
//	Mini Program codes generated via this API are permanently effective, but the quantity is limited. For details,
//	see Obtaining the QR Code: https://developers.weixin.qq.com/miniprogram/en/dev/framework/open-ability/qr-code.html
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.get.html
func (c *Client) GetQRCode(ctx context.Context, req *GetQRCodeRequest) (*GetQRCodeResponse, error) {
	resp := GetQRCodeResponse{}
	err := c.NewRequest().Post().
		RequestURI("/wxa/getwxacode").
		Body(req).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUnlimitedQRCodeRequest is a request to get the GetUnlimitedQRCode
// see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html#Request-Parameters
// chinese see: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0
type GetUnlimitedQRCodeRequest struct {
	// Scene is required
	// 	A maximum of 32 visible characters including numbers, uppercase and lowercase English letters,
	//	and some special characters !#$&'()*+,/:;=?@-._~ are supported. Other characters need to be encoded
	//	as valid characters. (Because % is not supported, Chinese characters cannot be processed via urlencode,
	//	so you must use other encoding schemes instead.)
	Scene string `json:"scene"`

	// Page This must be an existing page of a published Mini Program (otherwise, an error is reported),
	//	for example, pages/index/index. Do not prefix a root path with /. This field cannot contain parameters
	//	(parameters can be contained in the scene field).
	//	If this field is left empty, the user is redirected to the homepage by default.
	Page string `json:"page"`

	// Width The width of the QR code (in px). The minimum value is 280 px and the maximum value is 1280 px.
	Width int `json:"width"`

	// AutoColor The automatically configured line color. If the color is still black,
	//	it indicates that the setting of the dominant tone is not recommended. Its value is "false" by default.
	AutoColor bool `json:"auto_color"`

	// LineColor It takes effect when auto_color is "false". The color is specified via RGB,
	//	for example, {"r":"xxx","g":"xxx","b":"xxx"}, and indicated by a decimal value.
	LineColor QRCodeLineColor `json:"line_color"`

	// IsHyaline Specifies whether a transparent background is required.
	//	"true" indicates a transparent background is generated for the Mini Program.
	IsHyaline bool `json:"is_hyaline"`
}

// GetUnlimitedQRCodeResponse is the response of GetUnlimitedQRCode
//
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html#Return-Values
//	chinese see: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html#%E8%BF%94%E5%9B%9E%E5%80%BC
type GetUnlimitedQRCodeResponse struct {
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

// GetUnlimitedQRCode Obtains the Mini Program code in scenarios where a large number of codes are required.
//
//	Mini Program codes generated via this API are permanently effective.
//	There is no upper limit on the number of codes. For more information on how to use it,
//	see Obtaining the QR Code: https://developers.weixin.qq.com/miniprogram/en/dev/framework/open-ability/qr-code.html
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
//	chinese see: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) GetUnlimitedQRCode(ctx context.Context, req *GetUnlimitedQRCodeRequest) (*GetUnlimitedQRCodeResponse, error) {
	resp := GetUnlimitedQRCodeResponse{}
	err := c.NewRequest().Post().
		Body(req).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateQRCodeResponse is the response of CreateQRCode
type CreateQRCodeResponse struct {
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

// CreateQRCode Obtains Mini Program QR code in scenarios where a small number of codes are required.
//
//	Note that Mini Program codes generated via this API never expire, but the quantity is limited. For details,
//	see Obtaining the QR Code https://developers.weixin.qq.com/miniprogram/en/dev/framework/open-ability/qr-code.html.
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (c *Client) CreateQRCode(ctx context.Context, path string, width int) (*CreateQRCodeResponse, error) {
	resp := CreateQRCodeResponse{}
	err := c.NewRequest().Post().
		RequestURI("/wxa/getwxacode").
		Param("path", path).
		Param("width", strconv.Itoa(width)).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
