package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopOrdersApi struct {
}

var shopOrdersService = service.ServiceGroupApp.ShopServiceGroup.ShopOrdersService

// CreateShopOrders 创建shopOrders表
// @Tags ShopOrders
// @Summary 创建shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopOrders true "创建shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopOrders/createShopOrders [post]
func (shopOrdersApi *ShopOrdersApi) CreateShopOrders(c *gin.Context) {
	var shopOrders shop.ShopOrders
	err := c.ShouldBindJSON(&shopOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopOrdersService.CreateShopOrders(&shopOrders); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShopOrders 删除shopOrders表
// @Tags ShopOrders
// @Summary 删除shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopOrders true "删除shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopOrders/deleteShopOrders [delete]
func (shopOrdersApi *ShopOrdersApi) DeleteShopOrders(c *gin.Context) {
	var shopOrders shop.ShopOrders
	err := c.ShouldBindJSON(&shopOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopOrdersService.DeleteShopOrders(shopOrders); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShopOrdersByIds 批量删除shopOrders表
// @Tags ShopOrders
// @Summary 批量删除shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopOrders/deleteShopOrdersByIds [delete]
func (shopOrdersApi *ShopOrdersApi) DeleteShopOrdersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopOrdersService.DeleteShopOrdersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShopOrders 更新shopOrders表
// @Tags ShopOrders
// @Summary 更新shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopOrders true "更新shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopOrders/updateShopOrders [put]
func (shopOrdersApi *ShopOrdersApi) UpdateShopOrders(c *gin.Context) {
	var shopOrders shop.ShopOrders
	err := c.ShouldBindJSON(&shopOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopOrdersService.UpdateShopOrders(shopOrders); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShopOrders 用id查询shopOrders表
// @Tags ShopOrders
// @Summary 用id查询shopOrders表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopOrders true "用id查询shopOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopOrders/findShopOrders [get]
func (shopOrdersApi *ShopOrdersApi) FindShopOrders(c *gin.Context) {
	var shopOrders shop.ShopOrders
	err := c.ShouldBindQuery(&shopOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshopOrders, err := shopOrdersService.GetShopOrders(shopOrders.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopOrders": reshopOrders}, c)
	}
}

// GetShopOrdersList 分页获取shopOrders表列表
// @Tags ShopOrders
// @Summary 分页获取shopOrders表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.ShopOrdersSearch true "分页获取shopOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopOrders/getShopOrdersList [get]
func (shopOrdersApi *ShopOrdersApi) GetShopOrdersList(c *gin.Context) {
	var pageInfo shopReq.ShopOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopOrdersService.GetShopOrdersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
		return
	}
}

// RefundShopOrders 用订单号退款
// @Tags ShopOrders
// @Summary 用订单号退款
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopOrders true "用订单号退款"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopOrders/findShopOrders [get]
func (shopOrdersApi *ShopOrdersApi) RefundShopOrders(c *gin.Context) {
	var shopOrders shop.ShopOrders
	err := c.ShouldBindQuery(&shopOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(shopOrders.OutTradeNo) <= 0 {
		response.FailWithMessage("参数有误", c)
		return
	}
	order, err := shopOrdersService.RefundShopOrders(shopOrders.OutTradeNo)
	if err != nil {
		global.GVA_LOG.Error("退款失败!", zap.Error(err))
		response.FailWithMessage("退款失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reshopOrders": order}, "退款成功", c)
	}
}
