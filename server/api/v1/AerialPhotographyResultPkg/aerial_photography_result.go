package AerialPhotographyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg"
	AerialPhotographyResultPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AerialPhotographyResultApi struct {
}

var ALPhotographyResultService = service.ServiceGroupApp.AerialPhotographyResultPkgServiceGroup.AerialPhotographyResultService

// CreateAerialPhotographyResult 创建AerialPhotographyResult
// @Tags AerialPhotographyResult
// @Summary 创建AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AerialPhotographyResultPkg.AerialPhotographyResult true "创建AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ALPhotographyResult/createAerialPhotographyResult [post]
func (ALPhotographyResultApi *AerialPhotographyResultApi) CreateAerialPhotographyResult(c *gin.Context) {
	var ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult
	err := c.ShouldBindJSON(&ALPhotographyResult)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ALPhotographyResult.CreatedBy = utils.GetUserID(c)
	if err := ALPhotographyResultService.CreateAerialPhotographyResult(&ALPhotographyResult); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAerialPhotographyResult 删除AerialPhotographyResult
// @Tags AerialPhotographyResult
// @Summary 删除AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AerialPhotographyResultPkg.AerialPhotographyResult true "删除AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ALPhotographyResult/deleteAerialPhotographyResult [delete]
func (ALPhotographyResultApi *AerialPhotographyResultApi) DeleteAerialPhotographyResult(c *gin.Context) {
	var ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult
	err := c.ShouldBindJSON(&ALPhotographyResult)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ALPhotographyResult.DeletedBy = utils.GetUserID(c)
	if err := ALPhotographyResultService.DeleteAerialPhotographyResult(ALPhotographyResult); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAerialPhotographyResultByIds 批量删除AerialPhotographyResult
// @Tags AerialPhotographyResult
// @Summary 批量删除AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ALPhotographyResult/deleteAerialPhotographyResultByIds [delete]
func (ALPhotographyResultApi *AerialPhotographyResultApi) DeleteAerialPhotographyResultByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := ALPhotographyResultService.DeleteAerialPhotographyResultByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAerialPhotographyResult 更新AerialPhotographyResult
// @Tags AerialPhotographyResult
// @Summary 更新AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AerialPhotographyResultPkg.AerialPhotographyResult true "更新AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ALPhotographyResult/updateAerialPhotographyResult [put]
func (ALPhotographyResultApi *AerialPhotographyResultApi) UpdateAerialPhotographyResult(c *gin.Context) {
	var ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult
	err := c.ShouldBindJSON(&ALPhotographyResult)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ALPhotographyResult.UpdatedBy = utils.GetUserID(c)
	if err := ALPhotographyResultService.UpdateAerialPhotographyResult(ALPhotographyResult); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAerialPhotographyResult 用id查询AerialPhotographyResult
// @Tags AerialPhotographyResult
// @Summary 用id查询AerialPhotographyResult
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AerialPhotographyResultPkg.AerialPhotographyResult true "用id查询AerialPhotographyResult"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ALPhotographyResult/findAerialPhotographyResult [get]
func (ALPhotographyResultApi *AerialPhotographyResultApi) FindAerialPhotographyResult(c *gin.Context) {
	var ALPhotographyResult AerialPhotographyResultPkg.AerialPhotographyResult
	err := c.ShouldBindQuery(&ALPhotographyResult)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reALPhotographyResult, err := ALPhotographyResultService.GetAerialPhotographyResult(ALPhotographyResult.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reALPhotographyResult": reALPhotographyResult}, c)
	}
}

// GetAerialPhotographyResultList 分页获取AerialPhotographyResult列表
// @Tags AerialPhotographyResult
// @Summary 分页获取AerialPhotographyResult列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AerialPhotographyResultPkgReq.AerialPhotographyResultSearch true "分页获取AerialPhotographyResult列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ALPhotographyResult/getAerialPhotographyResultList [get]
func (ALPhotographyResultApi *AerialPhotographyResultApi) GetAerialPhotographyResultList(c *gin.Context) {
	var pageInfo AerialPhotographyResultPkgReq.AerialPhotographyResultSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ALPhotographyResultService.GetAerialPhotographyResultInfoList(pageInfo, c); err != nil {
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

// GetAerialPhotographyResultList 查询并处理航摄成果数据
// @Tags QueryAerialPhotographyResult
// @Summary 查询并处理航摄成果数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AerialPhotographyResultPkgReq.QueryAerialPhotographyResult true "查询并处理航摄成果数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ALPhotographyResult/queryAerialPhotographyResult [get]
func (ALPhotographyResultApi *AerialPhotographyResultApi) QueryAerialPhotographyResult(c *gin.Context) {
	if modelList, orthoList, err := ALPhotographyResultService.QueryAerialPhotographyResult(c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		returnObj := make(map[string]interface{})
		returnObj["modelList"] = modelList
		returnObj["orthoList"] = orthoList
		response.OkWithData(gin.H{"list": returnObj}, c)
	}
}
