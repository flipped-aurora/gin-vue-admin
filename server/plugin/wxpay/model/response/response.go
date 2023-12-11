package response

type ResqQueryOrder struct {
	Code        int    `json:"code"`
	TradeStatus string `json:"tradestatus"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
}
