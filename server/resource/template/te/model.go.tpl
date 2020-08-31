// 自动生成模板{{.StructName}}
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
      gorm.Model {{- range .Fields}}
            {{- if eq .FieldType "bool" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:'{{.Comment}}'{{- if .DataType -}};type:{{.DataType}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:'{{.Comment}}'{{- if .DataType -}};type:{{.DataType}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
            {{- end }}  {{- end }} 
}

{{ if .TableName }}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}