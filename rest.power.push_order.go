package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/gorequest"
)

type RestPowerPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestPowerPushOrderResult struct {
	Result RestPowerPushOrderResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func newRestPowerPushOrderResult(result RestPowerPushOrderResponse, body []byte, http gorequest.Response, err error) *RestPowerPushOrderResult {
	return &RestPowerPushOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerPushOrder 电费充值API
// https://open.wikeyun.cn/#/apiDocument/9/document/311
func (c *Client) RestPowerPushOrder(notMustParams ...gorequest.Params) *RestPowerPushOrderResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.StoreId) // 店铺ID
	// 请求
	request, err := c.request(apiUrl+"/rest/Power/pushOrder", params)
	// 定义
	var response RestPowerPushOrderResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestPowerPushOrderResult(response, request.ResponseBody, request, err)
}
