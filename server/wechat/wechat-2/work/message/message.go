package message

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 发送应用消息的接口地址
	sendURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

type (
	// SendRequestCommon 发送应用消息请求公共参数
	SendRequestCommon struct {
		// 指定接收消息的成员，成员ID列表（多个接收者用‘|’分隔，最多支持1000个）。 特殊情况：指定为"@all"，则向该企业应用的全部成员发送
		ToUser string `json:"touser"`
		// 指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。 当touser为"@all"时忽略本参数
		ToParty string `json:"toparty"`
		// 指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。 当touser为"@all"时忽略本参数
		ToTag string `json:"totag"`
		// 消息类型，此时固定为：text
		MsgType string `json:"msgtype"`
		// 企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
		AgentID string `json:"agentid"`
		// 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
		Safe int `json:"safe"`
		// 表示是否开启id转译，0表示否，1表示是，默认0。仅第三方应用需要用到，企业自建应用可以忽略。
		EnableIDTrans int `json:"enable_id_trans"`
		// 表示是否开启重复消息检查，0表示否，1表示是，默认0
		EnableDuplicateCheck int `json:"enable_duplicate_check"`
		// 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
		DuplicateCheckInterval int `json:"duplicate_check_interval"`
	}
	// SendResponse 发送应用消息响应参数
	SendResponse struct {
		util.CommonError
		InvalidUser    string `json:"invaliduser"`    // 不合法的userid，不区分大小写，统一转为小写
		InvalidParty   string `json:"invalidparty"`   // 不合法的partyid
		InvalidTag     string `json:"invalidtag"`     // 不合法的标签id
		UnlicensedUser string `json:"unlicenseduser"` // 没有基础接口许可(包含已过期)的userid
		MsgID          string `json:"msgid"`          // 消息id
		ResponseCode   string `json:"response_code"`
	}

	// SendTextRequest 发送文本消息的请求
	SendTextRequest struct {
		*SendRequestCommon
		Text TextField `json:"text"`
	}
	// TextField 文本消息参数
	TextField struct {
		// 消息内容，最长不超过2048个字节，超过将截断（支持id转译）
		Content string `json:"content"`
	}

	// SendImageRequest 发送图片消息的请求
	SendImageRequest struct {
		*SendRequestCommon
		Image ImageField `json:"image"`
	}
	// ImageField 图片消息参数
	ImageField struct {
		// 图片媒体文件id，可以调用上传临时素材接口获取
		MediaID string `json:"media_id"`
	}

	// SendVoiceRequest 发送语音消息的请求
	SendVoiceRequest struct {
		*SendRequestCommon
		Voice VoiceField `json:"voice"`
	}
	// VoiceField 语音消息参数
	VoiceField struct {
		// 语音文件id，可以调用上传临时素材接口获取
		MediaID string `json:"media_id"`
	}
)

// Send 发送应用消息
// @desc 实现企业微信发送应用消息接口：https://developer.work.weixin.qq.com/document/path/90236
func (r *Client) Send(apiName string, request interface{}) (*SendResponse, error) {
	// 获取accessToken
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 请求参数转 JSON 格式
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	// 发起http请求
	response, err := util.HTTPPost(fmt.Sprintf(sendURL, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	// 按照结构体解析返回值
	result := &SendResponse{}
	if err = util.DecodeWithError(response, result, apiName); err != nil {
		return nil, err
	}
	// 返回数据
	return result, nil
}

// SendText 发送文本消息
func (r *Client) SendText(request SendTextRequest) (*SendResponse, error) {
	// 发送文本消息MsgType参数固定为：text
	request.MsgType = "text"
	return r.Send("MessageSendText", request)
}

// SendImage 发送图片消息
func (r *Client) SendImage(request SendImageRequest) (*SendResponse, error) {
	// 发送图片消息MsgType参数固定为：image
	request.MsgType = "image"
	return r.Send("MessageSendImage", request)
}

// SendVoice 发送语音消息
func (r *Client) SendVoice(request SendVoiceRequest) (*SendResponse, error) {
	// 发送语音消息MsgType参数固定为：voice
	request.MsgType = "voice"
	return r.Send("MessageSendVoice", request)
}

// 以上实现了部分常用消息推送：SendText 发送文本消息、SendImage 发送图片消息、SendVoice 发送语音消息，
// 如需扩展其他消息类型，建议按照以上格式，扩展对应消息类型的参数即可
// 也可以直接使用Send方法，按照企业微信消息推送的接口文档传对应消息类型的参数来使用
