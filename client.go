package wikeyun

import (
	"errors"
	"go.dtapp.net/golog"
)

type ClientConfig struct {
	ApiUrl    string // 接口地址
	StoreId   int64  // 店铺ID
	AppKey    int64  // key
	AppSecret string // secret
	CurrentIp string // 当前ip
}

// Client 实例
type Client struct {
	config struct {
		apiUrl    string // 接口地址
		storeId   int64  // 店铺ID
		appKey    int64  // key
		appSecret string // secret
		clientIp  string // 当前Ip
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	if config.ApiUrl == "" {
		return nil, errors.New("接口地址不能为空")
	}
	c := &Client{}

	c.config.apiUrl = config.ApiUrl
	c.config.storeId = config.StoreId
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret
	c.config.clientIp = config.CurrentIp

	return c, nil
}
