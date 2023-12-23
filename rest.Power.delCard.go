package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestPowerDelCardResponse struct {
	Data string `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

type RestPowerDelCardResult struct {
	Result RestPowerDelCardResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestPowerDelCardResult(result RestPowerDelCardResponse, body []byte, http gorequest.Response) *RestPowerDelCardResult {
	return &RestPowerDelCardResult{Result: result, Body: body, Http: http}
}

// RestPowerDelCard 删除电费充值卡
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/330
func (c *Client) RestPowerDelCard(ctx context.Context, cardID int64, notMustParams ...gorequest.Params) (*RestPowerDelCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("card_id", cardID)
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/Power/delCard", params)
	if err != nil {
		return newRestPowerDelCardResult(RestPowerDelCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerDelCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerDelCardResult(response, request.ResponseBody, request), err
}
