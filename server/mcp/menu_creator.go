package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&MenuCreator{})
}

type MenuCreateRequest struct {
	ParentId    uint                   `json:"parentId"`
	Path        string                 `json:"path"`
	Name        string                 `json:"name"`
	Hidden      bool                   `json:"hidden"`
	Component   string                 `json:"component"`
	Sort        int                    `json:"sort"`
	Title       string                 `json:"title"`
	Icon        string                 `json:"icon"`
	KeepAlive   bool                   `json:"keepAlive"`
	DefaultMenu bool                   `json:"defaultMenu"`
	CloseTab    bool                   `json:"closeTab"`
	ActiveName  string                 `json:"activeName"`
	Parameters  []MenuParameterRequest `json:"parameters"`
	MenuBtn     []MenuButtonRequest    `json:"menuBtn"`
}

type MenuParameterRequest struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MenuButtonRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type MenuCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	MenuID  uint   `json:"menuId"`
	Name    string `json:"name"`
	Path    string `json:"path"`
}

type MenuCreator struct{}

func (m *MenuCreator) New() mcp.Tool {
	return mcp.NewTool("create_menu",
		mcp.WithDescription(`创建前端菜单记录，用于AI编辑器自动添加前端页面时自动创建对应的菜单项。

**重要限制：**
- 当使用gva_auto_generate工具且needCreatedModules=true时，模块创建会自动生成菜单项，不应调用此工具
- 仅在以下情况使用：1) 单独创建菜单（不涉及模块创建）；2) AI编辑器自动添加前端页面时`),
		mcp.WithNumber("parentId",
			mcp.Description("父菜单ID，0表示根菜单"),
			mcp.DefaultNumber(0),
		),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("路由path，如：userList"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("路由name，用于Vue Router，如：userList"),
		),
		mcp.WithBoolean("hidden",
			mcp.Description("是否在菜单列表中隐藏"),
		),
		mcp.WithString("component",
			mcp.Required(),
			mcp.Description("对应的前端Vue组件路径，如：view/user/list.vue"),
		),
		mcp.WithNumber("sort",
			mcp.Description("菜单排序号，数字越小越靠前"),
			mcp.DefaultNumber(1),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("菜单显示标题"),
		),
		mcp.WithString("icon",
			mcp.Description("菜单图标名称"),
			mcp.DefaultString("menu"),
		),
		mcp.WithBoolean("keepAlive",
			mcp.Description("是否缓存页面"),
		),
		mcp.WithBoolean("defaultMenu",
			mcp.Description("是否是基础路由"),
		),
		mcp.WithBoolean("closeTab",
			mcp.Description("是否自动关闭tab"),
		),
		mcp.WithString("activeName",
			mcp.Description("高亮菜单名称"),
		),
		mcp.WithString("parameters",
			mcp.Description("路由参数JSON字符串，格式：[{\"type\":\"params\",\"key\":\"id\",\"value\":\"1\"}]"),
		),
		mcp.WithString("menuBtn",
			mcp.Description("菜单按钮JSON字符串，格式：[{\"name\":\"add\",\"desc\":\"新增\"}]"),
		),
	)
}

func (m *MenuCreator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	path, ok := args["path"].(string)
	if !ok || path == "" {
		return nil, errors.New("path 参数是必需的")
	}
	name, ok := args["name"].(string)
	if !ok || name == "" {
		return nil, errors.New("name 参数是必需的")
	}
	component, ok := args["component"].(string)
	if !ok || component == "" {
		return nil, errors.New("component 参数是必需的")
	}
	title, ok := args["title"].(string)
	if !ok || title == "" {
		return nil, errors.New("title 参数是必需的")
	}

	parentID := uint(0)
	if value, ok := args["parentId"].(float64); ok {
		parentID = uint(value)
	}
	hidden, _ := args["hidden"].(bool)
	sort := 1
	if value, ok := args["sort"].(float64); ok {
		sort = int(value)
	}
	icon := "menu"
	if value, ok := args["icon"].(string); ok && value != "" {
		icon = value
	}
	keepAlive, _ := args["keepAlive"].(bool)
	defaultMenu, _ := args["defaultMenu"].(bool)
	closeTab, _ := args["closeTab"].(bool)
	activeName, _ := args["activeName"].(string)

	parameters := make([]system.SysBaseMenuParameter, 0)
	if parametersStr, ok := args["parameters"].(string); ok && parametersStr != "" {
		var paramReqs []MenuParameterRequest
		if err := json.Unmarshal([]byte(parametersStr), &paramReqs); err != nil {
			return nil, fmt.Errorf("parameters 参数格式错误: %v", err)
		}
		for _, param := range paramReqs {
			parameters = append(parameters, system.SysBaseMenuParameter{
				Type:  param.Type,
				Key:   param.Key,
				Value: param.Value,
			})
		}
	}

	menuBtns := make([]system.SysBaseMenuBtn, 0)
	if menuBtnStr, ok := args["menuBtn"].(string); ok && menuBtnStr != "" {
		var buttonReqs []MenuButtonRequest
		if err := json.Unmarshal([]byte(menuBtnStr), &buttonReqs); err != nil {
			return nil, fmt.Errorf("menuBtn 参数格式错误: %v", err)
		}
		for _, button := range buttonReqs {
			menuBtns = append(menuBtns, system.SysBaseMenuBtn{
				Name: button.Name,
				Desc: button.Desc,
			})
		}
	}

	menu := system.SysBaseMenu{
		ParentId:  parentID,
		Path:      path,
		Name:      name,
		Hidden:    hidden,
		Component: component,
		Sort:      sort,
		Meta: system.Meta{
			Title:       title,
			Icon:        icon,
			KeepAlive:   keepAlive,
			DefaultMenu: defaultMenu,
			CloseTab:    closeTab,
			ActiveName:  activeName,
		},
		Parameters: parameters,
		MenuBtn:    menuBtns,
	}

	if _, err := postUpstream[map[string]any](ctx, "/menu/addBaseMenu", menu); err != nil {
		return nil, fmt.Errorf("创建菜单失败: %v", err)
	}

	menuID := uint(0)
	if menuListResp, err := postUpstream[[]system.SysBaseMenu](ctx, "/menu/getMenuList", map[string]any{}); err == nil {
		menuID = findMenuID(menuListResp.Data, name, path)
	}

	return textResultWithJSON("菜单创建结果：", &MenuCreateResponse{
		Success: true,
		Message: fmt.Sprintf("成功创建菜单 %s", title),
		MenuID:  menuID,
		Name:    name,
		Path:    path,
	})
}

func findMenuID(menus []system.SysBaseMenu, name, path string) uint {
	for _, menu := range menus {
		if menu.Name == name && menu.Path == path {
			return menu.ID
		}
		if id := findMenuID(menu.Children, name, path); id != 0 {
			return id
		}
	}
	return 0
}
