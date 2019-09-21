package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/middleware"
	"main/model/dbModel"
	"main/model/modelInterface"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "可以什么都不填"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	err, menus := new(dbModel.Menu).GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, menuList, total := new(dbModel.BaseMenu).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list": menuList,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var addMenu dbModel.BaseMenu
	_ = c.BindJSON(&addMenu)
	err := addMenu.AddBaseMenu()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("添加成功，%v", err), gin.H{})
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "可以什么都不填"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	err, menus := new(dbModel.BaseMenu).GetBaseMenuTree()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}

type AddMenuAuthorityInfo struct {
	Menus       []dbModel.BaseMenu
	AuthorityId string
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var addMenuAuthorityInfo AddMenuAuthorityInfo
	_ = c.BindJSON(&addMenuAuthorityInfo)
	err := new(dbModel.Menu).AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("添加成功，%v", err), gin.H{})
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
// @Param data body modelInterface.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var authorityIdInfo AuthorityIdInfo
	_ = c.BindJSON(&authorityIdInfo)
	err, menus := new(dbModel.Menu).GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{"menus": menus})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"menus": menus})
	}
}


type IdInfo struct {
	Id float64
}
// @Tags menu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.IdInfo true "删除菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var idInfo IdInfo
	_ = c.BindJSON(&idInfo)
	err :=new(dbModel.BaseMenu).DeleteBaseMenu(idInfo.Id)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}