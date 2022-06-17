package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/gorequest"
)

type RestPowerCancelResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RestPowerCancelResult struct {
	Result RestPowerCancelResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func NewRestPowerCancelResult(result RestPowerCancelResponse, body []byte, http gorequest.Response, err error) *RestPowerCancelResult {
	return &RestPowerCancelResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerCancel 电费订单取消
// https://open.wikeyun.cn/#/apiDocument/9/document/323
func (c *Client) RestPowerCancel(orderNumber string) *RestPowerCancelResult {
	// 参数
	param := NewParams()
	param.Set("order_number", orderNumber) // 取消的单号，多个用英文逗号隔开
	params := c.NewParamsWith(param)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Power/cancel", params)
	// 定义
	var response RestPowerCancelResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestPowerCancelResult(response, request.ResponseBody, request, err)
}
