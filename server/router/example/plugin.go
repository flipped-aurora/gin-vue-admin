package example

import "github.com/gin-gonic/gin"

type Plugin struct {
}

func (*Plugin) Register(group *gin.RouterGroup) {
	group.GET("hello", func(context *gin.Context) {
		context.JSON(200, "hello world")
	})
}

func (*Plugin) RouterPath() string {
	return "group"
}
