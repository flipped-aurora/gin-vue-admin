// Package user migrate 用于微信公众号账号迁移，获取openID变化
// 参考文档：https://kf.qq.com/faq/1901177NrqMr190117nqYJze.html
package user

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	changeOpenIDURL = "https://api.weixin.qq.com/cgi-bin/changeopenid"
)

// ChangeOpenIDResult OpenID迁移变化
type ChangeOpenIDResult struct {
	OriOpenID string `json:"ori_openid"`
	NewOpenID string `json:"new_openid"`
	ErrMsg    string `json:"err_msg,omitempty"`
}

// ChangeOpenIDResultList OpenID迁移变化列表
type ChangeOpenIDResultList struct {
	util.CommonError
	List []ChangeOpenIDResult `json:"result_list"`
}

// ListChangeOpenIDs 返回指定OpenID变化列表
// fromAppID 为老账号AppID
// openIDs 为老账号的openID，openIDs限100个以内
// AccessToken 为新账号的AccessToken
func (user *User) ListChangeOpenIDs(fromAppID string, openIDs ...string) (list *ChangeOpenIDResultList, err error) {
	list = &ChangeOpenIDResultList{}
	// list.List = make([]ChangeOpenIDResult, 0)
	if len(openIDs) > 100 {
		err = errors.New("openIDs length must be lt 100")
		return
	}

	if fromAppID == "" {
		err = errors.New("fromAppID is required")
		return
	}

	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", changeOpenIDURL, accessToken)
	var resp []byte
	var req struct {
		FromAppID  string   `json:"from_appid"`
		OpenidList []string `json:"openid_list"`
	}
	req.FromAppID = fromAppID
	req.OpenidList = append(req.OpenidList, openIDs...)
	resp, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithError(resp, list, "ListChangeOpenIDs")
	if err != nil {
		return
	}

	return
}

// ListAllChangeOpenIDs  返回所有用户OpenID列表
// fromAppID 为老账号AppID
// openIDs 为老账号的openID
// AccessToken 为新账号的AccessToken
func (user *User) ListAllChangeOpenIDs(fromAppID string, openIDs ...string) (list []ChangeOpenIDResult, err error) {
	list = make([]ChangeOpenIDResult, 0)
	chunks := util.SliceChunk(openIDs, 100)
	for _, chunk := range chunks {
		result, err := user.ListChangeOpenIDs(fromAppID, chunk...)
		if err != nil {
			return list, err
		}
		list = append(list, result.List...)
	}
	return
}
