package wikeyun

// ConfigApp 配置
func (c *Client) ConfigApp(storeId, appKey int, appSecret string) *Client {
	c.config.storeId = storeId
	c.config.appKey = appKey
	c.config.appSecret = appSecret
	return c
}
