package orderinfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/orderinfo"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    orderinfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderinfo/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OrderInfoApi struct {
}

var orderInfoService = service.ServiceGroupApp.OrderinfoServiceGroup.OrderInfoService


// CreateOrderInfo 创建OrderInfo
// @Tags OrderInfo
// @Summary 创建OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderinfo.OrderInfo true "创建OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderInfo/createOrderInfo [post]
func (orderInfoApi *OrderInfoApi) CreateOrderInfo(c *gin.Context) {
	var orderInfo orderinfo.OrderInfo
	err := c.ShouldBindJSON(&orderInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderInfoService.CreateOrderInfo(&orderInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderInfo 删除OrderInfo
// @Tags OrderInfo
// @Summary 删除OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderinfo.OrderInfo true "删除OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderInfo/deleteOrderInfo [delete]
func (orderInfoApi *OrderInfoApi) DeleteOrderInfo(c *gin.Context) {
	var orderInfo orderinfo.OrderInfo
	err := c.ShouldBindJSON(&orderInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderInfoService.DeleteOrderInfo(orderInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderInfoByIds 批量删除OrderInfo
// @Tags OrderInfo
// @Summary 批量删除OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderInfo/deleteOrderInfoByIds [delete]
func (orderInfoApi *OrderInfoApi) DeleteOrderInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderInfoService.DeleteOrderInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderInfo 更新OrderInfo
// @Tags OrderInfo
// @Summary 更新OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderinfo.OrderInfo true "更新OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderInfo/updateOrderInfo [put]
func (orderInfoApi *OrderInfoApi) UpdateOrderInfo(c *gin.Context) {
	var orderInfo orderinfo.OrderInfo
	err := c.ShouldBindJSON(&orderInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderInfoService.UpdateOrderInfo(orderInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderInfo 用id查询OrderInfo
// @Tags OrderInfo
// @Summary 用id查询OrderInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderinfo.OrderInfo true "用id查询OrderInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderInfo/findOrderInfo [get]
func (orderInfoApi *OrderInfoApi) FindOrderInfo(c *gin.Context) {
	var orderInfo orderinfo.OrderInfo
	err := c.ShouldBindQuery(&orderInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorderInfo, err := orderInfoService.GetOrderInfo(orderInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderInfo": reorderInfo}, c)
	}
}

// GetOrderInfoList 分页获取OrderInfo列表
// @Tags OrderInfo
// @Summary 分页获取OrderInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderinfoReq.OrderInfoSearch true "分页获取OrderInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderInfo/getOrderInfoList [get]
func (orderInfoApi *OrderInfoApi) GetOrderInfoList(c *gin.Context) {
	var pageInfo orderinfoReq.OrderInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderInfoService.GetOrderInfoInfoList(pageInfo); err != nil {
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
