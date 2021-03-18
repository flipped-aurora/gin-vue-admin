// 自动生成模板{{.StructName}}
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
      global.GVA_MODEL {{- range .Fields}}
            {{- if eq .FieldType "bool" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- end }} {{- end }}
}

{{ if .TableName }}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
