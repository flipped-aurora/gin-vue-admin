package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// SysAutoCodeHistory 自动迁移代码记录,用于回滚,重放使用
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	Table            string             `json:"tableName" gorm:"column:table_name;comment:表名"`
	Package          string             `json:"package" gorm:"column:package;comment:模块名/插件名"`
	Request          string             `json:"request" gorm:"type:text;column:request;comment:前端传入的结构化信息"`
	StructName       string             `json:"structName" gorm:"column:struct_name;comment:结构体名称"`
	Abbreviation     string             `json:"abbreviation" gorm:"column:abbreviation;comment:结构体名称缩写"`
	BusinessDB       string             `json:"businessDb" gorm:"column:business_db;comment:业务库"`
	Description      string             `json:"description" gorm:"column:description;comment:Struct中文名称"`
	Templates        map[string]string  `json:"template" gorm:"serializer:json;type:text;column:templates;comment:模板信息"`
	Injections       map[string]string  `json:"injections" gorm:"serializer:json;type:text;column:Injections;comment:注入路径"`
	Flag             int                `json:"flag" gorm:"column:flag;comment:[0:创建,1:回滚]"`
	ApiIDs           []uint             `json:"apiIDs" gorm:"serializer:json;column:api_ids;comment:api表注册内容"`
	MenuID           uint               `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
	ExportTemplateID uint               `json:"exportTemplateID" gorm:"column:export_template_id;comment:导出模板ID"`
	AutoCodePackage  SysAutoCodePackage `json:"autoCodePackage" gorm:"foreignKey:ID;references:PackageID"`
	PackageID        uint               `json:"packageID" gorm:"column:package_id;comment:包ID"`
}

func (s *SysAutoCodeHistory) BeforeCreate(db *gorm.DB) error {
	templates := make(map[string]string, len(s.Templates))
	for key, value := range s.Templates {
		server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
		{
			hasServer := strings.Index(key, server)
			if hasServer != -1 {
				key = strings.TrimPrefix(key, server)
				keys := strings.Split(key, string(os.PathSeparator))
				key = path.Join(keys...)
			}
		} // key
		web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
		hasWeb := strings.Index(value, web)
		if hasWeb != -1 {
			value = strings.TrimPrefix(value, web)
			values := strings.Split(value, string(os.PathSeparator))
			value = path.Join(values...)
			templates[key] = value
			continue
		}
		hasServer := strings.Index(value, server)
		if hasServer != -1 {
			value = strings.TrimPrefix(value, server)
			values := strings.Split(value, string(os.PathSeparator))
			value = path.Join(values...)
			templates[key] = value
			continue
		}
	}
	s.Templates = templates
	return nil
}

func (s *SysAutoCodeHistory) TableName() string {
	return "sys_auto_code_histories"
}
