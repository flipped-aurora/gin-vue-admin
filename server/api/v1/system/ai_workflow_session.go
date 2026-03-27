package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AIWorkflowSessionApi struct{}

func (a *AIWorkflowSessionApi) Save(c *gin.Context) {
	var info systemReq.SysAIWorkflowSessionUpsert
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	session, err := aiWorkflowSessionService.Save(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		global.GVA_LOG.Error("保存 AI 工作流会话失败", zap.Error(err))
		response.FailWithMessage("保存会话失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"session": session}, "保存成功", c)
}

func (a *AIWorkflowSessionApi) GetList(c *gin.Context) {
	var info systemReq.SysAIWorkflowSessionSearch
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := aiWorkflowSessionService.GetList(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		global.GVA_LOG.Error("获取 AI 工作流会话列表失败", zap.Error(err))
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

func (a *AIWorkflowSessionApi) GetDetail(c *gin.Context) {
	var info commonReq.GetById
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	session, err := aiWorkflowSessionService.GetDetail(c.Request.Context(), utils.GetUserID(c), info.Uint())
	if err != nil {
		global.GVA_LOG.Error("获取 AI 工作流会话详情失败", zap.Error(err))
		response.FailWithMessage("获取会话详情失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"session": session}, "获取成功", c)
}

func (a *AIWorkflowSessionApi) Delete(c *gin.Context) {
	var info commonReq.GetById
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aiWorkflowSessionService.Delete(c.Request.Context(), utils.GetUserID(c), info.Uint()); err != nil {
		global.GVA_LOG.Error("删除 AI 工作流会话失败", zap.Error(err))
		response.FailWithMessage("删除会话失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (a *AIWorkflowSessionApi) DumpMarkdown(c *gin.Context) {
	var info systemReq.SysAIWorkflowMarkdownDump
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	result, err := aiWorkflowSessionService.DumpMarkdown(c.Request.Context(), utils.GetUserID(c), info)
	if err != nil {
		global.GVA_LOG.Error("AI 工作流 Markdown 落盘失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"result": result}, "落盘成功", c)
}
