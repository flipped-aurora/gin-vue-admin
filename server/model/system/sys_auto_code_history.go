package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
	"strings"
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
	Templates       map[string]string  `json:"template" gorm:"serializer:json;type:text;column:templates;comment:模板信息"`
	Injections      map[string]ast.Ast `json:"injections" gorm:"serializer:json;type:text;column:Injections;comment:注入路径"`
	Flag            int                `json:"flag" gorm:"column:flag;comment:[0:创建,1:回滚]"`
	ApiIDs          []uint             `json:"apiIDs" gorm:"serializer:json;column:api_ids;comment:api表注册内容"`
	MenuID          uint               `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
	AutoCodePackage SysAutoCodePackage `json:"autoCodePackage" gorm:"foreignKey:ID;references:PackageID"`
	PackageID       uint               `json:"packageID" gorm:"column:package_id;comment:包ID"`
}

func (s *SysAutoCodeHistory) BeforeCreate(db *gorm.DB) error {
	templates := make(map[string]string, len(s.Templates))
	for key, value := range s.Templates {
		{
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			hasServer := strings.Index(key, server)
			if hasServer != -1 {
				key = strings.TrimPrefix(key, server)
				keys := strings.Split(key, string(os.PathSeparator))
				key = path.Join(keys...)
			}
		} // key
		{
			web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
			hasWeb := strings.Index(value, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot()))
			if hasWeb != -1 {
				value = strings.TrimPrefix(value, web)
				values := strings.Split(key, string(os.PathSeparator))
				value = path.Join(values...)
			}
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			hasServer := strings.Index(key, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server))
			if hasServer != -1 {
				value = strings.TrimPrefix(value, server)
				values := strings.Split(key, string(os.PathSeparator))
				value = path.Join(values...)
			}
		} // value
		templates[key] = value
	}
	s.Templates = templates
	return nil
}

// AfterFirst 手动调用
func (s *SysAutoCodeHistory) AfterFirst() {
	templates := make(map[string]string, len(s.Templates))
	for key, value := range s.Templates {
		{
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			keys := strings.Split(key, "/")
			key = filepath.Join(keys...)
			key = strings.TrimPrefix(key, server)
		} // key
		{
			web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			values := strings.Split(value, "/")
			value = filepath.Join(values...)
			ext := path.Ext(value)
			switch ext {
			case ".js", ".vue":
				value = filepath.Join(web, value)
			case ".go":
				value = filepath.Join(server, value)
			}
		} // value
		templates[key] = value
	}
	s.Templates = templates
}

func (s *SysAutoCodeHistory) TableName() string {
	return "sys_auto_code_histories"
}
