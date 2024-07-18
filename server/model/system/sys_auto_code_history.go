package system

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
	Templates       Templates          `json:"template" gorm:"serializer:json;type:text;column:templates;comment:模板信息"`
	Injections      map[string]string  `json:"injections" gorm:"serializer:json;type:text;column:Injections;comment:注入路径"`
	Flag            int                `json:"flag" gorm:"column:flag;comment:[0:创建,1:回滚]"`
	ApiIDs          []uint             `json:"apiIDs" gorm:"serializer:json;column:api_ids;comment:api表注册内容"`
	MenuID          uint               `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
	AutoCodePackage SysAutoCodePackage `json:"autoCodePackage" gorm:"foreignKey:ID;references:PackageID"`
	PackageID       uint               `json:"packageID" gorm:"column:package_id;comment:包ID"`
}

type Templates map[string]string

func (t *Templates) Scan(value any) error {
	if value == nil {
		return errors.New("value is nil")
	}
	values, ok := value.([]byte)
	if ok {
		err := json.Unmarshal(values, t)
		if err != nil {
			return err
		}
	} // 反序列化
	if t == nil {
		return errors.New("templates is nil")
	}
	templates := make(map[string]string, len(*t))
	for key, template := range *t {
		{
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			keys := strings.Split(key, "/")
			key = filepath.Join(keys...)
			key = strings.TrimPrefix(key, server)
		} // key
		{
			web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			values := strings.Split(template, "/")
			template = filepath.Join(values...)
			ext := path.Ext(template)
			switch ext {
			case ".js", ".vue":
				template = filepath.Join(web, template)
			case ".go":
				template = filepath.Join(server, template)
			}
		} // value
		templates[key] = template
	}
	*t = templates
	return nil
}

func (t *Templates) Value() (driver.Value, error) {
	if t == nil {
		return nil, errors.New("templates is nil")
	}
	templates := make(map[string]string, len(*t))
	for key, value := range *t {
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
			values := strings.Split(key, string(os.PathSeparator))
			value = path.Join(values...)
			templates[key] = value
			continue
		}
		hasServer := strings.Index(key, server)
		if hasServer != -1 {
			value = strings.TrimPrefix(value, server)
			values := strings.Split(key, string(os.PathSeparator))
			value = path.Join(values...)
			templates[key] = value
			continue
		}
	}
	bytes, err := json.Marshal(templates)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *SysAutoCodeHistory) TableName() string {
	return "sys_auto_code_histories"
}
