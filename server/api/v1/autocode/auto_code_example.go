package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeExampleApi struct{}

var autoCodeExampleService = service.ServiceGroupApp.AutoCodeServiceGroup.AutoCodeExampleService

// @Tags AutoCodeExample
// @Summary 创建AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "AutoCodeExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCodeExample/createAutoCodeExample [post]
func (autoCodeExampleApi *AutoCodeExampleApi) CreateAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.CreateAutoCodeExample(autoCodeExample); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 删除AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "AutoCodeExample模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /autoCodeExample/deleteAutoCodeExample [delete]
func (autoCodeExampleApi *AutoCodeExampleApi) DeleteAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.DeleteAutoCodeExample(autoCodeExample); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 更新AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoCodeExample true "更新AutoCodeExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /autoCodeExample/updateAutoCodeExample [put]
func (autoCodeExampleApi *AutoCodeExampleApi) UpdateAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindJSON(&autoCodeExample)
	if err := autoCodeExampleService.UpdateAutoCodeExample(&autoCodeExample); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 用id查询AutoCodeExample
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.AutoCodeExample true "用id查询AutoCodeExample"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /autoCodeExample/findAutoCodeExample [get]
func (autoCodeExampleApi *AutoCodeExampleApi) FindAutoCodeExample(c *gin.Context) {
	var autoCodeExample autocode.AutoCodeExample
	_ = c.ShouldBindQuery(&autoCodeExample)
	if err := utils.Verify(autoCodeExample, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, reAutoCodeExample := autoCodeExampleService.GetAutoCodeExample(autoCodeExample.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reAutoCodeExample": reAutoCodeExample}, "查询成功", c)
	}
}

// @Tags AutoCodeExample
// @Summary 分页获取AutoCodeExample列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AutoCodeExampleSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCodeExample/getAutoCodeExampleList [get]
func (autoCodeExampleApi *AutoCodeExampleApi) GetAutoCodeExampleList(c *gin.Context) {
	var pageInfo autocodeReq.AutoCodeExampleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := autoCodeExampleService.GetAutoCodeExampleInfoList(pageInfo); err != nil {
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
