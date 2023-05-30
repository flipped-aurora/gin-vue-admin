package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.ClothingServiceGroup.OrderService
var wechatService = service.ServiceGroupApp.ClothingServiceGroup.WechatService
var aliService = service.ServiceGroupApp.ClothingServiceGroup.AliService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var createOrderReq clothingReq.CreateOrderReq
	err := c.ShouldBindJSON(&createOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	option, err := rechargeOptionService.GetRechargeOption(createOrderReq.OptID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var order clothing.Order
	order.OrderNo = utils.GenerateOrderNo()
	order.CompanyID = createOrderReq.CompanyID
	order.UserID = utils.GetUserID(c)
	order.Price = option.Price
	order.Amount = option.Amount
	order.Status = enum.Pending
	order.PayStatus = enum.Pending
	order.CreatedBy = utils.GetUserID(c)
	if err := orderService.CreateOrder(&order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order clothing.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.DeletedBy = utils.GetUserID(c)
	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := orderService.DeleteOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order clothing.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.UpdatedBy = utils.GetUserID(c)
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order clothing.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo clothingReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (orderApi *OrderApi) PayOrder(c *gin.Context) {
	var payOrderReq clothingReq.PayOrderReq
	err := c.ShouldBindJSON(&payOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order, err := orderService.GetOrderByOrderNo(payOrderReq.OrderNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var result interface{}
	Description := "合伙人资格"
	Attach := "合伙人资格购买"
	global.GVA_LOG.Sugar().Info("支付方式：", payOrderReq.PayType)
	switch payOrderReq.PayType {
	case enum.WechatJSAPI:
		createOrder, err := wechatService.CreateJSAPIOrder(payOrderReq.OpenID, global.GVA_CONFIG.Wechat.AppID, Description, Attach, order)
		if err != nil {
			global.GVA_LOG.Error("调起微信支付失败!", zap.Error(err))
			response.FailWithMessage("调起微信支付失败", c)
			return
		}
		order.PayNo = *createOrder.PrepayId
		result = createOrder
	case enum.WechatH5:
		createOrder, err := wechatService.CreateH5Order(global.GVA_CONFIG.Wechat.AppID, Description, Attach, order)
		if err != nil {
			global.GVA_LOG.Error("调起微信支付失败!", zap.Error(err))
			response.FailWithMessage("调起微信支付失败", c)
			return
		}
		order.PayNo = *createOrder.H5Url
		result = *createOrder.H5Url
	case enum.WechatApp:
		createOrder, err := wechatService.CreateAppOrder(global.GVA_CONFIG.Wechat.AppID, Description, Attach, order)
		if err != nil {
			global.GVA_LOG.Error("调起微信支付失败!", zap.Error(err))
			response.FailWithMessage("调起微信支付失败", c)
			return
		}
		order.PayNo = *createOrder.PrepayId
		result = createOrder
	case enum.AliH5:
		h5Order, err := aliService.CreateH5Order(Description, order)
		if err != nil {
			global.GVA_LOG.Error("调起支付宝支付失败!", zap.Error(err))
			response.FailWithMessage("调起支付宝支付失败", c)
			return
		}
		result = h5Order
	case enum.AliApp:
		appOrder, err := aliService.CreateAppOrder(Description, order)
		if err != nil {
			global.GVA_LOG.Error("调起支付宝支付失败!", zap.Error(err))
			response.FailWithMessage("调起支付宝支付失败", c)
			return
		}
		result = appOrder
	}
	global.GVA_LOG.Sugar().Info(result)
	order.PayType = payOrderReq.PayType
	now := time.Now()
	order.PayAt = &now
	if err := orderService.UpdateOrder(order); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}
