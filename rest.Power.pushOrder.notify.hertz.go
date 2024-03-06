package wikeyun

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type ResponseRestPowerPushOrderNotifyHertz struct {
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

// RestPowerPushOrderNotifyHertz 电费充值API - 回调通知
// https://open.wikeyun.cn/#/document/1/article/303
func (c *Client) RestPowerPushOrderNotifyHertz(ctx context.Context, h *app.RequestContext) (validateJson ResponseRestPowerPushOrderNotifyHertz, err error) {
	err = h.BindAndValidate(&validateJson)
	return validateJson, err
}
