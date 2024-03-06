package wikeyun

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type RestMovieHotCityResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CityId      int64  `json:"cityId"`      // 城市ID
		CityName    string `json:"cityName"`    // 城市名称
		Firstletter string `json:"firstletter"` // 城市名称首字母
		Ishot       int64  `json:"ishot"`       // 是否热门 1是 0否
	} `json:"data"` // 城市列表
}

type RestMovieHotCityResult struct {
	Result RestMovieHotCityResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestMovieHotCityResult(result RestMovieHotCityResponse, body []byte, http gorequest.Response) *RestMovieHotCityResult {
	return &RestMovieHotCityResult{Result: result, Body: body, Http: http}
}

// RestMovieHotCity 定位--获取热门城市
// https://open.wikeyun.cn/#/apiDocument/4/document/301
func (c *Client) RestMovieHotCity(ctx context.Context, notMustParams ...gorequest.Params) (*RestMovieHotCityResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.config.storeId) // 店铺ID
	// 请求
	request, err := c.request(ctx, c.config.apiUrl+"/rest/movie/hotCity", params)
	if err != nil {
		return newRestMovieHotCityResult(RestMovieHotCityResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestMovieHotCityResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestMovieHotCityResult(response, request.ResponseBody, request), err
}
