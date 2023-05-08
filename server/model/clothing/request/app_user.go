package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserSearch struct {
	clothing.AppUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartWages     *float64   `json:"startWages" form:"startWages"`
	EndWages       *float64   `json:"endWages" form:"endWages"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type Login struct {
	Username string `json:"username"` // 用户名
	PhoneNum string `json:"phoneNum"` // 手机号
	Password string `json:"password"` // 密码
}

type Register struct {
	PhoneNum        string `json:"phoneNum"`        // 手机号
	Username        string `json:"username"`        // 用户名
	Password        string `json:"password"`        // 密码
	ConfirmPassword string `json:"confirmPassword"` // 确认密码
	VerifyCode      string `json:"verifyCode"`      // 验证码
}

type CheckExist struct {
	Content string `json:"content" form:"content"` // 检测内容
	Type    int    `json:"type" form:"type"`       // 1:用户名；2：手机号
}

type UserFilter struct {
	request.PageInfo
	CompanyID int    `json:"companyID" form:"companyID"` // 公司id
	TeamID    int    `json:"teamID" form:"teamID"`       // 组id
	Nickname  string `json:"nickname" form:"nickname"`   // 昵称
	Username  string `json:"username" form:"username"`   // 用户名
	PhoneNum  string `json:"phoneNum" form:"phoneNum"`   // 手机号
}
