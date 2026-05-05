package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	_ = ctx
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "programmingAssistant",
			Name:      "AutoRoot",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      91,
			Meta:      model.Meta{Title: "编程辅助", Icon: "cpu"},
		},
		{
			Path:      "autoCode",
			Name:      "AutoCode",
			Hidden:    false,
			Component: "plugin/auto/view/autoCode/index.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "自动代码", Icon: "magic-stick"},
		},
		{
			Path:      "autoPkg",
			Name:      "AutoPkg",
			Hidden:    false,
			Component: "plugin/auto/view/autoPkg/autoPkg.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "自动化包", Icon: "files"},
		},
		{
			Path:      "formCreate",
			Name:      "formCreate",
			Hidden:    false,
			Component: "view/systemTools/formCreate/index.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true},
		},
		{
			Path:      "aiWorkflow",
			Name:      "AIWorkflow",
			Hidden:    false,
			Component: "plugin/auto/view/aiWorkflow/index.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "AI 工作流", Icon: "memo"},
		},
		{
			Path:      "autoCodeEdit/:id",
			Name:      "autoCodeEdit",
			Hidden:    true,
			Component: "view/systemTools/autoCode/index.vue",
			Sort:      0,
			Meta:      model.Meta{Title: "自动化代码-${id}", Icon: "magic-stick"},
		},
		{
			Path:      "exportTemplate",
			Name:      "exportTemplate",
			Hidden:    false,
			Component: "view/systemTools/exportTemplate/exportTemplate.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "导出模板", Icon: "reading"},
		},
		{
			Path:      "mcp",
			Name:      "MCP",
			Hidden:    false,
			Component: "plugin/auto/view/autoCode/mcp.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "MCP 工具", Icon: "monitor"},
		},
		{
			Path:      "mcpTest",
			Name:      "MCPTest",
			Hidden:    false,
			Component: "plugin/auto/view/autoCode/mcpTest.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "MCP 测试", Icon: "connection"},
		},
		{
			Path:      "installPlugin",
			Name:      "AutoInstallPlugin",
			Hidden:    false,
			Component: "plugin/auto/view/installPlugin/index.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "插件安装", Icon: "upload-filled"},
		},
		{
			Path:      "pubPlug",
			Name:      "PubPlug",
			Hidden:    false,
			Component: "plugin/auto/view/pubPlug/pubPlug.vue",
			Sort:      7,
			Meta:      model.Meta{Title: "插件打包", Icon: "box"},
		},
		{
			Path:      "skills",
			Name:      "Skills",
			Hidden:    false,
			Component: "plugin/auto/view/skills/index.vue",
			Sort:      8,
			Meta:      model.Meta{Title: "技能管理", Icon: "edit-pen"},
		},
		{
			Path:      "picture",
			Name:      "picture",
			Hidden:    false,
			Component: "view/systemTools/autoCode/picture.vue",
			Sort:      10,
			Meta:      model.Meta{Title: "AI页面绘制", Icon: "picture-filled"},
		},
		{
			Path:      "autoCodeAdmin",
			Name:      "AutoCodeAdmin",
			Hidden:    true,
			Component: "plugin/auto/view/autoCodeAdmin/index.vue",
			Sort:      9,
			Meta:      model.Meta{Title: "自动代码管理", Icon: "setting"},
		},
	}
	utils.RegisterMenus(entities...)
}
