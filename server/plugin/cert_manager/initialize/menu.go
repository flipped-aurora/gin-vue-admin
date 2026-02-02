package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	parent := model.SysBaseMenu{
		ParentId:  0,
		Path:      "observability",
		Name:      "observability",
		Hidden:    false,
		Component: "view/routerHolder.vue",
		Sort:      10,
		Meta:      model.Meta{Title: "可观测可视化", Icon: "monitor"},
	}
	child := model.SysBaseMenu{
		Path:      "certManager",
		Name:      "certManager",
		Hidden:    false,
		Component: "plugin/cert_manager/view/certCertificate.vue",
		Sort:      2,
		Meta:      model.Meta{Title: "域名证书监控", Icon: "certificate"},
	}
	categoryChild := model.SysBaseMenu{
		Path:      "certCategory",
		Name:      "certCategory",
		Hidden:    false,
		Component: "plugin/cert_manager/view/certCategory.vue",
		Sort:      1,
		Meta:      model.Meta{Title: "项目管理", Icon: "files"},
	}
	domainTreeChild := model.SysBaseMenu{
		Path:      "domainTree",
		Name:      "domainTree",
		Hidden:    false,
		Component: "plugin/cert_manager/view/domainTree.vue",
		Sort:      3,
		Meta:      model.Meta{Title: "子域名探测", Icon: "tree"},
	}
	utils.RegisterMenus(parent, categoryChild, child, domainTreeChild)

	// 强制更新已存在的菜单标题和排序，确保重构即时生效
	global.GVA_DB.WithContext(ctx).Model(&model.SysBaseMenu{}).Where("name = ?", "certCategory").Updates(map[string]interface{}{
		"title": "项目管理",
		"sort":  1,
	})
	global.GVA_DB.WithContext(ctx).Model(&model.SysBaseMenu{}).Where("name = ?", "certManager").Updates(map[string]interface{}{
		"sort": 2,
	})
	global.GVA_DB.WithContext(ctx).Model(&model.SysBaseMenu{}).Where("name = ?", "domainTree").Updates(map[string]interface{}{
		"title": "子域名探测",
		"sort":  3,
	})

	global.GVA_DB.WithContext(ctx).Model(&model.SysBaseMenu{}).Where("name = ?", "observability").Updates(map[string]any{
		"title": "可观测可视化",
	})

	// 删除已废弃的证书分布统计菜单
	global.GVA_DB.WithContext(ctx).Where("name = ?", "certStatistics").Delete(&model.SysBaseMenu{})
}
