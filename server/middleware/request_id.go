package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/tools"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

//通过设置唯一的 ID 之后, 就可以追踪一系列的请求了
func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" || requestId == "undefined" {
			key := tools.RequestClientIPContextLimit + "===" + c.ClientIP()
			//同一个IP一个小时最多开1000个上下文 考虑到同个局域网可能会同时会有100个用户
			limit := 1000
			second := 3600
			expiration := time.Second * time.Duration(second)
			if err := global.GVA_REDIS.SetLimitWithTime(key, limit, expiration); err != nil {
				response.MakeResponse(c).ErrorString(err.Error())
				return
			}
			//3.生成唯一uuid
			requestId = uuid.New().String()
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
