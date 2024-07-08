package initialize

import (
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
}
