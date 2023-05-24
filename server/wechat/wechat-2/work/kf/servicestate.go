package kf

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取会话状态
	serviceStateGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get?access_token=%s"
	// 变更会话状态
	serviceStateTransAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans?access_token=%s"
)

// ServiceStateGetOptions 获取会话状态请求参数
type ServiceStateGetOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
}

// ServiceStateGetSchema 获取会话状态响应内容
type ServiceStateGetSchema struct {
	util.CommonError
	ServiceState  int    `json:"service_state"`  // 当前的会话状态，状态定义参考概述中的表格
	ServiceUserID string `json:"service_userid"` // 接待人员的userid，仅当state=3时有效
}

// ServiceStateGet 获取会话状态
// 0	未处理	新会话接入（客户发信咨询）。可选择：1.直接用API自动回复消息。2.放进待接入池等待接待人员接待。3.指定接待人员（接待人员须处于“正在接待”中，下同）进行接待
// 1	由智能助手接待		可使用API回复消息。可选择转入待接入池或者指定接待人员处理
// 2	待接入池排队中		在待接入池中排队等待接待人员接入。可选择转为指定人员接待
// 3	由人工接待	人工接待中。可选择转接给其他接待人员处理或者结束会话
// 4	已结束	会话已经结束或未开始。不允许变更会话状态，客户重新发信咨询后会话状态变为“未处理”
// 注：一个微信用户向一个客服帐号发起咨询后，在48h内，或主动结束会话前（包括接待人员手动结束，或企业通过API结束会话），都算是一次会话
func (r *Client) ServiceStateGet(options ServiceStateGetOptions) (info ServiceStateGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(serviceStateGetAddr, accessToken), options); err != nil {
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

// ServiceStateTransOptions 变更会话状态请求参数
type ServiceStateTransOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
	ServiceState   int    `json:"service_state"`   // 变更的目标状态，状态定义和所允许的变更可参考概述中的流程图和表格
	ServicerUserID string `json:"servicer_userid"` // 接待人员的userid，当state=3时要求必填，接待人员须处于“正在接待”中
}

// ServiceStateTransSchema 变更会话状态响应内容
type ServiceStateTransSchema struct {
	util.CommonError
	MsgCode string `json:"msg_code"` // 用于发送响应事件消息的code，将会话初次变更为service_state为2和3时，返回回复语code，service_state为4时，返回结束语code。可用该code调用发送事件响应消息接口给客户发送事件响应消息
}

// ServiceStateTrans 变更会话状态
func (r *Client) ServiceStateTrans(options ServiceStateTransOptions) (info ServiceStateTransSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(serviceStateTransAddr, accessToken), options); err != nil {
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
