package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestMeta() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-Id")
		if !logger.SaneHeaderID(reqID) {
			reqID = uuid.NewString()
		}
		traceID, parentSpanID := extractTrace(c)
		if traceID == "" {
			// 无上游链路时本地生成,保证全链路 trace_id 永不为空
			traceID = logger.NewTraceID()
		}
		spanID := logger.NewSpanID()
		f := &logger.Fields{
			RequestID:    reqID,
			TraceID:      traceID,
			SpanID:       spanID,
			ParentSpanID: parentSpanID,
			DeviceID:     c.GetHeader("X-Device-Id"),
			ClientIP:     c.ClientIP(),
			HTTPMethod:   c.Request.Method,
			HTTPPath:     c.Request.URL.Path,
		}
		c.Request = c.Request.WithContext(logger.WithFields(c.Request.Context(), f))
		// 回写响应头:前端/网关拿到 trace_id 后报障可直接关联日志
		c.Header("X-Request-Id", reqID)
		c.Header("X-Trace-Id", traceID)
		// 透传的 X-Trace-Id 可能不是 W3C 格式,只有合法 trace-id 才能组装 traceparent
		if logger.IsValidTraceID(traceID) {
			c.Header("traceparent", logger.BuildTraceparent(traceID, spanID))
		}
		c.Next()
	}
}

// extractTrace 提取上游链路标识,优先级:合法 traceparent → X-Trace-Id。
// traceparent 严格校验(畸形头按无上游处理),X-Trace-Id 宽松透传以兼容
// 非 W3C 格式的自定义网关(但需通过 logger.SaneHeaderID 卫生校验);
// 均无则返回空,由调用方本地生成。
func extractTrace(c *gin.Context) (traceID, parentSpanID string) {
	if tid, psid, ok := logger.ParseTraceparent(c.GetHeader("traceparent")); ok {
		return tid, psid
	}
	if tid := c.GetHeader("X-Trace-Id"); logger.SaneHeaderID(tid) {
		return tid, ""
	}
	return "", ""
}
