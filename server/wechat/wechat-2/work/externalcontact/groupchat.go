package externalcontact

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// OpengIDToChatIDURL 客户群opengid转换URL
const OpengIDToChatIDURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/opengid_to_chatid"

type (
	//GroupChatListRequest 获取客户群列表的请求参数
	GroupChatListRequest struct {
		StatusFilter int         `json:"status_filter"` // 非必填	客户群跟进状态过滤。0 - 所有列表(即不过滤) 1 - 离职待继承 2 - 离职继承中 3 - 离职继承完成
		OwnerFilter  OwnerFilter `json:"owner_filter"`  //非必填	群主过滤。如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
		Cursor       string      `json:"cursor"`        //非必填	用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
		Limit        int         `json:"limit"`         //必填	分页，预期请求的数据量，取值范围 1 ~ 1000
	}

	//GroupChatList 客户群列表
	GroupChatList struct {
		ChatID string `json:"chat_id"`
		Status int    `json:"status"`
	}
	//GroupChatListResponse 获取客户群列表的返回值
	GroupChatListResponse struct {
		util.CommonError
		GroupChatList []GroupChatList `json:"group_chat_list"`
		NextCursor    string          `json:"next_cursor"` //游标
	}
)

// GetGroupChatList 获取客户群列表
// @see https://developer.work.weixin.qq.com/document/path/92120
func (r *Client) GetGroupChatList(req *GroupChatListRequest) (*GroupChatListResponse, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.PostJSON(fmt.Sprintf("%s/list?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &GroupChatListResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupChatList"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	//GroupChatDetailRequest 客户群详情 请求参数
	GroupChatDetailRequest struct {
		ChatID   string `json:"chat_id"`
		NeedName int    `json:"need_name"`
	}
	//Invitor 邀请者
	Invitor struct {
		UserID string `json:"userid"` //邀请者的userid
	}
	//GroupChatMember 群成员
	GroupChatMember struct {
		UserID        string  `json:"userid"`            //群成员id
		Type          int     `json:"type"`              //成员类型。 1 - 企业成员  2 - 外部联系人
		JoinTime      int     `json:"join_time"`         //入群时间
		JoinScene     int     `json:"join_scene"`        //入群方式 1 - 由群成员邀请入群（直接邀请入群） 2 - 由群成员邀请入群（通过邀请链接入群） 3 - 通过扫描群二维码入群
		Invitor       Invitor `json:"invitor,omitempty"` //邀请者。目前仅当是由本企业内部成员邀请入群时会返回该值
		GroupNickname string  `json:"group_nickname"`    //在群里的昵称
		Name          string  `json:"name"`              //名字。仅当 need_name = 1 时返回 如果是微信用户，则返回其在微信中设置的名字 如果是企业微信联系人，则返回其设置对外展示的别名或实名
		UnionID       string  `json:"unionid,omitempty"` //外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当群成员类型是微信用户（包括企业成员未添加好友），且企业绑定了微信开发者ID有此字段（查看绑定方法）。第三方不可获取，上游企业不可获取下游企业客户的unionid字段
	}
	//GroupChatAdmin 群管理员
	GroupChatAdmin struct {
		UserID string `json:"userid"` //群管理员userid
	}
	//GroupChat 客户群详情
	GroupChat struct {
		ChatID     string            `json:"chat_id"`     //客户群ID
		Name       string            `json:"name"`        //群名
		Owner      string            `json:"owner"`       //群主ID
		CreateTime int               `json:"create_time"` //群的创建时间
		Notice     string            `json:"notice"`      //群公告
		MemberList []GroupChatMember `json:"member_list"` //群成员列表
		AdminList  []GroupChatAdmin  `json:"admin_list"`  //群管理员列表
	}
	//GroupChatDetailResponse 客户群详情 返回值
	GroupChatDetailResponse struct {
		util.CommonError
		GroupChat GroupChat `json:"group_chat"` //客户群详情
	}
)

// GetGroupChatDetail 获取客户群详情
// @see https://developer.work.weixin.qq.com/document/path/92122
func (r *Client) GetGroupChatDetail(req *GroupChatDetailRequest) (*GroupChatDetailResponse, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.PostJSON(fmt.Sprintf("%s/get?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &GroupChatDetailResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupChatDetail"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	//OpengIDToChatIDRequest 客户群opengid转换 请求参数
	OpengIDToChatIDRequest struct {
		OpengID string `json:"opengid"`
	}
	//OpengIDToChatIDResponse 客户群opengid转换 返回值
	OpengIDToChatIDResponse struct {
		util.CommonError
		ChatID string `json:"chat_id"` //客户群ID
	}
)

// OpengIDToChatID 客户群opengid转换
// @see https://developer.work.weixin.qq.com/document/path/94828
func (r *Client) OpengIDToChatID(req *OpengIDToChatIDRequest) (*OpengIDToChatIDResponse, error) {
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", OpengIDToChatIDURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &OpengIDToChatIDResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupChatDetail"); err != nil {
		return nil, err
	}
	return result, nil
}
