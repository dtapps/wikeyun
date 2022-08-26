package wikeyun

import (
	"context"
	"go.dtapp.net/dorm"
	"go.dtapp.net/goip"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// Client 实例
type Client struct {
	gormClient    *dorm.GormClient  // 日志数据库
	mongoClient   *dorm.MongoClient // 日志数据库
	requestClient *gorequest.App    // 请求服务
	config        struct {
		clientIp  string // 当前Ip
		storeId   int    // 店铺ID
		appKey    int    // key
		appSecret string // secret
		debug     bool   // 日志开关
	}
	log struct {
		gorm        bool             // 日志
		gormClient  *golog.ApiClient // 日志服务
		mongo       bool             // 日志
		mongoClient *golog.ApiClient // 日志服务
	}
}

type gormClientFun func() *dorm.GormClient
type mongoClientFun func() (client *dorm.MongoClient, databaseName string)

// NewClient 创建实例化
// storeId 店铺ID
// appKey key
// appSecret secret
func NewClient(storeId, appKey int, appSecret string, gormClientFun gormClientFun, mongoClientFun mongoClientFun, debug bool) (*Client, error) {

	var err error
	c := &Client{}

	c.config.storeId = storeId
	c.config.appKey = appKey
	c.config.appSecret = appSecret

	c.requestClient = gorequest.NewHttp()

	gormClient := gormClientFun()
	if c.gormClient.Db != nil {
		c.log.gormClient, err = golog.NewApiGormClient(func() (client *dorm.GormClient, tableName string) {
			return gormClient, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
	}
	c.gormClient = gormClient

	mongoClient, databaseName := mongoClientFun()
	if c.mongoClient.Db != nil {
		c.log.mongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
	}
	c.mongoClient = mongoClient

	xip := goip.GetOutsideIp(context.Background())
	if xip != "" && xip != "0.0.0.0" {
		c.config.clientIp = xip
	}

	return c, nil
}
