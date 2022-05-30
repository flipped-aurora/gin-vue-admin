package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysBaseMenuBtn struct {
	global.GVA_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}

// TableName 系统基础菜单按钮表
func (s *SysBaseMenuBtn) TableName() string {
	return "sys_base_menu_btns"
}
