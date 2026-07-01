package system

import (
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type AIWorkflowSessionApi struct{}

// Save
// @Tags      AIWorkflowSession
// @Summary   保存AI工作流会话
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysAIWorkflowSessionUpsert  true  "AI工作流会话信息"
// @Success   200   {object}  response.Response{data=object,msg=string}  "保存AI工作流会话"
// @Router    /autoCode/saveAIWorkflowSession [post]
func (a *AIWorkflowSessionApi) Save(c *gin.Context) {
	var info systemReq.SysAIWorkflowSessionUpsert
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	session, err := aiWorkflowSessionService.Save(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("保存 AI 工作流会话失败")
		response.FailWithMessage("保存会话失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"session": session}, "保存成功", c)
}

// GetList
// @Tags      AIWorkflowSession
// @Summary   分页获取AI工作流会话列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysAIWorkflowSessionSearch  true  "页码, 每页大小, 分组Tab"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取AI工作流会话列表"
// @Router    /autoCode/getAIWorkflowSessionList [post]
func (a *AIWorkflowSessionApi) GetList(c *gin.Context) {
	var info systemReq.SysAIWorkflowSessionSearch
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := aiWorkflowSessionService.GetList(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取 AI 工作流会话列表失败")
		response.FailWithMessage("获取会话列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

// GetDetail
// @Tags      AIWorkflowSession
// @Summary   获取AI工作流会话详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                            true  "会话ID"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取AI工作流会话详情"
// @Router    /autoCode/getAIWorkflowSessionDetail [post]
func (a *AIWorkflowSessionApi) GetDetail(c *gin.Context) {
	var info commonReq.GetById
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	session, err := aiWorkflowSessionService.GetDetail(c.Request.Context(), utils.GetUserID(c), info.Uint())
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取 AI 工作流会话详情失败")
		response.FailWithMessage("获取会话详情失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"session": session}, "获取成功", c)
}

// Delete
// @Tags      AIWorkflowSession
// @Summary   删除AI工作流会话
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "会话ID"
// @Success   200   {object}  response.Response{msg=string}  "删除AI工作流会话"
// @Router    /autoCode/deleteAIWorkflowSession [post]
func (a *AIWorkflowSessionApi) Delete(c *gin.Context) {
	var info commonReq.GetById
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aiWorkflowSessionService.Delete(c.Request.Context(), utils.GetUserID(c), info.Uint()); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除 AI 工作流会话失败")
		response.FailWithMessage("删除会话失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// DumpMarkdown
// @Tags      AIWorkflowSession
// @Summary   导出AI工作流Markdown落盘
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SysAIWorkflowMarkdownDump        true  "AI工作流Markdown落盘信息"
// @Success   200   {object}  response.Response{data=object,msg=string}  "导出AI工作流Markdown落盘"
// @Router    /autoCode/dumpAIWorkflowMarkdown [post]
func (a *AIWorkflowSessionApi) DumpMarkdown(c *gin.Context) {
	var info systemReq.SysAIWorkflowMarkdownDump
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	result, err := aiWorkflowSessionService.DumpMarkdown(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("AI 工作流 Markdown 落盘失败")
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"result": result}, "落盘成功", c)
}
