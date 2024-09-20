package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EPayNotifyReq struct {
	g.Meta `path:"/pay/e_pay_notify" tags:"Pay" method:"get,post" summary:"易支付 异步通知"`
}

type EPayNotifyRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
