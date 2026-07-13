// server/api/v1/system/sys_timed_task.go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/sse"
	"github.com/gin-gonic/gin"
)

type TimedTaskApi struct{}

// CreateTimedTask
// @Tags      SysTimedTask
// @Summary   创建定时任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysTimedTask            true  "任务定义"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /timedTask/createTimedTask [post]
func (a *TimedTaskApi) CreateTimedTask(c *gin.Context) {
	var t system.SysTimedTask
	if err := c.ShouldBindJSON(&t); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := timedTaskService.CreateTimedTask(c.Request.Context(), &t); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("创建失败!")
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateTimedTask
// @Tags      SysTimedTask
// @Summary   更新定时任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysTimedTask            true  "任务定义(含ID)"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /timedTask/updateTimedTask [put]
func (a *TimedTaskApi) UpdateTimedTask(c *gin.Context) {
	var t system.SysTimedTask
	if err := c.ShouldBindJSON(&t); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := timedTaskService.UpdateTimedTask(c.Request.Context(), &t); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("更新失败!")
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteTimedTask
// @Tags      SysTimedTask
// @Summary   删除定时任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.TriggerTimedTask     true  "任务ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /timedTask/deleteTimedTask [delete]
func (a *TimedTaskApi) DeleteTimedTask(c *gin.Context) {
	var req systemReq.TriggerTimedTask
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := timedTaskService.DeleteTimedTask(c.Request.Context(), req.ID); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ToggleTimedTask
// @Tags      SysTimedTask
// @Summary   启用/停用定时任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.ToggleTimedTask      true  "ID与目标状态"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /timedTask/toggleTimedTask [post]
func (a *TimedTaskApi) ToggleTimedTask(c *gin.Context) {
	var req systemReq.ToggleTimedTask
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := timedTaskService.ToggleTimedTask(c.Request.Context(), req.ID, req.Enabled); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("启停失败!")
		response.FailWithMessage("操作失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// TriggerTimedTask
// @Tags      SysTimedTask
// @Summary   手动触发一次执行
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.TriggerTimedTask     true  "任务ID"
// @Success   200   {object}  response.Response{msg=string}  "已触发"
// @Router    /timedTask/triggerTimedTask [post]
func (a *TimedTaskApi) TriggerTimedTask(c *gin.Context) {
	var req systemReq.TriggerTimedTask
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := timedTaskService.TriggerTimedTask(c.Request.Context(), req.ID); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("手动触发失败!")
		response.FailWithMessage("触发失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("已触发, 执行结果见执行日志", c)
}

// GetTimedTaskList
// @Tags      SysTimedTask
// @Summary   分页获取定时任务列表(含下次执行时间)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     systemReq.SysTimedTaskSearch                            true  "分页与筛选"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /timedTask/getTimedTaskList [get]
func (a *TimedTaskApi) GetTimedTaskList(c *gin.Context) {
	var pageInfo systemReq.SysTimedTaskSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := timedTaskService.GetTimedTaskList(c.Request.Context(), pageInfo)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("获取失败!")
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

// GetTimedTaskLogList
// @Tags      SysTimedTask
// @Summary   分页获取执行日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     systemReq.SysTimedTaskLogSearch                         true  "分页与筛选"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /timedTask/getTimedTaskLogList [get]
func (a *TimedTaskApi) GetTimedTaskLogList(c *gin.Context) {
	var pageInfo systemReq.SysTimedTaskLogSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := timedTaskService.GetTimedTaskLogList(c.Request.Context(), pageInfo)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("timedTask").Err(err).Error("获取执行日志失败!")
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

// GetRegisteredMethods
// @Tags      SysTimedTask
// @Summary   获取已注册方法列表(供下拉选择)
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取成功"
// @Router    /timedTask/getRegisteredMethods [get]
func (a *TimedTaskApi) GetRegisteredMethods(c *gin.Context) {
	response.OkWithDetailed(gin.H{"methods": timedTaskService.GetRegisteredMethods()}, "获取成功", c)
}

// AlertStream 失败告警 SSE 订阅(全仓首个业务级 Hub.Stream 挂载点)。
// 注意: 本路由绝不能套 TimeoutMiddleware(见 utils/sse/hub.go Stream 注释);
// nginx 反代缓冲已由 Hub 设置 X-Accel-Buffering: no 处理。
// @Tags      SysTimedTask
// @Summary   订阅定时任务失败告警(SSE)
// @Security  ApiKeyAuth
// @Produce   text/event-stream
// @Success   200  {string}  string  "SSE 流"
// @Router    /timedTask/alertStream [get]
func (a *TimedTaskApi) AlertStream(c *gin.Context) {
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户身份", c)
		return
	}
	sse.Default().Stream(c, userID, 0)
}
