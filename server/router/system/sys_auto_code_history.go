package system

import (
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	{
		autoCodeHistoryRouter.POST("getMeta", autocodeHistoryApi.First)         // 根据id获取meta信息
		autoCodeHistoryRouter.POST("rollback", autocodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", autocodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", autocodeHistoryApi.GetList) // 获取回滚记录分页
	}
}
