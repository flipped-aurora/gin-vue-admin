package middleware

import (
	"bytes"
	"fmt"
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
			fmt.Println(global.GVA_CONFIG.Operation.SkipPaths)
			for _, v := range global.GVA_CONFIG.Operation.SkipPaths {
				if strings.Contains(param.Path, v) {
					fmt.Println(param.Path)
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
				ErrorMessage: string(body),
				UserId:       int(userId),
			})
			if err != nil {
				global.GVA_LOG.Error(err)
			}
			return ""
		},
		// 暂时没考虑好
		Output:    nil,
		SkipPaths: global.GVA_CONFIG.Operation.SkipPaths,
	})
}
