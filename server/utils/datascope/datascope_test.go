package datascope

import (
	"context"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// scopedDoc 受控业务表: 操作人列齐全 + 软删除, 对应代码生成器 AutoCreateResource 的产物形态
type scopedDoc struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	CreatedBy uint           `gorm:"column:created_by"`
	UpdatedBy uint           `gorm:"column:updated_by"`
	DeletedBy uint           `gorm:"column:deleted_by"`
	DeptId    uint           `gorm:"column:dept_id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// plainDoc 受控(有归属列)但没有 updated_by/deleted_by 的表
type plainDoc struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	CreatedBy uint           `gorm:"column:created_by"`
	DeptId    uint           `gorm:"column:dept_id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// sysFake 系统表(sys_ 前缀): 引擎必须完全跳过
type sysFake struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	CreatedBy uint `gorm:"column:created_by"`
	UpdatedBy uint `gorm:"column:updated_by"`
	DeletedBy uint `gorm:"column:deleted_by"`
	DeptId    uint `gorm:"column:dept_id"`
	DeletedAt gorm.DeletedAt
}

func (sysFake) TableName() string { return "sys_fakes" }

// embDoc 生产真实形态: 嵌入 global.GVA_MODEL + 生成器四列
type embDoc struct {
	global.GVA_MODEL
	Title     string
	CreatedBy uint `gorm:"column:created_by"`
	UpdatedBy uint `gorm:"column:updated_by"`
	DeletedBy uint `gorm:"column:deleted_by"`
	DeptId    uint `gorm:"column:dept_id"`
}

