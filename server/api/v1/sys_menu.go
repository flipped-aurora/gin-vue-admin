package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	err, menus := new(model.SysMenu).GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, menuList, total := new(model.SysBaseMenu).GetInfoList(pageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"list":     menuList,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取数据成功", c)
	}
}

// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var addMenu model.SysBaseMenu
	_ = c.ShouldBindJSON(&addMenu)
	err := addMenu.AddBaseMenu()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	err, menus := new(model.SysBaseMenu).GetBaseMenuTree()
	if err != nil {
		response.Result(response.ERROR, gin.H{"menus": menus}, fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", c)

	}
}

type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var addMenuAuthorityInfo AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&addMenuAuthorityInfo)

	err := new(model.SysMenu).AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityId)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "添加成功", c)
	}
}

type AuthorityIdInfo struct {
	AuthorityId string
}

// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var authorityIdInfo AuthorityIdInfo
	_ = c.ShouldBindJSON(&authorityIdInfo)
	err, menus := new(model.SysMenu).GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		response.Result(response.ERROR, gin.H{"menus": menus}, fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", c)
	}
}

type IdInfo struct {
	Id float64
}

// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.IdInfo true "删除菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var idInfo IdInfo
	_ = c.ShouldBindJSON(&idInfo)
	err := new(model.SysBaseMenu).DeleteBaseMenu(idInfo.Id)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "删除成功", c)

	}
}

// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	err := menu.UpdateBaseMenu()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("修改失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "修改成功", c)
	}
}

type GetById struct {
	Id float64 `json:"id"`
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(c *gin.Context) {
	var idInfo GetById
	_ = c.ShouldBindJSON(&idInfo)
	err, menu := new(model.SysBaseMenu).GetBaseMenuById(idInfo.Id)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("查询失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"menu": menu}, "查询成功", c)
	}
}
