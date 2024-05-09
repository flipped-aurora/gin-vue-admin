package {{ .Snake}}

import (
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
{{- if .HasGlobal }}
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/global"
{{- end }}
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/router"
	"github.com/gin-gonic/gin"
)

type {{ .PlugName}}Plugin struct {
}

func Create{{ .PlugName}}Plug({{- range .Global}} {{.Key}} {{.Type}}, {{- end }})*{{ .PlugName }}Plugin {
{{- if .HasGlobal }}
	{{- range .Global}}
	    global.GlobalConfig.{{.Key}} = {{.Key}}
	{{- end }}
{{ end }}

    gvaGlobal.GVA_DB.AutoMigrate() // 此处可以把插件依赖的数据库结构体自动创建表 需要填写对应的结构体

    // 下方会自动注册菜单 第一个参数为菜单一级路由信息一般为定义好的组名 第二个参数为真实使用的web页面路由信息
    // 具体值请根据实际情况修改
    	utils.RegisterMenus(
    		system.SysBaseMenu{
    			Path:      "{{ .Snake }}Group",
    			Name:      "{{ .Snake }}Group",
    			Hidden:    false,
    			Component: "view/routerHolder.vue",
    			Sort:      1000,
    			Meta: system.Meta{
    				Title: "{{ .PlugName }}组",
    				Icon:  "school",
    			},
    		},
    		system.SysBaseMenu{
    			Path:      "{{ .Snake }}",
    			Name:      "{{ .Snake }}",
    			Hidden:    false,
    			Component: "plugin/{{ .Snake }}/view/index.vue",
    			Sort:      0,
    			Meta: system.Meta{
    				Title: "{{ .PlugName }}",
    				Icon:  "school",
    			},
    		},
    	)

    	// 下方会自动注册api 以下格式为示例格式，请按照实际情况修改
    	utils.RegisterApis(
    		system.SysApi{
    			Path:        "/{{ .RouterGroup }}/routerName",
    			Description: "{{ .PlugName }}接口",
    			ApiGroup:    "{{ .RouterGroup }}",
    			Method:      "POST",
    		},
    	)

	return &{{ .PlugName}}Plugin{}
}

func (*{{ .PlugName}}Plugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.Init{{ .PlugName}}Router(group)
}

func (*{{ .PlugName}}Plugin) RouterPath() string {
	return "{{ .RouterGroup}}"
}
