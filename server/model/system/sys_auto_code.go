package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AutoCodeStruct 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName         string   `json:"structName"`         // Struct名称
	TableName          string   `json:"tableName"`          // 表名
	PackageName        string   `json:"packageName"`        // 文件名称
	HumpPackageName    string   `json:"humpPackageName"`    // go文件名称
	Abbreviation       string   `json:"abbreviation"`       // Struct简称
	Description        string   `json:"description"`        // Struct中文名称
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoMoveFile       bool     `json:"autoMoveFile"`       // 是否自动移动文件
	Fields             []*Field `json:"fields"`
	DictTypes          []string `json:"-"`
	Package            string   `json:"package"`
	PackageT           string   `json:"-"`
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
}

var AutoMoveErr error = errors.New("创建代码成功并移动文件成功")

type SysAutoCode struct {
	global.GVA_MODEL
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:展示名"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}

type AutoPlugReq struct {
	PlugName    string `json:"plugName"` // 必然大写开头
	Snake       string `json:"snake"`    // 后端自动转为 snake
	RouterGroup string `json:"routerGroup"`
	HasGlobal   bool   `json:"hasGlobal"`
	HasRequest  bool   `json:"hasRequest"`
	HasResponse bool   `json:"hasResponse"`
	NeedModel   bool   `json:"needModel"`
	Global      []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
		Desc string `json:"desc"`
	} `json:"global"`
	Request []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
		Desc string `json:"desc"`
	} `json:"request"`
	Response []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
		Desc string `json:"desc"`
	} `json:"response"`
}
