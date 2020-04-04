package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
)

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createExaCustomer [post]
func CreateExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	cu.SysUserID = waitUse.ID
	cu.SysUserAuthorityID = waitUse.AuthorityId
	err := cu.CreateExaCustomer()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("创建失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/deleteExaCustomer [post]
func DeleteExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	err := cu.DeleteExaCustomer()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "删除成功", c)
	}
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/updateExaCustomer [post]
func UpdateExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	err := cu.UpdateExaCustomer()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("更新失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "更新成功", c)
	}
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getExaCustomer [post]
func GetExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	err, customer := cu.GetExaCustomer()
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"customer": customer,
		}, "创建", c)
	}
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getExaCustomerList [post]
func GetExaCustomerList(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.CustomClaims)
	var cu model.ExaCustomer
	cu.SysUserAuthorityID = waitUse.AuthorityId
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, customerList, total := cu.GetInfoList(pageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("创建失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"customer": customerList,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "创建成功", c)
	}
}
