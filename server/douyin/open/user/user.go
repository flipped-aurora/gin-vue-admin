package user

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

const (
	userInfoURL      string = "https://open.douyin.com/oauth/oauth/userinfo?access_token=%s&open_id=%s"
	fansListURL      string = "https://open.douyin.com/fans/list?access_token=%s&open_id=%s&cursor=%d&count=%d"
	followingListURL string = "https://open.douyin.com/following/list?access_token=%s&open_id=%s&cursor=%d&count=%d"
)

// User 用户
type User struct {
	*context.Context
}

// NewUser .
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

// Info .
type Info struct {
	util.CommonError

	Avatar       string `json:"avatar"`
	City         string `json:"city"`
	Country      string `json:"country"`
	EAccountRole string `json:"e_account_role"`
	Gender       int32  `json:"gender"`
	Nickname     string `json:"nickname"`
	OpenID       string `json:"open_id"`
	Province     string `json:"province"`
	Unionid      string `json:"union_id"`
}

type userInfoRes struct {
	Message string `json:"message"`
	Data    Info   `json:"data"`
}

// GetUserInfo 获取用户信息.
func (user *User) GetUserInfo(openid string) (userInfo *Info, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(userInfoURL, accessToken, openid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result userInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	userInfo = &result.Data
	return
}

// ListInfo user list.
type ListInfo struct {
	util.CommonError

	Total   int64  `json:"total"`
	Cursor  int64  `json:"cursor"`
	HasMore bool   `json:"has_more"`
	List    []Info `json:"list"`
}

type infoRes struct {
	Message string   `json:"message"`
	Data    ListInfo `json:"data"`
}

// ListFans 粉丝列表
func (user *User) ListFans(openid string, cursor, count int64) (fans *ListInfo, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(fansListURL, accessToken, openid, cursor, count)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result infoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListFans error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	fans = &result.Data
	return
}

// ListFollowing 关注列表
func (user *User) ListFollowing(openid string, cursor, count int64) (following *ListInfo, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(followingListURL, accessToken, openid, cursor, count)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result infoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListFollowing error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	following = &result.Data
	return
}
