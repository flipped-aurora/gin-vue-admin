package system

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mcpTool "github.com/flipped-aurora/gin-vue-admin/server/mcp"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/client"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/mcp"
)

// MCP
// @Tags      AutoCodeMcp
// @Summary   创建MCP Tool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoMcpTool            true  "MCP Tool配置信息"
// @Success   200   {object}  response.Response{msg=string}  "创建MCP Tool"
// @Router    /autoCode/mcp [post]
func (a *AutoCodeTemplateApi) MCP(c *gin.Context) {
	var info request.AutoMcpTool
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	toolFilePath, err := autoCodeTemplateService.CreateMcp(c.Request.Context(), info)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		logger.WithCtx(c.Request.Context()).Mod("biz").Error(err.Error())
		return
	}
	response.OkWithMessage("创建成功,MCP Tool路径:"+toolFilePath, c)
}

// MCPStatus
// @Tags      AutoCodeMcp
// @Summary   获取MCP独立服务状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "获取MCP独立服务状态及服务配置"
// @Router    /autoCode/mcpStatus [post]
func (a *AutoCodeTemplateApi) MCPStatus(c *gin.Context) {
	response.OkWithData(gin.H{
		"status":          mcpTool.GetManagedStandaloneStatus(c.Request.Context()),
		"mcpServerConfig": buildMCPServerConfig(),
	}, c)
}

// MCPStart
// @Tags      AutoCodeMcp
// @Summary   启动MCP独立服务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "启动MCP独立服务"
// @Router    /autoCode/mcpStart [post]
func (a *AutoCodeTemplateApi) MCPStart(c *gin.Context) {
	status, err := mcpTool.StartManagedStandalone(c.Request.Context())
	if err != nil {
		response.FailWithDetailed(gin.H{
			"status":          status,
			"mcpServerConfig": buildMCPServerConfig(),
		}, err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"status":          status,
		"mcpServerConfig": buildMCPServerConfig(),
	}, "MCP独立服务已启动", c)
}

// MCPStop
// @Tags      AutoCodeMcp
// @Summary   停止MCP独立服务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "停止MCP独立服务"
// @Router    /autoCode/mcpStop [post]
func (a *AutoCodeTemplateApi) MCPStop(c *gin.Context) {
	status, err := mcpTool.StopManagedStandalone(c.Request.Context())
	if err != nil {
		response.FailWithDetailed(gin.H{
			"status":          status,
			"mcpServerConfig": buildMCPServerConfig(),
		}, err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"status":          status,
		"mcpServerConfig": buildMCPServerConfig(),
	}, "MCP独立服务已停用", c)
}

// MCPList
// @Tags      AutoCodeMcp
// @Summary   获取MCP工具列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "获取MCP工具列表"
// @Router    /autoCode/mcpList [post]
func (a *AutoCodeTemplateApi) MCPList(c *gin.Context) {
	baseURL := mcpTool.ResolveMCPServiceURL()
	testClient, err := client.NewClient(baseURL, "testClient", "v1.0.0", mcpServerName(), incomingMCPHeaders(c))
	if err != nil {
		response.FailWithDetailed(gin.H{
			"status":          mcpTool.GetManagedStandaloneStatus(c.Request.Context()),
			"mcpServerConfig": buildMCPServerConfig(),
		}, "连接MCP服务失败:"+err.Error(), c)
		return
	}
	defer testClient.Close()

	list, err := testClient.ListTools(c.Request.Context(), mcp.ListToolsRequest{})
	if err != nil {
		response.FailWithDetailed(gin.H{
			"status":          mcpTool.GetManagedStandaloneStatus(c.Request.Context()),
			"mcpServerConfig": buildMCPServerConfig(),
		}, "获取工具列表失败:"+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"status":          mcpTool.GetManagedStandaloneStatus(c.Request.Context()),
		"mcpServerConfig": buildMCPServerConfig(),
		"list":            list,
	}, c)
}

// MCPRoutes
// @Tags      AutoCodeMcp
// @Summary   获取系统路由列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "获取系统路由列表"
// @Router    /autoCode/mcpRoutes [post]
func (a *AutoCodeTemplateApi) MCPRoutes(c *gin.Context) {
	response.OkWithData(gin.H{
		"routes": global.GVA_ROUTERS,
	}, c)
}

// MCPTest
// @Tags      AutoCodeMcp
// @Summary   测试调用MCP工具
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object                                     true  "工具名称name与调用参数arguments"
// @Success   200   {object}  response.Response{data=[]interface{},msg=string}  "测试调用MCP工具返回内容"
// @Router    /autoCode/mcpTest [post]
func (a *AutoCodeTemplateApi) MCPTest(c *gin.Context) {
	var testRequest struct {
		Name      string                 `json:"name" binding:"required"`
		Arguments map[string]interface{} `json:"arguments" binding:"required"`
	}
	if err := c.ShouldBindJSON(&testRequest); err != nil {
		response.FailWithMessage("参数解析失败:"+err.Error(), c)
		return
	}

	baseURL := mcpTool.ResolveMCPServiceURL()
	testClient, err := client.NewClient(baseURL, "testClient", "v1.0.0", mcpServerName(), incomingMCPHeaders(c))
	if err != nil {
		response.FailWithMessage("连接MCP服务失败:"+err.Error(), c)
		return
	}
	defer testClient.Close()

	callRequest := mcp.CallToolRequest{}
	callRequest.Params.Name = testRequest.Name
	callRequest.Params.Arguments = testRequest.Arguments

	result, err := testClient.CallTool(c.Request.Context(), callRequest)
	if err != nil {
		response.FailWithMessage("工具调用失败:"+err.Error(), c)
		return
	}
	if len(result.Content) == 0 {
		response.FailWithMessage("工具未返回任何内容", c)
		return
	}

	response.OkWithData(result.Content, c)
}

func incomingMCPHeaders(c *gin.Context) map[string]string {
	headerName := mcpTool.ConfiguredAuthHeader()
	headerValue := c.GetHeader(headerName)
	if headerValue == "" {
		return nil
	}

	return map[string]string{
		headerName: headerValue,
	}
}

func buildMCPServerConfig() map[string]interface{} {
	baseURL := mcpTool.ResolveMCPServiceURL()
	authHeader := mcpTool.ConfiguredAuthHeader()
	serverName := mcpServerName()

	return map[string]interface{}{
		"mcpServers": map[string]interface{}{
			serverName: map[string]interface{}{
				"url": baseURL,
				"headers": map[string]string{
					authHeader: "${YOUR_GVA_TOKEN}",
				},
			},
		},
	}
}

func mcpServerName() string {
	if name := strings.TrimSpace(global.GVA_CONFIG.MCP.Name); name != "" {
		return name
	}
	return "gva"
}
