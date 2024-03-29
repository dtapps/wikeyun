package wikeyun

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 签名
	sign := c.sign(param)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, c.config.appKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign))

	// 设置FORM格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
