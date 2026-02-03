package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
)

// 注册工具
func init() {
	// 注册工具将在enter.go中统一处理
	RegisterTool(&ApiLister{})
}

// ApiInfo API信息结构
type ApiInfo struct {
	ID          uint   `json:"id,omitempty"`          // 数据库ID（仅数据库API有）
	Path        string `json:"path"`                  // API路径
	Description string `json:"description,omitempty"` // API描述
	ApiGroup    string `json:"apiGroup,omitempty"`    // API组
	Method      string `json:"method"`                // HTTP方法
	Source      string `json:"source"`                // 来源：database 或 gin
}

// ApiListResponse API列表响应结构
type ApiListResponse struct {
	Success      bool      `json:"success"`
	Message      string    `json:"message"`
	DatabaseApis []ApiInfo `json:"databaseApis"` // 数据库中的API
	GinApis      []ApiInfo `json:"ginApis"`      // gin框架中的API
	TotalCount   int       `json:"totalCount"`   // 总数量
}

// ApiLister API列表工具
type ApiLister struct{}

// New 创建API列表工具
func (a *ApiLister) New() mcp.Tool {
	return mcp.NewTool("list_all_apis",
		mcp.WithDescription(`获取系统所有API接口。返回databaseApis(数据库API)和ginApis(gin路由)两组数据。

**详细说明**: /gva-helper/lister`),
		mcp.WithString("_placeholder",
			mcp.Description("占位符，防止json schema校验失败"),
		),	
	)
}

// Handle 处理API列表请求
func (a *ApiLister) Handle(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// 获取数据库中的API
	databaseApis, err := a.getDatabaseApis()
	if err != nil {
		global.GVA_LOG.Error("获取数据库API失败", zap.Error(err))
		errorResponse := ApiListResponse{
			Success: false,
			Message: "获取数据库API失败: " + err.Error(),
		}
		resultJSON, _ := json.Marshal(errorResponse)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: string(resultJSON),
				},
			},
		}, nil
	}

	// 获取gin路由中的API
	ginApis, err := a.getGinApis()
	if err != nil {
		global.GVA_LOG.Error("获取gin路由API失败", zap.Error(err))
		errorResponse := ApiListResponse{
			Success: false,
			Message: "获取gin路由API失败: " + err.Error(),
		}
		resultJSON, _ := json.Marshal(errorResponse)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: string(resultJSON),
				},
			},
		}, nil
	}

	// 构建响应
	response := ApiListResponse{
		Success:      true,
		Message:      "获取API列表成功",
		DatabaseApis: databaseApis,
		GinApis:      ginApis,
		TotalCount:   len(databaseApis) + len(ginApis),
	}

	global.GVA_LOG.Info("API列表获取成功",
		zap.Int("数据库API数量", len(databaseApis)),
		zap.Int("gin路由API数量", len(ginApis)),
		zap.Int("总数量", response.TotalCount))

	resultJSON, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(resultJSON),
			},
		},
	}, nil
}

// getDatabaseApis 获取数据库中的所有API
func (a *ApiLister) getDatabaseApis() ([]ApiInfo, error) {
	var apis []system.SysApi
	err := global.GVA_DB.Model(&system.SysApi{}).Order("api_group ASC, path ASC").Find(&apis).Error
	if err != nil {
		return nil, err
	}

	// 转换为ApiInfo格式
	var result []ApiInfo
	for _, api := range apis {
		result = append(result, ApiInfo{
			ID:          api.ID,
			Path:        api.Path,
			Description: api.Description,
			ApiGroup:    api.ApiGroup,
			Method:      api.Method,
			Source:      "database",
		})
	}

	return result, nil
}

// getGinApis 获取gin路由中的所有API（包含被忽略的API）
func (a *ApiLister) getGinApis() ([]ApiInfo, error) {
	// 从gin路由信息中获取所有API
	var result []ApiInfo
	for _, route := range global.GVA_ROUTERS {
		result = append(result, ApiInfo{
			Path:   route.Path,
			Method: route.Method,
			Source: "gin",
		})
	}

	return result, nil
}
