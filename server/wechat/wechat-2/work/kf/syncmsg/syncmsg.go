package syncmsg

import "encoding/json"

// Message 同步的消息内容
type Message struct {
	MsgID              string `json:"msgid"`           // 消息ID
	OpenKFID           string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID     string `json:"external_userid"` // 客户UserID
	ReceptionistUserID string `json:"servicer_userid"` // 接待客服userID
	SendTime           uint64 `json:"send_time"`       // 消息发送时间
	Origin             uint32 `json:"origin"`          // 消息来源。3-客户回复的消息 4-系统推送的消 息
	MsgType            string `json:"msgtype"`         // 消息类型
	EventType          string `json:"event_type"`      // 事件类型
	OriginData         []byte `json:"origin_data"`     // 原始数据内容
}

// GetOriginMessage 获取原始消息
func (r Message) GetOriginMessage() (info []byte) {
	return r.OriginData
}

// GetTextMessage 获取文本消息
func (r Message) GetTextMessage() (info Text, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetImageMessage 获取图片消息
func (r Message) GetImageMessage() (info Image, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetVoiceMessage 获取语音消息
func (r Message) GetVoiceMessage() (info Voice, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetVideoMessage 获取视频消息
func (r Message) GetVideoMessage() (info Video, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetFileMessage 获取文件消息
func (r Message) GetFileMessage() (info File, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetLocationMessage 获取文件消息
func (r Message) GetLocationMessage() (info Location, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetLinkMessage 获取链接消息
func (r Message) GetLinkMessage() (info Link, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetBusinessCardMessage 获取名片消息
func (r Message) GetBusinessCardMessage() (info BusinessCard, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetMiniProgramMessage 获取小程序消息
func (r Message) GetMiniProgramMessage() (info MiniProgram, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetEnterSessionEvent 用户进入会话事件
func (r Message) GetEnterSessionEvent() (info EnterSessionEvent, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	info.OpenKFID = info.Event.OpenKFID
	info.ExternalUserID = info.Event.ExternalUserID
	return info, err
}

// GetMsgSendFailEvent 消息发送失败事件
func (r Message) GetMsgSendFailEvent() (info MsgSendFailEvent, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetReceptionistStatusChangeEvent 客服人员接待状态变更事件
func (r Message) GetReceptionistStatusChangeEvent() (info ReceptionistStatusChangeEvent, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	return info, err
}

// GetSessionStatusChangeEvent 会话状态变更事件
func (r Message) GetSessionStatusChangeEvent() (info SessionStatusChangeEvent, err error) {
	err = json.Unmarshal(r.OriginData, &info)
	info.OpenKFID = info.Event.OpenKFID
	info.ExternalUserID = info.Event.ExternalUserID
	return info, err
}
