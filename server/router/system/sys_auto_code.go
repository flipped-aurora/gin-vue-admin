package system

import "github.com/gin-gonic/gin"

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	publicAutoCodeRouter := RouterPublic.Group("autoCode")
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables)
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn)
	}
	{
		autoCodeRouter.POST("preview", autoCodeTemplateApi.Preview)
		autoCodeRouter.POST("createTemp", autoCodeTemplateApi.Create)
		autoCodeRouter.POST("addFunc", autoCodeTemplateApi.AddFunc)
	}
	{
		autoCodeRouter.POST("mcp", autoCodeTemplateApi.MCP)
		autoCodeRouter.POST("mcpList", autoCodeTemplateApi.MCPList)
		autoCodeRouter.POST("mcpRoutes", autoCodeTemplateApi.MCPRoutes)
		autoCodeRouter.POST("mcpTest", autoCodeTemplateApi.MCPTest)
	}
	{
		autoCodeRouter.POST("getPackage", autoCodePackageApi.All)
		autoCodeRouter.POST("delPackage", autoCodePackageApi.Delete)
		autoCodeRouter.POST("createPackage", autoCodePackageApi.Create)
	}
	{
		autoCodeRouter.GET("getTemplates", autoCodePackageApi.Templates)
	}
	{
		autoCodeRouter.POST("pubPlug", autoCodePluginApi.Packaged)
		autoCodeRouter.POST("installPlugin", autoCodePluginApi.Install)
		autoCodeRouter.POST("removePlugin", autoCodePluginApi.Remove)
		autoCodeRouter.GET("getPluginList", autoCodePluginApi.GetPluginList)
	}
	{
		publicAutoCodeRouter.POST("llmAuto", autoCodeApi.LLMAuto)
		publicAutoCodeRouter.POST("initMenu", autoCodePluginApi.InitMenu)
		publicAutoCodeRouter.POST("initAPI", autoCodePluginApi.InitAPI)
		publicAutoCodeRouter.POST("initDictionary", autoCodePluginApi.InitDictionary)
	}
}
