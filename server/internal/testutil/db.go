package testutil

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// memoryDBOpen 打开一个 sqlite :memory: 库，并兜底 GVA_LOG（datascope 等路径在
// 无身份告警时会经 GVA_LOG 输出，nil 会 panic）。失败时 t.Fatalf。
//
// 仅返回实例，不赋值任何全局变量、不注册任何 callback，由上层 helper 组合使用。
func memoryDBOpen(t testing.TB) *gorm.DB {
	t.Helper()
	InitNopLogger() // 幂等兜底 GVA_LOG

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("testutil: 打开 sqlite :memory: 失败: %v", err)
	}
	return db
}

// migrate 对传入的 model 列表执行 AutoMigrate。models 为空时跳过。
func migrate(t testing.TB, db *gorm.DB, models ...any) {
	t.Helper()
	if len(models) == 0 {
		return
	}
	if err := db.AutoMigrate(models...); err != nil {
		t.Fatalf("testutil: AutoMigrate 失败: %v", err)
	}
}

// DBOption 配置 NewMemoryDB / NewMemoryDBWithoutGlobal 的可选行为。
type DBOption func(*dbConfig)

type dbConfig struct {
	models           []any
	dataScope        bool
	skipAutoMigrate  bool
}

// WithModels 指定需要 AutoMigrate 的 model 列表。等价于把 models 作为可变参传入
// NewMemoryDB；用于需要链式表达时。注意：若 NewMemoryDB 已显式传入 models，则两者会被合并。
func WithModels(models ...any) DBOption {
	return func(c *dbConfig) { c.models = append(c.models, models...) }
}

// WithDataScopeCallbacks 对 db 注册数据权限全局回调
// （复用 utils/datascope.RegisterCallbacks，命名回调幂等，可安全重复调用）。
// 适用于测试目标涉及 datascope 自动过滤/归属盖章的场景。
func WithDataScopeCallbacks() DBOption {
	return func(c *dbConfig) { c.dataScope = true }
}

// WithoutAutoMigrate 跳过 AutoMigrate，仅打开空库由调用方自行建表。
func WithoutAutoMigrate() DBOption {
	return func(c *dbConfig) { c.skipAutoMigrate = true }
}

// applyDBOptions 合并可变参 models 与 DBOption，返回最终配置。
func applyDBOptions(models []any, opts []DBOption) dbConfig {
	c := dbConfig{models: append([]any(nil), models...)}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// NewMemoryDB 打开一个 sqlite :memory: 库，对传入的 models 执行 AutoMigrate，
// 赋值给 global.GVA_DB，并在 t.Cleanup 中还原旧值。
//
// 同时幂等地兜底 global.GVA_LOG = zap.NewNop()（仅当为 nil 时）。
// 返回新建的 *gorm.DB 供调用方直接使用。
//
// 注意：本函数会修改全局 GVA_DB，因此不应用于 t.Parallel() 的并行测试；
// 并行场景请使用 NewMemoryDBWithoutGlobal 持有独立实例。
//
// 用法：
//
//	db := testutil.NewMemoryDB(t, &SysUser{}, &SysAuthority{})
//	// 或使用 options：
//	db := testutil.NewMemoryDB(t, testutil.WithModels(&SysUser{}), testutil.WithDataScopeCallbacks())
func NewMemoryDB(t testing.TB, modelsAndOpts ...any) *gorm.DB {
	t.Helper()
	models, opts := splitModelsAndOpts(modelsAndOpts)
	cfg := applyDBOptions(models, opts)

	db := memoryDBOpen(t)
	if !cfg.skipAutoMigrate {
		migrate(t, db, cfg.models...)
	}
	if cfg.dataScope {
		datascope.RegisterCallbacks(db)
	}

	old := global.GVA_DB
	global.GVA_DB = db
	t.Cleanup(func() { global.GVA_DB = old })
	return db
}

// NewMemoryDBWithoutGlobal 与 NewMemoryDB 行为一致，但不赋值 global.GVA_DB，
// 仅返回独立实例。适用于 t.Parallel() 并行测试，避免全局变量数据竞争。
//
// 用法：
//
//	func TestX(t *testing.T) {
//	    t.Parallel()
//	    db := testutil.NewMemoryDBWithoutGlobal(t, &SysUser{})
//	    _ = db // 直接使用返回值，不依赖 global.GVA_DB
//	}
func NewMemoryDBWithoutGlobal(t testing.TB, modelsAndOpts ...any) *gorm.DB {
	t.Helper()
	models, opts := splitModelsAndOpts(modelsAndOpts)
	cfg := applyDBOptions(models, opts)

	db := memoryDBOpen(t)
	if !cfg.skipAutoMigrate {
		migrate(t, db, cfg.models...)
	}
	if cfg.dataScope {
		datascope.RegisterCallbacks(db)
	}
	return db
}

// splitModelsAndOpts 将可变参按类型分离为 model（any）与 DBOption。
// 允许调用方两种风格混用：
//
//	NewMemoryDB(t, &SysUser{}, testutil.WithDataScopeCallbacks())
func splitModelsAndOpts(in []any) (models []any, opts []DBOption) {
	for _, v := range in {
		if opt, ok := v.(DBOption); ok {
			opts = append(opts, opt)
			continue
		}
		models = append(models, v)
	}
	return models, opts
}
