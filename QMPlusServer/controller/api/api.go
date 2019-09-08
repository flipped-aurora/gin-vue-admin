package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/model/dbModel"
)

type CreateApiParams struct {
	AuthorityId uint   `json:"-"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

type DeleteApiParams struct {
	AuthorityId uint `json:"-"`
}

// @Tags Api
// @Summary 为指定角色创建api
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
// @Summary 删除指定角色api
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
