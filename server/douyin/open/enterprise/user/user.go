package user

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

// LeadsLevelType 用户状态 .
type LeadsLevelType int64

// ActionType 分类 .
type ActionType int64

const (
	// 获取意向用户列表
	leadsUserListURL string = "https://open.douyin.com/enterprise/leads/user/list?access_token=%s&open_id=%s&cursor=%d&count=%d&start_time=%d&end_time=%d&leads_level=%d&action_type=%d"
	// 获取意向用户详情
	leadsUserDetailURL string = "https://open.douyin.com/enterprise/leads/user/detail?access_token=%s&open_id=%s&user_id=%s"
	// 获取意向用户互动记录
	leadsUserActionListURL string = "https://open.douyin.com/enterprise/leads/user/action/list?access_token=%s&open_id=%s&user_id=%s&count=%d&cursor=%s&action_type=%d"
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

// LeadsUser 意向用户详情.
type LeadsUser struct {
	util.CommonError

	Avatar     string    `json:"avatar"`
	City       string    `json:"city"`
	Age        int64     `json:"age"`
	Gender     int64     `json:"gender"`
	IsFollow   bool      `json:"is_follow"`
	LeadsLevel int64     `json:"leads_level"`
	Nickname   string    `json:"nickname"`
	OpenID     string    `json:"open_id"`
	TagList    []TagInfo `json:"tag_list"`
	Telephone  string    `json:"telephone"`
	Wechat     string    `json:"wechat"`
}

// LeadsUserList user list.
type LeadsUserList struct {
	util.CommonError

	Total   int64       `json:"total"`
	Cursor  int64       `json:"cursor"`
	HasMore bool        `json:"has_more"`
	Users   []LeadsUser `json:"users"`
}

type listLeadsUserRes struct {
	Message string        `json:"message"`
	Data    LeadsUserList `json:"data"`
}

// ListLeadsUser 获取意向用户列表
func (user *User) ListLeadsUser(openid string, cursor, count int64, startTime, endTime, leadsLevel int64, actionType int64) (res *LeadsUserList, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsUserListURL, accessToken, openid, cursor, count, startTime, endTime, leadsLevel, actionType)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result listLeadsUserRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListLeadsUser error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	res = &result.Data
	return
}

type leadsUserDetailRes struct {
	Message string    `json:"message"`
	Data    LeadsUser `json:"data"`
}

// LeadsUserDetail 获取意向用户详情
func (user *User) LeadsUserDetail(openid, userid string) (leadsUser *LeadsUser, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsUserDetailURL, accessToken, openid, userid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result leadsUserDetailRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("LeadsUserDetail error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	leadsUser = &result.Data
	return
}

// LeadsUserAction user action.
type LeadsUserAction struct {
	util.CommonError

	Cursor  int64 `json:"cursor"`
	HasMore bool  `json:"has_more"`
	List    []struct {
		ActionFlag   int64  `json:"action_flag"`
		ActionSource string `json:"action_source"`
		ActionType   int64  `json:"action_type"`
		CreateTime   int64  `json:"create_time"`
		UserID       int64  `json:"user_id"`
	} `json:"list"`
}

type listLeadsUserActionRes struct {
	Message string          `json:"message"`
	Data    LeadsUserAction `json:"data"`
}

// ListLeadsUserAction 获取意向用户互动记录
func (user *User) ListLeadsUserAction(openid, userid string, count int64, cursor string, actionType int64) (leadsUserAction *LeadsUserAction, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}

	uri := fmt.Sprintf(leadsUserActionListURL, accessToken, openid, userid, count, cursor, actionType)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result listLeadsUserActionRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListLeadsUserAction error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	leadsUserAction = &result.Data
	return
}
