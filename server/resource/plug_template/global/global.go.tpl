package global

{{- if .HasGlobal }}

import "kirer.cn/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}