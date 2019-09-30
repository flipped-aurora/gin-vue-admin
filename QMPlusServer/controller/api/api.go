package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/model/dbModel"
	"main/model/modelInterface"
)

type CreateApiParams struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

type DeleteApiParams struct {
	Path uint `json:"path"`
}

// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api dbModel.Api
	_ = c.BindJSON(&api)
	err := api.CreateApi()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Api
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.DeleteApiParams true "删除api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var a dbModel.Api
	_ = c.BindJSON(&a)
	err := a.DeleteApi()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

type AuthAndPathIn struct {
	AuthorityId string        `json:"authorityId"`
	ApiIds        []uint     `json:"apiIds"`
}

// @Tags Api
// @Summary 创建api和角色关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.AuthAndPathIn true "创建api和角色关系"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/setAuthAndApi [post]
func SetAuthAndApi(c *gin.Context) {
	var authAndPathIn AuthAndPathIn
	_ = c.BindJSON(&authAndPathIn)
	err := new(dbModel.ApiAuthority).SetAuthAndApi(authAndPathIn.AuthorityId, authAndPathIn.ApiIds)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

// @Tags Api
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取API列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(dbModel.Api).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})

	}
}

// @Tags Api
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo GetById
	_ = c.BindJSON(&idInfo)
	err, api := new(dbModel.Api).GetApiById(idInfo.Id)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"api":     api,
		})

	}
}



// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updataApi [post]
 func UpdataApi(c *gin.Context) {
	 var api dbModel.Api
	 _ = c.BindJSON(&api)
	 err := api.UpdataApi()
	 if err != nil {
		 servers.ReportFormat(c, false, fmt.Sprintf("修改数据失败，%v", err), gin.H{})
	 } else {
		 servers.ReportFormat(c, true, "修改数据成功", gin.H{})
	 }
 }

// @Tags Api
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context){
	err,apis := new(dbModel.Api).GetAllApis()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"apis":     apis,
		})
	}
}