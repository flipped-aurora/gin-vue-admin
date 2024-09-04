package biz_apphub

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type BizApparmor struct {
	global.GVA_MODEL
	Host       string `json:"host" gorm:"column:host"`
	ToolPath   string `json:"tool_path" gorm:"column:tool_path"`
	ToolName   string `json:"tool_name" gorm:"column:tool_name"`
	ConfigName string `json:"config_name" gorm:"column:config_name"`
	Config     string `json:"config" gorm:"column:config"`
	Status     string `json:"status" gorm:"column:status"`
}

func (a *BizApparmor) TableName() string {
	return "biz_apparmor"
}
