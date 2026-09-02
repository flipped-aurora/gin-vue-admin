package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysParamsApi struct{}

// CreateSysParams 创建参数
// @Tags SysParams
// @Summary 创建参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysParams true "创建参数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysParams/createSysParams [post]
func (sysParamsApi *SysParamsApi) CreateSysParams(c *gin.Context) {
	var sysParams system.SysParams
	err := c.ShouldBindJSON(&sysParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysParamsService.CreateSysParams(&sysParams)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysParams 删除参数
// @Tags SysParams
// @Summary 删除参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysParams true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysParams/deleteSysParams [delete]
func (sysParamsApi *SysParamsApi) DeleteSysParams(c *gin.Context) {
	ID := c.Query("ID")
	err := sysParamsService.DeleteSysParams(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysParamsByIds 批量删除参数
// @Tags SysParams
// @Summary 批量删除参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysParams/deleteSysParamsByIds [delete]
func (sysParamsApi *SysParamsApi) DeleteSysParamsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := sysParamsService.DeleteSysParamsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysParams 更新参数
// @Tags SysParams
// @Summary 更新参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysParams true "更新参数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysParams/updateSysParams [put]
func (sysParamsApi *SysParamsApi) UpdateSysParams(c *gin.Context) {
	var sysParams system.SysParams
	err := c.ShouldBindJSON(&sysParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysParamsService.UpdateSysParams(sysParams)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysParams 用id查询参数
// @Tags SysParams
// @Summary 用id查询参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysParams true "用id查询参数"
// @Success 200 {object} response.Response{data=system.SysParams,msg=string} "查询成功"
// @Router /sysParams/findSysParams [get]
func (sysParamsApi *SysParamsApi) FindSysParams(c *gin.Context) {
	ID := c.Query("ID")
	resysParams, err := sysParamsService.GetSysParams(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysParams, c)
}

// GetSysParamsList 分页获取参数列表
// @Tags SysParams
// @Summary 分页获取参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysParamsSearch true "分页获取参数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysParams/getSysParamsList [get]
func (sysParamsApi *SysParamsApi) GetSysParamsList(c *gin.Context) {
	var pageInfo systemReq.SysParamsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysParamsService.GetSysParamsInfoList(pageInfo)
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

// GetSysParam 根据key获取参数value
// @Tags SysParams
// @Summary 根据key获取参数value
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param key query string true "key"
// @Success 200 {object} response.Response{data=system.SysParams,msg=string} "获取成功"
// @Router /sysParams/getSysParam [get]
func (sysParamsApi *SysParamsApi) GetSysParam(c *gin.Context) {
	k := c.Query("key")
	params, err := sysParamsService.GetSysParam(k)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(params, "获取成功", c)
}
