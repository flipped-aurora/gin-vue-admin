package reservation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/reservation"
	reservationReq "github.com/flipped-aurora/gin-vue-admin/server/model/reservation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EqtInfoApi struct {
}

var eqtInfoService = service.ServiceGroupApp.ReservationServiceGroup.EqtInfoService

// CreateEqtInfo 创建EqtInfo
// @Tags EqtInfo
// @Summary 创建EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reservation.EqtInfo true "创建EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eqtInfo/createEqtInfo [post]
func (eqtInfoApi *EqtInfoApi) CreateEqtInfo(c *gin.Context) {
	var eqtInfo reservation.EqtInfo
	err := c.ShouldBindJSON(&eqtInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	eqtInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"EqtClass":       {utils.NotEmpty()},
		"EqtCardNo":      {utils.NotEmpty()},
		"EqtName":        {utils.NotEmpty()},
		"EqtStatus":      {utils.NotEmpty()},
		"EqtStockStatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(eqtInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eqtInfoService.CreateEqtInfo(&eqtInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEqtInfo 删除EqtInfo
// @Tags EqtInfo
// @Summary 删除EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reservation.EqtInfo true "删除EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eqtInfo/deleteEqtInfo [delete]
func (eqtInfoApi *EqtInfoApi) DeleteEqtInfo(c *gin.Context) {
	var eqtInfo reservation.EqtInfo
	err := c.ShouldBindJSON(&eqtInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	eqtInfo.DeletedBy = utils.GetUserID(c)
	if err := eqtInfoService.DeleteEqtInfo(eqtInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEqtInfoByIds 批量删除EqtInfo
// @Tags EqtInfo
// @Summary 批量删除EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eqtInfo/deleteEqtInfoByIds [delete]
func (eqtInfoApi *EqtInfoApi) DeleteEqtInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := eqtInfoService.DeleteEqtInfoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEqtInfo 更新EqtInfo
// @Tags EqtInfo
// @Summary 更新EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reservation.EqtInfo true "更新EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eqtInfo/updateEqtInfo [put]
func (eqtInfoApi *EqtInfoApi) UpdateEqtInfo(c *gin.Context) {
	var eqtInfo reservation.EqtInfo
	err := c.ShouldBindJSON(&eqtInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	eqtInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"EqtClass":       {utils.NotEmpty()},
		"EqtCardNo":      {utils.NotEmpty()},
		"EqtName":        {utils.NotEmpty()},
		"EqtStatus":      {utils.NotEmpty()},
		"EqtStockStatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(eqtInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eqtInfoService.UpdateEqtInfo(eqtInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEqtInfo 用id查询EqtInfo
// @Tags EqtInfo
// @Summary 用id查询EqtInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reservation.EqtInfo true "用id查询EqtInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eqtInfo/findEqtInfo [get]
func (eqtInfoApi *EqtInfoApi) FindEqtInfo(c *gin.Context) {
	var eqtInfo reservation.EqtInfo
	err := c.ShouldBindQuery(&eqtInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeqtInfo, err := eqtInfoService.GetEqtInfo(eqtInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeqtInfo": reeqtInfo}, c)
	}
}

// GetEqtInfoList 分页获取EqtInfo列表
// @Tags EqtInfo
// @Summary 分页获取EqtInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reservationReq.EqtInfoSearch true "分页获取EqtInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eqtInfo/getEqtInfoList [get]
func (eqtInfoApi *EqtInfoApi) GetEqtInfoList(c *gin.Context) {
	var pageInfo reservationReq.EqtInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eqtInfoService.GetEqtInfoInfoList(pageInfo); err != nil {
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
