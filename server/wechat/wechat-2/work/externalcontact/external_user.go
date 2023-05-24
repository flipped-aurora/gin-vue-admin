package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// FetchExternalContactUserListURL 获取客户列表
	FetchExternalContactUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	// FetchExternalContactUserDetailURL 获取客户详情
	FetchExternalContactUserDetailURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	// FetchBatchExternalContactUserDetailURL 批量获取客户详情
	FetchBatchExternalContactUserDetailURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
	// UpdateUserRemarkURL 更新客户备注信息
	UpdateUserRemarkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
)

// ExternalUserListResponse 外部联系人列表响应
type ExternalUserListResponse struct {
	util.CommonError
	ExternalUserID []string `json:"external_userid"`
}

// GetExternalUserList 获取客户列表
// @see https://developer.work.weixin.qq.com/document/path/92113
func (r *Client) GetExternalUserList(userID string) ([]string, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.HTTPGet(fmt.Sprintf("%s?access_token=%v&userid=%v", FetchExternalContactUserListURL, accessToken, userID))
	if err != nil {
		return nil, err
	}
	var result ExternalUserListResponse
	err = util.DecodeWithError(response, &result, "GetExternalUserList")
	if err != nil {
		return nil, err
	}
	return result.ExternalUserID, nil
}

// ExternalUserDetailResponse 外部联系人详情响应
type ExternalUserDetailResponse struct {
	util.CommonError
	ExternalContact ExternalUser `json:"external_contact"`
	FollowUser      []FollowUser `json:"follow_user"`
	NextCursor      string       `json:"next_cursor"`
}

// ExternalUser 外部联系人
type ExternalUser struct {
	ExternalUserID  string `json:"external_userid"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Type            int64  `json:"type"`
	Gender          int64  `json:"gender"`
	UnionID         string `json:"unionid"`
	Position        string `json:"position"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	ExternalProfile string `json:"external_profile"`
}

// FollowUser 跟进用户（指企业内部用户）
type FollowUser struct {
	UserID         string        `json:"userid"`
	Remark         string        `json:"remark"`
	Description    string        `json:"description"`
	CreateTime     string        `json:"create_time"`
	Tags           []Tag         `json:"tags"`
	RemarkCorpName string        `json:"remark_corp_name"`
	RemarkMobiles  []string      `json:"remark_mobiles"`
	OperUserID     string        `json:"oper_userid"`
	AddWay         int64         `json:"add_way"`
	WeChatChannels WechatChannel `json:"wechat_channels"`
	State          string        `json:"state"`
}

// Tag 已绑定在外部联系人的标签
type Tag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	Type      int64  `json:"type"`
	TagID     string `json:"tag_id"`
}

// WechatChannel 视频号添加的场景
type WechatChannel struct {
	NickName string `json:"nickname"`
	Source   string `json:"source"`
}

// GetExternalUserDetail 获取外部联系人详情
// @see https://developer.work.weixin.qq.com/document/path/92114
func (r *Client) GetExternalUserDetail(externalUserID string, nextCursor ...string) (*ExternalUserDetailResponse, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	var cursor string
	if len(nextCursor) > 0 {
		cursor = nextCursor[0]
	}
	response, err = util.HTTPGet(fmt.Sprintf("%s?access_token=%v&external_userid=%v&cursor=%v", FetchExternalContactUserDetailURL, accessToken, externalUserID, cursor))
	if err != nil {
		return nil, err
	}
	result := &ExternalUserDetailResponse{}
	err = util.DecodeWithError(response, result, "get_external_user_detail")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchGetExternalUserDetailsRequest 批量获取外部联系人详情请求
type BatchGetExternalUserDetailsRequest struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
	Limit      int      `json:"limit,omitempty"`
}

// ExternalUserDetailListResponse 批量获取外部联系人详情响应
type ExternalUserDetailListResponse struct {
	util.CommonError
	ExternalContactList []ExternalUserForBatch `json:"external_contact_list"`
}

// ExternalUserForBatch 批量获取外部联系人客户列表
type ExternalUserForBatch struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// ExternalContact 批量获取外部联系人用户信息
type ExternalContact struct {
	ExternalUserID  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int64  `json:"type"`
	Gender          int64  `json:"gender"`
	UnionID         string `json:"unionid"`
	ExternalProfile string `json:"external_profile"`
}

// FollowInfo 批量获取外部联系人跟进人信息
type FollowInfo struct {
	UserID         string        `json:"userid"`
	Remark         string        `json:"remark"`
	Description    string        `json:"description"`
	CreateTime     int           `json:"create_time"`
	TagID          []string      `json:"tag_id"`
	RemarkCorpName string        `json:"remark_corp_name"`
	RemarkMobiles  []string      `json:"remark_mobiles"`
	OperUserID     string        `json:"oper_userid"`
	AddWay         int64         `json:"add_way"`
	WeChatChannels WechatChannel `json:"wechat_channels"`
}

// BatchGetExternalUserDetails 批量获取外部联系人详情
// @see https://developer.work.weixin.qq.com/document/path/92994
func (r *Client) BatchGetExternalUserDetails(request BatchGetExternalUserDetailsRequest) ([]ExternalUserForBatch, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err = util.HTTPPost(fmt.Sprintf("%s?access_token=%v", FetchBatchExternalContactUserDetailURL, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	var result ExternalUserDetailListResponse
	err = util.DecodeWithError(response, &result, "BatchGetExternalUserDetails")
	if err != nil {
		return nil, err
	}
	return result.ExternalContactList, nil
}

// UpdateUserRemarkRequest 修改客户备注信息请求体
type UpdateUserRemarkRequest struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid"`
}

// UpdateUserRemark 修改客户备注信息
// @see https://developer.work.weixin.qq.com/document/path/92115
func (r *Client) UpdateUserRemark(request UpdateUserRemarkRequest) error {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err = util.HTTPPost(fmt.Sprintf("%s?access_token=%v", UpdateUserRemarkURL, accessToken), string(jsonData))
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UpdateUserRemark")
}
