package system

type SysAuthorityBtn struct {
	AuthorityId      string         `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}

// TableName 系统角色按钮表
func (s *SysAuthorityBtn) TableName() string {
	return "sys_authority_btns"
}
