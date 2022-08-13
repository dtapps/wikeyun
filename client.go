package wikeyun

import (
	"context"
	"go.dtapp.net/dorm"
	"go.dtapp.net/goip"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	StoreId    int              // 店铺ID
	AppKey     int              // key
	AppSecret  string           // secret
	GormClient *dorm.GormClient // 日志数据库
	LogClient  *golog.GoLog     // 日志驱动
	LogDebug   bool             // 日志开关
}
type Client struct {
	client   *gorequest.App   // 请求客户端
	clientIp string           // Ip
	log      *golog.ApiClient // 日志服务
	config   *ConfigClient
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.client = gorequest.NewHttp()

	if c.config.GormClient.Db != nil {
		c.log, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	xip := goip.GetOutsideIp(context.Background())
	if xip != "" && xip != "0.0.0.0" {
		c.clientIp = xip
	}

	return c, nil
}
