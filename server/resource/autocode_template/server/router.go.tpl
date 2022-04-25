package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {
}

// Init{{.StructName}}Router 初始化 {{.StructName}} 路由信息
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.Abbreviation}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	{{.Abbreviation}}RouterWithoutRecord := Router.Group("{{.Abbreviation}}")
	var {{.Abbreviation}}Api = v1.ApiGroupApp.{{.PackageT}}ApiGroup.{{.StructName}}Api
	{
		{{.Abbreviation}}Router.POST("create{{.StructName}}", {{.Abbreviation}}Api.Create{{.StructName}})   // 新建{{.StructName}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}", {{.Abbreviation}}Api.Delete{{.StructName}}) // 删除{{.StructName}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}ByIds", {{.Abbreviation}}Api.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
		{{.Abbreviation}}Router.PUT("update{{.StructName}}", {{.Abbreviation}}Api.Update{{.StructName}})    // 更新{{.StructName}}
	}
	{
		{{.Abbreviation}}RouterWithoutRecord.GET("find{{.StructName}}", {{.Abbreviation}}Api.Find{{.StructName}})        // 根据ID获取{{.StructName}}
		{{.Abbreviation}}RouterWithoutRecord.GET("get{{.StructName}}List", {{.Abbreviation}}Api.Get{{.StructName}}List)  // 获取{{.StructName}}列表
	}
}
