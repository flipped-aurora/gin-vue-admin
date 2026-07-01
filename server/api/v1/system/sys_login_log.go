package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type LoginLogApi struct{}

// DeleteLoginLog
// @Tags      SysLoginLog
// @Summary   删除登录日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysLoginLog             true  "登录日志ID"
// @Success   200   {object}  response.Response{msg=string}  "删除登录日志"
// @Router    /sysLoginLog/deleteLoginLog [delete]
func (s *LoginLogApi) DeleteLoginLog(c *gin.Context) {
	var loginLog system.SysLoginLog
	err := c.ShouldBindJSON(&loginLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = loginLogService.DeleteLoginLog(c.Request.Context(), loginLog)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLoginLogByIds
// @Tags      SysLoginLog
// @Summary   批量删除登录日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "登录日志ID集合"
// @Success   200   {object}  response.Response{msg=string}  "批量删除登录日志"
// @Router    /sysLoginLog/deleteLoginLogByIds [delete]
func (s *LoginLogApi) DeleteLoginLogByIds(c *gin.Context) {
	var SDS request.IdsReq
	err := c.ShouldBindJSON(&SDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = loginLogService.DeleteLoginLogByIds(c.Request.Context(), SDS)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("批量删除失败!")
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// FindLoginLog
// @Tags      SysLoginLog
// @Summary   用id查询登录日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysLoginLog                            true  "登录日志ID"
// @Success   200   {object}  response.Response{data=system.SysLoginLog,msg=string}  "用id查询登录日志"
// @Router    /sysLoginLog/findLoginLog [get]
func (s *LoginLogApi) FindLoginLog(c *gin.Context) {
	var loginLog system.SysLoginLog
	err := c.ShouldBindQuery(&loginLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reLoginLog, err := loginLogService.GetLoginLog(c.Request.Context(), loginLog.ID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("查询失败!")
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(reLoginLog, "查询成功", c)
}

// GetLoginLogList
// @Tags      SysLoginLog
// @Summary   分页获取登录日志列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     systemReq.SysLoginLogSearch                          true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取登录日志列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysLoginLog/getLoginLogList [get]
func (s *LoginLogApi) GetLoginLogList(c *gin.Context) {
	var pageInfo systemReq.SysLoginLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := loginLogService.GetLoginLogInfoList(c.Request.Context(), pageInfo)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
