package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.WithCtx(c.Request.Context()).Mod("http").
						Field("error", fmt.Sprintf("%v", err)).
						Field("request", string(httpRequest)).
						Error(c.Request.URL.Path)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				// 仅记录一次结构化日志：Error 级日志已由 zap core 统一落 sys_error 表，
				// 此处不再直接写库，消除此前每次 panic 产生两条 sys_error 记录的双写。
				b := logger.WithCtx(c.Request.Context()).Mod("http").
					Field("error", fmt.Sprintf("%v", err)).
					Field("request", string(httpRequest))
				if stack {
					b = b.Field("stack", string(debug.Stack()))
				}
				b.Error("[Recovery from panic]")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
