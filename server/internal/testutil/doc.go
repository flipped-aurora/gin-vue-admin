// Package testutil 提供 gin-vue-admin 测试用例共享的单例初始化 helper。
//
// gva 本体大量使用裸全局变量（global.GVA_DB / GVA_LOG / GVA_CACHE / GVA_REDIS /
// GVA_CONFIG 等），各 _test.go 在 setup 阶段需要重复写样板代码：
// 打开 sqlite :memory:、AutoMigrate、GVA_LOG = zap.NewNop()、赋值并 t.Cleanup 还原。
// 本包把这些样板收敛为幂等、带 cleanup 的函数式 helper，写测试时一行即可完成初始化。
//
// # 基本用法
//
//	func TestFoo(t *testing.T) {
//	    db := testutil.NewMemoryDB(t, &SysUser{}, &SysAuthority{})
//	    _ = db // global.GVA_DB 已同步赋值
//	    // ...
//	}
//
// # 设计约束
//
//   - 默认 helper 会赋值对应的全局单例，并在 t.Cleanup 中还原旧值，
//     避免污染同一进程内其它测试。
//   - 这些全局变量是裸赋值、无锁（仅 GVA_DBList 等少数路径有锁）。
//     在 t.Parallel() 并行场景下请使用 *WithoutGlobal 变体（如
//     NewMemoryDBWithoutGlobal），每个测试持有独立实例，避免数据竞争。
//   - 本包不依赖 core.Viper()（其内部 flag.Parse 与 go test flag 冲突），
//     配置加载自建 viper 实例仅做 Unmarshal。
//   - 本包不调用 initialize.RegisterTables()（其失败会 os.Exit），
//     NewMemoryDB 让调用方按需传入 model 做精准 AutoMigrate。
package testutil
