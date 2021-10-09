package middleware

import (
	"bytes"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		requestBody := ""
		if body, err := c.GetRawData(); err != nil {
			global.GVA_LOG.Error("获取post body 报错", zap.Error(err))
		} else {
			// 再设置回去
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 关键点
			requestBody = utils.Bytes2str(body)
		}
		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		elapsed := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		path := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求ID
		requestId := c.Request.Header.Get("X-Request-Id")
		// 请求头
		headerCp := c.Request.Header.Clone()
		// 请求IP
		clientIP := c.ClientIP()
		responseString := c.GetString("response")
		if elapsed > time.Second*5 {
			// 服务器响应太慢的记录下来 大于5秒的
			global.GVA_LOG.Error("慢请求", zap.String("Path", path), zap.String("X-Request-Id", requestId),
				zap.Time("RequestTime", startTime),
				zap.String("ClientIP", clientIP),
				zap.Int("StatusCode", statusCode),
				zap.String("Method", reqMethod),
				zap.Duration("elapsed", elapsed),
				zap.Any("header", headerCp),
				zap.String("requestBody", requestBody),
				zap.String("response", responseString))
		} else {
			global.GVA_LOG.Info("请求进来了", zap.String("Path", path), zap.String("X-Request-Id", requestId),
				zap.Time("RequestTime", startTime),
				zap.String("ClientIP", clientIP),
				zap.Int("StatusCode", statusCode),
				zap.String("Method", reqMethod),
				zap.Duration("elapsed", elapsed),
				zap.Any("header", headerCp),
				zap.String("requestBody", requestBody),
				zap.String("response", responseString))
		}
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
