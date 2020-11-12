package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = getUserID(c)
	customer.SysUserAuthorityID = getUserAuthorityId(c)
	if err := service.CreateExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
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
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
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
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateExaCustomer(&customer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败!", c)
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
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func GetExaCustomer(c *gin.Context) {
	var customer model.ExaCustomer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, data := service.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.ExaCustomerResponse{Customer: data}, "获取成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 分页获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, customerList, total := service.GetCustomerInfoList(getUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
