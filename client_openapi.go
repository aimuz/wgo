package wgo

import "context"

func (c *Client) ClearQuota(ctx context.Context) error {
	resp := struct{}{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota").
		Param("appid", c.config.appid).
		Do(ctx).
		Into(&resp)
	return err
}

type GetAPIQuotaResponse struct {
	Quota GetAPIQuota `json:"quota"`
}

type GetAPIQuota struct {
	DailyLimit int `json:"daily_limit"`
	Used       int `json:"used"`
	Remain     int `json:"remain"`
}

func (c *Client) GetAPIQuota(ctx context.Context, cgiPath string) (*GetAPIQuotaResponse, error) {
	resp := GetAPIQuotaResponse{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota").
		Param("cgi_path", cgiPath).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

type GetRidInfoResponse struct {
	Request GetRidInfo `json:"request"`
}

type GetRidInfo struct {
	InvokeTime   int    `json:"invoke_time"`
	CostInMs     int    `json:"cost_in_ms"`
	RequestUrl   string `json:"request_url"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
	ClientIp     string `json:"client_ip"`
}

func (c *Client) GetRidInfo(ctx context.Context, rid string) (*GetRidInfoResponse, error) {
	resp := GetRidInfoResponse{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/openapi/rid/get").
		Param("rid", rid).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) ClearQuotaByAppSecret(ctx context.Context) error {
	resp := struct{}{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota/v2").
		Param("appid", c.config.appid).
		Param("appsecret", c.config.secret).
		Do(ctx).
		Into(&resp)
	return err
}
