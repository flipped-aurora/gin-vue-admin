package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Info(global.Translate("initialize.projectNotInitialized"))
		return
	}
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
}
