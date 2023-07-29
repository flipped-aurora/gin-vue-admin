package initialize

import (
	"net/http"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))

	Router.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		PublicGroup.POST("/getGvaToken", v1.ApiGroupApp.UserTeemlinkPkgApiGroup.GetGvaToken)
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitChatGptRouter(PrivateGroup)

		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	{
		NestInfoRouter := router.RouterGroupApp.NestInfo
		NestInfoRouter.InitNestInfoRouter(PrivateGroup)
	}
	{
		nestrolepkgRouter := router.RouterGroupApp.Nestrolepkg
		nestrolepkgRouter.InitNestRoleRouter(PrivateGroup)
	}
	{

	}
	{
		NestExecRecordPkgRouter := router.RouterGroupApp.NestExecRecordPkg
		NestExecRecordPkgRouter.InitNestExecRecordRouter(PrivateGroup)
	}
	{

	}
	{

	}
	{
		FlyResultPkgRouter := router.RouterGroupApp.FlyResultPkg
		FlyResultPkgRouter.InitFlyResultRouter(PrivateGroup)
		UserTeemlinkPkgRouter := router.RouterGroupApp.UserTeemlinkPkg
		UserTeemlinkPkgRouter.InitUserTeemlinkRouter(PrivateGroup)
	}
	{
		NestAirlinePkgRouter := router.RouterGroupApp.NestAirlinePkg
		NestAirlinePkgRouter.InitNestAirlineRouter(PrivateGroup)
	}
	{
		AerialPhotographyResultPkgRouter := router.RouterGroupApp.AerialPhotographyResultPkg
		AerialPhotographyResultPkgRouter.InitAerialPhotographyResultRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
