package router

import (
	"github.com/eyotang/game-proxy/server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGameProxyRouter(r *gin.RouterGroup) {
	gameProxyRouter := r.Group("game/:id")
	{
		connRouter := gameProxyRouter.Group("connection")
		{
			connRouter.POST("open", v1.ServedPlugin, v1.OpenConnection)
			connRouter.POST("close", v1.OwnedToken, v1.CloseConnection)
		}
		gameProxyRouter.POST("destroy", v1.ServedPlugin, v1.Destroy)
		gameProxyRouter.POST("request/:request", v1.OwnedToken, v1.GameRequest)
	}
}
