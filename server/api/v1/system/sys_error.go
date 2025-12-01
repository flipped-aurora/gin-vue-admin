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
// @Router /sysError/createSysError [post]
func (sysErrorApi *SysErrorApi) CreateSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var sysError system.SysError
	err := c.ShouldBindJSON(&sysError)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysErrorService.CreateSysError(ctx, &sysError)
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
// @Router /sysError/deleteSysError [delete]
func (sysErrorApi *SysErrorApi) DeleteSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := sysErrorService.DeleteSysError(ctx, ID)
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
// @Router /sysError/deleteSysErrorByIds [delete]
func (sysErrorApi *SysErrorApi) DeleteSysErrorByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := sysErrorService.DeleteSysErrorByIds(ctx, IDs)
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
// @Router /sysError/updateSysError [put]
func (sysErrorApi *SysErrorApi) UpdateSysError(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var sysError system.SysError
	err := c.ShouldBindJSON(&sysError)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysErrorService.UpdateSysError(ctx, sysError)
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
// @Router /sysError/findSysError [get]
func (sysErrorApi *SysErrorApi) FindSysError(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resysError, err := sysErrorService.GetSysError(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysError, c)
}

// GetSysErrorList 分页获取错误日志列表
// @Tags SysError
// @Summary 分页获取错误日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysErrorSearch true "分页获取错误日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysError/getSysErrorList [get]
func (sysErrorApi *SysErrorApi) GetSysErrorList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.SysErrorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysErrorService.GetSysErrorInfoList(ctx, pageInfo)
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

// GetSysErrorSolution 触发错误日志的异步处理
// @Tags SysError
// @Summary 根据ID触发处理：标记为处理中，1分钟后自动改为处理完成
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query string true "错误日志ID"
// @Success 200 {object} response.Response{msg=string} "处理已提交"
// @Router /sysError/getSysErrorSolution [get]
func (sysErrorApi *SysErrorApi) GetSysErrorSolution(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 兼容 id 与 ID 两种参数
	ID := c.Query("id")
	if ID == "" {
		response.FailWithMessage("缺少参数: id", c)
		return
	}

	err := sysErrorService.GetSysErrorSolution(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("处理触发失败!", zap.Error(err))
		response.FailWithMessage("处理触发失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("已提交至AI处理", c)
}
