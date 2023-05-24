package addresslist

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// UserSimpleListURL 获取部门成员
	UserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d"
	// UserGetURL 读取成员
	UserGetURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
	// UserListIDURL 获取成员ID列表
	UserListIDURL = "https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token=%s"
)

type (
	// UserSimpleListResponse 获取部门成员响应
	UserSimpleListResponse struct {
		util.CommonError
		UserList []*UserList
	}
	// UserList 部门成员
	UserList struct {
		UserID     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserID string `json:"open_userid"`
	}
)

// UserSimpleList 获取部门成员
// @see https://developer.work.weixin.qq.com/document/path/90200
func (r *Client) UserSimpleList(departmentID int) ([]*UserList, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(UserSimpleListURL, accessToken, departmentID)); err != nil {
		return nil, err
	}
	result := &UserSimpleListResponse{}
	err = util.DecodeWithError(response, result, "UserSimpleList")
	if err != nil {
		return nil, err
	}
	return result.UserList, nil
}

// UserGetResponse 获取部门成员响应
type UserGetResponse struct {
	util.CommonError
	UserID         string   `json:"userid"`            // 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节；第三方应用返回的值为open_userid
	Name           string   `json:"name"`              // 成员名称；第三方不可获取，调用时返回userid以代替name；代开发自建应用需要管理员授权才返回；对于非第三方创建的成员，第三方通讯录应用也不可获取；未返回name的情况需要通过通讯录展示组件来展示名字
	Department     []int    `json:"department"`        // 成员所属部门id列表，仅返回该应用有查看权限的部门id；成员授权模式下，固定返回根部门id，即固定为1。对授权了“组织架构信息”权限的第三方应用，返回成员所属的全部部门id
	Order          []int    `json:"order"`             // 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。值范围是[0, 2^32)。成员授权模式下不返回该字段
	Position       string   `json:"position"`          // 职务信息；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Mobile         string   `json:"mobile"`            // 手机号码，代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Gender         string   `json:"gender"`            // 性别。0表示未定义，1表示男性，2表示女性。代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段。注：不可获取指返回值0
	Email          string   `json:"email"`             // 邮箱，代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	BizMail        string   `json:"biz_mail"`          // 企业邮箱，代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	IsLeaderInDept []int    `json:"is_leader_in_dept"` // 表示在所在的部门内是否为部门负责人，数量与department一致；第三方通讯录应用或者授权了“组织架构信息-应用可获取企业的部门组织架构信息-部门负责人”权限的第三方应用可获取；对于非第三方创建的成员，第三方通讯录应用不可获取；上游企业不可获取下游企业成员该字段
	DirectLeader   []string `json:"direct_leader"`     // 直属上级UserID，返回在应用可见范围内的直属上级列表，最多有五个直属上级；第三方通讯录应用或者授权了“组织架构信息-应用可获取可见范围内成员组织架构信息-直属上级”权限的第三方应用可获取；对于非第三方创建的成员，第三方通讯录应用不可获取；上游企业不可获取下游企业成员该字段；代开发自建应用不可获取该字段
	Avatar         string   `json:"avatar"`            // 头像url。 代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	ThumbAvatar    string   `json:"thumb_avatar"`      // 头像缩略图url。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Telephone      string   `json:"telephone"`         // 座机。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Alias          string   `json:"alias"`             // 别名；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Address        string   `json:"address"`           // 地址。代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	OpenUserid     string   `json:"open_userid"`       // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	MainDepartment int      `json:"main_department"`   // 主部门，仅当应用对主部门有查看权限时返回。
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"` // 扩展属性，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	Status           int    `json:"status"`            // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。 已激活代表已激活企业微信或已关注微信插件（原企业号）。未激活代表既未激活企业微信又未关注微信插件（原企业号）。
	QrCode           string `json:"qr_code"`           // 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；代开发自建应用需要管理员授权且成员oauth2授权获取；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	ExternalPosition string `json:"external_position"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"` // 成员对外属性，字段详情见对外属性；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
}

// UserGet 获取部门成员
// @see https://developer.work.weixin.qq.com/document/path/90196
func (r *Client) UserGet(UserID string) (*UserGetResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(UserGetURL, accessToken, UserID)); err != nil {
		return nil, err
	}
	result := &UserGetResponse{}
	err = util.DecodeWithError(response, result, "UserGet")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserListIDRequest 获取成员ID列表请求
type UserListIDRequest struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// UserListIDResponse 获取成员ID列表响应
type UserListIDResponse struct {
	util.CommonError
	NextCursor string      `json:"next_cursor"`
	DeptUser   []*DeptUser `json:"dept_user"`
}

// DeptUser 用户-部门关系
type DeptUser struct {
	UserID     string `json:"userid"`
	Department int    `json:"department"`
}

// UserListID 获取成员ID列表
// see https://developer.work.weixin.qq.com/document/path/96067
func (r *Client) UserListID(req *UserListIDRequest) (*UserListIDResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(UserListIDURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &UserListIDResponse{}
	if err = util.DecodeWithError(response, result, "UserListID"); err != nil {
		return nil, err
	}
	return result, nil
}
