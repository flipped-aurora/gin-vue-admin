package global

{{- if .HasGlobal }}

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}