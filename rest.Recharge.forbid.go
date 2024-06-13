package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
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

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Recharge/forbid")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID
	params.Set("status", status)             // 1 禁用 0启用

	// 请求
	request, err := c.request(ctx, "rest/Recharge/forbid", params)
	if err != nil {
		return newRestRechargeForbidResult(RestRechargeForbidResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response RestRechargeForbidResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newRestRechargeForbidResult(response, request.ResponseBody, request), err
}
