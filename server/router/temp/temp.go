package temp

import (
	apiTemp "github.com/flipped-aurora/gin-vue-admin/server/api/v1/temp"
	"github.com/gin-gonic/gin"
)

type TempRouter struct {
}

func (temp *TempRouter) InitTempRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tempRouter := PublicRouter.Group("ttt")
	tempApi := apiTemp.NewTemapApi()
	{
		tempRouter.GET("ttt", tempApi.CreateTemp)
		tempRouter.GET("ttt2", tempApi.FindFunc)
		tempRouter.GET("qqq", tempApi.Q自己定义api)
	}
}
