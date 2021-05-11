package router

import (
	"github.com/eyotang/game-api-admin/server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGameAPIRouter(r *gin.RouterGroup) {
	gameApiRouter := r.Group("game/:id")
	{
		connRouter := gameApiRouter.Group("connection")
		{
			connRouter.POST("open", v1.OpenConnection)
			connRouter.POST("close", v1.CloseConnection)
		}
		gameApiRouter.POST("request/:request", v1.GameRequest)
	}
}
