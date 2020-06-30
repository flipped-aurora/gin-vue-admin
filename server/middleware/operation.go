package middleware

import (
	"bytes"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

var body []byte
var userId uint

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

		//TODO parse token , userId <-
	}
}

func OperationRecord() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// 防止加载查询参数，再次过滤
			for _, v := range global.GVA_CONFIG.Operation.SkipPaths {
				if strings.Contains(param.Path, v) {
					return ""
				}
			}
			err := service.CreateSysOperationRecord(model.SysOperationRecord{
				Ip:           param.ClientIP,
				Method:       param.Method,
				Path:         param.Path,
				Status:       param.StatusCode,
				Latency:      param.Latency,
				Agent:        param.Request.UserAgent(),
				ErrorMessage: param.ErrorMessage,
				Body:         string(body),
				UserId:       int(userId),
			})
			if err != nil {
				global.GVA_LOG.Error(err)
			}
			return ""
		},
		Output:    nil,
		SkipPaths: global.GVA_CONFIG.Operation.SkipPaths,
	})
}
