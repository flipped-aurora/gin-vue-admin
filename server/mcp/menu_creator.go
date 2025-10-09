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
	RegisterTool(&MenuCreator{})
}

// MenuCreateRequest 菜单创建请求结构
type MenuCreateRequest struct {
	ParentId    uint                   `json:"parentId,omitempty" jsonschema_description:"父菜单ID，0表示根菜单" jsonschema:"default=0"` // 父菜单ID，0表示根菜单
	Path        string                 `json:"path" jsonschema_description:"路由path"`                                            // 路由path
	Name        string                 `json:"name" jsonschema_description:"路由name"`                                            // 路由name
	Hidden      bool                   `json:"hidden,omitempty" jsonschema_description:"是否在列表隐藏"`                               // 是否在列表隐藏
	Component   string                 `json:"component" jsonschema_description:"对应前端文件路径"`                                     // 对应前端文件路径
	Sort        int                    `json:"sort,omitempty" jsonschema_description:"排序标记" jsonschema:"default=1"`             // 排序标记
	Title       string                 `json:"title" jsonschema_description:"菜单名"`                                              // 菜单名
	Icon        string                 `json:"icon,omitempty" jsonschema_description:"菜单图标" jsonschema:"default=menu"`          // 菜单图标
	KeepAlive   bool                   `json:"keepAlive,omitempty" jsonschema_description:"是否缓存"`                               // 是否缓存
	DefaultMenu bool                   `json:"defaultMenu,omitempty" jsonschema_description:"是否是基础路由"`                          // 是否是基础路由
	CloseTab    bool                   `json:"closeTab,omitempty" jsonschema_description:"自动关闭tab"`                             // 自动关闭tab
	ActiveName  string                 `json:"activeName,omitempty" jsonschema_description:"高亮菜单"`                              // 高亮菜单
	Parameters  []MenuParameterRequest `json:"parameters,omitempty" jsonschema_description:"路由参数"`                              // 路由参数
	MenuBtn     []MenuButtonRequest    `json:"menuBtn,omitempty" jsonschema_description:"菜单按钮"`                                 // 菜单按钮
}

// MenuParameterRequest 菜单参数请求结构
type MenuParameterRequest struct {
	Type  string `json:"type" jsonschema_description:"参数类型：params或query" jsonschema:"enum=params,enum=query"` // 参数类型：params或query
	Key   string `json:"key" jsonschema_description:"参数key"`                                                  // 参数key
	Value string `json:"value" jsonschema_description:"参数值"`                                                  // 参数值
}

// MenuButtonRequest 菜单按钮请求结构
type MenuButtonRequest struct {
	Name string `json:"name" jsonschema_description:"按钮名称"` // 按钮名称
	Desc string `json:"desc" jsonschema_description:"按钮描述"` // 按钮描述
}

// MenuCreateResponse 菜单创建响应结构
type MenuCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	MenuID  uint   `json:"menuId"`
	Name    string `json:"name"`
	Path    string `json:"path"`
}

// MenuCreator 菜单创建工具
type MenuCreator struct{}

// New 创建菜单创建工具
func (m *MenuCreator) New() mcp.Tool {
	return mcp.NewTool("create_menu",
		mcp.WithDescription(`创建前端菜单记录，用于AI编辑器自动添加前端页面时自动创建对应的菜单项。

**重要限制：**
- 当使用gva_auto_generate工具且needCreatedModules=true时，模块创建会自动生成菜单项，不应调用此工具
- 仅在以下情况使用：1) 单独创建菜单（不涉及模块创建）；2) AI编辑器自动添加前端页面时`),
		mcp.WithInputSchema[MenuCreateRequest](),
	)
}

// Handle 处理菜单创建请求
func (m *MenuCreator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 解析请求参数
	var args MenuCreateRequest
	if err := request.BindArguments(&args); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("参数格式错误: %v", err)), nil
	}
	// 必需参数
	if args.Path == "" {
		return nil, errors.New("path 参数是必需的")
	}

	if args.Name == "" {
		return nil, errors.New("name 参数是必需的")
	}

	if args.Component == "" {
		return nil, errors.New("component 参数是必需的")
	}

	if args.Title == "" {
		return nil, errors.New("title 参数是必需的")
	}

	if args.Sort == 0 {
		args.Sort = 1
	}
	// 解析参数和按钮
	var parameters []system.SysBaseMenuParameter
	for _, param := range args.Parameters {
		parameters = append(parameters, system.SysBaseMenuParameter{
			Type:  param.Type,
			Key:   param.Key,
			Value: param.Value,
		})
	}
	var menuBtn []system.SysBaseMenuBtn
	for _, btn := range args.MenuBtn {
		menuBtn = append(menuBtn, system.SysBaseMenuBtn{
			Name: btn.Name,
			Desc: btn.Desc,
		})
	}
	// 构建菜单对象
	menu := system.SysBaseMenu{
		ParentId:  args.ParentId,
		Path:      args.Path,
		Name:      args.Name,
		Hidden:    args.Hidden,
		Component: args.Component,
		Sort:      args.Sort,
		Meta: system.Meta{
			Title:       args.Title,
			Icon:        args.Icon,
			KeepAlive:   args.KeepAlive,
			DefaultMenu: args.DefaultMenu,
			CloseTab:    args.CloseTab,
			ActiveName:  args.ActiveName,
		},
		Parameters: parameters,
		MenuBtn:    menuBtn,
	}

	// 创建菜单
	menuService := service.ServiceGroupApp.SystemServiceGroup.MenuService
	err := menuService.AddBaseMenu(menu)
	if err != nil {
		return mcp.NewToolResultErrorf("创建菜单失败: %v", err), nil
	}

	// 获取创建的菜单ID
	var createdMenu system.SysBaseMenu
	err = global.GVA_DB.Where("name = ? AND path = ?", args.Name, args.Path).First(&createdMenu).Error
	if err != nil {
		global.GVA_LOG.Warn("获取创建的菜单ID失败", zap.Error(err))
	}

	// 构建响应
	response := &MenuCreateResponse{
		Success: true,
		Message: fmt.Sprintf("成功创建菜单 %s", args.Title),
		MenuID:  createdMenu.ID,
		Name:    args.Name,
		Path:    args.Path,
	}

	resultJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorf("序列化结果失败: %v", err), nil
	}

	// 添加权限分配提醒
	permissionReminder := "\n\n⚠️ 重要提醒：\n" +
		"菜单创建完成后，请前往【系统管理】->【角色管理】中为相关角色分配新创建的菜单权限，" +
		"以确保用户能够正常访问新菜单。\n" +
		"具体步骤：\n" +
		"1. 进入角色管理页面\n" +
		"2. 选择需要授权的角色\n" +
		"3. 在菜单权限中勾选新创建的菜单项\n" +
		"4. 保存权限配置"

	return mcp.NewToolResultText(fmt.Sprintf("菜单创建结果：\n\n%s%s", string(resultJSON), permissionReminder)), nil
}
