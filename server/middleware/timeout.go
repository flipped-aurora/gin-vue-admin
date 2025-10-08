package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// TimeoutMiddleware 创建超时中间件
// 入参 timeout 设置超时时间（例如：time.Second * 5）
// 使用示例 xxx.Get("path",middleware.TimeoutMiddleware(30*time.Second),HandleFunc)
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		// 使用 buffered channel 避免 goroutine 泄漏
		done := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					select {
					case panicChan <- p:
					default:
					}
				}
				select {
				case done <- struct{}{}:
				default:
				}
			}()
			c.Next()
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-done:
			return
		case <-ctx.Done():
			// 确保服务器超时设置足够长
			c.Header("Connection", "close")
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"code": 504,
				"msg":  "请求超时",
			})
			return
		}
	}
}
