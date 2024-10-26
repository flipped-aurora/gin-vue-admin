package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		xiaoRouter := router.RouterGroupApp.Xiao
		xiaoRouter.InitCliLoadRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliUserRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliTreeRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliOrderRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliMainorderRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliMainprofitRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliProfitRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliWithdrawRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliMainwithRouter(privateGroup, publicGroup)
		xiaoRouter.InitCliSetvipRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		xiaoRouter.InitCliSetRouter(privateGroup, publicGroup)
	}
}
