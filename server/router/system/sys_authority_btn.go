package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router *gin.RouterGroup) {
	//authorityRouter := Router.Group("authorityBtn").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouterWithoutRecord.POST("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn)
	}
}
