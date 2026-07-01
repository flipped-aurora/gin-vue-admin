package middleware

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestMeta() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-Id")
		if reqID == "" {
			reqID = uuid.NewString()
		}
		f := &logger.Fields{
			RequestID:  reqID,
			TraceID:    extractTraceID(c),
			DeviceID:   c.GetHeader("X-Device-Id"),
			ClientIP:   c.ClientIP(),
			HTTPMethod: c.Request.Method,
			HTTPPath:   c.Request.URL.Path,
		}
		c.Request = c.Request.WithContext(logger.WithFields(c.Request.Context(), f))
		c.Header("X-Request-Id", reqID)
		c.Next()
	}
}

// extractTraceID 仅从上游 header 透传（traceparent 取 trace-id 段，或 X-Trace-Id），无则空
func extractTraceID(c *gin.Context) string {
	if tp := c.GetHeader("traceparent"); tp != "" {
		// 格式: version-traceid-spanid-flags
		parts := strings.Split(tp, "-")
		if len(parts) >= 2 {
			return parts[1]
		}
	}
	return c.GetHeader("X-Trace-Id")
}
