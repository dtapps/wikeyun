package wikeyun

import (
	"go.dtapp.net/golog"
)

// ConfigApp 配置
func (c *Client) ConfigApp(storeId, appKey int64, appSecret string) *Client {
	c.config.storeId = storeId
	c.config.appKey = appKey
	c.config.appSecret = appSecret
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}
