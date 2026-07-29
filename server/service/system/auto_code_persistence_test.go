package system

import (
	"context"
	"errors"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

func newAutoCodePersistenceDB(t *testing.T) *gorm.DB {
	t.Helper()
	return testutil.NewMemoryDBWithoutGlobal(t,
		&model.SysApi{},
		&model.SysBaseMenu{},
		&model.SysBaseMenuBtn{},
		&model.SysExportTemplate{},
		&model.SysAutoCodeHistory{},
	)
}

func validAutoCodePersistenceInfo() request.AutoCode {
	return request.AutoCode{
		Package:             "example",
		TableName:           "orders",
		BusinessDB:          "business",
		StructName:          "Order",
		PackageName:         "order",
		HumpPackageName:     "order",
		Abbreviation:        "order",
		Description:         "订单",
		AutoCreateApiToSql:  true,
		AutoCreateMenuToSql: true,
		AutoCreateBtnAuth:   true,
		HasExcel:            true,
		Fields: []*request.AutoCodeField{
			{ColumnName: "number", FieldDesc: "订单号", Excel: true},
		},
	}
}

func assertAutoCodeTableCount(t *testing.T, db *gorm.DB, value any, want int64) {
	t.Helper()
	var got int64
	if err := db.Model(value).Count(&got).Error; err != nil {
		t.Fatalf("Count(%T) error = %v", value, err)
	}
	if got != want {
		t.Fatalf("Count(%T) = %d, want %d", value, got, want)
	}
}

func TestPersistAutoCodeMetadataRollsBackAllRowsWhenHistoryFails(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	wantFailure := errors.New("history failed")
	if err := db.Callback().Create().Before("gorm:create").Register("test:fail_history", func(tx *gorm.DB) {
		if tx.Statement.Schema != nil && tx.Statement.Schema.Table == "sys_auto_code_histories" {
			tx.AddError(wantFailure)
		}
	}); err != nil {
		t.Fatalf("register callback: %v", err)
	}
	t.Cleanup(func() { _ = db.Callback().Create().Remove("test:fail_history") })

	info := validAutoCodePersistenceInfo()
	err := persistAutoCodeMetadata(context.Background(), db, info, "package", info.History())
	if !errors.Is(err, wantFailure) {
		t.Fatalf("persistAutoCodeMetadata() error = %v, want %v", err, wantFailure)
	}
	assertAutoCodeTableCount(t, db, &model.SysApi{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysBaseMenu{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysBaseMenuBtn{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysExportTemplate{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysAutoCodeHistory{}, 0)
}

func TestPersistAutoCodeMetadataCommitsCompleteCreateSet(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	info := validAutoCodePersistenceInfo()

	if err := persistAutoCodeMetadata(context.Background(), db, info, "package", info.History()); err != nil {
		t.Fatalf("persistAutoCodeMetadata() error = %v", err)
	}
	assertAutoCodeTableCount(t, db, &model.SysApi{}, int64(len(info.Apis())))
	assertAutoCodeTableCount(t, db, &model.SysBaseMenu{}, 1)
	assertAutoCodeTableCount(t, db, &model.SysBaseMenuBtn{}, 8)
	assertAutoCodeTableCount(t, db, &model.SysExportTemplate{}, 1)
	assertAutoCodeTableCount(t, db, &model.SysAutoCodeHistory{}, 1)

	var history model.SysAutoCodeHistory
	if err := db.First(&history).Error; err != nil {
		t.Fatalf("load history: %v", err)
	}
	if len(history.ApiIDs) != len(info.Apis()) {
		t.Fatalf("history API IDs = %v, want %d IDs", history.ApiIDs, len(info.Apis()))
	}
	if history.MenuID == 0 || history.ExportTemplateID == 0 {
		t.Fatalf("history relation IDs = menu:%d export:%d, want non-zero", history.MenuID, history.ExportTemplateID)
	}
	var orphanButtons int64
	if err := db.Model(&model.SysBaseMenuBtn{}).Where("sys_base_menu_id = ?", 0).Count(&orphanButtons).Error; err != nil {
		t.Fatalf("count orphan menu buttons: %v", err)
	}
	if orphanButtons != 0 {
		t.Fatalf("orphan menu buttons = %d, want 0", orphanButtons)
	}
}

func TestPersistAutoCodeMetadataReusesExactMenu(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	info := validAutoCodePersistenceInfo()
	info.AutoCreateApiToSql = false
	info.AutoCreateBtnAuth = false
	info.HasExcel = false
	existing := info.Menu("package")
	if err := db.Create(&existing).Error; err != nil {
		t.Fatalf("create existing menu: %v", err)
	}

	if err := persistAutoCodeMetadata(context.Background(), db, info, "package", info.History()); err != nil {
		t.Fatalf("persistAutoCodeMetadata() error = %v", err)
	}
	assertAutoCodeTableCount(t, db, &model.SysBaseMenu{}, 1)
	var history model.SysAutoCodeHistory
	if err := db.First(&history).Error; err != nil {
		t.Fatalf("load history: %v", err)
	}
	if history.MenuID != existing.ID {
		t.Fatalf("history menu ID = %d, want %d", history.MenuID, existing.ID)
	}
}

func TestPersistAutoCodeMetadataContinuesAfterRequestCancellation(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	info := validAutoCodePersistenceInfo()
	info.AutoCreateApiToSql = false
	info.AutoCreateMenuToSql = false
	info.HasExcel = false

	requestCtx, cancel := context.WithCancel(context.Background())
	createCtx := autoCodeCreateContext(requestCtx)
	cancel()
	if err := persistAutoCodeMetadata(createCtx, db, info, "package", info.History()); err != nil {
		t.Fatalf("persistAutoCodeMetadata() after request cancellation error = %v", err)
	}
	assertAutoCodeTableCount(t, db, &model.SysAutoCodeHistory{}, 1)
}

func TestPersistAutoCodeMetadataRejectsSameNameDifferentMenu(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	info := validAutoCodePersistenceInfo()
	existing := info.Menu("package")
	existing.Component = "view/other/other.vue"
	if err := db.Create(&existing).Error; err != nil {
		t.Fatalf("create conflicting menu: %v", err)
	}

	err := persistAutoCodeMetadata(context.Background(), db, info, "package", info.History())
	if !errors.Is(err, errAutoCodeMenuConflict) {
		t.Fatalf("persistAutoCodeMetadata() error = %v, want %v", err, errAutoCodeMenuConflict)
	}
	assertAutoCodeTableCount(t, db, &model.SysApi{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysExportTemplate{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysAutoCodeHistory{}, 0)
	assertAutoCodeTableCount(t, db, &model.SysBaseMenu{}, 1)
}

func TestAutoCodeHistoryPreservesAbbreviation(t *testing.T) {
	info := validAutoCodePersistenceInfo()
	historyRequest := info.History()
	if got := historyRequest.Create().Abbreviation; got != info.Abbreviation {
		t.Fatalf("history abbreviation = %q, want %q", got, info.Abbreviation)
	}
}

func TestAutoCodeIdentityExistsReturnsQueryError(t *testing.T) {
	db := newAutoCodePersistenceDB(t)
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("db.DB() error = %v", err)
	}
	if err = sqlDB.Close(); err != nil {
		t.Fatalf("close database: %v", err)
	}
	if _, err = autoCodeIdentityExists(context.Background(), db, validAutoCodePersistenceInfo()); err == nil {
		t.Fatal("autoCodeIdentityExists() error = nil, want query error")
	}
}
