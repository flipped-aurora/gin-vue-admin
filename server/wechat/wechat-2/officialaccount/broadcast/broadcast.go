package broadcast

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	sendURLByTag      = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall"
	sendURLByOpenID   = "https://api.weixin.qq.com/cgi-bin/message/mass/send"
	deleteSendURL     = "https://api.weixin.qq.com/cgi-bin/message/mass/delete"
	previewSendURL    = "https://api.weixin.qq.com/cgi-bin/message/mass/preview"
	massStatusSendURL = "https://api.weixin.qq.com/cgi-bin/message/mass/get"
	getSpeedSendURL   = "https://api.weixin.qq.com/cgi-bin/message/mass/speed/get"
	setSpeedSendURL   = "https://api.weixin.qq.com/cgi-bin/message/mass/speed/set"
)

// MsgType 发送消息类型
type MsgType string

const (
	// MsgTypeNews 图文消息
	MsgTypeNews MsgType = "mpnews"
	// MsgTypeText 文本
	MsgTypeText MsgType = "text"
	// MsgTypeVoice 语音/音频
	MsgTypeVoice MsgType = "voice"
	// MsgTypeImage 图片
	MsgTypeImage MsgType = "image"
	// MsgTypeVideo 视频
	MsgTypeVideo MsgType = "mpvideo"
	// MsgTypeWxCard 卡券
	MsgTypeWxCard MsgType = "wxcard"
)

// Broadcast 群发消息
type Broadcast struct {
	*context.Context
	preview bool
}

// NewBroadcast new
func NewBroadcast(ctx *context.Context) *Broadcast {
	return &Broadcast{ctx, false}
}

// User 发送的用户
type User struct {
	TagID  int64
	OpenID []string
}

// Result 群发返回结果
type Result struct {
	util.CommonError
	MsgID     int64  `json:"msg_id"`
	MsgDataID int64  `json:"msg_data_id"`
	MsgStatus string `json:"msg_status"`
}

// SpeedResult 群发速度返回结果
type SpeedResult struct {
	util.CommonError
	Speed     int64 `json:"speed"`
	RealSpeed int64 `json:"realspeed"`
}

// sendRequest 发送请求的数据
type sendRequest struct {
	// 根据tag获全部发送
	Filter map[string]interface{} `json:"filter,omitempty"`
	// 根据OpenID发送
	ToUser interface{} `json:"touser,omitempty"`
	// 发送文本
	Text map[string]interface{} `json:"text,omitempty"`
	// 发送图文消息
	Mpnews map[string]interface{} `json:"mpnews,omitempty"`
	// 发送语音
	Voice map[string]interface{} `json:"voice,omitempty"`
	// 发送图片
	Images *Image `json:"images,omitempty"`
	// 发送卡券
	WxCard            map[string]interface{} `json:"wxcard,omitempty"`
	MsgType           MsgType                `json:"msgtype"`
	SendIgnoreReprint int32                  `json:"send_ignore_reprint,omitempty"`
}

// Image 发送图片
type Image struct {
	MediaIDs           []string `json:"media_ids"`
	Recommend          string   `json:"recommend"`
	NeedOpenComment    int32    `json:"need_open_comment"`
	OnlyFansCanComment int32    `json:"only_fans_can_comment"`
}

// SendText 群发文本
// user 为nil，表示全员发送
// &User{TagID:2} 根据tag发送
// &User{OpenID:[]string("xxx","xxx")} 根据openid发送
func (broadcast *Broadcast) SendText(user *User, content string) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeText,
	}
	req.Text = map[string]interface{}{
		"content": content,
	}
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendText")
	return res, err
}

// SendNews 发送图文
func (broadcast *Broadcast) SendNews(user *User, mediaID string, ignoreReprint bool) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeNews,
	}
	if ignoreReprint {
		req.SendIgnoreReprint = 1
	}
	req.Mpnews = map[string]interface{}{
		"media_id": mediaID,
	}
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendNews")
	return res, err
}

// SendVoice 发送语音
func (broadcast *Broadcast) SendVoice(user *User, mediaID string) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeVoice,
	}
	req.Voice = map[string]interface{}{
		"media_id": mediaID,
	}
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendVoice")
	return res, err
}

// SendImage 发送图片
func (broadcast *Broadcast) SendImage(user *User, images *Image) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeImage,
	}
	req.Images = images
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendImage")
	return res, err
}

// SendVideo 发送视频
func (broadcast *Broadcast) SendVideo(user *User, mediaID string, title, description string) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeVideo,
	}
	req.Voice = map[string]interface{}{
		"media_id":    mediaID,
		"title":       title,
		"description": description,
	}
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendVideo")
	return res, err
}

// SendWxCard 发送卡券
func (broadcast *Broadcast) SendWxCard(user *User, cardID string) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := &sendRequest{
		ToUser:  nil,
		MsgType: MsgTypeWxCard,
	}
	req.WxCard = map[string]interface{}{
		"card_id": cardID,
	}
	req, sendURL := broadcast.chooseTagOrOpenID(user, req)
	url := fmt.Sprintf("%s?access_token=%s", sendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "SendWxCard")
	return res, err
}

// Delete 删除群发消息
func (broadcast *Broadcast) Delete(msgID int64, articleIDx int64) error {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return err
	}
	req := map[string]interface{}{
		"msg_id":      msgID,
		"article_idx": articleIDx,
	}
	url := fmt.Sprintf("%s?access_token=%s", deleteSendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "Delete")
}

// Preview 预览
func (broadcast *Broadcast) Preview() *Broadcast {
	broadcast.preview = true
	return broadcast
}

// GetMassStatus 获取群发状态
func (broadcast *Broadcast) GetMassStatus(msgID string) (*Result, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := map[string]interface{}{
		"msg_id": msgID,
	}
	url := fmt.Sprintf("%s?access_token=%s", massStatusSendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &Result{}
	err = util.DecodeWithError(data, res, "GetMassStatus")
	return res, err
}

// GetSpeed 获取群发速度
func (broadcast *Broadcast) GetSpeed() (*SpeedResult, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := map[string]interface{}{}
	url := fmt.Sprintf("%s?access_token=%s", getSpeedSendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &SpeedResult{}
	err = util.DecodeWithError(data, res, "GetSpeed")
	return res, err
}

// SetSpeed 设置群发速度
func (broadcast *Broadcast) SetSpeed(speed int) (*SpeedResult, error) {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req := map[string]interface{}{
		"speed": speed,
	}
	url := fmt.Sprintf("%s?access_token=%s", setSpeedSendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	res := &SpeedResult{}
	err = util.DecodeWithError(data, res, "SetSpeed")
	return res, err
}

func (broadcast *Broadcast) chooseTagOrOpenID(user *User, req *sendRequest) (ret *sendRequest, url string) {
	sendURL := ""
	if user == nil {
		req.Filter = map[string]interface{}{
			"is_to_all": true,
		}
		sendURL = sendURLByTag
	} else {
		if broadcast.preview {
			// 预览 默认发给第一个用户
			if len(user.OpenID) != 0 {
				req.ToUser = user.OpenID[0]
				sendURL = previewSendURL
			}
		} else {
			if user.TagID != 0 {
				req.Filter = map[string]interface{}{
					"is_to_all": false,
					"tag_id":    user.TagID,
				}
				sendURL = sendURLByTag
			}
			if len(user.OpenID) != 0 {
				req.ToUser = user.OpenID
				sendURL = sendURLByOpenID
			}
		}
	}
	return req, sendURL
}
