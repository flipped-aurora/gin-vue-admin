package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Info = new(info)

type info struct{}

// Init 初始化 公告 路由信息
func (r *info) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("info").Use(middleware.OperationRecord())
		group.POST("createInfo", apiInfo.CreateInfo)             // 新建公告
		group.DELETE("deleteInfo", apiInfo.DeleteInfo)           // 删除公告
		group.DELETE("deleteInfoByIds", apiInfo.DeleteInfoByIds) // 批量删除公告
		group.PUT("updateInfo", apiInfo.UpdateInfo)              // 更新公告
	}
	{
		group := private.Group("info")
		group.GET("findInfo", apiInfo.FindInfo)       // 根据ID获取公告
		group.GET("getInfoList", apiInfo.GetInfoList) // 获取公告列表
	}
	{
		group := public.Group("info")
		group.GET("getInfoDataSource", apiInfo.GetInfoDataSource) // 获取公告数据源
		group.GET("getInfoPublic", apiInfo.GetInfoPublic)         // 获取公告列表
	}
}
