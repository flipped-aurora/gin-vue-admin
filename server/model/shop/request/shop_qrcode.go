package request

type RequestMhtData struct {
	Ts           string `json:"ts"`
	Key          string `json:"key"`
	PatternId    string `json:"patternId"`
	Macid        string `json:"macid"`
	AgencyNo     string `json:"agencyNo"`
	HlMerchantId string `json:"hlMerchantId"`
	OutTreadNo   string `json:"outTreadNo"`
}

// Attach 棉花糖机请求参数
type Attach struct {
	Attach string `form:"attach" json:"attach"` // 附加参数
}

// ToPay 棉花糖机请求参数
type ToPay struct {
	Money string `form:"money" json:"money"` // 为了兼容自定义付款添加的字段
}
