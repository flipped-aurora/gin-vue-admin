package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// OnlineUserSearch 在线用户搜索请求
type OnlineUserSearch struct {
	request.PageInfo
	Username string `json:"username" form:"username"` // 用户名搜索
	NickName string `json:"nickName" form:"nickName"` // 昵称搜索
	RoleId   uint   `json:"roleId" form:"roleId"`     // 角色ID过滤
}
