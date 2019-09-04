package middleware

import (
	"github.com/gin-gonic/gin"
	"main/init/qmlog"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		requestParams := buf[0:n]
		qmlog.QMLog.Infof("| %3d | %13v | %15s | %s  %s |%s|",
			statusCode,
			latency,
			clientIP,
			method, path, requestParams,
		)
	}
}
