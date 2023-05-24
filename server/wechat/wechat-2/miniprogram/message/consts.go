package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
	// MsgTypeText 文本消息
	MsgTypeText MsgType = "text"
	// MsgTypeImage 图片消息
	MsgTypeImage = "image"
	// MsgTypeLink 图文链接
	MsgTypeLink = "link"
	// MsgTypeMiniProgramPage 小程序卡片
	MsgTypeMiniProgramPage = "miniprogrampage"
)

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

// MiniProgramMixMessage 小程序回调的消息结构
type MiniProgramMixMessage struct {
	CommonToken

	MsgID int64 `xml:"MsgId"`

	// 文本消息
	Content string `xml:"Content"`

	// 图片消息
	PicURL  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`

	// 小程序卡片消息
	Title        string `xml:"Title"`
	AppID        string `xml:"AppId"`
	PagePath     string `xml:"PagePath"`
	ThumbURL     string `xml:"ThumbUrl"`
	ThumbMediaID string `xml:"ThumbMediaId"`

	// 进入会话事件
	Event       string `xml:"Event"`
	SessionFrom string `xml:"SessionFrom"`
}
