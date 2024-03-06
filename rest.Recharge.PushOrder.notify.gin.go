package wikeyun

import (
	"context"
	"github.com/gin-gonic/gin"
)

// ResponseRestRechargePushOrderNotifyGin 声明接收的变量
type ResponseRestRechargePushOrderNotifyGin struct {
	Status      int64   `form:"status" json:"status" uri:"status" binding:"omitempty"`                   // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	Mobile      string  `form:"mobile" json:"mobile" uri:"mobile" binding:"omitempty"`                   // 充值手机号
	OrderNo     string  `form:"order_no" json:"order_no" uri:"order_no" binding:"omitempty"`             // 第三方单号
	OrderNumber string  `form:"order_number" json:"order_number" uri:"order_number" binding:"omitempty"` // 微客云平台单号
	Amount      string  `form:"amount" json:"amount" uri:"amount" binding:"omitempty"`                   // 充值金额，如50，100，200可选
	Fanli       float64 `form:"fanli" json:"fanli" uri:"fanli" binding:"omitempty"`                      // 返利金额
	CostPrice   float64 `form:"cost_price" json:"cost_price" uri:"cost_price" binding:"omitempty"`       // 成本价格
	Sign        string  `form:"sign" json:"sign" uri:"sign" binding:"omitempty"`                         // 加密内容
}

// RestRechargePushOrderNotifyGin 话费充值推送 - 回调通知
// https://open.wikeyun.cn/#/document/1/article/302
func (c *Client) RestRechargePushOrderNotifyGin(ctx context.Context, g *gin.Context) (validateJson ResponseRestRechargePushOrderNotifyGin, err error) {
	err = g.ShouldBind(&validateJson)
	return validateJson, err
}
