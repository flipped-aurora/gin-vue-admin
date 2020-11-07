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

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	if err, menus := service.GetMenuTree(getUserAuthorityId(c)); err != nil {
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
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"]); err != nil {
		response.FailWithMessage(err.Error(), c)
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
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
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
	if err := utils.Verify(menu, MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.Rules{"Title": {utils.NotEmpty()}}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddBaseMenu(menu); err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	if err, menus := service.GetBaseMenuTree(); err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.SysBaseMenusResponse{Menus: menus}, c)
	}
}

// @Tags AuthorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var params request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.Rules{"AuthorityId": {"notEmpty"}}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddMenuAuthority(params.Menus, params.AuthorityId); err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags AuthorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.Rules{"AuthorityId": {"notEmpty"}}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := service.GetMenuAuthority(&param); err != nil {
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
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var param request.GetById
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.CustomizeMap["IdVerify"]); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBaseMenu(param.Id); err != nil {
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
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
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
	if err := utils.Verify(menu, MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.Rules{"Title": {utils.NotEmpty()}}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateBaseMenu(menu); err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败：%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.Rules{"Id": {"notEmpty"}}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := service.GetBaseMenuById(idInfo.Id); err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.SysBaseMenuResponse{Menu: menu}, c)
	}
}
