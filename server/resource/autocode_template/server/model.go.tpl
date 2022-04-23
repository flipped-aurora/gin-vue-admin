// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// {{.StructName}} 结构体
// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
      global.GVA_MODEL {{- range .Fields}}
            {{- if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- end }} {{- end }}
}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
