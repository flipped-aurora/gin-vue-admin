package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysAutoCodeHistory 自动迁移代码记录,用于回滚,重放使用
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	Table           string             `json:"tableName" gorm:"column:table_name;comment:表名"`
	Package         string             `json:"package" gorm:"column:package;comment:模块名/插件名"`
	Request         string             `json:"request" gorm:"type:text;column:request;comment:前端传入的结构化信息"`
	StructName      string             `json:"structName" gorm:"column:struct_name;comment:结构体名称"`
	BusinessDB      string             `json:"businessDb" gorm:"column:business_db;comment:业务库"`
	Description     string             `json:"description" gorm:"column:description;comment:Struct中文名称"`
	Enter           map[string]string  `json:"enter" gorm:"serializer:json;type:text;column:enter;comment:注入路径"`
	Template        map[string]string  `json:"template" gorm:"serializer:json;type:text;column:template;comment:模板信息"`
	Flag            int                `json:"flag" gorm:"column:flag;comment:[0:创建,1:回滚]"`
	ApiIDs          []uint             `json:"apiIDs" gorm:"serializer:json;column:api_ids;comment:api表注册内容"`
	MenuID          uint               `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
	AutoCodePackage SysAutoCodePackage `json:"autoCodePackage" gorm:"foreignKey:ID;references:PackageID"`
}

func (s *SysAutoCodeHistory) TableName() string {
	return "sys_auto_code_histories"
}
