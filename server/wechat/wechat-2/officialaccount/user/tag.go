package user

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	tagCreateURL         = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	tagGetURL            = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
	tagUpdateURL         = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"
	tagDeleteURL         = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"
	tagUserListURL       = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"
	tagBatchtaggingURL   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"
	tagBatchuntaggingURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s"
	tagUserTidListURL    = "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s"
)

// TagInfo 标签信息
type TagInfo struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// TagOpenIDList 标签用户列表
type TagOpenIDList struct {
	Count int `json:"count"`
	Data  struct {
		OpenIDs []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

// CreateTag 创建标签
func (user *User) CreateTag(tagName string) (tagInfo *TagInfo, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(tagCreateURL, accessToken)
	var response []byte
	var request struct {
		Tag struct {
			Name string `json:"name"`
		} `json:"tag"`
	}
	request.Tag.Name = tagName
	response, err = util.PostJSON(uri, &request)
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
		Tag *TagInfo `json:"tag"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateTag Error , errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}
	return result.Tag, nil
}

// DeleteTag  删除标签
func (user *User) DeleteTag(tagID int32) (err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(tagDeleteURL, accessToken)
	var request struct {
		Tag struct {
			ID int32 `json:"id"`
		} `json:"tag"`
	}
	request.Tag.ID = tagID
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(resp, "DeleteTag")
}

// UpdateTag  编辑标签
func (user *User) UpdateTag(tagID int32, tagName string) (err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(tagUpdateURL, accessToken)
	var request struct {
		Tag struct {
			ID   int32  `json:"id"`
			Name string `json:"name"`
		} `json:"tag"`
	}
	request.Tag.ID = tagID
	request.Tag.Name = tagName
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(resp, "UpdateTag")
}

// GetTag 获取公众号已创建的标签
func (user *User) GetTag() (tags []*TagInfo, err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(tagGetURL, accessToken)
	response, err := util.HTTPGet(url)
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
		Tags []*TagInfo `json:"tags"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return result.Tags, nil
}

// OpenIDListByTag 获取标签下粉丝列表
func (user *User) OpenIDListByTag(tagID int32, nextOpenID ...string) (userList *TagOpenIDList, err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(tagUserListURL, accessToken)
	var request = struct {
		ID     int32  `json:"tagid"`
		OpenID string `json:"next_openid"`
	}{
		ID: tagID,
	}
	if len(nextOpenID) > 0 {
		request.OpenID = nextOpenID[0]
	}
	response, err := util.PostJSON(url, &request)
	if err != nil {
		return nil, err
	}
	userList = new(TagOpenIDList)
	err = json.Unmarshal(response, &userList)
	if err != nil {
		return
	}
	return
}

// BatchTag 批量为用户打标签
func (user *User) BatchTag(openIDList []string, tagID int32) (err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}
	if len(openIDList) == 0 {
		return
	}
	var request = struct {
		OpenIDList []string `json:"openid_list"`
		TagID      int32    `json:"tagid"`
	}{
		OpenIDList: openIDList,
		TagID:      tagID,
	}
	url := fmt.Sprintf(tagBatchtaggingURL, accessToken)
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(resp, "BatchTag")
}

// BatchUntag 批量为用户取消标签
func (user *User) BatchUntag(openIDList []string, tagID int32) (err error) {
	if len(openIDList) == 0 {
		return
	}
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(tagBatchuntaggingURL, accessToken)
	var request = struct {
		OpenIDList []string `json:"openid_list"`
		TagID      int32    `json:"tagid"`
	}{
		OpenIDList: openIDList,
		TagID:      tagID,
	}
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(resp, "BatchUntag")
}

// UserTidList 获取用户身上的标签列表
func (user *User) UserTidList(openID string) (tagIDList []int32, err error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(tagUserTidListURL, accessToken)
	var request = struct {
		OpenID string `json:"openid"`
		TagID  int32  `json:"tagid"`
	}{
		OpenID: openID,
	}
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
		TagIDList []int32 `json:"tagid_list"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UserTidList Error , errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}
	return result.TagIDList, nil
}
