package wikeyun

import (
	"go.dtapp.net/godecimal"
	"go.dtapp.net/gojson"
	"net/http"
)

// ResponseRestRechargePushOrderNotifyHttp 声明接收的变量
type ResponseRestRechargePushOrderNotifyHttp struct {
	Status      int64   `json:"status"`       // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	Mobile      string  `json:"mobile"`       // 充值手机号
	OrderNo     string  `json:"order_no"`     // 第三方单号
	OrderNumber string  `json:"order_number"` // 微客云平台单号
	Amount      string  `json:"amount"`       // 充值金额，如50，100，200可选
	Fanli       float64 `json:"fanli"`        // 返利金额
	CostPrice   float64 `json:"cost_price"`   // 成本价格
	Sign        string  `json:"sign"`         // 加密内容
}

// RestRechargePushOrderNotifyHttp 话费充值推送 - 回调通知
// https://open.wikeyun.cn/#/document/1/article/302
func (c *Client) RestRechargePushOrderNotifyHttp(w http.ResponseWriter, r *http.Request) (validateJson ResponseRestRechargePushOrderNotifyHttp, err error) {
	if r.Method == http.MethodPost {
		err = gojson.NewDecoder(r.Body).Decode(&validateJson)
	} else if r.Method == http.MethodGet {
		validateJson.Status = godecimal.NewString(r.URL.Query().Get("status")).Int64()
		validateJson.Mobile = r.URL.Query().Get("mobile")
		validateJson.OrderNo = r.URL.Query().Get("order_no")
		validateJson.OrderNumber = r.URL.Query().Get("order_number")
		validateJson.Amount = r.URL.Query().Get("amount")
		validateJson.Fanli = godecimal.NewString(r.URL.Query().Get("fanli")).Float64()
		validateJson.CostPrice = godecimal.NewString(r.URL.Query().Get("cost_price")).Float64()
		validateJson.Sign = r.URL.Query().Get("sign")
	}
	return validateJson, err
}
