package system

import (
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkdb", dbApi.CheckDB) // 检测是否需要初始化数据库
	}
}
