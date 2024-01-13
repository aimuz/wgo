package wgo

import "context"

// GetPhoneNumberResponse 获取手机号响应
type GetPhoneNumberResponse struct {
	PhoneInfo PhoneInfo `json:"phone_info"`
}

// PhoneInfo 用户手机号信息
type PhoneInfo struct {
	// PhoneNumber 用户绑定的手机号（国外手机号会有区号）
	PhoneNumber string `json:"phoneNumber"`

	// PurePhoneNumber 没有区号的手机号
	PurePhoneNumber string `json:"purePhoneNumber"`

	// CountryCode 区号
	CountryCode string `json:"countryCode"`

	// Watermark 数据水印
	Watermark PhoneInfoWatermark `json:"watermark"`
}

// PhoneInfoWatermark 数据水印
type PhoneInfoWatermark struct {
	// 用户获取手机号操作的时间戳
	Timestamp int `json:"timestamp"`
	// 小程序appid
	AppID string `json:"appid"`
}

type getPhoneNumber struct {
	Code string `json:"code"`
}

// GetPhoneNumber 获取手机号
// 该接口用于将code换取用户手机号。 说明，每个 code 只能使用一次，code 的有效期为 5min。
func (c *Client) GetPhoneNumber(ctx context.Context, code string) (*GetPhoneNumberResponse, error) {
	resp := GetPhoneNumberResponse{}
	err := c.NewRequest().Post().
		RequestURI("/wxa/business/getuserphonenumber").
		Body(getPhoneNumber{Code: code}).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, err
}
