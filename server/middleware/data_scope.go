package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
	"github.com/gin-gonic/gin"
)

// DataScope 数据权限中间件: 在 JWTAuth 之后, 依据 claims 构建数据权限身份,
// 并注入 c.Request.Context(), 供 Service 层统一 WithContext(ctx) 透传到 GORM 回调消费。
// 这补上了历史缺口: 此前 jwt.go 只 c.Set("claims"), 身份没进 request.Context()。
func DataScope() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserID(c)
		if userID == 0 {
			c.Next()
			return
		}
		authorityID := utils.GetUserAuthorityId(c)
		id, err := service.ServiceGroupApp.SystemServiceGroup.DataScopeService.
			BuildIdentity(c.Request.Context(), userID, authorityID)
		if err == nil && id != nil {
			c.Request = c.Request.WithContext(datascope.WithIdentity(c.Request.Context(), id))
		}
		c.Next()
	}
}
