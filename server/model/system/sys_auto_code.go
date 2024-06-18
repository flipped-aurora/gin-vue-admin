package system

import (
	"go/token"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AutoCodeStruct 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName          string                 `json:"structName"`          // Struct名称
	TableName           string                 `json:"tableName"`           // 表名
	PackageName         string                 `json:"packageName"`         // 文件名称
	HumpPackageName     string                 `json:"humpPackageName"`     // go文件名称
	Abbreviation        string                 `json:"abbreviation"`        // Struct简称
	Description         string                 `json:"description"`         // Struct中文名称
	AutoCreateApiToSql  bool                   `json:"autoCreateApiToSql"`  // 是否自动创建api
	AutoCreateMenuToSql bool                   `json:"autoCreateMenuToSql"` // 是否自动创建menu
	AutoMigrate         bool                   `json:"autoMigrate"`         // 是否自动迁移表结构
	AutoCreateResource  bool                   `json:"autoCreateResource"`  // 是否自动创建资源标识
	BusinessDB          string                 `json:"businessDB"`          // 业务数据库
	GvaModel            bool                   `json:"gvaModel"`            // 是否使用gva默认Model
	Fields              []*Field               `json:"fields"`
	PrimaryField        *Field                 `json:"primaryField"`
	HasTimer            bool                   `json:"-"`
	HasSearchTimer      bool                   `json:"-"`
	DictTypes           []string               `json:"-"`
	Package             string                 `json:"package"`
	PackageT            string                 `json:"-"`
	NeedSort            bool                   `json:"-"`
	HasPic              bool                   `json:"-"`
	HasRichText         bool                   `json:"-"`
	HasFile             bool                   `json:"-"`
	NeedJSON            bool                   `json:"-"`
	FrontFields         []*Field               `json:"-"`
	HasDataSource       bool                   `json:"-"`
	DataSourceMap       map[string]*DataSource `json:"-"`
}

type DataSource struct {
	Association int    `json:"association"` // 关联关系 1 一对一 2 一对多
	Table       string `json:"table"`
	Label       string `json:"label"`
	Value       string `json:"value"`
}

func (a *AutoCodeStruct) Pretreatment() {
	a.KeyWord()
	a.SuffixTest()
}

// KeyWord 是go关键字的处理加上 _ ，防止编译报错
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) KeyWord() {
	if token.IsKeyword(a.Abbreviation) {
		a.Abbreviation = a.Abbreviation + "_"
	}
}

// SuffixTest 处理_test 后缀
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) SuffixTest() {
	if strings.HasSuffix(a.HumpPackageName, "test") {
		a.HumpPackageName = a.HumpPackageName + "_"
	}
}

type Field struct {
	FieldName       string      `json:"fieldName"`       // Field名
	FieldDesc       string      `json:"fieldDesc"`       // 中文名
	FieldType       string      `json:"fieldType"`       // Field数据类型
	FieldJson       string      `json:"fieldJson"`       // FieldJson
	DataTypeLong    string      `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string      `json:"comment"`         // 数据库字段描述
	ColumnName      string      `json:"columnName"`      // 数据库字段
	FieldSearchType string      `json:"fieldSearchType"` // 搜索条件
	FieldSearchHide bool        `json:"fieldSearchHide"` // 是否隐藏查询条件
	DictType        string      `json:"dictType"`        // 字典
	Front           bool        `json:"front"`           // 是否前端可见
	Require         bool        `json:"require"`         // 是否必填
	DefaultValue    string      `json:"defaultValue"`    // 是否必填
	ErrorText       string      `json:"errorText"`       // 校验失败文字
	Clearable       bool        `json:"clearable"`       // 是否可清空
	Sort            bool        `json:"sort"`            // 是否增加排序
	PrimaryKey      bool        `json:"primaryKey"`      // 是否主键
	DataSource      *DataSource `json:"dataSource"`      // 数据源
	CheckDataSource bool        `json:"checkDataSource"` // 是否检查数据源
	FieldIndexType  string      `json:"fieldIndexType"`  // 索引类型
}

type SysAutoCode struct {
	global.GVA_MODEL
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:展示名"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}

type AutoPlugReq struct {
	PlugName    string         `json:"plugName"` // 必然大写开头
	Snake       string         `json:"snake"`    // 后端自动转为 snake
	RouterGroup string         `json:"routerGroup"`
	HasGlobal   bool           `json:"hasGlobal"`
	HasRequest  bool           `json:"hasRequest"`
	HasResponse bool           `json:"hasResponse"`
	NeedModel   bool           `json:"needModel"`
	Global      []AutoPlugInfo `json:"global,omitempty"`
	Request     []AutoPlugInfo `json:"request,omitempty"`
	Response    []AutoPlugInfo `json:"response,omitempty"`
}

func (a *AutoPlugReq) CheckList() {
	a.Global = bind(a.Global)
	a.Request = bind(a.Request)
	a.Response = bind(a.Response)

}
func bind(req []AutoPlugInfo) []AutoPlugInfo {
	var r []AutoPlugInfo
	for _, info := range req {
		if info.Effective() {
			r = append(r, info)
		}
	}
	return r
}

type AutoPlugInfo struct {
	Key  string `json:"key"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}

func (a AutoPlugInfo) Effective() bool {
	return a.Key != "" && a.Type != "" && a.Desc != ""
}
