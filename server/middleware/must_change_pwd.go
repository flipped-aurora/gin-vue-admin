package middleware

import (
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// mustChangePwdAllowList 强制改密状态下允许访问的接口后缀
var mustChangePwdAllowList = []string{
	"/user/changePassword",
	"/user/getUserInfo",
	"/jwt/jsonInBlacklist",
}

// MustChangePwdGuard 当 jwt 携带 MustChangePwd=true 时 仅放行改密/用户信息/登出 其余返回账号状态冲突
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
		c.JSON(http.StatusConflict, response.Response{
			Code: response.PASSWORD_CHANGE_REQUIRED,
			Data: gin.H{"needChangePassword": true},
			Msg:  "密码已过期，请先修改密码",
		})
		c.Abort()
	}
}
