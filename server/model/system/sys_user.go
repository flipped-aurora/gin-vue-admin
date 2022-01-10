package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID" example:"c0523f1a-3358-495f-bdc6-20f173ffc513"`                                                                     // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名" example:"admin"`                                                                                                 // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码" example:"0a828a36-ad1c-43b0-8b6a-3bad4bbbaf80"`                                                                       // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称" example:"别名"`                                                                                        // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题" example:"dark"`                                                                                    // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像" example:"https://qmplusimg.henrongyi.top/gva_header.jpg"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色" example:"#fff"`                                                                                     // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色" example:"#1890ff"`                                                                             // 活跃颜色
	AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID" example:"888"`                                                                                   // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
