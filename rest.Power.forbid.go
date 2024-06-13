package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
)

type RestPowerForbidResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Time string   `json:"time"`
	Data struct{} `json:"data"`
}

type RestPowerForbidResult struct {
	Result RestPowerForbidResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newRestPowerForbidResult(result RestPowerForbidResponse, body []byte, http gorequest.Response) *RestPowerForbidResult {
	return &RestPowerForbidResult{Result: result, Body: body, Http: http}
}

// RestPowerForbid 禁启用非API渠道电费充值
// status = 1 禁用 0启用
// https://open.wikeyun.cn/#/apiDocument/9/document/446
func (c *Client) RestPowerForbid(ctx context.Context, status int64, notMustParams ...gorequest.Params) (*RestPowerForbidResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Power/forbid")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID
	params.Set("status", status)             // 1 禁用 0启用

	// 请求
	request, err := c.request(ctx, "rest/Power/forbid", params)
	if err != nil {
		return newRestPowerForbidResult(RestPowerForbidResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response RestPowerForbidResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newRestPowerForbidResult(response, request.ResponseBody, request), err
}
