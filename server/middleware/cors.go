package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("AccessToken", "X-CSRF-Token", "Authorization",
		"Token", "X-Token", "X-User-Id")
	corsConfig.AddExposeHeaders("Content-Length", "Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers", "Content-Type")
	return cors.New(corsConfig)
}
