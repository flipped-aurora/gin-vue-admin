package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysCli struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"column:name;size:128;not null;uniqueIndex;comment:CLI唯一标识"`
	Command     string `json:"command" gorm:"column:command;size:128;not null;default:'';comment:CLI主命令"`
	DisplayName string `json:"displayName" gorm:"column:display_name;size:128;not null;comment:CLI展示名称"`
	Version     string `json:"version" gorm:"column:version;size:64;not null;default:v1;comment:CLI版本"`
	Description string `json:"description" gorm:"column:description;type:text;comment:CLI描述"`
	Status      string `json:"status" gorm:"column:status;size:32;not null;default:enabled;comment:CLI状态"`
	AuthMode    string `json:"authMode" gorm:"column:auth_mode;size:32;not null;default:jwt;comment:认证方式"`
}

func (SysCli) TableName() string {
	return "sys_clis"
}

type SysCliApi struct {
	global.GVA_MODEL
	CliID       uint   `json:"cliId" gorm:"column:cli_id;not null;uniqueIndex:idx_cli_api;comment:CLI ID"`
	ApiID       uint   `json:"apiId" gorm:"column:api_id;not null;uniqueIndex:idx_cli_api;comment:API ID"`
	CommandName string `json:"commandName" gorm:"column:command_name;size:128;comment:命令名覆盖"`
	Enabled     bool   `json:"enabled" gorm:"column:enabled;not null;default:true;comment:是否启用"`
	Sort        int    `json:"sort" gorm:"column:sort;not null;default:0;comment:排序"`
}

func (SysCliApi) TableName() string {
	return "sys_cli_apis"
}
