package kf

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 发送消息
	sendMsgAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=%s"
)

// SendMsgSchema 发送消息响应内容
type SendMsgSchema struct {
	util.CommonError
	MsgID string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节, 字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// SendMsg 发送消息
// 当微信客户处于“新接入待处理”或“由智能助手接待”状态下，可调用该接口给用户发送消息。
// 注意仅当微信客户在主动发送消息给客服后的48小时内，企业可发送消息给客户，最多可发送5条消息；若用户继续发送消息，企业可再次下发消息。
// 支持发送消息类型：文本、图片、语音、视频、文件、图文、小程序、菜单消息、地理位置。
// 目前该接口允许下发消息条数和下发时限如下：
//
// 用户动作	允许下发条数限制	下发时限
// 用户发送消息	5条	48 小时
func (r *Client) SendMsg(options interface{}) (info SendMsgSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(sendMsgAddr, accessToken), options); err != nil {
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
