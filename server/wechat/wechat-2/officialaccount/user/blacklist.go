// Package user blacklist 公众号用户黑名单管理
// 参考文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
package user

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取公众号的黑名单列表
	getblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"
	// 拉黑用户
	batchblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"
	// 取消拉黑用户
	batchunblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s"
)

// GetBlackList 获取公众号的黑名单列表
// 该接口每次调用最多可拉取 1000 个OpenID，当列表数较多时，可以通过多次拉取的方式来满足需求。
// 参数 beginOpenid：当 begin_openid 为空时，默认从开头拉取。
func (user *User) GetBlackList(beginOpenid ...string) (userlist *OpenidList, err error) {
	if len(beginOpenid) > 1 {
		return nil, errors.New("参数 beginOpenid 错误：请传递 1 个openID，若需要从头开始拉取列表请留空。")
	}
	// 获取 AccessToken
	var accessToken string
	if accessToken, err = user.GetAccessToken(); err != nil {
		return
	}

	// 处理 request 内容
	request := map[string]string{"begin_openid": ""}
	if len(beginOpenid) == 1 {
		request["begin_openid"] = beginOpenid[0]
	}

	// 调用接口
	var resp []byte
	url := fmt.Sprintf(getblacklistURL, accessToken)
	if resp, err = util.PostJSON(url, &request); err != nil {
		return nil, err
	}

	// 处理返回
	userlist = &OpenidList{}
	if err = util.DecodeWithError(resp, userlist, "GetBlackList"); err != nil {
		return nil, err
	}

	return
}

// GetAllBlackList 获取公众号的所有黑名单列表
func (user *User) GetAllBlackList() (openIDList []string, err error) {
	var (
		beginOpenid string
		count       int
		userlist    *OpenidList
	)

	for {
		// 获取列表（每次1k条）
		if userlist, err = user.GetBlackList(beginOpenid); err != nil {
			return nil, err
		}
		openIDList = append(openIDList, userlist.Data.OpenIDs...) // 存储本次获得的OpenIDs
		count += userlist.Count                                   // 记录获得的总数量
		beginOpenid = userlist.NextOpenID                         // 记录下次循环的起始openID
		if count >= userlist.Total {
			break // 获得的数量=total，结束循环
		}
	}

	return
}

// BatchBlackList 拉黑用户
// 参数 openidList：需要拉入黑名单的用户的openid，每次拉黑最多允许20个
func (user *User) BatchBlackList(openidList ...string) (err error) {
	return user.batch(batchblacklistURL, "BatchBlackList", openidList...)
}

// BatchUnBlackList 取消拉黑用户
// 参数 openidList：需要取消拉入黑名单的用户的openid，每次拉黑最多允许20个
func (user *User) BatchUnBlackList(openidList ...string) (err error) {
	return user.batch(batchunblacklistURL, "BatchUnBlackList", openidList...)
}

// batch 公共方法
func (user *User) batch(url, apiName string, openidList ...string) (err error) {
	// 检查参数
	if len(openidList) == 0 || len(openidList) > 20 {
		return errors.New("参数 openidList 错误：每次操作黑名单用户数量为1-20个。")
	}

	// 获取 AccessToken
	var accessToken string
	if accessToken, err = user.GetAccessToken(); err != nil {
		return
	}

	// 处理 request 内容
	request := map[string][]string{"openid_list": openidList}

	// 调用接口
	var resp []byte
	url = fmt.Sprintf(url, accessToken)
	if resp, err = util.PostJSON(url, &request); err != nil {
		return
	}

	return util.DecodeWithCommonError(resp, apiName)
}
