package system

import "errors"

// 初始版本自动化代码工具
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
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataType        string `json:"dataType"`        // 数据库字段类型
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
}

var AutoMoveErr error = errors.New("创建代码成功并移动文件成功")
