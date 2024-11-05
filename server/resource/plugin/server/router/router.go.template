package router

import (
	{{if .OnlyTemplate }} // {{end}}"{{.Module}}/middleware"
	"github.com/gin-gonic/gin"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {}

// Init 初始化 {{.Description}} 路由信息
func (r *{{.Abbreviation}}) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
{{- if not .OnlyTemplate }}
	{
	    group := private.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
		group.POST("create{{.StructName}}", api{{.StructName}}.Create{{.StructName}})   // 新建{{.Description}}
		group.DELETE("delete{{.StructName}}", api{{.StructName}}.Delete{{.StructName}}) // 删除{{.Description}}
		group.DELETE("delete{{.StructName}}ByIds", api{{.StructName}}.Delete{{.StructName}}ByIds) // 批量删除{{.Description}}
		group.PUT("update{{.StructName}}", api{{.StructName}}.Update{{.StructName}})    // 更新{{.Description}}
	}
	{
	    group := private.Group("{{.Abbreviation}}")
		group.GET("find{{.StructName}}", api{{.StructName}}.Find{{.StructName}})        // 根据ID获取{{.Description}}
		group.GET("get{{.StructName}}List", api{{.StructName}}.Get{{.StructName}}List)  // 获取{{.Description}}列表
	}
	{
	    group := public.Group("{{.Abbreviation}}")
    	{{- if .HasDataSource}}
	    group.GET("get{{.StructName}}DataSource", api{{.StructName}}.Get{{.StructName}}DataSource)  // 获取{{.Description}}数据源
	    {{- end}}
	    group.GET("get{{.StructName}}Public", api{{.StructName}}.Get{{.StructName}}Public)  // {{.Description}}开放接口
	}
{{- else}}
     // {
	 //   group := private.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	 // }
	 // {
     //   group := private.Group("{{.Abbreviation}}")
     // }
    {
	    group := public.Group("{{.Abbreviation}}")
	    group.GET("get{{.StructName}}Public", api{{.StructName}}.Get{{.StructName}}Public)  // {{.Description}}开放接口
    }
{{- end}}
}
