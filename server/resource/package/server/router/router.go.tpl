package {{.Package}}

import (
	{{if .OnlyTemplate}}// {{ end}}"{{.Module}}/middleware"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {}

// Init{{.StructName}}Router 初始化 {{.Description}} 路由信息
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	{{- if not .OnlyTemplate}}
	{{.Abbreviation}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	{{.Abbreviation}}RouterWithoutRecord := Router.Group("{{.Abbreviation}}")
	{{- else }}
	// {{.Abbreviation}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
    // {{.Abbreviation}}RouterWithoutRecord := Router.Group("{{.Abbreviation}}")
	{{- end}}
	{{.Abbreviation}}RouterWithoutAuth := PublicRouter.Group("{{.Abbreviation}}")
	{{- if not .OnlyTemplate}}
	{
		{{.Abbreviation}}Router.POST("create{{.StructName}}", {{.Abbreviation}}Api.Create{{.StructName}})   // 新建{{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}", {{.Abbreviation}}Api.Delete{{.StructName}}) // 删除{{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}ByIds", {{.Abbreviation}}Api.Delete{{.StructName}}ByIds) // 批量删除{{.Description}}
		{{.Abbreviation}}Router.PUT("update{{.StructName}}", {{.Abbreviation}}Api.Update{{.StructName}})    // 更新{{.Description}}
	}
	{
		{{.Abbreviation}}RouterWithoutRecord.GET("find{{.StructName}}", {{.Abbreviation}}Api.Find{{.StructName}})        // 根据ID获取{{.Description}}
		{{.Abbreviation}}RouterWithoutRecord.GET("get{{.StructName}}List", {{.Abbreviation}}Api.Get{{.StructName}}List)  // 获取{{.Description}}列表
	}
	{
	{{- if .HasDataSource}}
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}DataSource", {{.Abbreviation}}Api.Get{{.StructName}}DataSource)  // 获取{{.Description}}数据源
	{{- end}}
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}Public", {{.Abbreviation}}Api.Get{{.StructName}}Public)  // {{.Description}}开放接口
	}
	{{- else}}
	{
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}Public", {{.Abbreviation}}Api.Get{{.StructName}}Public)  // {{.Description}}开放接口
	}
    {{ end }}
}
