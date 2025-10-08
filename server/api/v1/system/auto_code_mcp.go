package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/client"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/mcp"
)

// Create
// @Tags      mcp
// @Summary   自动McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoMcpTool  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/mcp [post]
func (a *AutoCodeTemplateApi) MCP(c *gin.Context) {
	var info request.AutoMcpTool
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	toolFilePath, err := autoCodeTemplateService.CreateMcp(c.Request.Context(), info)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		global.GVA_LOG.Error(err.Error())
		return
	}
	response.OkWithMessage("创建成功,MCP Tool路径:"+toolFilePath, c)
}

// Create
// @Tags      mcp
// @Summary   自动McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoMcpTool  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/mcpList [post]
func (a *AutoCodeTemplateApi) MCPList(c *gin.Context) {

	baseUrl := fmt.Sprintf("http://127.0.0.1:%d%s", global.GVA_CONFIG.System.Addr, global.GVA_CONFIG.MCP.SSEPath)

	testClient, err := client.NewClient(baseUrl, "testClient", "v1.0.0", global.GVA_CONFIG.MCP.Name)
	defer testClient.Close()
	toolsRequest := mcp.ListToolsRequest{}

	list, err := testClient.ListTools(c.Request.Context(), toolsRequest)

	if err != nil {
		response.FailWithMessage("创建失败", c)
		global.GVA_LOG.Error(err.Error())
		return
	}

	mcpServerConfig := map[string]interface{}{
		"mcpServers": map[string]interface{}{
			global.GVA_CONFIG.MCP.Name: map[string]string{
				"url": baseUrl,
			},
		},
	}
	response.OkWithData(gin.H{
		"mcpServerConfig": mcpServerConfig,
		"list":            list,
	}, c)
}

// Create
// @Tags      mcp
// @Summary   测试McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object  true  "调用MCP Tool的参数"
// @Success   200   {object}  response.Response  "{"success":true,"data":{},"msg":"测试成功"}"
// @Router    /autoCode/mcpTest [post]
func (a *AutoCodeTemplateApi) MCPTest(c *gin.Context) {
	// 定义接口请求结构
	var testRequest struct {
		Name      string                 `json:"name" binding:"required"`      // 工具名称
		Arguments map[string]interface{} `json:"arguments" binding:"required"` // 工具参数
	}

	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&testRequest); err != nil {
		response.FailWithMessage("参数解析失败:"+err.Error(), c)
		return
	}

	// 创建MCP客户端
	baseUrl := fmt.Sprintf("http://127.0.0.1:%d%s", global.GVA_CONFIG.System.Addr, global.GVA_CONFIG.MCP.SSEPath)
	testClient, err := client.NewClient(baseUrl, "testClient", "v1.0.0", global.GVA_CONFIG.MCP.Name)
	if err != nil {
		response.FailWithMessage("创建MCP客户端失败:"+err.Error(), c)
		return
	}
	defer testClient.Close()

	ctx := c.Request.Context()

	// 初始化MCP连接
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "testClient",
		Version: "v1.0.0",
	}

	_, err = testClient.Initialize(ctx, initRequest)
	if err != nil {
		response.FailWithMessage("初始化MCP连接失败:"+err.Error(), c)
		return
	}

	// 构建工具调用请求
	request := mcp.CallToolRequest{}
	request.Params.Name = testRequest.Name
	request.Params.Arguments = testRequest.Arguments

	// 调用工具
	result, err := testClient.CallTool(ctx, request)
	if err != nil {
		response.FailWithMessage("工具调用失败:"+err.Error(), c)
		return
	}

	// 处理响应结果
	if len(result.Content) == 0 {
		response.FailWithMessage("工具未返回任何内容", c)
		return
	}

	// 返回结果
	response.OkWithData(result.Content, c)
}
