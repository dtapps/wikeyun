package wikeyun

import (
	"go.dtapp.net/godecimal"
	"go.dtapp.net/gojson"
	"net/http"
)

type ResponseRestPowerPushOrderNotifyHttp struct {
	Status        int64   `form:"status" json:"status" query:"status"`                              // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	ArrivedAmount string  `form:"arrived_amount" json:"arrived_amount" query:"arrived_amount"`      // 到账金额
	OrderNo       string  `form:"order_no" json:"order_no" query:"order_no"`                        // 第三方单号
	OrderNumber   string  `form:"order_number" json:"order_number" query:"order_number"`            // 微客云平台单号
	Amount        string  `form:"amount" json:"amount" query:"amount"`                              // 充值金额，如50，100，200可选
	Fanli         float64 `form:"fanli" json:"fanli" query:"fanli"`                                 // 返利金额
	CostPrice     float64 `form:"cost_price" json:"cost_price" xml:"cost_price" query:"cost_price"` // 成本价格
	Sign          string  `form:"sign" json:"sign" query:"sign"`                                    // 加密内容
	FailReason    string  `form:"failReason" json:"failReason" query:"failReason"`                  // 失败原因，有些渠道没有返回，不是很准确，但电费失败大部分原因都是户号不对或者地区不对或者缴费金额小于欠费金额
}

// RestPowerPushOrderNotifyHttp 电费充值API - 回调通知
// https://open.wikeyun.cn/#/document/1/article/303
func (c *Client) RestPowerPushOrderNotifyHttp(w http.ResponseWriter, r *http.Request) (validateJson ResponseRestPowerPushOrderNotifyHttp, err error) {
	if r.Method == http.MethodPost {
		err = gojson.NewDecoder(r.Body).Decode(&validateJson)
	} else if r.Method == http.MethodGet {
		validateJson.Status = godecimal.NewString(r.URL.Query().Get("status")).Int64()
		validateJson.ArrivedAmount = r.URL.Query().Get("arrived_amount")
		validateJson.OrderNo = r.URL.Query().Get("order_no")
		validateJson.OrderNumber = r.URL.Query().Get("order_number")
		validateJson.Amount = r.URL.Query().Get("amount")
		validateJson.Fanli = godecimal.NewString(r.URL.Query().Get("fanli")).Float64()
		validateJson.CostPrice = godecimal.NewString(r.URL.Query().Get("cost_price")).Float64()
		validateJson.Sign = r.URL.Query().Get("sign")
		validateJson.FailReason = r.URL.Query().Get("failReason")
	}
	return validateJson, err
}
