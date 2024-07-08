package request

import (
	"encoding/json"
	"fmt"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	"go/token"
	"strings"
)

type AutoCode struct {
	Package             string                 `json:"package"`
	PackageT            string                 `json:"-"`
	TableName           string                 `json:"tableName" example:"表名"`
	BusinessDB          string                 `json:"businessDB" example:"业务数据库"`
	StructName          string                 `json:"structName" example:"Struct名称"`
	PackageName         string                 `json:"packageName" example:"文件名称"`
	Description         string                 `json:"description" example:"Struct中文名称"`
	Abbreviation        string                 `json:"abbreviation" example:"Struct简称"`
	HumpPackageName     string                 `json:"humpPackageName" example:"go文件名称"`
	GvaModel            bool                   `json:"gvaModel" example:"是否使用gva默认Model"`
	AutoMigrate         bool                   `json:"autoMigrate" example:"是否自动迁移表结构"`
	AutoCreateResource  bool                   `json:"autoCreateResource" example:"是否自动创建资源标识"`
	AutoCreateApiToSql  bool                   `json:"autoCreateApiToSql" example:"是否自动创建api"`
	AutoCreateMenuToSql bool                   `json:"autoCreateMenuToSql" example:"是否自动创建menu"`
	Fields              []*AutoCodeField       `json:"fields"`
	DictTypes           []string               `json:"-"`
	FrontFields         []*AutoCodeField       `json:"-"`
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
}

type DataSource struct {
	Table       string `json:"table"`
	Label       string `json:"label"`
	Value       string `json:"value"`
	Association int    `json:"association"` // 关联关系 1 一对一 2 一对多
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

func (r *AutoCode) Menu() model.SysBaseMenu {
	return model.SysBaseMenu{
		ParentId:  0,
		Path:      r.Abbreviation,
		Name:      r.Abbreviation,
		Component: fmt.Sprintf("view/%s/%s/%s.vue", r.Package, r.PackageName, r.PackageName),
		Meta: model.Meta{
			Title: r.Description,
		},
	}
}

// Pretreatment 预处理
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *AutoCode) Pretreatment() error {
	if token.IsKeyword(r.Abbreviation) {
		r.Abbreviation = r.Abbreviation + "_"
	} // go 关键字处理
	if strings.HasSuffix(r.HumpPackageName, "test") {
		r.HumpPackageName = r.HumpPackageName + "_"
	} // test
	length := len(r.Fields)
	dict := make(map[string]string, length)
	r.FrontFields = make([]*AutoCodeField, 0, length)
	r.DataSourceMap = make(map[string]*DataSource, length)
	for i := 0; i < length; i++ {
		if r.Fields[i].DictType != "" {
			dict[r.Fields[i].DictType] = ""
		}
		if r.Fields[i].Sort {
			r.NeedSort = true
		}
		if r.Fields[i].Front {
			r.FrontFields = append(r.FrontFields, r.Fields[i])
		}
		switch r.Fields[i].FieldType {
		case "file":
			r.HasFile = true
			r.NeedJSON = true
		case "json":
			r.NeedJSON = true
		case "array":
			r.NeedJSON = true
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
	if r.Package == "" {
		return errors.New("Package为空!")
	} // 增加判断：Package不为空
	r.PackageT = utils.FirstUpper(r.Package)
	return nil
}

func (r *AutoCode) History() model.SysAutoCodeHistory {
	bytes, _ := json.Marshal(r)
	return model.SysAutoCodeHistory{
		Table:       r.TableName,
		Package:     r.Package,
		Request:     string(bytes),
		StructName:  r.StructName,
		BusinessDB:  r.BusinessDB,
		Description: r.Description,
	}
}

type AutoCodeField struct {
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
