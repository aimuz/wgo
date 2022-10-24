package wgo

import "context"

type GetPhoneNumberResponse struct {
	PhoneInfo PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string             `json:"phoneNumber"`
	PurePhoneNumber string             `json:"purePhoneNumber"`
	CountryCode     int                `json:"countryCode"`
	Watermark       PhoneInfoWatermark `json:"watermark"`
}

type PhoneInfoWatermark struct {
	Timestamp int    `json:"timestamp"`
	Appid     string `json:"appid"`
}

func (c *Client) GetPhoneNumber(ctx context.Context, code string) (*GetPhoneNumberResponse, error) {
	resp := GetPhoneNumberResponse{}
	err := c.NewRequest().Get().
		RequestURI("/wxa/business/getuserphonenumber").
		Param("code", code).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
