package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestOilAddCardResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

type RestOilAddCardResult struct {
	Result RestOilAddCardResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newRestOilAddCardResult(result RestOilAddCardResponse, body []byte, http gorequest.Response) *RestOilAddCardResult {
	return &RestOilAddCardResult{Result: result, Body: body, Http: http}
}

// RestOilAddCard 添加油卡充值卡
// card_id = 充值卡ID
// card_num = 卡号
// name = 姓名
// phone = 手机号
// card_type = 类型 0中石化 1中石油
// user_num = 身份证号
// https://open.wikeyun.cn/#/apiDocument/9/document/370
func (c *Client) RestOilAddCard(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilAddCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Oil/addCard", params)
	if err != nil {
		return newRestOilAddCardResult(RestOilAddCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestOilAddCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestOilAddCardResult(response, request.ResponseBody, request), err
}
