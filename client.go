package wikeyun

import (
	"context"
	"go.dtapp.net/goip"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	StoreId          int                // 店铺ID
	AppKey           int                // key
	AppSecret        string             // secret
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
	CurrentIp        string             // 当前ip
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	currentIp     string         // 当前ip
	config        struct {
		clientIp  string // 当前Ip
		storeId   int    // 店铺ID
		appKey    int    // key
		appSecret string // secret
	}
	log struct {
		gorm   bool             // 日志开关
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.currentIp = config.CurrentIp

	c.config.storeId = config.StoreId
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.gorm = true
	}

	xip := goip.GetOutsideIp(context.Background())
	if xip != "" && xip != "0.0.0.0" {
		c.config.clientIp = xip
	}

	return c, nil
}
