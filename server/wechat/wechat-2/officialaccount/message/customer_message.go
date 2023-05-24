package message

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	customerSendMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

// Manager 消息管理者，可以发送消息
type Manager struct {
	*context.Context
}

// NewMessageManager 实例化消息管理者
func NewMessageManager(context *context.Context) *Manager {
	return &Manager{
		context,
	}
}

// CustomerMessage  客服消息
type CustomerMessage struct {
	ToUser          string                `json:"touser"`                    // 接受者OpenID
	Msgtype         MsgType               `json:"msgtype"`                   // 客服消息类型
	Text            *MediaText            `json:"text,omitempty"`            // 可选
	Image           *MediaResource        `json:"image,omitempty"`           // 可选
	Voice           *MediaResource        `json:"voice,omitempty"`           // 可选
	Video           *MediaVideo           `json:"video,omitempty"`           // 可选
	Music           *MediaMusic           `json:"music,omitempty"`           // 可选
	News            *MediaNews            `json:"news,omitempty"`            // 可选
	Mpnews          *MediaResource        `json:"mpnews,omitempty"`          // 可选
	Wxcard          *MediaWxcard          `json:"wxcard,omitempty"`          // 可选
	Msgmenu         *MediaMsgmenu         `json:"msgmenu,omitempty"`         // 可选
	Miniprogrampage *MediaMiniprogrampage `json:"miniprogrampage,omitempty"` // 可选
}

// NewCustomerTextMessage 文本消息结构体构造方法
func NewCustomerTextMessage(toUser, text string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeText,
		Text: &MediaText{
			text,
		},
	}
}

// NewCustomerImgMessage 图片消息的构造方法
func NewCustomerImgMessage(toUser, mediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeImage,
		Image: &MediaResource{
			mediaID,
		},
	}
}

// NewCustomerVoiceMessage 语音消息的构造方法
func NewCustomerVoiceMessage(toUser, mediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeVoice,
		Voice: &MediaResource{
			mediaID,
		},
	}
}

// NewCustomerMiniprogrampageMessage 小程序卡片消息的构造方法
func NewCustomerMiniprogrampageMessage(toUser, title, appID, pagePath, thumbMediaID string) *CustomerMessage {
	return &CustomerMessage{
		ToUser:  toUser,
		Msgtype: MsgTypeMiniprogrampage,
		Miniprogrampage: &MediaMiniprogrampage{
			Title:        title,
			AppID:        appID,
			Pagepath:     pagePath,
			ThumbMediaID: thumbMediaID,
		},
	}
}

// MediaText 文本消息的文字
type MediaText struct {
	Content string `json:"content"`
}

// MediaResource  消息使用的永久素材id
type MediaResource struct {
	MediaID string `json:"media_id"`
}

// MediaVideo 视频消息包含的内容
type MediaVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// MediaMusic 音乐消息包括的内容
type MediaMusic struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Musicurl     string `json:"musicurl"`
	Hqmusicurl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// MediaNews 图文消息的内容
type MediaNews struct {
	Articles []MediaArticles `json:"articles"`
}

// MediaArticles 图文消息的内容的文章列表中的单独一条
type MediaArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

// MediaMsgmenu 菜单消息的内容
type MediaMsgmenu struct {
	HeadContent string        `json:"head_content"`
	List        []MsgmenuItem `json:"list"`
	TailContent string        `json:"tail_content"`
}

// MsgmenuItem 菜单消息的菜单按钮
type MsgmenuItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// MediaWxcard 卡券的id
type MediaWxcard struct {
	CardID string `json:"card_id"`
}

// MediaMiniprogrampage 小程序消息
type MediaMiniprogrampage struct {
	Title        string `json:"title"`
	AppID        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
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
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("customer msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err
	}

	return nil
}
