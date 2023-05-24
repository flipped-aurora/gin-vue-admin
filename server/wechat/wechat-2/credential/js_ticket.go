package credential

// JsTicketHandle js ticket获取
type JsTicketHandle interface {
	// GetTicket 获取ticket
	GetTicket(accessToken string) (ticket string, err error)
}
