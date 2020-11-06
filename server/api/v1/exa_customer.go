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

// @Tags ExaCustomer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户用户名, 客户手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/customer [post]
func CreateExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	if CustomerVerifyErr := utils.Verify(cu, CustomerVerify); CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	cu.SysUserID = getUserID(c)
	cu.SysUserAuthorityID = getUserAuthorityId(c)
	if err := service.CreateExaCustomer(cu); err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/customer [delete]
func DeleteExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	if CustomerVerifyErr := utils.Verify(cu.GVA_MODEL, utils.Rules{"ID": {utils.NotEmpty()}}); CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	if err := service.DeleteExaCustomer(cu); err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID, 客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/customer [put]
func UpdateExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	if IdCustomerVerifyErr := utils.Verify(cu.GVA_MODEL, utils.Rules{"ID": {utils.NotEmpty()}}); IdCustomerVerifyErr != nil {
		response.FailWithMessage(IdCustomerVerifyErr.Error(), c)
		return
	}
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	if CustomerVerifyErr := utils.Verify(cu, CustomerVerify); CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	if err := service.UpdateExaCustomer(&cu); err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败：%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /customer/customer [get]
func GetExaCustomer(c *gin.Context) {
	var cu model.ExaCustomer
	_ = c.ShouldBindQuery(&cu)
	if IdCustomerVerifyErr := utils.Verify(cu.GVA_MODEL, utils.Rules{"ID": {utils.NotEmpty()}}); IdCustomerVerifyErr != nil {
		response.FailWithMessage(IdCustomerVerifyErr.Error(), c)
		return
	}
	err, customer := service.GetExaCustomer(cu.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.ExaCustomerResponse{Customer: customer}, c)
	}
}

// @Tags ExaCustomer
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /customer/customerList [get]
func GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"]); PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, customerList, total := service.GetCustomerInfoList(getUserAuthorityId(c), pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
