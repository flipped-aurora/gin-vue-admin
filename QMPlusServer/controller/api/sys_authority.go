package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/model/modelInterface"
	"main/model/sysModel"
)

type CreateAuthorityPatams struct {
	AuthorityId   string   `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
}

// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "创建角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var auth sysModel.SysAuthority
	_ = c.ShouldBind(&auth)
	err, authBack := auth.CreateAuthority()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{
			"authority": authBack,
		})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"authority": authBack,
		})
	}
}

type DeleteAuthorityPatams struct {
	AuthorityId uint `json:"authorityId"`
}

// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.DeleteAuthorityPatams true "删除角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var a sysModel.SysAuthority
	_ = c.BindJSON(&a)
	//删除角色之前需要判断是否有用户正在使用此角色
	err := a.DeleteAuthority()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(c *gin.Context){
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(sysModel.SysAuthority).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}


type GetAuthorityId struct {
	AuthorityId string `json:"authorityId"`
}

