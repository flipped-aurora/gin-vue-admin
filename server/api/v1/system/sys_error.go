package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysErrorApi struct{}

// CreateSysError 创建错误日志
// @Tags SysError
// @Summary 创建错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "创建错误日志"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysErrpr/createSysError [post]
func (sysErrprApi *SysErrorApi) CreateSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var sysErrpr system.SysError
	err := c.ShouldBindJSON(&sysErrpr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysErrprService.CreateSysError(ctx, &sysErrpr)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysError 删除错误日志
// @Tags SysError
// @Summary 删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "删除错误日志"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysErrpr/deleteSysError [delete]
func (sysErrprApi *SysErrorApi) DeleteSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := sysErrprService.DeleteSysError(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysErrorByIds 批量删除错误日志
// @Tags SysError
// @Summary 批量删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysErrpr/deleteSysErrorByIds [delete]
func (sysErrprApi *SysErrorApi) DeleteSysErrorByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := sysErrprService.DeleteSysErrorByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysError 更新错误日志
// @Tags SysError
// @Summary 更新错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "更新错误日志"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysErrpr/updateSysError [put]
func (sysErrprApi *SysErrorApi) UpdateSysError(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var sysErrpr system.SysError
	err := c.ShouldBindJSON(&sysErrpr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysErrprService.UpdateSysError(ctx, sysErrpr)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysError 用id查询错误日志
// @Tags SysError
// @Summary 用id查询错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询错误日志"
// @Success 200 {object} response.Response{data=system.SysError,msg=string} "查询成功"
// @Router /sysErrpr/findSysError [get]
func (sysErrprApi *SysErrorApi) FindSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resysErrpr, err := sysErrprService.GetSysError(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysErrpr, c)
}

// GetSysErrorList 分页获取错误日志列表
// @Tags SysError
// @Summary 分页获取错误日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysErrorSearch true "分页获取错误日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysErrpr/getSysErrorList [get]
func (sysErrprApi *SysErrorApi) GetSysErrorList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.SysErrorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysErrprService.GetSysErrorInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
