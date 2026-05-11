package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	notifier := GetDBReadyNotifier()

	// 保存插件初始化参数
	notifier.SetPluginParams(&PluginParams{
		PrivateGroup: PrivateGroup,
		PublicRouter: PublicRouter,
		Engine:       engine,
	})

	if global.GVA_DB == nil {
		global.GVA_LOG.Info("项目暂未初始化，插件将在数据库初始化完成后自动注册")
		// 订阅数据库就绪事件
		notifier.Subscribe(func() {
			global.GVA_LOG.Info("数据库已就绪，开始注册插件")
			params := notifier.GetPluginParams()
			if params != nil {
				bizPluginV1(params.PrivateGroup, params.PublicRouter)
				bizPluginV2(params.Engine)
				// 重新同步全局路由表，包含插件注册的路由
				global.GVA_ROUTERS = params.Engine.Routes()
				global.GVA_LOG.Info("插件注册完成")
			}
		})
		return
	}

	// 数据库已存在，直接注册
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
}
