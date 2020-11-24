package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		MenuRouter.POST("getMenu", v1.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getMenuList", v1.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuById)   // 根据id获取菜单
	}
	return MenuRouter
}
