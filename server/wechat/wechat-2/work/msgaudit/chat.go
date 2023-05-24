package msgaudit

import "encoding/json"

// ChatDataResponse 会话存档消息响应数据
type ChatDataResponse struct {
	Error
	ChatDataList []ChatData `json:"chatdata,omitempty"`
}

// IsError 判断是否正确响应
func (c ChatDataResponse) IsError() bool {
	return c.ErrCode != 0
}

// ChatData 会话存档原始数据
type ChatData struct {
	Seq              uint64 `json:"seq,omitempty"`                // 消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	MsgID            string `json:"msgid,omitempty"`              // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	PublickeyVer     uint32 `json:"publickey_ver,omitempty"`      // 加密此条消息使用的公钥版本号。
	EncryptRandomKey string `json:"encrypt_random_key,omitempty"` // 使用publickey_ver指定版本的公钥进行非对称加密后base64加密的内容，需要业务方先base64 decode处理后，再使用指定版本的私钥进行解密，得出内容。
	EncryptChatMsg   string `json:"encrypt_chat_msg,omitempty"`   // 消息密文。需要业务方使用将encrypt_random_key解密得到的内容，与encrypt_chat_msg，传入sdk接口DecryptData,得到消息明文。
}

// ChatMessage 会话存档消息
type ChatMessage struct {
	ID         string   // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	From       string   // 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	ToList     []string // 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	Action     string   // 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	Type       string   // 消息类型
	originData []byte   // 原始消息对象
}

// GetOriginMessage 获取消息原始数据
func (c ChatMessage) GetOriginMessage() (msg map[string]interface{}, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetTextMessage 获取文本消息
func (c ChatMessage) GetTextMessage() (msg TextMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetImageMessage 获取图片消息
func (c ChatMessage) GetImageMessage() (msg ImageMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetRevokeMessage 获取撤回消息
func (c ChatMessage) GetRevokeMessage() (msg RevokeMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetAgreeMessage 获取同意会话聊天内容
func (c ChatMessage) GetAgreeMessage() (msg AgreeMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetVoiceMessage 获取语音消息
func (c ChatMessage) GetVoiceMessage() (msg VoiceMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetVideoMessage 获取视频消息
func (c ChatMessage) GetVideoMessage() (msg VideoMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetCardMessage 获取名片消息
func (c ChatMessage) GetCardMessage() (msg CardMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetLocationMessage 获取位置消息
func (c ChatMessage) GetLocationMessage() (msg LocationMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetEmotionMessage 获取表情消息
func (c ChatMessage) GetEmotionMessage() (msg EmotionMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetFileMessage 获取文件消息
func (c ChatMessage) GetFileMessage() (msg FileMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetLinkMessage 获取链接消息
func (c ChatMessage) GetLinkMessage() (msg LinkMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetWeappMessage 获取小程序消息
func (c ChatMessage) GetWeappMessage() (msg WeappMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetChatRecordMessage 获取会话记录消息
func (c ChatMessage) GetChatRecordMessage() (msg ChatRecordMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetTodoMessage 获取待办消息
func (c ChatMessage) GetTodoMessage() (msg TodoMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetVoteMessage 获取投票消息
func (c ChatMessage) GetVoteMessage() (msg VoteMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetCollectMessage 获取填表消息
func (c ChatMessage) GetCollectMessage() (msg CollectMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetRedpacketMessage 获取红包消息
func (c ChatMessage) GetRedpacketMessage() (msg RedpacketMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetMeetingMessage 获取会议邀请消息
func (c ChatMessage) GetMeetingMessage() (msg MeetingMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetDocMessage 获取在线文档消息
func (c ChatMessage) GetDocMessage() (msg DocMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetMarkdownMessage 获取MarkDown格式消息
func (c ChatMessage) GetMarkdownMessage() (msg MarkdownMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetNewsMessage 获取图文消息
func (c ChatMessage) GetNewsMessage() (msg NewsMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetCalendarMessage 获取日程消息
func (c ChatMessage) GetCalendarMessage() (msg CalendarMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetMixedMessage 获取混合消息
func (c ChatMessage) GetMixedMessage() (msg MixedMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetMeetingVoiceCallMessage 获取音频存档消息
func (c ChatMessage) GetMeetingVoiceCallMessage() (msg MeetingVoiceCallMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetVoipDocShareMessage 获取音频共享消息
func (c ChatMessage) GetVoipDocShareMessage() (msg VoipDocShareMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetExternalRedPacketMessage 获取互通红包消息
func (c ChatMessage) GetExternalRedPacketMessage() (msg ExternalRedPacketMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetSphFeedMessage 获取视频号消息
func (c ChatMessage) GetSphFeedMessage() (msg SphFeedMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}

// GetSwitchMessage 获取切换企业日志
func (c ChatMessage) GetSwitchMessage() (msg SwitchMessage, err error) {
	err = json.Unmarshal(c.originData, &msg)
	return msg, err
}
