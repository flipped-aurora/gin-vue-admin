package enum

const (
	Pending = iota
	Success
	Fail
)

const (
	WechatJSAPI = iota + 1
	WechatH5
	WechatApp
	AliH5
	AliApp
)
