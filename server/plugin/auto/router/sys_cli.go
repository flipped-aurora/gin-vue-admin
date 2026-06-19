package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliRouter struct{}

func (s *CliRouter) InitCliRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	cliRouter := Router.Group("cli").Use(middleware.OperationRecord())
	cliRouterWithoutRecord := Router.Group("cli")
	{
		cliRouter.POST("createCli", cliApi.CreateCli)
		cliRouter.POST("updateCli", cliApi.UpdateCli)
		cliRouter.POST("deleteCli", cliApi.DeleteCli)
		cliRouter.POST("addCliApis", cliApi.AddCliApis)
		cliRouter.POST("removeCliApis", cliApi.RemoveCliApis)
	}
	{
		cliRouterWithoutRecord.POST("getCliList", cliApi.GetCliList)
		cliRouterWithoutRecord.POST("getCliDetail", cliApi.GetCliDetail)
		cliRouterWithoutRecord.POST("previewManifest", cliApi.PreviewManifest)
		cliRouterWithoutRecord.POST("previewApiCommand", cliApi.PreviewApiCommand)
		cliRouterWithoutRecord.POST("downloadManifest", cliApi.DownloadManifest)
		cliRouterWithoutRecord.POST("buildCli", cliApi.BuildCliBinary)
		cliRouterWithoutRecord.POST("downloadSkill", cliApi.DownloadCliSkill)
	}
	_ = RouterPub
}
