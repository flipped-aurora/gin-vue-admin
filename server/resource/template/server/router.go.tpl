package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	{
		{{.StructName}}Router.POST("create{{.StructName}}", v1.Create{{.StructName}})   // 新建{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}", v1.Delete{{.StructName}}) // 删除{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}ByIds", v1.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
		{{.StructName}}Router.PUT("update{{.StructName}}", v1.Update{{.StructName}})    // 更新{{.StructName}}
		{{.StructName}}Router.GET("find{{.StructName}}", v1.Find{{.StructName}})        // 根据ID获取{{.StructName}}
		{{.StructName}}Router.GET("get{{.StructName}}List", v1.Get{{.StructName}}List)  // 获取{{.StructName}}列表
	}
}
