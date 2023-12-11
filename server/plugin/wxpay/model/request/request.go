package request

// Native 微信下单参数 Native 下单
type Native struct {
	UserID  int    `form:"userid" json:"userid"`                      // 下单用户ID
	GoodsID string `form:"goodsid" josn:"goodsid" binding:"required"` // 商品ID用于查询商品价格
	Attach  string `form:"attach" json:"attach"`                      // 附加参数
	CallUrl string `form:"callurl" json:"callUrl"`                    //第三方回调call
}

// ReqTradeNo 微信查询订单
type ReqTradeNo struct {
	OrderId string `form:"orderid" binding:"required,min=5,max=24"` // 订单ID
}

// Jsapi 微信下单参数 Jsapi 下单
type Jsapi struct {
	UserID  int    `form:"userid" json:"userid"`                    // 下单用户ID
	GoodsID string `form:"goodsid" josn:"goodsid"`                  // 商品ID
	Attach  string `form:"attach" json:"attach" binding:"required"` // 附加参数
	CallUrl string `form:"callurl" json:"callUrl"`                  // 第三方回调call
	OpenId  string `form:"openId" json:"openId" binding:"required"` // 微信用户openid
	OrderId string `form:"orderid" `                                // 订单ID
}

// ToPayJsapi 微信下单参数 Jsapi 扫码下单
type ToPayJsapi struct {
	Money  int64  `form:"money" json:"money" binding:"required"`   // 金额 ，自定义订单使用
	OpenId string `form:"openId" json:"openId" binding:"required"` // 微信用户openid
}

// JsapiGetCode jsapi 获取用户授权code
type JsapiGetCode struct {
	GoodsID string `form:"patternId" json:"patternId" ` // 商品ID
	Macid   string `form:"macid" josn:"macid"`          // 商品ID
}

// JsapiGetCode jsapi 获取用户授权token
type JsapiGetToken struct {
	Attach string `form:"attach" json:"attach" ` // 附加数据，在查询API和支付通知中原样返回
	Code   string `form:"code" josn:"code" `     // 商品ID用于查询商品价格
	State  string `form:"state" josn:"state"`    // 商品ID用于查询商品价格
}
