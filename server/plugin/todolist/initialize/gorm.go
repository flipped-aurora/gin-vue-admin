package initialize

import (
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/model"
	"go.uber.org/zap"
)

// 注册所需数据库表
func RegisterTables() {
	err := gvaGlobal.GVA_DB.AutoMigrate(
		&model.TodoList{},
	)
	if err != nil {
		gvaGlobal.GVA_LOG.Error("todolist plugin register table failed", zap.Error(err))
	}
}
