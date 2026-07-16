package middleware

import (
	"net/http"
	"strings"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// mustChangePwdAllowList 强制改密状态下允许访问的接口后缀
var mustChangePwdAllowList = []string{
	"/user/changePassword",
	"/user/getUserInfo",
	"/jwt/jsonInBlacklist",
}

// MustChangePwdGuard 当 jwt 携带 MustChangePwd=true 时 仅放行改密/用户信息/登出 其余 403
func MustChangePwdGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, exists := c.Get("claims")
		if !exists {
			c.Next()
			return
		}
		claims, ok := raw.(*systemReq.CustomClaims)
		if !ok || !claims.MustChangePwd {
			c.Next()
			return
		}
		path := c.Request.URL.Path
		for _, allow := range mustChangePwdAllowList {
			if strings.HasSuffix(path, allow) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"code": 7,
			"data": gin.H{"needChangePassword": true},
			"msg":  "密码已过期，请先修改密码",
		})
		c.Abort()
	}
}
