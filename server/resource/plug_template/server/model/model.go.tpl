package model

{{- if .HasRequest }}
type Request struct {
    {{- range .Request }}
        {{ .Key }} {{ .Type }} {{- if ne .Desc "" }} // {{ .Desc }} {{ end -}}
    {{- end }}
}
{{ end -}}

{{- if .HasResponse }}
type Response struct {
    {{- range .Response }}
        {{ .Key }} {{ .Type }} {{- if ne .Desc "" }} // {{ .Desc }} {{ end -}}
    {{- end }}
}
{{ end -}}
