package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.POST("copyAuthority", authorityApi.CopyAuthority)       // 拷贝角色
		authorityRouter.POST("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
