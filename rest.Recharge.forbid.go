package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestRechargeForbidResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Time string   `json:"time"`
	Data struct{} `json:"data"`
}

type RestRechargeForbidResult struct {
	Result RestRechargeForbidResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestRechargeForbidResult(result RestRechargeForbidResponse, body []byte, http gorequest.Response) *RestRechargeForbidResult {
	return &RestRechargeForbidResult{Result: result, Body: body, Http: http}
}

// RestRechargeForbid 禁启用非API渠道下单
// status = 1 禁用 0启用
// https://open.wikeyun.cn/#/apiDocument/9/document/445
func (c *Client) RestRechargeForbid(ctx context.Context, status int64, notMustParams ...gorequest.Params) (*RestRechargeForbidResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID
	params.Set("status", status)             // 1 禁用 0启用
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Recharge/forbid", params)
	if err != nil {
		return newRestRechargeForbidResult(RestRechargeForbidResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestRechargeForbidResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeForbidResult(response, request.ResponseBody, request), err
}
