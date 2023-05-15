package wgo

import "context"

// GetPhoneNumberResponse ...
type GetPhoneNumberResponse struct {
	PhoneInfo PhoneInfo `json:"phone_info"`
}

// PhoneInfo ...
type PhoneInfo struct {
	PhoneNumber     string             `json:"phoneNumber"`
	PurePhoneNumber string             `json:"purePhoneNumber"`
	CountryCode     int                `json:"countryCode"`
	Watermark       PhoneInfoWatermark `json:"watermark"`
}

// PhoneInfoWatermark ...
type PhoneInfoWatermark struct {
	Timestamp int    `json:"timestamp"`
	AppID     string `json:"appid"`
}

// GetPhoneNumber ...
func (c *Client) GetPhoneNumber(ctx context.Context, code string) (*GetPhoneNumberResponse, error) {
	resp := GetPhoneNumberResponse{}
	err := c.NewRequest().Get().
		RequestURI("/wxa/business/getuserphonenumber").
		Param("code", code).
		Do(ctx).
		Into(NewJSONValidator(&resp))
	if err != nil {
		return nil, err
	}
	return &resp, err
}
