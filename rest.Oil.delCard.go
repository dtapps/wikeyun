package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestOilDelCardResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

type RestOilDelCardResult struct {
	Result RestOilDelCardResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newRestOilDelCardResult(result RestOilDelCardResponse, body []byte, http gorequest.Response) *RestOilDelCardResult {
	return &RestOilDelCardResult{Result: result, Body: body, Http: http}
}

// RestOilDelCard 删除油卡充值卡
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/372
func (c *Client) RestOilDelCard(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilDelCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Oil/delCard", params)
	if err != nil {
		return newRestOilDelCardResult(RestOilDelCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestOilDelCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestOilDelCardResult(response, request.ResponseBody, request), err
}
