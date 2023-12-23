package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestOilCancelResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

type RestOilCancelResult struct {
	Result RestOilCancelResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newRestOilCancelResult(result RestOilCancelResponse, body []byte, http gorequest.Response) *RestOilCancelResult {
	return &RestOilCancelResult{Result: result, Body: body, Http: http}
}

// RestOilCancel 油卡订单取消
// order_number = 取消的单号，多个用英文逗号隔开
// https://open.wikeyun.cn/#/apiDocument/9/document/369
func (c *Client) RestOilCancel(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilCancelResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Oil/cancel", params)
	if err != nil {
		return newRestOilCancelResult(RestOilCancelResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestOilCancelResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestOilCancelResult(response, request.ResponseBody, request), err
}
