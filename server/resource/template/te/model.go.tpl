// 自动生成模板{{.StructName}}
package model

import (
	"github.com/jinzhu/gorm"
)

type {{.StructName}} struct {
      gorm.Model {{range .Fields}}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" {{if .ColumnName}} gorm:"column:{{.ColumnName}}"{{end}}`{{ end }}
}
