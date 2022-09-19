package wikeyun

import "go.dtapp.net/golog"

func (c *Client) GetStoreId() int {
	return c.config.storeId
}

func (c *Client) GetAppKey() int {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetClientIp() string {
	return c.config.clientIp
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.client
}
