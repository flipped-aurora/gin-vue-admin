package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/tools"
	"github.com/gin-gonic/gin"
	"time"
)

// ip限制
func IPLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := tools.RequestClientIPLimit + "===" + c.ClientIP()
		limit := global.GVA_CONFIG.System.LimitCountIP
		second := global.GVA_CONFIG.System.LimitTimeIP
		expiration := time.Duration(second) * time.Second
		if err := global.GVA_REDIS.SetLimitWithTime(key, limit, expiration); err != nil {
			response.MakeResponse(c).Error(response.ExceedLimitRequestIP)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
