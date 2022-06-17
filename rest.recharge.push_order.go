package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/gorequest"
)

type RestRechargePushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestRechargePushOrderResult struct {
	Result RestRechargePushOrderResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func NewRestRechargePushOrderResult(result RestRechargePushOrderResponse, body []byte, http gorequest.Response, err error) *RestRechargePushOrderResult {
	return &RestRechargePushOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargePushOrder 话费充值推送
// https://open.wikeyun.cn/#/apiDocument/9/document/298
func (c *Client) RestRechargePushOrder(notMustParams ...Params) *RestRechargePushOrderResult {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.StoreId) // 店铺ID
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Recharge/pushOrder", params)
	// 定义
	var response RestRechargePushOrderResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestRechargePushOrderResult(response, request.ResponseBody, request, err)
}
