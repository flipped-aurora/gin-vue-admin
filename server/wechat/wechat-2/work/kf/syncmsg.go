package kf

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/kf/syncmsg"
)

const (
	// 获取消息
	syncMsgAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg?access_token=%s"
)

// SyncMsgOptions 获取消息查询参数
type SyncMsgOptions struct {
	Cursor string `json:"cursor"` // 上一次调用时返回的next_cursor，第一次拉取可以不填, 不多于64字节
	Token  string `json:"token"`  // 回调事件返回的token字段，10分钟内有效；可不填，如果不填接口有严格的频率限制, 不多于128字节
	Limit  uint   `json:"limit"`  // 期望请求的数据量，默认值和最大值都为1000, 注意：可能会出现返回条数少于limit的情况，需结合返回的has_more字段判断是否继续请求。
}

// SyncMsgSchema 获取消息查询响应内容
type syncMsgSchema struct {
	ErrCode    int32                    `json:"errcode"`     // 返回码
	ErrMsg     string                   `json:"errmsg"`      // 错误码描述
	NextCursor string                   `json:"next_cursor"` // 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。强烈建议对改该字段入库保存，每次请求读取带上，请求结束后更新。避免因意外丢，导致必须从头开始拉取，引起消息延迟。
	HasMore    uint32                   `json:"has_more"`    // 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList    []map[string]interface{} `json:"msg_list"`    // 消息列表
}

// SyncMsgSchema 获取消息查询响应内容
type SyncMsgSchema struct {
	ErrCode    int32             `json:"errcode"`     // 返回码
	ErrMsg     string            `json:"errmsg"`      // 错误码描述
	NextCursor string            `json:"next_cursor"` // 下次调用带上该值则从该key值往后拉，用于增量拉取
	HasMore    uint32            `json:"has_more"`    // 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList    []syncmsg.Message `json:"msg_list"`    // 消息列表
}

// SyncMsg 获取消息
func (r *Client) SyncMsg(options SyncMsgOptions) (info SyncMsgSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(syncMsgAddr, accessToken), options); err != nil {
		return
	}
	originInfo := syncMsgSchema{}
	if err = json.Unmarshal(data, &originInfo); err != nil {
		return
	}
	if originInfo.ErrCode != 0 {
		return info, errors.New(originInfo.ErrMsg)
	}
	msgList := make([]syncmsg.Message, 0)
	if len(originInfo.MsgList) > 0 {
		for _, msg := range originInfo.MsgList {
			newMsg := syncmsg.Message{}
			if val, ok := msg["msgid"].(string); ok {
				newMsg.MsgID = val
			}
			if val, ok := msg["open_kfid"].(string); ok {
				newMsg.OpenKFID = val
			}
			if val, ok := msg["external_userid"].(string); ok {
				newMsg.ExternalUserID = val
			}
			if val, ok := msg["send_time"].(float64); ok {
				newMsg.SendTime = uint64(val)
			}
			if val, ok := msg["origin"].(float64); ok {
				newMsg.Origin = uint32(val)
			}

			if val, ok := msg["msgtype"].(string); ok {
				newMsg.MsgType = val
			}
			if newMsg.MsgType == "event" {
				if event, ok := msg["event"].(map[string]interface{}); ok {
					if eType, ok := event["event_type"].(string); ok {
						newMsg.EventType = eType
					}
				}
			}
			originData, err := json.Marshal(msg)
			if err != nil {
				return info, err
			}
			newMsg.OriginData = originData
			msgList = append(msgList, newMsg)
		}
	}
	return SyncMsgSchema{
		ErrCode:    originInfo.ErrCode,
		ErrMsg:     originInfo.ErrMsg,
		NextCursor: originInfo.NextCursor,
		HasMore:    originInfo.HasMore,
		MsgList:    msgList,
	}, nil
}
