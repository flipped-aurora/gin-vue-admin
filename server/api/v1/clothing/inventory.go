package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type InventoryApi struct {
}

var inventoryService = service.ServiceGroupApp.ClothingServiceGroup.InventoryService


// CreateInventory 创建Inventory
// @Tags Inventory
// @Summary 创建Inventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Inventory true "创建Inventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventory/createInventory [post]
func (inventoryApi *InventoryApi) CreateInventory(c *gin.Context) {
	var inventory clothing.Inventory
	err := c.ShouldBindJSON(&inventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventoryService.CreateInventory(&inventory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInventory 删除Inventory
// @Tags Inventory
// @Summary 删除Inventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Inventory true "删除Inventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventory/deleteInventory [delete]
func (inventoryApi *InventoryApi) DeleteInventory(c *gin.Context) {
	var inventory clothing.Inventory
	err := c.ShouldBindJSON(&inventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventoryService.DeleteInventory(inventory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInventoryByIds 批量删除Inventory
// @Tags Inventory
// @Summary 批量删除Inventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Inventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inventory/deleteInventoryByIds [delete]
func (inventoryApi *InventoryApi) DeleteInventoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventoryService.DeleteInventoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInventory 更新Inventory
// @Tags Inventory
// @Summary 更新Inventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Inventory true "更新Inventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inventory/updateInventory [put]
func (inventoryApi *InventoryApi) UpdateInventory(c *gin.Context) {
	var inventory clothing.Inventory
	err := c.ShouldBindJSON(&inventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventoryService.UpdateInventory(inventory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInventory 用id查询Inventory
// @Tags Inventory
// @Summary 用id查询Inventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Inventory true "用id查询Inventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inventory/findInventory [get]
func (inventoryApi *InventoryApi) FindInventory(c *gin.Context) {
	var inventory clothing.Inventory
	err := c.ShouldBindQuery(&inventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinventory, err := inventoryService.GetInventory(inventory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinventory": reinventory}, c)
	}
}

// GetInventoryList 分页获取Inventory列表
// @Tags Inventory
// @Summary 分页获取Inventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.InventorySearch true "分页获取Inventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventory/getInventoryList [get]
func (inventoryApi *InventoryApi) GetInventoryList(c *gin.Context) {
	var pageInfo clothingReq.InventorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := inventoryService.GetInventoryInfoList(pageInfo); err != nil {
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
