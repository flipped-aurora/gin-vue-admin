package mcpTool

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&MenuLister{})
}

type MenuListResponse struct {
	Success     bool                 `json:"success"`
	Message     string               `json:"message"`
	Menus       []system.SysBaseMenu `json:"menus"`
	TotalCount  int                  `json:"totalCount"` // 全部菜单节点数(含各级子菜单)
	RootCount   int                  `json:"rootCount"`  // 顶层菜单数量
	Description string               `json:"description"`
}

type MenuLister struct{}

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
		mcp.WithString("_placeholder",
			mcp.Description("占位符，防止json schema校验失败"),
		),
	)
}

func (m *MenuLister) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := postUpstream[[]system.SysBaseMenu](ctx, "/menu/getMenuList", map[string]any{})
	if err != nil {
		return nil, err
	}

	// getMenuList 返回的是树,resp.Data 只是顶层菜单;摊平后取节点总数,避免 totalCount 只数顶层而误导
	menuIndex := make(map[uint]system.SysBaseMenu, len(resp.Data))
	flattenBaseMenus(resp.Data, menuIndex)

	return textResultWithJSON("", MenuListResponse{
		Success:     true,
		Message:     "获取菜单列表成功",
		Menus:       resp.Data,
		TotalCount:  len(menuIndex),
		RootCount:   len(resp.Data),
		Description: "系统中所有菜单信息的标准列表（树形，menus 为顶层，子菜单在各自 children 内），包含路由配置和组件信息",
	})
}
