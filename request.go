package wikeyun

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	sign := c.sign(params)

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, c.GetAppKey(), sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign))

	// 设置FORM格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.log.status {
		go c.log.client.Middleware(ctx, request, Version)
	}

	return request, err
}
