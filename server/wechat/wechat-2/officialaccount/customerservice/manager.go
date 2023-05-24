package customerservice

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// TypingStatus 输入状态类型
type TypingStatus string

const (
	customerServiceListURL       = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	customerServiceOnlineListURL = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	customerServiceAddURL        = "https://api.weixin.qq.com/customservice/kfaccount/add"
	customerServiceUpdateURL     = "https://api.weixin.qq.com/customservice/kfaccount/update"
	customerServiceDeleteURL     = "https://api.weixin.qq.com/customservice/kfaccount/del"
	customerServiceInviteURL     = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	customerServiceUploadHeadImg = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	customerServiceTypingURL     = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

const (
	// Typing 表示正在输入状态
	Typing TypingStatus = "Typing"
	// CancelTyping 表示取消正在输入状态
	CancelTyping TypingStatus = "CancelTyping"
)

// Manager 客服管理者，可以管理客服
type Manager struct {
	*context.Context
}

// NewCustomerServiceManager 实例化客服管理
func NewCustomerServiceManager(ctx *context.Context) *Manager {
	csm := new(Manager)
	csm.Context = ctx
	return csm
}

// KeFuInfo 客服基本信息
type KeFuInfo struct {
	KfAccount     string `json:"kf_account"`         // 完整客服帐号，格式为：帐号前缀@公众号微信号
	KfNick        string `json:"kf_nick"`            // 客服昵称
	KfID          int    `json:"kf_id"`              // 客服编号
	KfHeadImgURL  string `json:"kf_headimgurl"`      // 客服头像
	KfWX          string `json:"kf_wx"`              // 如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
	InviteWX      string `json:"invite_wx"`          // 如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
	InviteExpTime int    `json:"invite_expire_time"` // 如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
	InviteStatus  string `json:"invite_status"`      // 邀请的状态，有等待确认“waiting”，被拒绝“rejected”， 过期“expired”
}

type resKeFuList struct {
	util.CommonError
	KfList []*KeFuInfo `json:"kf_list"`
}

// List 获取所有客服基本信息
func (csm *Manager) List() (customerServiceList []*KeFuInfo, err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resKeFuList
	err = util.DecodeWithError(response, &res, "ListCustomerService")
	if err != nil {
		return
	}
	customerServiceList = res.KfList
	return
}

// KeFuOnlineInfo 客服在线信息
type KeFuOnlineInfo struct {
	KfAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KfID         int    `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

type resKeFuOnlineList struct {
	util.CommonError
	KfOnlineList []*KeFuOnlineInfo `json:"kf_online_list"`
}

// OnlineList 获取在线客服列表
func (csm *Manager) OnlineList() (customerServiceOnlineList []*KeFuOnlineInfo, err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceOnlineListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resKeFuOnlineList
	err = util.DecodeWithError(response, &res, "ListOnlineCustomerService")
	if err != nil {
		return
	}
	customerServiceOnlineList = res.KfOnlineList
	return
}

// Add 添加客服账号
func (csm *Manager) Add(kfAccount, nickName string) (err error) {
	// kfAccount：完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
	// nickName：客服昵称，最长16个字
	// 参数此处均不做校验
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceAddURL, accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		NickName  string `json:"nickname"`
	}{
		KfAccount: kfAccount,
		NickName:  nickName,
	}
	var response []byte
	response, err = util.PostJSON(uri, data)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "AddCustomerService")
	return
}

// Update 修改客服账号
func (csm *Manager) Update(kfAccount, nickName string) (err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceUpdateURL, accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		NickName  string `json:"nickname"`
	}{
		KfAccount: kfAccount,
		NickName:  nickName,
	}
	var response []byte
	response, err = util.PostJSON(uri, data)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "UpdateCustomerService")
	return
}

// Delete 删除客服帐号
func (csm *Manager) Delete(kfAccount string) (err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceDeleteURL, accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
	}{
		KfAccount: kfAccount,
	}
	var response []byte
	response, err = util.PostJSON(uri, data)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "DeleteCustomerService")
	return
}

// InviteBind 邀请绑定客服帐号和微信号
func (csm *Manager) InviteBind(kfAccount, inviteWX string) (err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceInviteURL, accessToken)
	data := struct {
		KfAccount string `json:"kf_account"`
		InviteWX  string `json:"invite_wx"`
	}{
		KfAccount: kfAccount,
		InviteWX:  inviteWX,
	}
	var response []byte
	response, err = util.PostJSON(uri, data)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "InviteBindCustomerService")
	return
}

// UploadHeadImg 上传客服头像
func (csm *Manager) UploadHeadImg(kfAccount, fileName string) (err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&kf_account=%s", customerServiceUploadHeadImg, accessToken, kfAccount)
	var response []byte
	response, err = util.PostFile("media", fileName, uri)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "UploadCustomerServiceHeadImg")
	return
}

// SendTypingStatus 下发客服输入状态给用户
func (csm *Manager) SendTypingStatus(openid string, cmd TypingStatus) (err error) {
	var accessToken string
	accessToken, err = csm.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customerServiceTypingURL, accessToken)
	data := struct {
		ToUser  string `json:"touser"`
		Command string `json:"command"`
	}{
		ToUser:  openid,
		Command: string(cmd),
	}
	var response []byte
	response, err = util.PostJSON(uri, data)
	if err != nil {
		return
	}
	err = util.DecodeWithCommonError(response, "SendTypingStatus")
	return
}
