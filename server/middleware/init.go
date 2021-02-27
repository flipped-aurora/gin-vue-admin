package middleware

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"

	"github.com/gin-gonic/gin"
)

func InitCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GVA_DB == nil {
			// 未初始化
			response.FailWithDetailed(gin.H{"database": true}, "数据库未初始化", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
