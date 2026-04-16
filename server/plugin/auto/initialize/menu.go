package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
			Path:      "aiWorkflow",
			Name:      "AIWorkflow",
			Hidden:    false,
			Component: "plugin/auto/view/aiWorkflow/index.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "AI 工作流", Icon: "memo"},
		},
		{
			Path:      "mcp",
			Name:      "MCP",
			Hidden:    false,
			Component: "plugin/auto/view/autoCode/mcp.vue",
			Sort:      4,
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
			Path:      "autoCodeAdmin",
			Name:      "AutoCodeAdmin",
			Hidden:    true,
			Component: "plugin/auto/view/autoCodeAdmin/index.vue",
			Sort:      9,
			Meta:      model.Meta{Title: "自动代码管理", Icon: "setting"},
		},
	}
	utils.RegisterMenus(entities...)

	root := entities[0]
	if err := global.GVA_DB.Where("name = ?", root.Name).First(&root).Error; err != nil {
		return
	}
	for _, child := range entities[1:] {
		child.ParentId = root.ID
		global.GVA_DB.Model(&model.SysBaseMenu{}).Where("name = ?", child.Name).Updates(map[string]interface{}{
			"parent_id":    child.ParentId,
			"path":         child.Path,
			"component":    child.Component,
			"sort":         child.Sort,
			"hidden":       child.Hidden,
			"meta":         child.Meta,
			"keep_alive":   child.KeepAlive,
			"default_menu": child.DefaultMenu,
		})
	}
}
