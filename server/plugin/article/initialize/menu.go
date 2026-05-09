package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  9,
			Path:      "article",
			Name:      "article",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      20,
			Meta:      model.Meta{Title: "文章管理", Icon: "document"},
		},
		{
			Path:      "articleList",
			Name:      "articleList",
			Hidden:    false,
			Component: "plugin/article/view/list.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "文章列表", Icon: "list"},
		},
		{
			Path:      "articleCategory",
			Name:      "articleCategory",
			Hidden:    false,
			Component: "plugin/article/view/category.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "文章分类", Icon: "folder-opened"},
		},
	}
	utils.RegisterMenus(entities...)
}
