package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
)

// RegisterDataScopeCallbacks 为主库及所有多库连接注册数据权限 GORM 回调,
// 并接上审计事件异步落表(sys_data_access_logs)。
// 需在 DB(含 GVA_DBList)初始化完成后调用。
func RegisterDataScopeCallbacks() {
	datascope.RegisterCallbacks(global.GVA_DB)
	for _, db := range global.GVA_DBList {
		datascope.RegisterCallbacks(db)
	}
	auditSvc := &service.ServiceGroupApp.SystemServiceGroup.DataAccessLogService
	datascope.SetAuditHook(auditSvc.Enqueue)
	auditSvc.StartWriter()
}
