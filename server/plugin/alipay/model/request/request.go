package request

type Request struct {
	patternId string // 商品id
	macid     string // 商户id
}

type ResAuthCode struct {
	Attach string `form:"attach" json:"attach"` // 附加参数
}

type ResAuthToken struct {
	Attach string `form:"attach" json:"attach"`       // 附加参数.
	Code   string `form:"auth_code" json:"auth_code"` // 支付宝获取的临时code 用来换取token
}

type ResAuthJsApi struct {
	UserId string `form:"userId" json:"userId"`       //支付宝用户id
	OpenId string `form:"ailopenId" json:"ailopenId"` //支付宝用户新版ID
	Attach string `form:"attach" json:"attach"`       // 附加参数
	Money  int64  `form:"money" json:"money"`         // 金额 ，自定义订单使用
}

// ReqTradeNo 退款
type ReqTradeNo struct {
	OutTradeNo string `form:"OutTradeNo" binding:"required,min=5,max=24"` // 订单ID
}
