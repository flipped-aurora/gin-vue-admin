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

type ShopGoodsApi struct {
}

var shopGoodsService = service.ServiceGroupApp.ShopServiceGroup.ShopGoodsService

// CreateShopGoods 创建shopGoods表
// @Tags ShopGoods
// @Summary 创建shopGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "创建shopGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopGoods/createShopGoods [post]
func (shopGoodsApi *ShopGoodsApi) CreateShopGoods(c *gin.Context) {
	var shopGoods shop.ShopGoods
	err := c.ShouldBindJSON(&shopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopGoodsService.CreateShopGoods(&shopGoods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShopGoods 删除shopGoods表
// @Tags ShopGoods
// @Summary 删除shopGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "删除shopGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopGoods/deleteShopGoods [delete]
func (shopGoodsApi *ShopGoodsApi) DeleteShopGoods(c *gin.Context) {
	var shopGoods shop.ShopGoods
	err := c.ShouldBindJSON(&shopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopGoodsService.DeleteShopGoods(shopGoods); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShopGoodsByIds 批量删除shopGoods表
// @Tags ShopGoods
// @Summary 批量删除shopGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shopGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopGoods/deleteShopGoodsByIds [delete]
func (shopGoodsApi *ShopGoodsApi) DeleteShopGoodsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopGoodsService.DeleteShopGoodsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShopGoods 更新shopGoods表
// @Tags ShopGoods
// @Summary 更新shopGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "更新shopGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopGoods/updateShopGoods [put]
func (shopGoodsApi *ShopGoodsApi) UpdateShopGoods(c *gin.Context) {
	var shopGoods shop.ShopGoods
	err := c.ShouldBindJSON(&shopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopGoodsService.UpdateShopGoods(shopGoods); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShopGoods 用id查询shopGoods表
// @Tags ShopGoods
// @Summary 用id查询shopGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopGoods true "用id查询shopGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopGoods/findShopGoods [get]
func (shopGoodsApi *ShopGoodsApi) FindShopGoods(c *gin.Context) {
	var shopGoods shop.ShopGoods
	err := c.ShouldBindQuery(&shopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshopGoods, err := shopGoodsService.GetShopGoods(shopGoods.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopGoods": reshopGoods}, c)
	}
}

// GetShopGoodsList 分页获取shopGoods表列表
// @Tags ShopGoods
// @Summary 分页获取shopGoods表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.ShopGoodsSearch true "分页获取shopGoods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopGoods/getShopGoodsList [get]
func (shopGoodsApi *ShopGoodsApi) GetShopGoodsList(c *gin.Context) {
	var pageInfo shopReq.ShopGoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopGoodsService.GetShopGoodsInfoList(pageInfo); err != nil {
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
