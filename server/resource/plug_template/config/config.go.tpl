package config

{{- if .HasGlobal }}
type {{ .PlugName }} struct {
    {{- range .Global }}
        {{ .Key }} {{ .Type }} {{- if ne .Desc "" }} // {{ .Desc }} {{ end -}}
    {{- end }}
}
{{ end -}}