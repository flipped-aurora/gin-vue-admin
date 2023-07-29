package NestAirlinePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg"
	NestAirlinePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NestAirlineApi struct {
}

var NtAirlineService = service.ServiceGroupApp.NestAirlinePkgServiceGroup.NestAirlineService

// CreateNestAirline 创建NestAirline
// @Tags NestAirline
// @Summary 创建NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestAirlinePkg.NestAirline true "创建NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtAirline/createNestAirline [post]
func (NtAirlineApi *NestAirlineApi) CreateNestAirline(c *gin.Context) {
	var NtAirline NestAirlinePkg.NestAirline
	err := c.ShouldBindJSON(&NtAirline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	NtAirline.CreatedBy = utils.GetUserID(c)
	if err := NtAirlineService.CreateNestAirline(&NtAirline); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNestAirline 删除NestAirline
// @Tags NestAirline
// @Summary 删除NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestAirlinePkg.NestAirline true "删除NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NtAirline/deleteNestAirline [delete]
func (NtAirlineApi *NestAirlineApi) DeleteNestAirline(c *gin.Context) {
	var NtAirline NestAirlinePkg.NestAirline
	err := c.ShouldBindJSON(&NtAirline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	NtAirline.DeletedBy = utils.GetUserID(c)
	if err := NtAirlineService.DeleteNestAirline(NtAirline); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNestAirlineByIds 批量删除NestAirline
// @Tags NestAirline
// @Summary 批量删除NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /NtAirline/deleteNestAirlineByIds [delete]
func (NtAirlineApi *NestAirlineApi) DeleteNestAirlineByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := NtAirlineService.DeleteNestAirlineByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNestAirline 更新NestAirline
// @Tags NestAirline
// @Summary 更新NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestAirlinePkg.NestAirline true "更新NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NtAirline/updateNestAirline [put]
func (NtAirlineApi *NestAirlineApi) UpdateNestAirline(c *gin.Context) {
	var NtAirline NestAirlinePkg.NestAirline
	err := c.ShouldBindJSON(&NtAirline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	NtAirline.UpdatedBy = utils.GetUserID(c)
	if err := NtAirlineService.UpdateNestAirline(NtAirline); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNestAirline 用id查询NestAirline
// @Tags NestAirline
// @Summary 用id查询NestAirline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestAirlinePkg.NestAirline true "用id查询NestAirline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NtAirline/findNestAirline [get]
func (NtAirlineApi *NestAirlineApi) FindNestAirline(c *gin.Context) {
	var NtAirline NestAirlinePkg.NestAirline
	err := c.ShouldBindQuery(&NtAirline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reNtAirline, err := NtAirlineService.GetNestAirline(NtAirline.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reNtAirline": reNtAirline}, c)
	}
}

// GetNestAirlineList 分页获取NestAirline列表
// @Tags NestAirline
// @Summary 分页获取NestAirline列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestAirlinePkgReq.NestAirlineSearch true "分页获取NestAirline列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NtAirline/getNestAirlineList [get]
func (NtAirlineApi *NestAirlineApi) GetNestAirlineList(c *gin.Context) {
	var pageInfo NestAirlinePkgReq.NestAirlineSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := NtAirlineService.GetNestAirlineInfoList(pageInfo, c); err != nil {
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
