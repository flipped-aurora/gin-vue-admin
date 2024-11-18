package request

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"go/token"
	"strings"
)

type AutoCode struct {
	Package             string                 `json:"package"`
	PackageT            string                 `json:"-"`
	TableName           string                 `json:"tableName" example:"表名"`              // 表名
	BusinessDB          string                 `json:"businessDB" example:"业务数据库"`          // 业务数据库
	StructName          string                 `json:"structName" example:"Struct名称"`       // Struct名称
	PackageName         string                 `json:"packageName" example:"文件名称"`          // 文件名称
	Description         string                 `json:"description" example:"Struct中文名称"`    // Struct中文名称
	Abbreviation        string                 `json:"abbreviation" example:"Struct简称"`     // Struct简称
	HumpPackageName     string                 `json:"humpPackageName" example:"go文件名称"`    // go文件名称
	GvaModel            bool                   `json:"gvaModel" example:"false"`            // 是否使用gva默认Model
	AutoMigrate         bool                   `json:"autoMigrate" example:"false"`         // 是否自动迁移表结构
	AutoCreateResource  bool                   `json:"autoCreateResource" example:"false"`  // 是否自动创建资源标识
	AutoCreateApiToSql  bool                   `json:"autoCreateApiToSql" example:"false"`  // 是否自动创建api
	AutoCreateMenuToSql bool                   `json:"autoCreateMenuToSql" example:"false"` // 是否自动创建menu
	AutoCreateBtnAuth   bool                   `json:"autoCreateBtnAuth" example:"false"`   // 是否自动创建按钮权限
	OnlyTemplate        bool                   `json:"onlyTemplate" example:"false"`        // 是否只生成模板
	IsTree              bool                   `json:"isTree" example:"false"`              // 是否树形结构
	TreeJson            string                 `json:"treeJson" example:"展示的树json字段"`       // 展示的树json字段
	IsAdd               bool                   `json:"isAdd" example:"false"`               // 是否新增
	Fields              []*AutoCodeField       `json:"fields"`
	Module              string                 `json:"-"`
	DictTypes           []string               `json:"-"`
	PrimaryField        *AutoCodeField         `json:"primaryField"`
	DataSourceMap       map[string]*DataSource `json:"-"`
	HasPic              bool                   `json:"-"`
	HasFile             bool                   `json:"-"`
	HasTimer            bool                   `json:"-"`
	NeedSort            bool                   `json:"-"`
	NeedJSON            bool                   `json:"-"`
	HasRichText         bool                   `json:"-"`
	HasDataSource       bool                   `json:"-"`
	HasSearchTimer      bool                   `json:"-"`
	HasArray            bool                   `json:"-"`
	HasExcel            bool                   `json:"-"`
}

type DataSource struct {
	DBName       string `json:"dbName"`
	Table        string `json:"table"`
	Label        string `json:"label"`
	Value        string `json:"value"`
	Association  int    `json:"association"` // 关联关系 1 一对一 2 一对多
	HasDeletedAt bool   `json:"hasDeletedAt"`
}

func (r *AutoCode) Apis() []model.SysApi {
	return []model.SysApi{
		{
			Path:        "/" + r.Abbreviation + "/" + "create" + r.StructName,
			Description: "新增" + r.Description,
			ApiGroup:    r.Description,
			Method:      "POST",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "delete" + r.StructName,
			Description: "删除" + r.Description,
			ApiGroup:    r.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "delete" + r.StructName + "ByIds",
			Description: "批量删除" + r.Description,
			ApiGroup:    r.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "update" + r.StructName,
			Description: "更新" + r.Description,
			ApiGroup:    r.Description,
			Method:      "PUT",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "find" + r.StructName,
			Description: "根据ID获取" + r.Description,
			ApiGroup:    r.Description,
			Method:      "GET",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "get" + r.StructName + "List",
			Description: "获取" + r.Description + "列表",
			ApiGroup:    r.Description,
			Method:      "GET",
		},
	}
}

func (r *AutoCode) Menu(template string) model.SysBaseMenu {
	component := fmt.Sprintf("view/%s/%s/%s.vue", r.Package, r.PackageName, r.PackageName)
	if template != "package" {
		component = fmt.Sprintf("plugin/%s/view/%s.vue", r.Package, r.PackageName)
	}
	return model.SysBaseMenu{
		ParentId:  0,
		Path:      r.Abbreviation,
		Name:      r.Abbreviation,
		Component: component,
		Meta: model.Meta{
			Title: r.Description,
		},
	}
}

