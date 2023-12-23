package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestOilPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string  `json:"order_number"`
		Amount      int64   `json:"amount"`
		Fanli       float64 `json:"fanli"`
		CostPrice   float64 `json:"cost_price"`
	} `json:"data"`
}

type RestOilPushOrderResult struct {
	Result RestOilPushOrderResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestOilPushOrderResult(result RestOilPushOrderResponse, body []byte, http gorequest.Response) *RestOilPushOrderResult {
	return &RestOilPushOrderResult{Result: result, Body: body, Http: http}
}

// RestOilPushOrder 油卡充值
// store_id = 店铺ID
// order_no = 商户单号
// amount = 充值金额
// recharge_type = 充值类型 1快充 0慢充
// notify_url = 回调通知地址，用于订单状态通知
// cardId = 卡号ID，通过新增获取
// https://open.wikeyun.cn/#/apiDocument/9/document/367
func (c *Client) RestOilPushOrder(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilPushOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Oil/pushOrder", params)
	if err != nil {
		return newRestOilPushOrderResult(RestOilPushOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestOilPushOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestOilPushOrderResult(response, request.ResponseBody, request), err
}
