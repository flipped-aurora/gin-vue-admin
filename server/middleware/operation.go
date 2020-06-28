package middleware

import (
	"bytes"
	"fmt"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var body []byte

func RecordRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.GVA_LOG.Error(err)
			}

			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			body = nil
		}
	}
}

func OperationRecord() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {

			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				string(body),
				param.ErrorMessage,
			)
		},
		// 暂时没考虑好
		Output:    nil,
		SkipPaths: global.GVA_CONFIG.Operation.SkipPaths,
	})
}
