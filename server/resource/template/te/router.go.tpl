package router

import (
	"gin-vue-admin/api/v1"
    "gin-vue-admin/middleware"
    "github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		{{.StructName}}Router.POST("create{{.StructName}}", v1.Create{{.StructName}})     // 新建{{.StructName}}
		{{.StructName}}Router.POST("delete{{.StructName}}", v1.Delete{{.StructName}})   //删除{{.StructName}}
		{{.StructName}}Router.POST("update{{.StructName}}", v1.Update{{.StructName}})   //更新{{.StructName}}
		{{.StructName}}Router.POST("find{{.StructName}}", v1.Find{{.StructName}})           // 根据ID获取{{.StructName}}
		{{.StructName}}Router.POST("get{{.StructName}}List", v1.Get{{.StructName}}List) //获取{{.StructName}}列表
	}
}
