package testutil

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// dummyModel 仅供测试 AutoMigrate 生效检查
type dummyModel struct {
	ID    uint   `gorm:"primarykey"`
	Title string
}

func (dummyModel) TableName() string { return "testutil_dummy_models" }

func TestNewMemoryDB_AssignsGlobalAndMigrates(t *testing.T) {
	old := global.GVA_DB
	t.Cleanup(func() { global.GVA_DB = old })

	db := NewMemoryDB(t, &dummyModel{})
	if db == nil {
		t.Fatal("NewMemoryDB 返回 nil")
	}
	if global.GVA_DB == nil {
		t.Fatal("global.GVA_DB 未被赋值")
	}
	if global.GVA_DB != db {
		t.Fatal("global.GVA_DB 与返回值不一致")
	}

	// 验证 AutoMigrate 已建表
	if !db.Migrator().HasTable(&dummyModel{}) {
		t.Fatal("AutoMigrate 未生效：表未创建")
	}

	// 验证写入可用
	if err := db.Create(&dummyModel{Title: "hello"}).Error; err != nil {
		t.Fatalf("写入失败: %v", err)
	}
}

func TestNewMemoryDB_RestoresOldValueOnCleanup(t *testing.T) {
	// 先放一个已知的旧值
	sentinel := &gorm.DB{}
	old := global.GVA_DB
	global.GVA_DB = sentinel
	t.Cleanup(func() { global.GVA_DB = old })

	sub := &fakeT{TB: t}
	NewMemoryDB(sub, &dummyModel{})
	sub.runCleanup()

	if global.GVA_DB != sentinel {
		t.Fatal("cleanup 未还原 global.GVA_DB 到旧值")
	}
}

func TestNewMemoryDB_OptionsChain(t *testing.T) {
	old := global.GVA_DB
	t.Cleanup(func() { global.GVA_DB = old })

	// 混用 model 与 DBOption
	db := NewMemoryDB(t, &dummyModel{}, WithModels(&dummyModel{}), WithDataScopeCallbacks())
	if !db.Migrator().HasTable(&dummyModel{}) {
		t.Fatal("AutoMigrate 未生效")
	}
	// data_scope 回调幂等注册，这里仅检查不 panic 即可
}

func TestNewMemoryDB_SkipAutoMigrate(t *testing.T) {
	old := global.GVA_DB
	t.Cleanup(func() { global.GVA_DB = old })

	db := NewMemoryDB(t, WithoutAutoMigrate())
	if db.Migrator().HasTable(&dummyModel{}) {
		t.Fatal("WithoutAutoMigrate 未生效：不应建任何表")
	}
}

func TestNewMemoryDBWithoutGlobal_DoesNotTouchGlobal(t *testing.T) {
	sentinel := &gorm.DB{}
	old := global.GVA_DB
	global.GVA_DB = sentinel
	t.Cleanup(func() { global.GVA_DB = old })

	db := NewMemoryDBWithoutGlobal(t, &dummyModel{})
	if db == nil {
		t.Fatal("NewMemoryDBWithoutGlobal 返回 nil")
	}
	if global.GVA_DB != sentinel {
		t.Fatal("NewMemoryDBWithoutGlobal 不应修改 global.GVA_DB")
	}
	if !db.Migrator().HasTable(&dummyModel{}) {
		t.Fatal("AutoMigrate 未生效")
	}
}

func TestSplitModelsAndOpts(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m, o := splitModelsAndOpts(nil)
		if len(m) != 0 || len(o) != 0 {
			t.Fatal("空输入应返回空切片")
		}
	})
	t.Run("mixed", func(t *testing.T) {
		m, o := splitModelsAndOpts([]any{
			&dummyModel{},
			WithDataScopeCallbacks(),
			WithModels(&dummyModel{}),
		})
		// 直接 model 1 个；WithModels 内部的 model 不在此层展开（它在 option 内部）
		if len(m) != 1 {
			t.Fatalf("model 数量错误: %d", len(m))
		}
		// WithDataScopeCallbacks + WithModels = 2 个 DBOption
		if len(o) != 2 {
			t.Fatalf("option 数量错误: %d", len(o))
		}
	})
}

// fakeT 包装 testing.T，并记录 t.Cleanup 注册的函数，便于测试显式触发 cleanup。
type fakeT struct {
	testing.TB
	cleanups []func()
}

func (f *fakeT) Cleanup(fn func()) {
	f.cleanups = append(f.cleanups, fn)
}

func (f *fakeT) runCleanup() {
	for _, fn := range f.cleanups {
		fn()
	}
	f.cleanups = nil
}
