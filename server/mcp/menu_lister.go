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
	RegisterTool(&MenuLister{})
}

// MenuListResponse 菜单列表响应结构
type MenuListResponse struct {
	Success     bool                  `json:"success"`
	Message     string                `json:"message"`
	Menus       []system.SysBaseMenu  `json:"menus"`
	TotalCount  int                   `json:"totalCount"`
	Description string                `json:"description"`
}

// MenuLister 菜单列表工具
type MenuLister struct{}

// New 创建菜单列表工具
func (m *MenuLister) New() mcp.Tool {
	return mcp.NewTool("list_all_menus",
		mcp.WithDescription(`获取系统中所有菜单信息，包括菜单树结构、路由信息、组件路径等，用于前端编写vue-router时正确跳转

**功能说明：**
- 返回完整的菜单树形结构
- 包含路由配置信息（path、name、component）
- 包含菜单元数据（title、icon、keepAlive等）
- 包含菜单参数和按钮配置
- 支持父子菜单关系展示

**使用场景：**
- 前端路由配置：获取所有菜单信息用于配置vue-router
- 菜单权限管理：了解系统中所有可用的菜单项
- 导航组件开发：构建动态导航菜单
- 系统架构分析：了解系统的菜单结构和页面组织`),
	)
}

// Handle 处理菜单列表请求
func (m *MenuLister) Handle(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取所有基础菜单
	allMenus, err := m.getAllMenus()
	if err != nil {
		global.GVA_LOG.Error("获取菜单列表失败", zap.Error(err))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("获取菜单列表失败: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	// 构建返回结果
	response := MenuListResponse{
		Success:     true,
		Message:     "获取菜单列表成功",
		Menus:       allMenus,
		TotalCount:  len(allMenus),
		Description: "系统中所有菜单信息的标准列表，包含路由配置和组件信息",
	}

	// 序列化响应
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		global.GVA_LOG.Error("序列化菜单响应失败", zap.Error(err))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("序列化响应失败: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(responseJSON),
			},
		},
	}, nil
}

// getAllMenus 获取所有基础菜单
func (m *MenuLister) getAllMenus() ([]system.SysBaseMenu, error) {
	var menus []system.SysBaseMenu
	err := global.GVA_DB.Order("sort").Preload("Parameters").Preload("MenuBtn").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

