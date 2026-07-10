package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
	DeletedAt     *time.Time     `sql:"index"`
	AuthorityId   uint           `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      *uint          `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	Children      []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus  []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users         []SysUser      `json:"-" gorm:"many2many:sys_user_authority;"`
	DataScope     int            `json:"dataScope" gorm:"default:1;comment:数据范围 1全部 2本部门及子级 3本部门 4仅本人"` // 数据范围(数据权限)
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`           // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
