package initialize

import (
	"testing"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/stretchr/testify/require"
)

func TestEnsureLogViewerDataIsIdempotent(t *testing.T) {
	db := testutil.NewMemoryDBWithoutGlobal(t,
		&system.SysApi{},
		&system.SysBaseMenu{},
		&system.SysAuthorityMenu{},
		&adapter.CasbinRule{},
	)
	monitor := system.SysBaseMenu{Path: "monitor", Name: "monitor", Component: "view/routerHolder.vue"}
	require.NoError(t, db.Create(&monitor).Error)

	require.NoError(t, ensureLogViewerData(db))
	require.NoError(t, ensureLogViewerData(db))

	var apiCount int64
	require.NoError(t, db.Model(&system.SysApi{}).Where("path LIKE ?", "/logViewer/%").Count(&apiCount).Error)
	require.Equal(t, int64(3), apiCount)

	var menu system.SysBaseMenu
	require.NoError(t, db.Where("name = ?", "logViewer").First(&menu).Error)
	require.Equal(t, monitor.ID, menu.ParentId)

	var menuRoleCount int64
	require.NoError(t, db.Model(&system.SysAuthorityMenu{}).
		Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?", menu.ID, "888").
		Count(&menuRoleCount).Error)
	require.Equal(t, int64(1), menuRoleCount)

	var policyCount int64
	require.NoError(t, db.Model(&adapter.CasbinRule{}).
		Where("ptype = ? AND v0 = ? AND v1 LIKE ? AND v2 = ?", "p", "888", "/logViewer/%", "GET").
		Count(&policyCount).Error)
	require.Equal(t, int64(3), policyCount)
}

func TestEnsureLogViewerDataRequiresMonitorMenu(t *testing.T) {
	db := testutil.NewMemoryDBWithoutGlobal(t,
		&system.SysApi{},
		&system.SysBaseMenu{},
		&system.SysAuthorityMenu{},
		&adapter.CasbinRule{},
	)

	require.Error(t, ensureLogViewerData(db))
}
