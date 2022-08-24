package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    {{- range .Fields}}
        {{- if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
    Start{{.FieldName}}  *{{.FieldType}}  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
    End{{.FieldName}}  *{{.FieldType}}  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
        {{- end }}
       {{- end }}
    request.PageInfo
}
