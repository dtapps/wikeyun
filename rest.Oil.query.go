package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
)

type RestOilQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string  `json:"order_number"`
		OrderNo     string  `json:"order_no"`
		CardId      int64   `json:"card_id"`
		Amount      int64   `json:"amount"`
		CostPrice   float64 `json:"cost_price"`
		Fanli       float64 `json:"fanli"`
		Status      int64   `json:"status"`
	} `json:"data"`
}

type RestOilQueryResult struct {
	Result RestOilQueryResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newRestOilQueryResult(result RestOilQueryResponse, body []byte, http gorequest.Response) *RestOilQueryResult {
	return &RestOilQueryResult{Result: result, Body: body, Http: http}
}

// RestOilQuery 油卡订单查询
// order_number = 平台单号，与商户单号二选一
// order_no = 商户单号
// https://open.wikeyun.cn/#/apiDocument/9/document/368
func (c *Client) RestOilQuery(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Oil/query")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID

	// 请求
	request, err := c.request(ctx, "rest/Oil/query", params)
	if err != nil {
		return newRestOilQueryResult(RestOilQueryResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response RestOilQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newRestOilQueryResult(response, request.ResponseBody, request), err
}
