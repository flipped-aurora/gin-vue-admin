package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	{{ if or .HasSearchTimer .GvaModel}}"time"{{ end }}
	{{ if .NeedJSON }}"gorm.io/datatypes"{{ end }}
)

type {{.StructName}}Search struct{
    {{ if .GvaModel }}
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    {{ end }}
    {{- range .Fields}}
        {{- if ne .FieldSearchType ""}}
            {{- if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
                Start{{.FieldName}}  *{{.FieldType}}  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
                End{{.FieldName}}  *{{.FieldType}}  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
             {{- else }}
                {{- if eq .FieldType "enum" }}
                      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}"`
                            {{- else if eq .FieldType "picture" }}
                      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- else if eq .FieldType "video" }}
                      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                             {{- else if eq .FieldType "file" }}
                      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- else if eq .FieldType "pictures" }}
                      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- else if eq .FieldType "richtext" }}
                      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- else if ne .FieldType "string" }}
                      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- else }}
                      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
                            {{- end }}
            {{- end }}
        {{- end}}
       {{- end }}
    request.PageInfo
    {{- if .NeedSort}}
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    {{- end}}
}
