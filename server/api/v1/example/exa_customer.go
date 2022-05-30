package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct{}

// @Tags ExaCustomer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.ExaCustomer true "客户用户名, 客户手机号码"
// @Success 200 {object} response.Response{msg=string} "创建客户"
// @Router /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if err := customerService.CreateExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
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
// @Param data body example.ExaCustomer true "客户ID"
// @Success 200 {object} response.Response{msg=string} "删除客户"
// @Router /customer/customer [delete]
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.DeleteExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
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
// @Param data body example.ExaCustomer true "客户ID, 客户信息"
// @Success 200 {object} response.Response{msg=string} "更新客户信息"
// @Router /customer/customer [put]
func (e *CustomerApi) UpdateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.UpdateExaCustomer(&customer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query example.ExaCustomer true "客户ID"
// @Success 200 {object} response.Response{data=exampleRes.ExaCustomerResponse,msg=string} "获取单一客户信息,返回包括客户详情"
// @Router /customer/customer [get]
func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(exampleRes.ExaCustomerResponse{Customer: data}, "获取成功", c)
	}
}

// @Tags ExaCustomer
// @Summary 分页获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router /customer/customerList [get]
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
