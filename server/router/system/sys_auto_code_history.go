package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

var AutoCodeHistoryRouterApp = new(AutoCodeHistoryRouter)

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	{
		api := v1.ApiGroupApp.SystemApiGroup.AutocodeHistory
		autoCodeHistoryRouter.POST("getMeta", api.First)         // 根据id获取meta信息
		autoCodeHistoryRouter.POST("rollback", api.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", api.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", api.GetList) // 获取回滚记录分页
	}
}
