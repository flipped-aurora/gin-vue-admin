package device

// MsgDevice 设备消息响应
type MsgDevice struct {
	DeviceType string
	DeviceID   string
	SessionID  string
	OpenID     string
}
