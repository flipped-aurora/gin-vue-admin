package initialize

import (
	"fmt"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

var logViewerApis = []system.SysApi{
	{ApiGroup: "文件日志", Method: "GET", Path: "/logViewer/dates", Description: "获取存在日志的日期"},
	{ApiGroup: "文件日志", Method: "GET", Path: "/logViewer/files", Description: "获取日期下的日志文件"},
	{ApiGroup: "文件日志", Method: "GET", Path: "/logViewer/content", Description: "分块读取日志文件内容"},
}

// EnsureLogViewerData adds feature metadata to databases created before the log viewer existed.
func EnsureLogViewerData() error {
	if global.GVA_DB == nil {
		return nil
	}
	return ensureLogViewerData(global.GVA_DB)
}

func ensureLogViewerData(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, api := range logViewerApis {
			if err := tx.Where("path = ? AND method = ?", api.Path, api.Method).Attrs(api).FirstOrCreate(&system.SysApi{}).Error; err != nil {
				return fmt.Errorf("ensure log viewer API %s: %w", api.Path, err)
			}
		}

		var monitor system.SysBaseMenu
		if err := tx.Where("name = ?", "monitor").First(&monitor).Error; err != nil {
			return fmt.Errorf("find monitor menu: %w", err)
		}
		menu := system.SysBaseMenu{
			MenuLevel: 1,
			ParentId:  monitor.ID,
			Path:      "logViewer",
			Name:      "logViewer",
			Component: "view/systemTools/logViewer/index.vue",
			Sort:      8,
			Meta: system.Meta{
				Title: "文件日志",
				Icon:  "document",
			},
		}
		if err := tx.Where("name = ?", menu.Name).Attrs(menu).FirstOrCreate(&menu).Error; err != nil {
			return fmt.Errorf("ensure log viewer menu: %w", err)
		}

		menuRole := system.SysAuthorityMenu{MenuId: fmt.Sprint(menu.ID), AuthorityId: "888"}
		if err := tx.Where(menuRole).FirstOrCreate(&menuRole).Error; err != nil {
			return fmt.Errorf("ensure log viewer admin menu: %w", err)
		}

		for _, api := range logViewerApis {
			policy := adapter.CasbinRule{Ptype: "p", V0: "888", V1: api.Path, V2: api.Method}
			if err := tx.Where(policy).FirstOrCreate(&policy).Error; err != nil {
				return fmt.Errorf("ensure log viewer policy %s: %w", api.Path, err)
			}
		}
		return nil
	})
}
