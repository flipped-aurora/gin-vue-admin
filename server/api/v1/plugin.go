package v1

import (
	"github.com/eyotang/gin-vue-admin/server/global"
    "github.com/eyotang/gin-vue-admin/server/model"
    "github.com/eyotang/gin-vue-admin/server/model/request"
    "github.com/eyotang/gin-vue-admin/server/model/response"
    "github.com/eyotang/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags ProductPlugin
// @Summary 创建ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "创建ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productPlugin/createProductPlugin [post]
func CreateProductPlugin(c *gin.Context) {
	var productPlugin model.ProductPlugin
	_ = c.ShouldBindJSON(&productPlugin)
	if err := service.CreateProductPlugin(productPlugin); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ProductPlugin
// @Summary 删除ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "删除ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productPlugin/deleteProductPlugin [delete]
func DeleteProductPlugin(c *gin.Context) {
	var productPlugin model.ProductPlugin
	_ = c.ShouldBindJSON(&productPlugin)
	if err := service.DeleteProductPlugin(productPlugin); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ProductPlugin
// @Summary 批量删除ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /productPlugin/deleteProductPluginByIds [delete]
func DeleteProductPluginByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteProductPluginByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ProductPlugin
// @Summary 更新ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "更新ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productPlugin/updateProductPlugin [put]
func UpdateProductPlugin(c *gin.Context) {
	var productPlugin model.ProductPlugin
	_ = c.ShouldBindJSON(&productPlugin)
	if err := service.UpdateProductPlugin(productPlugin); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ProductPlugin
// @Summary 用id查询ProductPlugin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductPlugin true "用id查询ProductPlugin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productPlugin/findProductPlugin [get]
func FindProductPlugin(c *gin.Context) {
	var productPlugin model.ProductPlugin
	_ = c.ShouldBindQuery(&productPlugin)
	if err, reproductPlugin := service.GetProductPlugin(productPlugin.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproductPlugin": reproductPlugin}, c)
	}
}

// @Tags ProductPlugin
// @Summary 分页获取ProductPlugin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProductPluginSearch true "分页获取ProductPlugin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productPlugin/getProductPluginList [get]
func GetProductPluginList(c *gin.Context) {
	var pageInfo request.ProductPluginSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetProductPluginInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
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
