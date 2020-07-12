package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err, menus := service.GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.SysMenusResponse{Menus: menus}, c)
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, menuList, total := service.GetInfoList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	MenuVerify := utils.Rules{
		"Path":      {utils.NotEmpty()},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0")},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		response.FailWithMessage(MetaVerifyErr.Error(), c)
		return
	}
	err := service.AddBaseMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	err, menus := service.GetBaseMenuTree()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.SysBaseMenusResponse{Menus: menus}, c)
	}
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var addMenuAuthorityInfo request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&addMenuAuthorityInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(addMenuAuthorityInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err := service.AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var authorityIdInfo request.AuthorityIdInfo
	_ = c.ShouldBindJSON(&authorityIdInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(authorityIdInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err, menus := service.GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.SysMenusResponse{Menus: menus}, fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.DeleteBaseMenu(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)

	}
}

// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	MenuVerify := utils.Rules{
		"Path":      {"notEmpty"},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0")},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		response.FailWithMessage(MetaVerifyErr.Error(), c)
		return
	}
	err := service.UpdateBaseMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改失败：%v", err), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	MenuVerify := utils.Rules{
		"Id": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(idInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err, menu := service.GetBaseMenuById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败：%v", err), c)
	} else {
		response.OkWithData(resp.SysBaseMenuResponse{Menu: menu}, c)
	}
}
