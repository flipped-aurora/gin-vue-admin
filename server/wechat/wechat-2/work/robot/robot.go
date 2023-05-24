package robot

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// WebhookSendURL 机器人发送群组消息
	WebhookSendURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

// RobotBroadcast 群机器人消息发送
// @see https://developer.work.weixin.qq.com/document/path/91770
func (r *Client) RobotBroadcast(webhookKey string, options interface{}) (info util.CommonError, err error) {
	var data []byte
	if data, err = util.PostJSON(fmt.Sprintf(WebhookSendURL, webhookKey), options); err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, err
	}
	return info, nil
}