// Pretreatment 预处理
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *AutoCode) Pretreatment() error {
	r.Module = global.GVA_CONFIG.AutoCode.Module
	if token.IsKeyword(r.Abbreviation) {
		r.Abbreviation = r.Abbreviation + "_"
	} // go 关键字处理
	if strings.HasSuffix(r.HumpPackageName, "test") {
		r.HumpPackageName = r.HumpPackageName + "_"
	} // test
	length := len(r.Fields)
	dict := make(map[string]string, length)
	r.DataSourceMap = make(map[string]*DataSource, length)
	for i := 0; i < length; i++ {
		if r.Fields[i].Excel {
			r.HasExcel = true
		}
		if r.Fields[i].DictType != "" {
			dict[r.Fields[i].DictType] = ""
		}
		if r.Fields[i].Sort {
			r.NeedSort = true
		}
		switch r.Fields[i].FieldType {
		case "file":
			r.HasFile = true
			r.NeedJSON = true
		case "json":
			r.NeedJSON = true
		case "array":
			r.NeedJSON = true
			r.HasArray = true
		case "video":
			r.HasPic = true
		case "richtext":
			r.HasRichText = true
		case "picture":
			r.HasPic = true
		case "pictures":
			r.HasPic = true
			r.NeedJSON = true
		case "time.Time":
			r.HasTimer = true
			if r.Fields[i].FieldSearchType != "" {
				r.HasSearchTimer = true
			}
		}
		if r.Fields[i].DataSource != nil {
			if r.Fields[i].DataSource.Table != "" && r.Fields[i].DataSource.Label != "" && r.Fields[i].DataSource.Value != "" {
				r.HasDataSource = true
				r.Fields[i].CheckDataSource = true
				r.DataSourceMap[r.Fields[i].FieldJson] = r.Fields[i].DataSource
			}
		}
		if !r.GvaModel && r.PrimaryField == nil && r.Fields[i].PrimaryKey {
			r.PrimaryField = r.Fields[i]
		} // 自定义主键
	}
	{
		for key := range dict {
			r.DictTypes = append(r.DictTypes, key)
		}
	} // DictTypes => 字典
	{
		if r.GvaModel {
			r.PrimaryField = &AutoCodeField{
				FieldName:    "ID",
				FieldType:    "uint",
				FieldDesc:    "ID",
				FieldJson:    "ID",
				DataTypeLong: "20",
				Comment:      "主键ID",
				ColumnName:   "id",
			}
		}
	} // GvaModel
	{
		if r.IsAdd && r.PrimaryField == nil {
			r.PrimaryField = new(AutoCodeField)
		}
	} // 新增字段模式下不关注主键
	if r.Package == "" {
		return errors.New("Package为空!")
	} // 增加判断：Package不为空
	packages := []rune(r.Package)
	if len(packages) > 0 {
		if packages[0] >= 97 && packages[0] <= 122 {
			packages[0] = packages[0] - 32
		}
		r.PackageT = string(packages)
	} // PackageT 是 Package 的首字母大写
	return nil
}

func (r *AutoCode) History() SysAutoHistoryCreate {
	bytes, _ := json.Marshal(r)
	return SysAutoHistoryCreate{
		Table:       r.TableName,
		Package:     r.Package,
		Request:     string(bytes),
		StructName:  r.StructName,
		BusinessDB:  r.BusinessDB,
		Description: r.Description,
	}
}

type AutoCodeField struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	FieldSearchHide bool   `json:"fieldSearchHide"` // 是否隐藏查询条件
	DictType        string `json:"dictType"`        // 字典
	//Front           bool        `json:"front"`           // 是否前端可见
	Form            bool        `json:"form"`            // 是否前端新建/编辑
	Table           bool        `json:"table"`           // 是否前端表格列
	Desc            bool        `json:"desc"`            // 是否前端详情
	Excel           bool        `json:"excel"`           // 是否导入/导出
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

type AutoFunc struct {
	Package         string `json:"package"`
	FuncName        string `json:"funcName"`        // 方法名称
	Router          string `json:"router"`          // 路由名称
	FuncDesc        string `json:"funcDesc"`        // 方法介绍
	BusinessDB      string `json:"businessDB"`      // 业务库
	StructName      string `json:"structName"`      // Struct名称
	PackageName     string `json:"packageName"`     // 文件名称
	Description     string `json:"description"`     // Struct中文名称
	Abbreviation    string `json:"abbreviation"`    // Struct简称
	HumpPackageName string `json:"humpPackageName"` // go文件名称
	Method          string `json:"method"`          // 方法
	IsPlugin        bool   `json:"isPlugin"`        // 是否插件
	IsAuth          bool   `json:"isAuth"`          // 是否鉴权
	IsPreview       bool   `json:"isPreview"`       // 是否预览
	IsAi            bool   `json:"isAi"`            // 是否AI
	ApiFunc         string `json:"apiFunc"`         // API方法
	ServerFunc      string `json:"serverFunc"`      // 服务方法
	JsFunc          string `json:"jsFunc"`          // JS方法
}

type InitMenu struct {
	PlugName   string `json:"plugName"`
	ParentMenu string `json:"parentMenu"`
	Menus      []uint `json:"menus"`
}

type InitApi struct {
	PlugName string `json:"plugName"`
	APIs     []uint `json:"apis"`
}

type LLMAutoCode struct {
	Prompt string `json:"prompt" form:"prompt" gorm:"column:prompt;comment:提示语;type:text;"` //提示语
	Mode   string `json:"mode" form:"mode" gorm:"column:mode;comment:模式;type:text;"`        //模式
}
