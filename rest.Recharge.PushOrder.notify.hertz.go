package wikeyun

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// ResponseRestRechargePushOrderNotifyHertz 声明接收的变量
type ResponseRestRechargePushOrderNotifyHertz struct {
	Status      int64   `form:"status" json:"status" query:"status"`                   // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	Mobile      string  `form:"mobile" json:"mobile" query:"mobile"`                   // 充值手机号
	OrderNo     string  `form:"order_no" json:"order_no" query:"order_no"`             // 第三方单号
	OrderNumber string  `form:"order_number" json:"order_number" query:"order_number"` // 微客云平台单号
	Amount      string  `form:"amount" json:"amount" query:"amount"`                   // 充值金额，如50，100，200可选
	Fanli       float64 `form:"fanli" json:"fanli" query:"fanli"`                      // 返利金额
	CostPrice   float64 `form:"cost_price" json:"cost_price" query:"cost_price"`       // 成本价格
	Sign        string  `form:"sign" json:"sign" query:"sign"`                         // 加密内容
}

// RestRechargePushOrderNotifyHertz 话费充值推送 - 回调通知
// https://open.wikeyun.cn/#/document/1/article/302
func (c *Client) RestRechargePushOrderNotifyHertz(ctx context.Context, h *app.RequestContext) (validateJson ResponseRestRechargePushOrderNotifyHertz, err error) {
	err = h.BindAndValidate(&validateJson)
	return validateJson, err
}
