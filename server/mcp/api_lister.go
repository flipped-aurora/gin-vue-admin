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
		mcp.WithDescription(`获取系统中所有的API接口，分为两组：

**功能说明：**
- 返回数据库中已注册的API列表
- 返回gin框架中实际注册的路由API列表
- 帮助前端判断是使用现有API还是需要创建新的API,如果api在前端未使用且需要前端调用的时候，请到api文件夹下对应模块的js中添加方法并暴露给当前业务调用

**返回数据结构：**
- databaseApis: 数据库中的API记录（包含ID、描述、分组等完整信息）
- ginApis: gin路由中的API（仅包含路径和方法），需要AI根据路径自行揣摩路径的业务含义，例如：/api/user/:id 表示根据用户ID获取用户信息`),
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
