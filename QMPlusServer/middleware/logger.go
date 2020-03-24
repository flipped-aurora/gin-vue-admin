package middleware

import (
	"bytes"
	"gin-vue-admin/init/log"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddlewareFactory(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// request time
		start := time.Now()
		// request path
		path := c.Request.URL.Path
		logFlag := true
		if strings.Contains(path, "swagger") {
			logFlag = false
		}
		// request ip
		clientIP := c.ClientIP()
		// method
		method := c.Request.Method
		// copy request content
		req, _ := httputil.DumpRequest(c.Request, true)
		if logFlag {
			logger.Debug(
				"Request:", method, clientIP, path, string(req))
		}
		// replace writer
		cusWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = cusWriter
		// handle request
		c.Next()
		// ending time
		end := time.Now()
		//execute time
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		if logFlag {
			logger.Debug(
				"Response:",
				statusCode,
				latency,
				cusWriter.body.String())
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
