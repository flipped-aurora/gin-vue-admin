package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "创建角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	err, authBack := auth.CreateAuthority()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"authority": authBack,
		}, fmt.Sprintf("创建成功，%v", err), c)
	}
}

// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var a model.SysAuthority
	_ = c.ShouldBindJSON(&a)
	//删除角色之前需要判断是否有用户正在使用此角色
	err := a.DeleteAuthority()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "删除失败", c)
	}
}

// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(c *gin.Context) {
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(model.SysAuthority).GetInfoList(pageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取数据成功", c)
	}
}

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func SetDataAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	err := auth.SetDataAuthority()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("设置关联失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "获取数据成功", c)
	}
}
