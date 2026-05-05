package utils

import (
	"sync"

	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// GetCasbin 获取casbin实例
func GetCasbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查 casbin 配置和数据表", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败", zap.Error(err))
			return
		}
		enforcer, err := casbin.NewSyncedCachedEnforcer(m, a)
		if err != nil {
			zap.L().Error("casbin enforcer 初始化失败", zap.Error(err))
			return
		}
		enforcer.SetExpireTime(60 * 60)
		if err = enforcer.LoadPolicy(); err != nil {
			zap.L().Error("casbin 策略加载失败", zap.Error(err))
			return
		}
		syncedCachedEnforcer = enforcer
	})
	return syncedCachedEnforcer
}
