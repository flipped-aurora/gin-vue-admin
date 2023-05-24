package kf

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 发送事件响应消息
	sendMsgOnEventAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg_on_event?access_token=%s"
)

// SendMsgOnEventSchema 发送事件响应消息
type SendMsgOnEventSchema struct {
	util.CommonError
	MsgID string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节, 字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// SendMsgOnEvent 发送事件响应消息
// 当特定的事件回调消息包含code字段，或通过接口变更到特定的会话状态，会返回code字段。
// 开发者可以此code为凭证，调用该接口给用户发送相应事件场景下的消息，如客服欢迎语、客服提示语和会话结束语等。
// 除”用户进入会话事件”以外，响应消息仅支持会话处于获取该code的会话状态时发送，如将会话转入待接入池时获得的code仅能在会话状态为”待接入池排队中“时发送。
//
// 目前支持的事件场景和相关约束如下：
//
// 事件场景	允许下发条数	code有效期	支持的消息类型	获取code途径
// 用户进入会话，用于发送客服欢迎语	1条	20秒	文本、菜单	事件回调
// 进入接待池，用于发送排队提示语等	1条	48小时	文本	转接会话接口
// 从接待池接入会话，用于发送非工作时间的提示语或超时未回复的提示语等	1条	48小时	文本	事件回调、转接会话接口
// 结束会话，用于发送结束会话提示语或满意度评价等	1条	20秒	文本、菜单	事件回调、转接会话接口
//
// 「进入会话事件」响应消息：
// 如果满足通过API下发欢迎语条件（条件为：1. 企业没有在管理端配置了原生欢迎语；2. 用户在过去48小时里未收过欢迎语，且未向该用户发过消息），则用户进入会话事件会额外返回一个welcome_code，开发者以此为凭据调用接口（填到该接口code参数），即可向客户发送客服欢迎语。
func (r *Client) SendMsgOnEvent(options interface{}) (info SendMsgOnEventSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(sendMsgOnEventAddr, accessToken), options); err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
