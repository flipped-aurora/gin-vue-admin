package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type SysApiToken struct {
	global.GVA_MODEL
	UserID      uint      `json:"userId" gorm:"comment:用户ID"`
	User        SysUser   `json:"user" form:"-" gorm:"foreignKey:UserID;"` // 所属用户(Preload 填充);form:"-" 阻止 gin 查询绑定递归进 SysUser
	AuthorityID uint      `json:"authorityId" gorm:"comment:角色ID"`
	Token       string    `json:"token" gorm:"type:text;comment:Token"`
	Status      bool      `json:"status" gorm:"default:true;comment:状态"` // true有效 false无效
	ExpiresAt   time.Time `json:"expiresAt" gorm:"comment:过期时间"`
	Remark      string    `json:"remark" gorm:"comment:备注"`
}
