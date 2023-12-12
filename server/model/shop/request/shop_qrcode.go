package request

type RequestMhtData struct {
	Ts           string `form:"ts" json:"ts"`
	Key          string `form:"key" json:"key"`
	PatternId    string `form:"patternId" json:"patternId"`
	Macid        string `form:"macid" json:"macid"`
	AgencyNo     string `form:"agencyNo" json:"agencyNo"`
	HlMerchantId string `form:"hlMerchantId" json:"hlMerchantId"`
	OutTreadNo   string `form:"outTreadNo" json:"outTreadNo"`
}

// Attach 棉花糖机请求参数
type Attach struct {
	Attach string `form:"attach" json:"attach"` // 附加参数
}

// ToPay 棉花糖机请求参数
type ToPay struct {
	Money string `form:"money" json:"money"` // 为了兼容自定义付款添加的字段
}
