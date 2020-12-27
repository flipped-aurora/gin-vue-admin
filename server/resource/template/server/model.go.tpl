// 自动生成模板{{.StructName}}
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type {{.StructName}}Base struct {
      {{- range .Fields}}
            {{- if eq .FieldType "bool" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- end }} {{- end }}
}

type {{.StructName}} struct {
      global.GVA_MODEL
      {{.StructName}}Base
}

{{ if .TableName }}
func ({{.StructName}}) TableName() string {
    return "{{.TableName}}"
}
{{ end }}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type {{.StructName}}Workflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	{{.StructName}}   `json:"business"`
// }

// func ({{.StructName}}) TableName() string {
// 	return "{{.TableName}}"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["{{.Abbreviation}}"] = func() model.GVA_Workflow {
//   return new(model.{{.StructName}}Workflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["{{.Abbreviation}}"] = func() interface{} {
// 	return new(model.{{.StructName}})
// }
