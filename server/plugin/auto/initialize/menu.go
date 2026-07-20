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
			Path:      "systemTools",
			Name:      "systemTools",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      91,
			Meta:      model.Meta{Title: "编程辅助", Icon: "cpu"},
		},
		{
			Path:      "autoCode",
			Name:      "autoCode",
			Hidden:    false,
			Component: "plugin/auto/view/autoCode/index.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "代码生成器", Icon: "magic-stick"},
		},
		{
			Path:      "autoCodeAdmin",
			Name:      "AutoCodeAdmin",
			Hidden:    false,
			Component: "plugin/auto/view/autoCodeAdmin/index.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "自动代码管理", Icon: "file-code-2-gva"},
		},
		{
			Path:      "autoPkg",
			Name:      "autoPkg",
			Hidden:    false,
			Component: "plugin/auto/view/autoPkg/autoPkg.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "模板配置", Icon: "files"},
		},
		{
			Path:      "formCreate",
			Name:      "formCreate",
			Hidden:    false,
			Component: "plugin/auto/view/formCreate/index.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true},
		},
		{
			Path:      "autoCodeEdit/:id",
			Name:      "autoCodeEdit",
			Hidden:    true,
			Component: "plugin/auto/view/autoCode/index.vue",
			Sort:      0,
			Meta:      model.Meta{Title: "自动化代码-${id}", Icon: "magic-stick"},
		},
		{
			Path:      "exportTemplate",
			Name:      "exportTemplate",
			Hidden:    false,
			Component: "plugin/auto/view/exportTemplate/exportTemplate.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "导出模板", Icon: "reading"},
		},
	}
	utils.RegisterMenus(entities...)
}
