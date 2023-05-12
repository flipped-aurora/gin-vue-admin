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
)

type CroppingRecordApi struct {
}

var croppingRecordService = service.ServiceGroupApp.ClothingServiceGroup.CroppingRecordService

// CreateCroppingRecord 创建CroppingRecord
// @Tags CroppingRecord
// @Summary 创建CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.CroppingRecord true "创建CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /croppingRecord/createCroppingRecord [post]
func (croppingRecordApi *CroppingRecordApi) CreateCroppingRecord(c *gin.Context) {
	var croppingRecord clothing.CroppingRecord
	err := c.ShouldBindJSON(&croppingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	croppingRecord.CreatedBy = utils.GetUserID(c)
	croppingRecord.UserID = utils.GetUserID(c)
	croppingRecord.Step = enum.CroppingPending
	if !userRoleService.CheckManager(croppingRecord.CreatedBy, croppingRecord.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := croppingRecordService.CreateCroppingRecord(&croppingRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (croppingRecordApi *CroppingRecordApi) DeleteCroppingRecord(c *gin.Context) {
	var croppingRecord clothing.CroppingRecord
	err := c.ShouldBindJSON(&croppingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	croppingRecord.DeletedBy = utils.GetUserID(c)
	if err := croppingRecordService.DeleteCroppingRecord(croppingRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (croppingRecordApi *CroppingRecordApi) DeleteCroppingRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := croppingRecordService.DeleteCroppingRecordByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCroppingRecord 更新CroppingRecord
// @Tags CroppingRecord
// @Summary 更新CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.CroppingRecord true "更新CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /croppingRecord/updateCroppingRecord [put]
func (croppingRecordApi *CroppingRecordApi) UpdateCroppingRecord(c *gin.Context) {
	var croppingRecord clothing.CroppingRecord
	err := c.ShouldBindJSON(&croppingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	croppingRecord.UpdatedBy = utils.GetUserID(c)
	if !userRoleService.CheckManager(croppingRecord.CreatedBy, croppingRecord.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if err := croppingRecordService.UpdateCroppingRecord(croppingRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCroppingRecord 用id查询CroppingRecord
// @Tags CroppingRecord
// @Summary 用id查询CroppingRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.CroppingRecord true "用id查询CroppingRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /croppingRecord/findCroppingRecord [get]
func (croppingRecordApi *CroppingRecordApi) FindCroppingRecord(c *gin.Context) {
	var croppingRecord clothing.CroppingRecord
	err := c.ShouldBindQuery(&croppingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recroppingRecord, err := croppingRecordService.GetCroppingRecord(croppingRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recroppingRecord": recroppingRecord}, c)
	}
}

// GetCroppingRecordList 分页获取CroppingRecord列表
// @Tags CroppingRecord
// @Summary 分页获取CroppingRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.CroppingRecordSearch true "分页获取CroppingRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /croppingRecord/getCroppingRecordList [get]
func (croppingRecordApi *CroppingRecordApi) GetCroppingRecordList(c *gin.Context) {
	var pageInfo clothingReq.CroppingRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !userRoleService.CheckStaff(utils.GetUserID(c), pageInfo.CompanyID) {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("权限不足", c)
		return
	}
	if list, total, err := croppingRecordService.GetCroppingRecordInfoList(pageInfo); err != nil {
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
