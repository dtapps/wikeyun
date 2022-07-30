package wikeyun

import "go.dtapp.net/gorequest"

// RestOilOrderPush 充值下单
func (c *Client) RestOilOrderPush(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Oil/pushOrder", params)
	return request.ResponseBody, err
}
