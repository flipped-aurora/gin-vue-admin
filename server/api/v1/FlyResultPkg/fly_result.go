package FlyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg"
	FlyResultPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FlyResultApi struct {
}

var FlyRtService = service.ServiceGroupApp.FlyResultPkgServiceGroup.FlyResultService

// CreateFlyResult 创建FlyResult
// @Tags FlyResult
// @Summary 创建FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "创建FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/createFlyResult [post]
func (FlyRtApi *FlyResultApi) CreateFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.CreatedBy = utils.GetUserID(c)
	if err := FlyRtService.CreateFlyResult(&FlyRt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFlyResult 删除FlyResult
// @Tags FlyResult
// @Summary 删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FlyRt/deleteFlyResult [delete]
func (FlyRtApi *FlyResultApi) DeleteFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.DeletedBy = utils.GetUserID(c)
	if err := FlyRtService.DeleteFlyResult(FlyRt); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFlyResultByIds 批量删除FlyResult
// @Tags FlyResult
// @Summary 批量删除FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /FlyRt/deleteFlyResultByIds [delete]
func (FlyRtApi *FlyResultApi) DeleteFlyResultByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := FlyRtService.DeleteFlyResultByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFlyResult 更新FlyResult
// @Tags FlyResult
// @Summary 更新FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FlyResultPkg.FlyResult true "更新FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FlyRt/updateFlyResult [put]
func (FlyRtApi *FlyResultApi) UpdateFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindJSON(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	FlyRt.UpdatedBy = utils.GetUserID(c)
	if err := FlyRtService.UpdateFlyResult(FlyRt); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFlyResult 用id查询FlyResult
// @Tags FlyResult
// @Summary 用id查询FlyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkg.FlyResult true "用id查询FlyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FlyRt/findFlyResult [get]
func (FlyRtApi *FlyResultApi) FindFlyResult(c *gin.Context) {
	var FlyRt FlyResultPkg.FlyResult
	err := c.ShouldBindQuery(&FlyRt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reFlyRt, err := FlyRtService.GetFlyResult(FlyRt.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reFlyRt": reFlyRt}, c)
	}
}

// GetFlyResultList 分页获取FlyResult列表
// @Tags FlyResult
// @Summary 分页获取FlyResult列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FlyResultPkgReq.FlyResultSearch true "分页获取FlyResult列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FlyRt/getFlyResultList [get]
func (FlyRtApi *FlyResultApi) GetFlyResultList(c *gin.Context) {
	var pageInfo FlyResultPkgReq.FlyResultSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := FlyRtService.GetFlyResultInfoList(pageInfo); err != nil {
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
