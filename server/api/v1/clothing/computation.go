package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ComputationApi struct {
}

var computationService = service.ServiceGroupApp.ClothingServiceGroup.ComputationService

// CreateComputation 创建Computation
// @Tags Computation
// @Summary 创建Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Computation true "创建Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /computation/createComputation [post]
func (computationApi *ComputationApi) CreateComputation(c *gin.Context) {
	var computation clothing.Computation
	err := c.ShouldBindJSON(&computation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	computation.CreatedBy = utils.GetUserID(c)
	if err := computationService.CreateComputation(&computation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComputation 删除Computation
// @Tags Computation
// @Summary 删除Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Computation true "删除Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /computation/deleteComputation [delete]
func (computationApi *ComputationApi) DeleteComputation(c *gin.Context) {
	var computation clothing.Computation
	err := c.ShouldBindJSON(&computation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	computation.DeletedBy = utils.GetUserID(c)
	if err := computationService.DeleteComputation(computation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteComputationByIds 批量删除Computation
// @Tags Computation
// @Summary 批量删除Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /computation/deleteComputationByIds [delete]
func (computationApi *ComputationApi) DeleteComputationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := computationService.DeleteComputationByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComputation 更新Computation
// @Tags Computation
// @Summary 更新Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Computation true "更新Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /computation/updateComputation [put]
func (computationApi *ComputationApi) UpdateComputation(c *gin.Context) {
	var computation clothing.Computation
	err := c.ShouldBindJSON(&computation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	computation.UpdatedBy = utils.GetUserID(c)
	if err := computationService.UpdateComputation(computation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComputation 用id查询Computation
// @Tags Computation
// @Summary 用id查询Computation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Computation true "用id查询Computation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /computation/findComputation [get]
func (computationApi *ComputationApi) FindComputation(c *gin.Context) {
	var computation clothing.Computation
	err := c.ShouldBindQuery(&computation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recomputation, err := computationService.GetComputation(computation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomputation": recomputation}, c)
	}
}

// GetComputationList 分页获取Computation列表
// @Tags Computation
// @Summary 分页获取Computation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.ComputationSearch true "分页获取Computation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /computation/getComputationList [get]
func (computationApi *ComputationApi) GetComputationList(c *gin.Context) {
	var pageInfo clothingReq.ComputationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := computationService.GetComputationInfoList(pageInfo); err != nil {
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

func (computationApi *ComputationApi) DoComputation(c *gin.Context) {
	var computation clothing.Computation
	err := c.ShouldBindJSON(&computation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	computation.CreatedBy = utils.GetUserID(c)
	if err := computationService.DoComputation(&computation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