// ptrDoc 指针形式软删除字段(原生 gorm 视为完整软删除表)
type ptrDoc struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	DeletedBy uint            `gorm:"column:deleted_by"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

// zvDoc 带 ZEROVALUE 标签的软删除字段(存活行=零值时间而非 NULL), 引擎必须不接管
type zvDoc struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	DeletedBy uint           `gorm:"column:deleted_by"`
	DeletedAt gorm.DeletedAt `gorm:"zerovalue:1970-01-01 00:00:01"`
}

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	if global.GVA_LOG == nil {
		global.GVA_LOG = zap.NewNop() // 无身份告警路径经 logger 走 GVA_LOG, 测试环境兜底
	}
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&scopedDoc{}, &plainDoc{}, &sysFake{}, &embDoc{}, &ptrDoc{}, &zvDoc{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	RegisterCallbacks(db)
	return db
}

func ctxAs(userID, deptID uint, scope int) context.Context {
	return WithIdentity(context.Background(), &Identity{
		UserID:         userID,
		AuthorityID:    100,
		DeptID:         deptID,
		DeptIDs:        []uint{deptID},
		VisibleDeptIDs: []uint{deptID},
		Scope:          scope,
	})
}

func TestCreateStampsOwnership(t *testing.T) {
	db := newTestDB(t)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctxAs(5, 3, ScopeAll)).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if doc.CreatedBy != 5 || doc.DeptId != 3 {
		t.Fatalf("create 盖章失败: created_by=%d dept_id=%d", doc.CreatedBy, doc.DeptId)
	}
}

func TestUpdateStampsUpdatedBy(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}

	// 单列 map 式更新
	if err := db.WithContext(ctx).Model(&scopedDoc{}).Where("id = ?", doc.ID).Update("title", "b").Error; err != nil {
		t.Fatal(err)
	}
	var got scopedDoc
	if err := db.WithContext(ctx).First(&got, doc.ID).Error; err != nil {
		t.Fatal(err)
	}
	if got.UpdatedBy != 5 {
		t.Fatalf("updated_by = %d, want 5", got.UpdatedBy)
	}

	// 换一个用户走 Save + Omit 归属列(与 service 层标准写法一致)
	got.Title = "c"
	if err := db.WithContext(ctxAs(7, 4, ScopeAll)).Omit("dept_id", "created_by").Save(&got).Error; err != nil {
		t.Fatal(err)
	}
	var got2 scopedDoc
	if err := db.WithContext(ctx).First(&got2, doc.ID).Error; err != nil {
		t.Fatal(err)
	}
	if got2.UpdatedBy != 7 {
		t.Fatalf("save 后 updated_by = %d, want 7", got2.UpdatedBy)
	}
	if got2.CreatedBy != 5 || got2.DeptId != 3 {
		t.Fatalf("归属列被覆盖: created_by=%d dept_id=%d", got2.CreatedBy, got2.DeptId)
	}
}

func TestSoftDeleteStampsDeletedBy(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}

	res := db.WithContext(ctx).Delete(&doc)
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	if res.RowsAffected != 1 {
		t.Fatalf("rows = %d, want 1", res.RowsAffected)
	}
	// 调用方结构体被回填(与 gorm 软删除对 DeletedAt 的行为一致)
	if !doc.DeletedAt.Valid || doc.DeletedBy != 5 {
		t.Fatalf("dest 未回填: deleted_at.valid=%v deleted_by=%d", doc.DeletedAt.Valid, doc.DeletedBy)
	}

	// 常规查询不可见
	var count int64
	db.WithContext(ctx).Model(&scopedDoc{}).Where("id = ?", doc.ID).Count(&count)
	if count != 0 {
		t.Fatal("软删除后仍可见")
	}
	// Unscoped 可见, 且 deleted_at/deleted_by 同一条 UPDATE 落库
	var got scopedDoc
	if err := db.WithContext(ctx).Unscoped().First(&got, doc.ID).Error; err != nil {
		t.Fatal(err)
	}
	if !got.DeletedAt.Valid {
		t.Fatal("deleted_at 未设置")
	}
	if got.DeletedBy != 5 {
		t.Fatalf("deleted_by = %d, want 5", got.DeletedBy)
	}
}

func TestSoftDeleteBatchWithConds(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	for _, title := range []string{"a", "b"} {
		if err := db.WithContext(ctx).Create(&scopedDoc{Title: title}).Error; err != nil {
			t.Fatal(err)
		}
	}
	res := db.WithContext(ctx).Where("dept_id = ?", 3).Delete(&scopedDoc{})
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	if res.RowsAffected != 2 {
		t.Fatalf("rows = %d, want 2", res.RowsAffected)
	}
	var docs []scopedDoc
	db.WithContext(ctx).Unscoped().Find(&docs)
	for _, d := range docs {
		if d.DeletedBy != 5 || !d.DeletedAt.Valid {
			t.Fatalf("批量删除未盖章: %+v", d)
		}
	}
}

// TestSoftDeleteRespectsScope 关键安全用例: 盖章回调抢先 Build SQL,
// 必须仍包含 data_scope:delete 已追加的数据范围 WHERE, 否则过滤被旁路
func TestSoftDeleteRespectsScope(t *testing.T) {
	db := newTestDB(t)
	doc := scopedDoc{Title: "of5"}
	if err := db.WithContext(ctxAs(5, 3, ScopeAll)).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}

	// 用户9, "仅本人"档, 尝试删用户5创建的行
	res := db.WithContext(ctxAs(9, 4, ScopeSelf)).Delete(&scopedDoc{}, doc.ID)
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	if res.RowsAffected != 0 {
		t.Fatalf("数据范围被旁路! rows = %d", res.RowsAffected)
	}
	var got scopedDoc
	if err := db.WithContext(ctxAs(5, 3, ScopeAll)).First(&got, doc.ID).Error; err != nil {
		t.Fatalf("行不应被删除: %v", err)
	}
	if got.DeletedBy != 0 {
		t.Fatalf("不应盖章: deleted_by = %d", got.DeletedBy)
	}

	// 本人删除自己的行 → 成功且盖章
	res = db.WithContext(ctxAs(5, 3, ScopeSelf)).Delete(&scopedDoc{}, doc.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("本人删除失败: err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got2 scopedDoc
	db.WithContext(ctxAs(5, 3, ScopeAll)).Unscoped().First(&got2, doc.ID)
	if got2.DeletedBy != 5 {
		t.Fatalf("deleted_by = %d, want 5", got2.DeletedBy)
	}
}

func TestSoftDeleteWithoutIdentity(t *testing.T) {
	db := newTestDB(t)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctxAs(5, 3, ScopeAll)).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	// 无身份 ctx: 走标准软删除(护栏放行), 不盖章
	res := db.Delete(&scopedDoc{}, doc.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got scopedDoc
	db.Unscoped().First(&got, doc.ID)
	if !got.DeletedAt.Valid {
		t.Fatal("应正常软删除")
	}
	if got.DeletedBy != 0 {
		t.Fatalf("无身份不应盖章: deleted_by = %d", got.DeletedBy)
	}
}

func TestSoftDeleteSystemContext(t *testing.T) {
	db := newTestDB(t)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctxAs(5, 3, ScopeAll)).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(WithSystem(context.Background())).Delete(&scopedDoc{}, doc.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got scopedDoc
	db.WithContext(WithSystem(context.Background())).Unscoped().First(&got, doc.ID)
	if !got.DeletedAt.Valid || got.DeletedBy != 0 {
		t.Fatalf("系统上下文应软删除且不盖章: valid=%v deleted_by=%d", got.DeletedAt.Valid, got.DeletedBy)
	}
}

func TestUnscopedHardDelete(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(ctx).Unscoped().Delete(&scopedDoc{}, doc.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var count int64
	db.WithContext(ctx).Unscoped().Model(&scopedDoc{}).Where("id = ?", doc.ID).Count(&count)
	if count != 0 {
		t.Fatal("Unscoped 删除应为物理删除")
	}
}

func TestDeleteWithoutConditionsStillErrors(t *testing.T) {
	db := newTestDB(t)
	// ScopeAll 不追加范围条件, 无主键无条件删除必须仍触发 gorm 的缺 WHERE 护栏
	err := db.WithContext(ctxAs(5, 3, ScopeAll)).Delete(&scopedDoc{}).Error
	if err == nil {
		t.Fatal("无条件删除应报错(ErrMissingWhereClause)")
	}
}

func TestPlainTableSoftDeleteUnaffected(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := plainDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(ctx).Delete(&plainDoc{}, doc.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got plainDoc
	db.WithContext(ctx).Unscoped().First(&got, doc.ID)
	if !got.DeletedAt.Valid {
		t.Fatal("无 deleted_by 列的表应走标准软删除")
	}
}

// TestUnconditionedWriteBlockedUnderScope 非 ScopeAll 档下, 漏写条件的 update/delete
// 必须报错(ErrMissingWhereClause), 而不是把引擎注入的范围 WHERE 当条件、删/改光整个可见范围
func TestUnconditionedWriteBlockedUnderScope(t *testing.T) {
	db := newTestDB(t)
	owner := ctxAs(5, 3, ScopeAll)
	for _, title := range []string{"a", "b"} {
		if err := db.WithContext(owner).Create(&scopedDoc{Title: title}).Error; err != nil {
			t.Fatal(err)
		}
	}
	self := ctxAs(5, 3, ScopeSelf)
	if err := db.WithContext(self).Delete(&scopedDoc{}).Error; err == nil {
		t.Fatal("无条件删除应报错, 而非静默软删可见范围")
	}
	if err := db.WithContext(self).Model(&scopedDoc{}).Update("title", "X").Error; err == nil {
		t.Fatal("无条件更新应报错, 而非静默改写可见范围")
	}
	var count int64
	db.WithContext(owner).Model(&scopedDoc{}).Count(&count)
	if count != 2 {
		t.Fatalf("行不应被动过: count=%d", count)
	}
	var hit int64
	db.WithContext(owner).Model(&scopedDoc{}).Where("title = ?", "X").Count(&hit)
	if hit != 0 {
		t.Fatal("行被改写")
	}
}

// TestAllowGlobalUpdateDeleteStamps 显式声明的全量软删除照常执行且盖章
func TestAllowGlobalUpdateDeleteStamps(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	for _, title := range []string{"a", "b"} {
		if err := db.WithContext(ctx).Create(&scopedDoc{Title: title}).Error; err != nil {
			t.Fatal(err)
		}
	}
	res := db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&scopedDoc{})
	if res.Error != nil || res.RowsAffected != 2 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var docs []scopedDoc
	db.WithContext(ctx).Unscoped().Find(&docs)
	for _, d := range docs {
		if !d.DeletedAt.Valid || d.DeletedBy != 5 {
			t.Fatalf("全量软删应盖章: %+v", d)
		}
	}
}

// TestUpdateColumnSkipsStamp UpdateColumn(SkipHooks)语义是"只改指定列",
// 与 gorm 跳过 AutoUpdateTime 一致, updated_by 也不盖
func TestUpdateColumnSkipsStamp(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.WithContext(ctxAs(9, 4, ScopeAll)).Model(&scopedDoc{}).Where("id = ?", doc.ID).UpdateColumn("title", "z").Error; err != nil {
		t.Fatal(err)
	}
	var got scopedDoc
	db.WithContext(ctx).First(&got, doc.ID)
	if got.Title != "z" {
		t.Fatal("UpdateColumn 未生效")
	}
	if got.UpdatedBy != 0 {
		t.Fatalf("SkipHooks 更新不应盖章: updated_by=%d", got.UpdatedBy)
	}
}

// TestByValueModelUpdateNotBroken Model(值)+Updates(结构体值) 是 vanilla gorm 下合法写法,
// 盖章不能把它打挂(该形态跳过盖章)
func TestByValueModelUpdateNotBroken(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.WithContext(ctx).Model(scopedDoc{}).Where("id = ?", doc.ID).Updates(scopedDoc{Title: "vv"}).Error; err != nil {
		t.Fatalf("Model(值)+Updates(结构体) 不应被盖章打挂: %v", err)
	}
	var got scopedDoc
	db.WithContext(ctx).First(&got, doc.ID)
	if got.Title != "vv" {
		t.Fatal("更新未生效")
	}
}

// TestSelectRestrictedUpdateStillStamps Select 白名单更新时 updated_by 被补进白名单, 不被静默剔除
func TestSelectRestrictedUpdateStillStamps(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.WithContext(ctxAs(7, 3, ScopeAll)).Model(&scopedDoc{}).Select("title").
		Where("id = ?", doc.ID).Updates(map[string]interface{}{"title": "sel"}).Error; err != nil {
		t.Fatal(err)
	}
	var got scopedDoc
	db.WithContext(ctx).First(&got, doc.ID)
	if got.Title != "sel" || got.UpdatedBy != 7 {
		t.Fatalf("Select 白名单下盖章丢失: title=%s updated_by=%d", got.Title, got.UpdatedBy)
	}
}

// TestEmbeddedGVAModelStamps 生产真实形态(嵌入 global.GVA_MODEL)全链路盖章
func TestEmbeddedGVAModelStamps(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := embDoc{Title: "e"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if doc.CreatedBy != 5 || doc.DeptId != 3 {
		t.Fatalf("create 盖章失败: %+v", doc)
	}
	if err := db.WithContext(ctx).Model(&embDoc{}).Where("id = ?", doc.ID).Update("title", "e2").Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(ctx).Delete(&doc)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got embDoc
	if err := db.WithContext(ctx).Unscoped().First(&got, doc.ID).Error; err != nil {
		t.Fatal(err)
	}
	if got.UpdatedBy != 5 || got.DeletedBy != 5 || !got.DeletedAt.Valid {
		t.Fatalf("嵌入形态盖章失败: updated_by=%d deleted_by=%d deleted_at.valid=%v",
			got.UpdatedBy, got.DeletedBy, got.DeletedAt.Valid)
	}
}

// TestPointerDeletedAtStamps 指针形式软删除字段同样接管盖章
func TestPointerDeletedAtStamps(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	d := ptrDoc{Title: "p"}
	if err := db.WithContext(ctx).Create(&d).Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(ctx).Delete(&d)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got ptrDoc
	if err := db.WithContext(ctx).Unscoped().First(&got, d.ID).Error; err != nil {
		t.Fatal(err)
	}
	if got.DeletedAt == nil || !got.DeletedAt.Valid || got.DeletedBy != 5 {
		t.Fatalf("指针软删字段盖章失败: %+v", got)
	}
}

// TestZeroValueTagFallsBackToNative ZEROVALUE 标签语义与 IS NULL 口径冲突, 必须不接管
func TestZeroValueTagFallsBackToNative(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	tx := db.WithContext(ctx).Session(&gorm.Session{DryRun: true}).Where("title = ?", "x").Delete(&zvDoc{})
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	sql := tx.Statement.SQL.String()
	if strings.Contains(sql, "deleted_by") {
		t.Fatalf("ZEROVALUE 表不应被接管: %s", sql)
	}
	if !strings.Contains(sql, "deleted_at") {
		t.Fatalf("应为原生软删除 UPDATE: %s", sql)
	}
}

// TestUpdateIdentityGuards 无身份/系统上下文的更新不盖章(与 delete 侧对称)
func TestUpdateIdentityGuards(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeAll)
	doc := scopedDoc{Title: "a"}
	if err := db.WithContext(ctx).Create(&doc).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Model(&scopedDoc{}).Where("id = ?", doc.ID).Update("title", "n1").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.WithContext(WithSystem(context.Background())).Model(&scopedDoc{}).Where("id = ?", doc.ID).Update("title", "n2").Error; err != nil {
		t.Fatal(err)
	}
	var got scopedDoc
	db.WithContext(ctx).First(&got, doc.ID)
	if got.UpdatedBy != 0 {
		t.Fatalf("无身份/系统上下文不应盖章: updated_by=%d", got.UpdatedBy)
	}
}

func TestSystemTableSkipsAllStamps(t *testing.T) {
	db := newTestDB(t)
	ctx := ctxAs(5, 3, ScopeSelf)
	f := sysFake{Title: "x"}
	if err := db.WithContext(ctx).Create(&f).Error; err != nil {
		t.Fatal(err)
	}
	if f.CreatedBy != 0 || f.DeptId != 0 {
		t.Fatalf("sys_ 表不应盖章: %+v", f)
	}
	if err := db.WithContext(ctx).Model(&sysFake{}).Where("id = ?", f.ID).Update("title", "y").Error; err != nil {
		t.Fatal(err)
	}
	res := db.WithContext(ctx).Delete(&sysFake{}, f.ID)
	if res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("sys_ 表应跳过范围过滤: err=%v rows=%d", res.Error, res.RowsAffected)
	}
	var got sysFake
	db.WithContext(ctx).Unscoped().First(&got, f.ID)
	if got.UpdatedBy != 0 || got.DeletedBy != 0 {
		t.Fatalf("sys_ 表不应盖章: updated_by=%d deleted_by=%d", got.UpdatedBy, got.DeletedBy)
	}
	if !got.DeletedAt.Valid {
		t.Fatal("sys_ 表应正常软删除")
	}
}
