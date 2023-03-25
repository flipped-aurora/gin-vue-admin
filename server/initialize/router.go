package initialize

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理
		systemRouter.InitChatGptRouter(PrivateGroup)             // chatGpt接口

		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
