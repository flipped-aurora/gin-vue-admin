package plug

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .PlugName}}/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .PlugName}}/router"
	"github.com/gin-gonic/gin"
)

type {{ .PlugName}}Plugin struct {
}

func Create{{ .PlugName}}Plug() ({{- range .Global}} {{.Key}} {{.Type}} {{- end }})*{{ .PlugName}}Plugin {
	{{- range .Global}}
	    global.GlobalConfig.{{.Key}} = {{.Key}}
	{{- end }}
	return &{{ .PlugName}}Plugin{}
}

func (*{{ .PlugName}}Plugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.Init{{ .PlugName}}Router(group)
}

func (*{{ .PlugName}}Plugin) RouterPath() string {
	return "{{ .RouterGroup}}"
}
