package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

// gin context 缓存键（供 operation.go 复用，避免重复读 body / 包装 writer）
const (
	ctxReqBodyKey    = "gva_req_body"
	ctxRespBufferKey = "gva_resp_buffer"
)

type captureWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (w captureWriter) Write(b []byte) (int, error) {
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 唯一 body 读取点：非 multipart 才读，超长截断
		reqBody := ""
		if !strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			raw, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(raw))
			reqBody = logger.Truncate(string(raw), logger.AccessLogMaxBytes())
		} else {
			reqBody = logger.MultipartMark
		}
		c.Set(ctxReqBodyKey, reqBody)

		// 唯一 writer 包装点：缓冲区指针在 pre 阶段缓存，供 operation 在 post 读取
		buf := &bytes.Buffer{}
		c.Writer = captureWriter{ResponseWriter: c.Writer, buf: buf}
		c.Set(ctxRespBufferKey, buf)
		c.Writer.Header().Set("X-Gva-Version", global.Version)

		c.Next()

		b := logger.WithCtx(c.Request.Context()).Mod("http").
			Field("http_status", c.Writer.Status()).
			Field("latency_ms", time.Since(start).Milliseconds()).
			Field("ua", c.Request.UserAgent()).
			Field("req_query", c.Request.URL.RawQuery)

		zc := global.GVA_CONFIG.Zap
		if zc.AccessReqHeaders {
			b = b.Field("req_headers", logger.SanitizeHeaders(c.Request.Header))
		}
		if zc.AccessReqBody {
			b = b.Field("req_body", reqBody)
		}
		if zc.AccessRespData {
			b = b.Field("resp_data", logger.Truncate(buf.String(), logger.AccessLogMaxBytes()))
		}
		if errs := c.Errors.ByType(gin.ErrorTypePrivate).String(); errs != "" {
			b = b.Field("error_msg", strings.TrimRight(errs, "\n"))
		}
		b.Info("请求完成")
	}
}
