package wgo

import "context"

// GetQRCodeRequest is a request to get the GetQRCode
type GetQRCodeRequest struct {
	Path      string             `json:"path"`
	Width     int                `json:"width"`
	AutoColor bool               `json:"auto_color"`
	LineColor GetQRCodeLineColor `json:"line_color"`
	IsHyaline bool               `json:"is_hyaline"`
}

// GetQRCodeLineColor It takes effect when auto_color is "false".
//
//	The color is specified via RGB, for example, {"r":"xxx","g":"xxx","b":"xxx"}, and indicated by a decimal value.
type GetQRCodeLineColor struct {
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
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUnlimitedQRCode Obtains the Mini Program code in scenarios where a large number of codes are required.
//
//	Mini Program codes generated via this API are permanently effective.
//	There is no upper limit on the number of codes. For more information on how to use it,
//	see Obtaining the QR Code: https://developers.weixin.qq.com/miniprogram/en/dev/framework/open-ability/qr-code.html
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
//
// TODO implement GetUnlimitedQRCode
func (c *Client) GetUnlimitedQRCode() {
	// TODO implement GetUnlimitedQRCode
	path := "/cgi-bin/token"
	_ = path
}

// CreateQRCode Obtains Mini Program QR code in scenarios where a small number of codes are required.
//
//	Note that Mini Program codes generated via this API never expire, but the quantity is limited. For details,
//	see Obtaining the QR Code https://developers.weixin.qq.com/miniprogram/en/dev/framework/open-ability/qr-code.html.
//	see: https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
//
// TODO implement CreateQRCode
func (c *Client) CreateQRCode() {
	path := "/cgi-bin/token"
	_ = path
}
