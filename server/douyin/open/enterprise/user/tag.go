package user

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

const (
	leadsTagListURL       string = "https://open.douyin.com/enterprise/leads/tag/list?access_token=%s&open_id=%s&cursor=%d&count=%d"
	leadsTagUserListURL   string = "https://open.douyin.com/enterprise/leads/tag/user/list?access_token=%s&open_id=%s&cursor=%d&count=%d&tag_id=%s"
	leadsTagUserUpdateURL string = "https://open.douyin.com/enterprise/leads/tag/user/update?access_token=%s&open_id=%s"
	leadsTagCreateURL     string = "https://open.douyin.com/enterprise/leads/tag/create?access_token=%s&open_id=%s"
	leadsTagUpdateURL     string = "https://open.douyin.com/enterprise/leads/tag/update?access_token=%s&open_id=%s"
)

// TagInfo user tag.
type TagInfo struct {
	TagID   string `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// TagList .
type TagList struct {
	util.CommonError

	Cursor  int64     `json:"cursor"`
	HasMore bool      `json:"has_more"`
	List    []TagInfo `json:"list"`
}

type listLeadsTagRes struct {
	Message string  `json:"message"`
	Data    TagList `json:"data"`
}

// ListLeadsTag 获取标签列表.
func (user *User) ListLeadsTag(openid string, cursor, count int64) (res *TagList, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsTagListURL, accessToken, openid, cursor, count)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result listLeadsTagRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListLeadsTag error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	res = &result.Data
	return
}

// LeadsTagUserList .
type LeadsTagUserList struct {
	util.CommonError

	Cursor  int64    `json:"cursor"`
	HasMore bool     `json:"has_more"`
	List    []string `json:"list"`
}

type listLeadsTagUserRes struct {
	Message string           `json:"message"`
	Data    LeadsTagUserList `json:"data"`
}

// ListLeadsTagUser 获取打标签的用户列表.
func (user *User) ListLeadsTagUser(openid, tagid string, cursor, count int64) (leadsTagUser *LeadsTagUserList, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsTagUserListURL, accessToken, openid, cursor, count, tagid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result listLeadsTagUserRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("ListLeadsTag error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	leadsTagUser = &result.Data
	return
}

type createTagReq struct {
	TagName string `json:"tag_name"`
}

type createTagRes struct {
	Message string `json:"message"`
	Data    struct {
		util.CommonError

		TagID string `json:"tag_id"`
	} `json:"data"`
}

// CreateTag 创建标签.
func (user *User) CreateTag(openid, tagName string) (tagID string, err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsTagCreateURL, accessToken, openid)
	req := &createTagReq{
		TagName: tagName,
	}
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var result createTagRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("CreateTag error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	tagID = result.Data.TagID
	return
}

type updateTagReq struct {
	TagName string `json:"tag_name"`
	TagID   string `json:"tag_id"`
}

type updateTagRes struct {
	Message string `json:"message"`
	Data    struct {
		util.CommonError

		TagID string `json:"tag_id"`
	} `json:"data"`
}

// UpdateTag 更新标签.
func (user *User) UpdateTag(openid, tagName, tagID string) (err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsTagUpdateURL, accessToken, openid)
	req := &updateTagReq{
		TagID:   tagID,
		TagName: tagName,
	}
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var result updateTagRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("UpdateTag error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	return
}

// UpdateUserTagReq .
type UpdateUserTagReq struct {
	Bind   bool   `json:"bind"`
	TagID  string `json:"tag_id"`
	UserID string `json:"user_id"`
}

type updateUserTagRes struct {
	Message string `json:"message"`
	Data    struct {
		util.CommonError
	} `json:"data"`
}

// UpdateUserTag 给用户设置标签.
func (user *User) UpdateUserTag(openid string, req UpdateUserTagReq) (err error) {
	accessToken, err := user.GetAccessToken(openid)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(leadsTagUserUpdateURL, accessToken, openid)
	var response []byte
	response, err = util.PostJSON(uri, &req)
	if err != nil {
		return
	}
	var result updateUserTagRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("UpdateUserTag error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	return
}
