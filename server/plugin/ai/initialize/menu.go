package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	_ = ctx
	// 第一个元素为父菜单（AI），其余元素由 RegisterMenus 自动挂为其子菜单
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "ai",
			Name:      "ai",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "AI 工坊", Icon: "magic-stick"},
		},
		{
			Path:      "mcpTool",
			Name:      "mcpTool",
			Hidden:    false,
			Component: "plugin/ai/view/mcp/mcp.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "Mcp Tools模板", Icon: "monitor"},
		},
		{
			Path:      "mcpTest",
			Name:      "mcpTest",
			Hidden:    false,
			Component: "plugin/ai/view/mcp/mcpTest.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "Mcp Tools管理", Icon: "connection"},
		},
		{
			Path:      "skills",
			Name:      "Skills",
			Hidden:    false,
			Component: "plugin/ai/view/skills/index.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "Skills管理", Icon: "edit-pen"},
		},
		{
			Path:      "cli",
			Name:      "Cli",
			Hidden:    false,
			Component: "plugin/ai/view/cli/index.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "AI CLI管理", Icon: "cpu", KeepAlive: true},
		},
		{
			Path:      "picture",
			Name:      "picture",
			Hidden:    false,
			Component: "plugin/ai/view/picture/picture.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "AI页面绘制", Icon: "picture-filled"},
		},
		{
			Path:      "aiWorkflow",
			Name:      "aiWorkflow",
			Hidden:    false,
			Component: "plugin/ai/view/aiWorkflow/index.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "AI 工作流", Icon: "memo"},
		},
	}
	utils.RegisterMenus(entities...)
}
