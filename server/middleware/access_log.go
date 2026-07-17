package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

// gin context 缓存键（供 operation.go 复用，避免重复读 body / 包装 writer）
const (
	ctxReqBodyKey    = "gva_req_body"
	ctxRespBufferKey = "gva_resp_buffer"
)

// respCaptureLimit 响应捕获缓冲上限:SSE 等长/无界流的响应体只捕获前 1MB,
// 防止缓冲随连接存活无限增长;正常 API 响应几乎不受影响。
// 刻意不按请求头(如 Accept)决定是否捕获——客户端可控的头不能成为绕过
// 响应审计(操作记录 Resp/访问日志 resp_data)的开关。
const respCaptureLimit = 1 << 20

type captureWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (w captureWriter) Write(b []byte) (int, error) {
	if remain := respCaptureLimit - w.buf.Len(); remain > 0 {
		if len(b) <= remain {
			w.buf.Write(b)
		} else {
			w.buf.Write(b[:remain])
		}
	}
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 唯一 body 读取点：非 multipart 才读，超长截断
		// bytes_in 优先用实读长度(chunked 请求 ContentLength 为 -1),multipart 退回 ContentLength
		reqBody := ""
		bytesIn := c.Request.ContentLength
		if !strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			raw, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(raw))
			bytesIn = int64(len(raw))
			sanitized := logger.SanitizeBody(c.GetHeader("Content-Type"), string(raw))
			reqBody = logger.Truncate(sanitized, logger.AccessLogMaxBytes())
		} else {
			reqBody = logger.MultipartMark
		}
		if bytesIn < 0 {
			bytesIn = 0 // 未知长度(chunked multipart 等)归零,避免聚合出现负值
		}
		c.Set(ctxReqBodyKey, reqBody)

		// 唯一 writer 包装点：缓冲区指针在 pre 阶段缓存，供 operation 在 post 读取。
		// 捕获有 respCaptureLimit 封顶,长流(SSE 等)不会撑爆内存
		buf := &bytes.Buffer{}
		c.Writer = captureWriter{ResponseWriter: c.Writer, buf: buf}
		c.Set(ctxRespBufferKey, buf)
		c.Writer.Header().Set("X-Gva-Version", global.Version)

		c.Next()

		// http_route 用路由模板(低基数):指标聚合的分组键;原始 URL 在 http_path(高基数)
		route := c.FullPath()
		if route == "" {
			route = "unmatched" // 未命中路由(404 等)统一兜底,防止基数爆炸
		}
		userID, authorityID := currentIdentity(c)
		privateErrs := c.Errors.ByType(gin.ErrorTypePrivate).String()
		hasErr := c.Writer.Status() >= http.StatusInternalServerError || privateErrs != ""
		bytesOut := int64(c.Writer.Size())
		if bytesOut < 0 {
			bytesOut = 0 // gin 未写入 body 时 Size() 为 -1(如仅 c.Status/304)
		}

		b := logger.WithCtx(c.Request.Context()).Mod("http").
			Field("http_route", route).
			Field("http_status", c.Writer.Status()).
			Field("latency_ms", time.Since(start).Milliseconds()).
			Field("bytes_in", bytesIn).
			Field("bytes_out", bytesOut).
			Field("user_id", userID).
			Field("authority_id", authorityID).
			Field("error", hasErr).
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
			respBody := logger.SanitizeBody(c.Writer.Header().Get("Content-Type"), buf.String())
			b = b.Field("resp_data", logger.Truncate(respBody, logger.AccessLogMaxBytes()))
		}
		if privateErrs != "" {
			b = b.Field("error_msg", strings.TrimRight(privateErrs, "\n"))
		}
		b.Info("请求完成")
	}
}

// currentIdentity 仅读取 JWTAuth 已解析放入 gin context 的 claims,
// 不主动解析 token:公开路由无 token,主动解析会每请求产生一条无意义的错误日志。
// 未认证请求返回 0/0。
func currentIdentity(c *gin.Context) (userID, authorityID uint) {
	if v, ok := c.Get("claims"); ok {
		if cl, ok := v.(*systemReq.CustomClaims); ok {
			return cl.BaseClaims.ID, cl.AuthorityId
		}
	}
	return 0, 0
}
