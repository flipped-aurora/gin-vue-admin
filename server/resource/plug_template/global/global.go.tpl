package global

{{- if .HasGlobal }}

import "github.com/gzpz/golf-sales-system/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}