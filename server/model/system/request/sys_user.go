package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Register User register structure
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

// Login User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type ResetPassword struct {
	ID       uint   `json:"ID" form:"ID"`
	Password string `json:"password" form:"password" gorm:"comment:用户登录密码"` // 用户登录密码
}

// SetUserAuth Modify user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// SetUserAuthorities Modify user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

// SetUserDepartments 设置用户归属部门(多部门)与主部门
type SetUserDepartments struct {
	ID            uint   `json:"ID"`            // 用户ID
	DeptIds       []uint `json:"deptIds"`       // 归属部门ID集合(数据可见范围)
	PrimaryDeptId uint   `json:"primaryDeptId"` // 主部门ID(数据归属/盖章),为空时取集合首个
}

// SetUserPositions 设置用户岗位(多岗位)
type SetUserPositions struct {
	ID          uint   `json:"ID"`          // 用户ID
	PositionIds []uint `json:"positionIds"` // 岗位ID集合
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

type GetUserList struct {
	common.PageInfo
	Username string `json:"username" form:"username"`
	NickName string `json:"nickName" form:"nickName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}

// SetRoleUsers 通过角色ID全量覆盖关联用户列表
type SetRoleUsers struct {
	AuthorityId uint   `json:"authorityId" form:"authorityId"` // 角色ID
	UserIds     []uint `json:"userIds" form:"userIds"`         // 用户ID列表
}

// SetDataScope 设置角色数据权限
type SetDataScope struct {
	AuthorityId uint   `json:"authorityId" form:"authorityId"` // 角色ID
	DataScope   int    `json:"dataScope" form:"dataScope"`     // 数据权限 1全部 2本部门及以下 3本部门 4仅本人 5自定义部门
	DeptIds     []uint `json:"deptIds" form:"deptIds"`         // 自定义部门集(档位为5时生效)
}
