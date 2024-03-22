package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LinksRouter struct {
}

// InitLinksRouter 初始化 Links 路由信息
func (s *LinksRouter) InitLinksRouter(Router *gin.RouterGroup) {
	linksRouter := Router.Group("links").Use(middleware.OperationRecord())
	linksRouterWithoutRecord := Router.Group("links")
	var linksApi = v1.ApiGroupApp.WebcmsApiGroup.LinksApi
	{
		linksRouter.POST("createLinks", linksApi.CreateLinks)             // 新建Links
		linksRouter.DELETE("deleteLinks", linksApi.DeleteLinks)           // 删除Links
		linksRouter.DELETE("deleteLinksByIds", linksApi.DeleteLinksByIds) // 批量删除Links
		linksRouter.PUT("updateLinks", linksApi.UpdateLinks)              // 更新Links
	}
	{
		linksRouterWithoutRecord.GET("findLinks", linksApi.FindLinks)       // 根据ID获取Links
		linksRouterWithoutRecord.GET("getLinksList", linksApi.GetLinksList) // 获取Links列表
	}
}
