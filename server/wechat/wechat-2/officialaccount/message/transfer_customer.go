package message

// TransferCustomer 转发客服消息
type TransferCustomer struct {
	CommonToken

	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

// TransInfo 转发到指定客服
type TransInfo struct {
	KfAccount string `xml:"KfAccount"`
}

// NewTransferCustomer 实例化
func NewTransferCustomer(kfAccount string) *TransferCustomer {
	tc := new(TransferCustomer)
	if kfAccount != "" {
		transInfo := new(TransInfo)
		transInfo.KfAccount = kfAccount
		tc.TransInfo = transInfo
	}
	return tc
}
