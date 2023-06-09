package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/api"
	"github.com/gin-gonic/gin"
)

type LdapRouter struct{}

func (s *LdapRouter) InitLdapRouter(Router *gin.RouterGroup) {
	ldapRouter := Router
	LoginApi := api.ApiGroupApp.LdapApi.LdapLogin
	{
		ldapRouter.POST("login", LoginApi) // ldap 登陆
	}
}
