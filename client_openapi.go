package wgo

import "context"

func (c *Client) ClearQuota(ctx context.Context) error {
	ret := struct{}{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota").
		Param("appid", c.config.appid).
		Do(ctx).
		Into(&ret)
	return err
}

type GetAPIQuota struct {
	Quota struct {
		DailyLimit int `json:"daily_limit"`
		Used       int `json:"used"`
		Remain     int `json:"remain"`
	} `json:"quota"`
}

func (c *Client) GetAPIQuota(ctx context.Context, cgiPath string) (*GetAPIQuota, error) {
	ret := GetAPIQuota{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota").
		Param("cgi_path", cgiPath).
		Do(ctx).
		Into(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

type GetRidInfo struct {
	Request GetRidInfoRequest `json:"request"`
}

type GetRidInfoRequest struct {
	InvokeTime   int    `json:"invoke_time"`
	CostInMs     int    `json:"cost_in_ms"`
	RequestUrl   string `json:"request_url"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
	ClientIp     string `json:"client_ip"`
}

func (c *Client) GetRidInfo(ctx context.Context, rid string) (*GetRidInfo, error) {
	ret := GetRidInfo{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/openapi/rid/get").
		Param("rid", rid).
		Do(ctx).
		Into(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *Client) ClearQuotaByAppSecret(ctx context.Context) error {
	ret := struct{}{}
	err := c.NewRequest().Post().
		RequestURI("/cgi-bin/clear_quota/v2").
		Param("appid", c.config.appid).
		Param("appsecret", c.config.secret).
		Do(ctx).
		Into(&ret)
	return err

}
