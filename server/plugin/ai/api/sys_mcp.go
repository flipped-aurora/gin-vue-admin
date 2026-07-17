package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type McpApi struct{}

// CreateMcp 创建MCP
// @Tags McpApi
// @Summary 创建MCP
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.CreateSysMcpRequest true "MCP基础信息"
// @Success 200 {object} response.Response{data=object,msg=string} "创建成功"
// @Router /mcpApi/createMcp [post]
func (a *McpApi) CreateMcp(c *gin.Context) {
	var req autoReq.CreateSysMcpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.CreateMcp(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("创建MCP失败!")
		response.FailWithMessage("创建MCP失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "创建成功", c)
}

// GetMcpList 获取MCP列表
// @Tags McpApi
// @Summary 获取MCP列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.SysMcpSearch true "查询条件"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/getMcpList [post]
func (a *McpApi) GetMcpList(c *gin.Context) {
	var req autoReq.SysMcpSearch
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.GetMcpList(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取MCP列表失败!")
		response.FailWithMessage("获取MCP列表失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: res.List, Total: res.Total, Page: res.Page, PageSize: res.PageSize}, "获取成功", c)
}

// GetMcpDetail 获取MCP详情
// @Tags McpApi
// @Summary 获取MCP详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.FindSysMcpRequest true "MCP ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/getMcpDetail [post]
func (a *McpApi) GetMcpDetail(c *gin.Context) {
	var req autoReq.FindSysMcpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.GetMcpDetail(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取MCP详情失败!")
		response.FailWithMessage("获取MCP详情失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "获取成功", c)
}

// UpdateMcp 更新MCP
// @Tags McpApi
// @Summary 更新MCP
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.UpdateSysMcpRequest true "MCP基础信息"
// @Success 200 {object} response.Response{data=object,msg=string} "更新成功"
// @Router /mcpApi/updateMcp [post]
func (a *McpApi) UpdateMcp(c *gin.Context) {
	var req autoReq.UpdateSysMcpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.UpdateMcp(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("更新MCP失败!")
		response.FailWithMessage("更新MCP失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "更新成功", c)
}

// DeleteMcp 删除MCP
// @Tags McpApi
// @Summary 删除MCP
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.DeleteSysMcpRequest true "MCP ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mcpApi/deleteMcp [post]
func (a *McpApi) DeleteMcp(c *gin.Context) {
	var req autoReq.DeleteSysMcpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mcpService.DeleteMcp(c.Request.Context(), req); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除MCP失败!")
		response.FailWithMessage("删除MCP失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// AddMcpApis 给MCP绑定API
// @Tags McpApi
// @Summary 给MCP绑定API
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.AddSysMcpApisRequest true "绑定信息"
// @Success 200 {object} response.Response{data=object,msg=string} "绑定成功"
// @Router /mcpApi/addMcpApis [post]
func (a *McpApi) AddMcpApis(c *gin.Context) {
	var req autoReq.AddSysMcpApisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.AddMcpApis(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("绑定API失败!")
		response.FailWithMessage("绑定API失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "绑定成功", c)
}

// RemoveMcpApis 解绑MCP的API
// @Tags McpApi
// @Summary 解绑MCP的API
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body autoReq.RemoveSysMcpApisRequest true "解绑信息"
// @Success 200 {object} response.Response{data=object,msg=string} "解绑成功"
// @Router /mcpApi/removeMcpApis [post]
func (a *McpApi) RemoveMcpApis(c *gin.Context) {
	var req autoReq.RemoveSysMcpApisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := mcpService.RemoveMcpApis(c.Request.Context(), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("解绑API失败!")
		response.FailWithMessage("解绑API失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "解绑成功", c)
}

// ListMcpBindingsPublic 返回所有启用绑定的能力定义（command 列表），供 MCP 进程启动时读取。
// 此接口为公开接口（不走 JWTAuth），因为它是 MCP 动态注册的元数据源，本身是"要对外暴露的 tool 清单"。
// @Tags McpApi
// @Summary 列出MCP动态tool的能力定义
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/listBindingsPublic [get]
func (a *McpApi) ListMcpBindingsPublic(c *gin.Context) {
	// 用独立的 context（公开接口可能无用户 token，但 service 层需要 ctx 做日志/DB）
	commands, err := mcpService.BuildCommandsFromBindings(c.Request.Context())
	if err != nil {
		logger.Bg().Mod("biz").Err(err).Error("获取MCP绑定能力定义失败!")
		response.FailWithMessage("获取MCP绑定能力定义失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"commands": commands, "server": gin.H{
		"baseUrl":    global.GVA_CONFIG.MCP.UpstreamBaseURL,
		"authHeader": global.GVA_CONFIG.MCP.AuthHeader,
	}}, "获取成功", c)
}

// PreviewMcpManifest 预览单个 MCP 的能力定义（commands 列表）
// @Tags McpApi
// @Summary 预览MCP能力定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "mcpId"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/previewManifest [post]
func (a *McpApi) PreviewMcpManifest(c *gin.Context) {
	var req struct {
		McpID uint `json:"mcpId" form:"mcpId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	commands, err := mcpService.PreviewMcpManifest(c.Request.Context(), req.McpID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("预览MCP能力定义失败!")
		response.FailWithMessage("预览MCP能力定义失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"commands": commands}, "获取成功", c)
}

// ListMcpPromptsPublic 返回所有 MCP 的编排 prompt 定义，供 MCP 进程启动时注册成 MCP prompt。
// 公开接口（不走 JWTAuth），AI 客户端调 prompts/list 时可见这些编排剧本。
// @Tags McpApi
// @Summary 列出MCP编排prompt
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/listPromptsPublic [get]
func (a *McpApi) ListMcpPromptsPublic(c *gin.Context) {
	prompts, err := mcpService.BuildMcpPrompts(c.Request.Context())
	if err != nil {
		logger.Bg().Mod("biz").Err(err).Error("获取MCP编排prompt失败!")
		response.FailWithMessage("获取MCP编排prompt失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"prompts": prompts}, "获取成功", c)
}

// PreviewMcpPrompt 预览单个 MCP 的编排 prompt（markdown 内容）
// @Tags McpApi
// @Summary 预览MCP编排prompt
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "mcpId"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/previewPrompt [post]
func (a *McpApi) PreviewMcpPrompt(c *gin.Context) {
	var req struct {
		McpID uint `json:"mcpId" form:"mcpId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prompt, err := mcpService.PreviewMcpPrompt(c.Request.Context(), req.McpID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("预览MCP编排prompt失败!")
		response.FailWithMessage("预览MCP编排prompt失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(prompt, "获取成功", c)
}

// PreviewApiCommand 按单个 apiId 实时生成能力定义（供"自动生成"按钮）
// @Tags McpApi
// @Summary 按API生成能力定义
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body object true "apiId"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpApi/previewApiCommand [post]
func (a *McpApi) PreviewApiCommand(c *gin.Context) {
	var req struct {
		ApiID uint `json:"apiId" form:"apiId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmd, err := mcpService.PreviewApiCommand(c.Request.Context(), req.ApiID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("按API生成能力定义失败!")
		response.FailWithMessage("按API生成能力定义失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(cmd, "获取成功", c)
}
