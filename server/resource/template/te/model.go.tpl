// 自动生成模板{{.StructName}}
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
      gorm.Model {{range .Fields}}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:'{{.Comment}}'"`{{ end }}
}
