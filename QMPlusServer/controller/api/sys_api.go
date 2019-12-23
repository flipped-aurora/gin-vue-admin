package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

type CreateApiParams struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

type DeleteApiParams struct {
	ID uint `json:"id"`
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api sysModel.SysApi
	_ = c.BindJSON(&api)
	err := api.CreateApi()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var a sysModel.SysApi
	_ = c.BindJSON(&a)
	err := a.DeleteApi()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

type AuthAndPathIn struct {
	AuthorityId string `json:"authorityId"`
	ApiIds      []uint `json:"apiIds"`
}

//条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	// 此结构体仅本方法使用
	type searchParams struct {
		sysModel.SysApi
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.SysApi.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		})

	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo GetById
	_ = c.BindJSON(&idInfo)
	err, api := new(sysModel.SysApi).GetApiById(idInfo.Id)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"api": api,
		})

	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updataApi [post]
func UpdataApi(c *gin.Context) {
	var api sysModel.SysApi
	_ = c.BindJSON(&api)
	err := api.UpdataApi()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改数据成功", gin.H{})
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context) {
	err, apis := new(sysModel.SysApi).GetAllApis()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"apis": apis,
		})
	}
}
