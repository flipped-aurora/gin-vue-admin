package kf

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取配置的专员与客户群
	upgradeServiceConfigAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/get_upgrade_service_config?access_token=%s"
	// 为客户升级为专员或客户群服务
	upgradeService = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service?access_token=%s"
	// 为客户取消推荐
	upgradeServiceCancel = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/cancel_upgrade_service?access_token=%s"
)

// UpgradeServiceConfigSchema 获取配置的专员与客户群
type UpgradeServiceConfigSchema struct {
	util.CommonError
	MemberRange struct {
		UserIDList       []string `json:"userid_list"`        // 专员userid列表
		DepartmentIDList []string `json:"department_id_list"` // 专员部门列表
	} `json:"member_range"` // 专员服务配置范围
	GroupChatRange struct {
		ChatIDList []string `json:"chat_id_list"` // 客户群列表
	} `json:"groupchat_range"` // 客户群配置范围
}

// UpgradeServiceConfig 获取配置的专员与客户群
func (r *Client) UpgradeServiceConfig() (info UpgradeServiceConfigSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.HTTPGet(fmt.Sprintf(upgradeServiceConfigAddr, accessToken)); err != nil {
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

// UpgradeServiceOptions 为客户升级为专员或客户群服务请求参数
type UpgradeServiceOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
	Type           int    `json:"type"`            // 表示是升级到专员服务还是客户群服务。1:专员服务。2:客户群服务
	Member         struct {
		UserID  string `json:"userid"`  // 服务专员的userid
		Wording string `json:"wording"` // 推荐语
	} `json:"member"` // 推荐的服务专员，type等于1时有效
	GroupChat struct {
		ChatID  string `json:"chat_id"` // 客户群id
		Wording string `json:"wording"` // 推荐语
	} `json:"groupchat"` // 推荐的客户群，type等于2时有效
}

// UpgradeService 为客户升级为专员或客户群服务
func (r *Client) UpgradeService(options UpgradeServiceOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(upgradeService, accessToken), options); err != nil {
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

// UpgradeMemberServiceOptions 为客户升级为专员服务请求参数
type UpgradeMemberServiceOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
	Type           int    `json:"type"`            // 表示是升级到专员服务还是客户群服务。1:专员服务
	Member         struct {
		UserID  string `json:"userid"`  // 服务专员的userid
		Wording string `json:"wording"` // 推荐语
	} `json:"member"` // 推荐的服务专员，type等于1时有效
}

// UpgradeMemberService 为客户升级为专员服务
func (r *Client) UpgradeMemberService(options UpgradeMemberServiceOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(upgradeService, accessToken), options); err != nil {
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

// UpgradeServiceGroupChatOptions 为客户升级为客户群服务请求参数
type UpgradeServiceGroupChatOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
	Type           int    `json:"type"`            // 表示是升级到专员服务还是客户群服务。2:客户群服务
	GroupChat      struct {
		ChatID  string `json:"chat_id"` // 客户群id
		Wording string `json:"wording"` // 推荐语
	} `json:"groupchat"` // 推荐的客户群，type等于2时有效
}

// UpgradeGroupChatService 为客户升级为客户群服务
func (r *Client) UpgradeGroupChatService(options UpgradeServiceGroupChatOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(upgradeService, accessToken), options)
	if err != nil {
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

// UpgradeServiceCancelOptions 为客户取消推荐
type UpgradeServiceCancelOptions struct {
	OpenKFID       string `json:"open_kfid"`       // 客服帐号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
}

// UpgradeServiceCancel 为客户取消推荐
func (r *Client) UpgradeServiceCancel(options UpgradeServiceCancelOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(upgradeServiceCancel, accessToken), options); err != nil {
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
