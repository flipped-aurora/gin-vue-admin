{{- if .IsAdd}}
// 在结构体中新增如下字段
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
         {{ GenerateSearchField . }}
    {{- end}}
{{- end }}
{{- if .NeedSort}}
Sort  string `json:"sort" form:"sort"`
Order string `json:"order" form:"order"`
{{- end}}
{{- else }}
package request
{{- if not .OnlyTemplate}}
import (
	"{{.Module}}/model/common/request"
	{{ if or .HasSearchTimer .GvaModel}}"time"{{ end }}
)
{{- end}}
type {{.StructName}}Search struct{
{{- if not .OnlyTemplate}}

{{- if .GvaModel }}
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
{{- end }}
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
       {{ GenerateSearchField . }}
    {{- end}}
{{- end }}
    request.PageInfo
    {{- if .NeedSort}}
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    {{- end}}
{{- end }}
}
{{- end }}