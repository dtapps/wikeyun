package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestPowerEditCardResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RestPowerEditCardResult struct {
	Result RestPowerEditCardResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newRestPowerEditCardResult(result RestPowerEditCardResponse, body []byte, http gorequest.Response) *RestPowerEditCardResult {
	return &RestPowerEditCardResult{Result: result, Body: body, Http: http}
}

// RestPowerEditCard 编辑电费充值卡
// card_id = 充值卡ID
// card_num = 卡号
// province = 省份
// city = 城市
// type = 0国家电网 1南方电网
// remark = 备注
// user_ext = 南网必填，请输入用户信息，身份证后六位 / 营业执照后六位 / 银行卡后六位 ，三者选任意一个即可
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (c *Client) RestPowerEditCard(ctx context.Context, cardID int64, cardNum string, province string, city string, Type int64, notMustParams ...gorequest.Params) (*RestPowerEditCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("card_id", cardID)    // 充值卡ID
	params.Set("card_num", cardNum)  // 卡号
	params.Set("province", province) // 省份
	params.Set("city", city)         // 城市
	params.Set("type", Type)         // 0国家电网 1南方电网
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/editCard", params)
	if err != nil {
		return newRestPowerEditCardResult(RestPowerEditCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerEditCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerEditCardResult(response, request.ResponseBody, request), err
}
