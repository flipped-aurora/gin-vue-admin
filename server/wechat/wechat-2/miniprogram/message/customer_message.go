package message

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	customerSendMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

// Manager 消息管理者，可以发送消息
type Manager struct {
	*context.Context
}

// NewCustomerMessageManager 实例化消息管理者
func NewCustomerMessageManager(context *context.Context) *Manager {
	return &Manager{
		context,
	}
}

// MediaText 文本消息的文字
type MediaText struct {
	Content string `json:"content"`
}

// MediaResource  消息使用的临时素材id
type MediaResource struct {
	MediaID string `json:"media_id"`
}

// MediaMiniprogrampage 小程序卡片
type MediaMiniprogrampage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// MediaLink 发送图文链接
type MediaLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// CustomerMessage  客服消息
type CustomerMessage struct {
	ToUser          string                `json:"touser"`                    // 接受者OpenID
	Msgtype         MsgType               `json:"msgtype"`                   // 客服消息类型
	Text            *MediaText            `json:"text,omitempty"`            // 可选
	Image           *MediaResource        `json:"image,omitempty"`           // 可选
	Link            *MediaLink            `json:"link,omitempty"`            // 可选
	Miniprogrampage *MediaMiniprogrampage `json:"miniprogrampage,omitempty"` // 可选
}

// NewCustomerTextMessage 文本消息结构体构造方法
func NewCustomerTextMessage(toUser, text string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeText,
		Text: &MediaText{
			Content: text,
		},
	}
}

// NewCustomerImgMessage 图片消息的构造方法
func NewCustomerImgMessage(toUser, mediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeImage,
		Image: &MediaResource{
			MediaID: mediaID,
		},
	}
}

// NewCustomerLinkMessage 图文链接消息的构造方法
func NewCustomerLinkMessage(toUser, title, description, url, thumbURL string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeLink,
		Link: &MediaLink{
			Title:       title,
			Description: description,
			URL:         url,
			ThumbURL:    thumbURL,
		},
	}
}

// NewCustomerMiniprogrampageMessage 小程序卡片消息的构造方法
func NewCustomerMiniprogrampageMessage(toUser, title, pagepath, thumbMediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeMiniProgramPage,
		Miniprogrampage: &MediaMiniprogrampage{
			Title:        title,
			Pagepath:     pagepath,
			ThumbMediaID: thumbMediaID,
		},
	}
}

// Send 发送客服消息
func (manager *Manager) Send(msg *CustomerMessage) error {
	accessToken, err := manager.Context.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerSendMessage, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "SendCustomerMessage")
}
