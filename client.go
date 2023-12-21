package wikeyun

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	StoreId   int64  // 店铺ID
	AppKey    int64  // key
	AppSecret string // secret
	CurrentIp string // 当前ip
}

// Client 实例
type Client struct {
	config struct {
		clientIp  string // 当前Ip
		storeId   int64  // 店铺ID
		appKey    int64  // key
		appSecret string // secret
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.clientIp = config.CurrentIp

	c.config.storeId = config.StoreId
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	return c, nil
}
