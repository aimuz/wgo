package wgo

import "context"

type GetQRCodeRequest struct {
	Path      string             `json:"path"`
	Width     int                `json:"width"`
	AutoColor bool               `json:"auto_color"`
	LineColor GetQRCodeLineColor `json:"line_color"`
	IsHyaline bool               `json:"is_hyaline"`
}

type GetQRCodeLineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type GetQRCodeResponse struct {
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

func (c *Client) GetQRCode(ctx context.Context, req *GetQRCodeRequest) (*GetQRCodeResponse, error) {
	resp := GetQRCodeResponse{}
	err := c.NewRequest().Post().
		Body(req).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUnlimitedQRCode() {
	// TODO implement GetUnlimitedQRCode
	path := "/cgi-bin/token"
	_ = path
}

func (c *Client) CreateQRCode() {
	// TODO implement CreateQRCode
	path := "/cgi-bin/token"
	_ = path
}
