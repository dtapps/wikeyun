package wikeyun

import (
	"errors"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/trace"
)

type ClientConfig struct {
	ApiUrl    string // 接口地址
	StoreId   int64  // 店铺ID
	AppKey    int64  // key
	AppSecret string // secret
}

// Client 实例
type Client struct {
	config struct {
		apiUrl    string // 接口地址
		storeId   int64  // 店铺ID
		appKey    int64  // key
		appSecret string // secret
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	if config.ApiUrl == "" {
		return nil, errors.New("接口地址不能为空")
	}
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.apiUrl = config.ApiUrl
	c.config.storeId = config.StoreId
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	c.trace = true
	return c, nil
}
