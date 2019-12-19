package middleware

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

//拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*CustomClaims)
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := waitUse.AuthorityId
		e := sysModel.Casbin()
		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			servers.ReportFormat(c, false, fmt.Sprintf("权限不足"), gin.H{})
			c.Abort()
			return
		}
	}
}
