package externalcontact

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// FetchFollowUserListURL 获取配置了客户联系功能的成员列表
	FetchFollowUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list"
)

// followerUserResponse 客户联系功能的成员列表响应
type followerUserResponse struct {
	util.CommonError
	FollowUser []string `json:"follow_user"`
}

// GetFollowUserList 获取配置了客户联系功能的成员列表
// @see https://developer.work.weixin.qq.com/document/path/92571
func (r *Client) GetFollowUserList() ([]string, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.HTTPGet(fmt.Sprintf("%s?access_token=%s", FetchFollowUserListURL, accessToken))
	if err != nil {
		return nil, err
	}
	var result followerUserResponse
	err = util.DecodeWithError(response, &result, "GetFollowUserList")
	if err != nil {
		return nil, err
	}
	return result.FollowUser, nil
}
