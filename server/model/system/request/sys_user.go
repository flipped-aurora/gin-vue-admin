package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Register Admin User register structure
type Register struct {
	Username     string `json:"userName" example:"用户名"`
	Password     string `json:"passWord" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
}

// Register Regular User register structure
type RegisterRegular struct {
	Username  string `json:"userName" example:"用户名"`
	Password  string `json:"passWord" example:"密码"`
	NickName  string `json:"nickName" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	//AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	//Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	//AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	Province     string `json:"province" example:"省份"`
	City         string `json:"city" example:"城市"`
	Organization string `json:"organization" example:"单位"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	CyberCoin    int                   `json:"cyberCoin" gorm:"default:0;comment:用户积分"`                                              // 用户积分
	Province     string                `json:"province" gorm:"default:0;comment:用户所在省份"`                                             // 用户所在省份
	City         string                `json:"city" gorm:"default:0;comment:用户所在城市"`                                                 // 用户所在城市
	Organization string                `json:"organization" gorm:"default:0;comment:用户所在单位"`                                         // 用户所在单位
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

type ChangeSelfInfo struct {
	ID       uint   `gorm:"primarykey"`                                // 主键ID
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Phone    string `json:"phone"  gorm:"comment:用户手机号"`               // 用户手机号
	//AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email string `json:"email"  gorm:"comment:用户邮箱"` // 用户邮箱
	//CyberCoin    int                   `json:"cyberCoin" gorm:"default:0;comment:用户积分"`                                              // 用户积分
	Province     string `json:"province" gorm:"default:0;comment:用户所在省份"`                                             // 用户所在省份
	City         string `json:"city" gorm:"default:0;comment:用户所在城市"`                                                 // 用户所在城市
	Organization string `json:"organization" gorm:"default:0;comment:用户所在单位"`                                         // 用户所在单位
	HeaderImg    string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	//Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	//Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

type ChangeUserCyberCoin struct {
	ID        uint `gorm:"primarykey"`                              // 主键ID
	CyberCoin int  `json:"cyberCoin" gorm:"default:0;comment:用户积分"` // 用户积分
}
