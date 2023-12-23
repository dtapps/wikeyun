package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestUserQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
		Money  string `json:"money"`
		Mobile string `json:"mobile"`
	} `json:"data"`
}

type RestUserQueryResult struct {
	Result RestUserQueryResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newRestUserQueryResult(result RestUserQueryResponse, body []byte, http gorequest.Response) *RestUserQueryResult {
	return &RestUserQueryResult{Result: result, Body: body, Http: http}
}

// RestUserQuery 用户信息
// https://open.wikeyun.cn/#/apiDocument/10/document/336
func (c *Client) RestUserQuery(ctx context.Context, notMustParams ...gorequest.Params) (*RestUserQueryResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/User/query", params)
	if err != nil {
		return newRestUserQueryResult(RestUserQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestUserQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestUserQueryResult(response, request.ResponseBody, request), err
}
