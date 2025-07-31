package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
)

// 注册工具
func init() {
	RegisterTool(&ApiCreator{})
}

// ApiCreateRequest API创建请求结构
type ApiCreateRequest struct {
	Path        string `json:"path"`        // API路径
	Description string `json:"description"` // API中文描述
	ApiGroup    string `json:"apiGroup"`    // API组
	Method      string `json:"method"`      // HTTP方法
}

// ApiCreateResponse API创建响应结构
type ApiCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ApiID   uint   `json:"apiId"`
	Path    string `json:"path"`
	Method  string `json:"method"`
}

// ApiCreator API创建工具
type ApiCreator struct{}

// New 创建API创建工具
func (a *ApiCreator) New() mcp.Tool {
	return mcp.NewTool("create_api",
		mcp.WithDescription("创建后端API记录，用于在生成后端接口时自动创建对应的API权限记录，只要创建了API层，router下的文件产生了路径变化等，都需要调用此mcp。"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("API路径，如：/user/create"),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("API中文描述，如：创建用户"),
		),
		mcp.WithString("apiGroup",
			mcp.Required(),
			mcp.Description("API组名称，用于分类管理，如：用户管理"),
		),
		mcp.WithString("method",
			mcp.Description("HTTP方法"),
			mcp.DefaultString("POST"),
		),
		mcp.WithString("apis",
			mcp.Description("批量创建API的JSON字符串，格式：[{\"path\":\"/user/create\",\"description\":\"创建用户\",\"apiGroup\":\"用户管理\",\"method\":\"POST\"}]"),
		),
	)
}

// Handle 处理API创建请求
func (a *ApiCreator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	var apis []ApiCreateRequest

	// 检查是否是批量创建
	if apisStr, ok := args["apis"].(string); ok && apisStr != "" {
		if err := json.Unmarshal([]byte(apisStr), &apis); err != nil {
			return nil, fmt.Errorf("apis 参数格式错误: %v", err)
		}
	} else {
		// 单个API创建
		path, ok := args["path"].(string)
		if !ok || path == "" {
			return nil, errors.New("path 参数是必需的")
		}

		description, ok := args["description"].(string)
		if !ok || description == "" {
			return nil, errors.New("description 参数是必需的")
		}

		apiGroup, ok := args["apiGroup"].(string)
		if !ok || apiGroup == "" {
			return nil, errors.New("apiGroup 参数是必需的")
		}

		method := "POST"
		if val, ok := args["method"].(string); ok && val != "" {
			method = val
		}

		apis = append(apis, ApiCreateRequest{
			Path:        path,
			Description: description,
			ApiGroup:    apiGroup,
			Method:      method,
		})
	}

	if len(apis) == 0 {
		return nil, errors.New("没有要创建的API")
	}

	// 创建API记录
	apiService := service.ServiceGroupApp.SystemServiceGroup.ApiService
	var responses []ApiCreateResponse
	successCount := 0

	for _, apiReq := range apis {
		api := system.SysApi{
			Path:        apiReq.Path,
			Description: apiReq.Description,
			ApiGroup:    apiReq.ApiGroup,
			Method:      apiReq.Method,
		}

		err := apiService.CreateApi(api)
		if err != nil {
			global.GVA_LOG.Warn("创建API失败",
				zap.String("path", apiReq.Path),
				zap.String("method", apiReq.Method),
				zap.Error(err))

			responses = append(responses, ApiCreateResponse{
				Success: false,
				Message: fmt.Sprintf("创建API失败: %v", err),
				Path:    apiReq.Path,
				Method:  apiReq.Method,
			})
		} else {
			// 获取创建的API ID
			var createdApi system.SysApi
			err = global.GVA_DB.Where("path = ? AND method = ?", apiReq.Path, apiReq.Method).First(&createdApi).Error
			if err != nil {
				global.GVA_LOG.Warn("获取创建的API ID失败", zap.Error(err))
			}

			responses = append(responses, ApiCreateResponse{
				Success: true,
				Message: fmt.Sprintf("成功创建API %s %s", apiReq.Method, apiReq.Path),
				ApiID:   createdApi.ID,
				Path:    apiReq.Path,
				Method:  apiReq.Method,
			})
			successCount++
		}
	}

	// 构建总体响应
	var resultMessage string
	if len(apis) == 1 {
		resultMessage = responses[0].Message
	} else {
		resultMessage = fmt.Sprintf("批量创建API完成，成功 %d 个，失败 %d 个", successCount, len(apis)-successCount)
	}

	result := map[string]interface{}{
		"success":      successCount > 0,
		"message":      resultMessage,
		"totalCount":   len(apis),
		"successCount": successCount,
		"failedCount":  len(apis) - successCount,
		"details":      responses,
	}

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	// 添加权限分配提醒
	permissionReminder := "\n\n⚠️ 重要提醒：\n" +
		"API创建完成后，请前往【系统管理】->【角色管理】中为相关角色分配新创建的API权限，" +
		"以确保用户能够正常访问新接口。\n" +
		"具体步骤：\n" +
		"1. 进入角色管理页面\n" +
		"2. 选择需要授权的角色\n" +
		"3. 在API权限中勾选新创建的API接口\n" +
		"4. 保存权限配置"

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("API创建结果：\n\n%s%s", string(resultJSON), permissionReminder),
			},
		},
	}, nil
}
