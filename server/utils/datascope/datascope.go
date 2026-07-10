// Package datascope 统一数据权限引擎: 通过 GORM 全局回调实现"行级默认安全"。
//
// 设计要点:
//   - 身份由 DataScope 中间件注入 request.Context(), 回调从 db.Statement.Context 读取
//   - "漏写=默认安全": 只有带 dept_id/created_by 归属列的业务表才会被约束; 系统表自动跳过
//   - 操作人盖章按列存在与否自动参与: 创建盖 created_by/dept_id, 更新盖 updated_by,
//     软删除把 deleted_by 并进同一条 UPDATE
//   - 与 Casbin 正交: Casbin 管"能不能调接口", 本引擎管"能看/改/删哪些行"
package datascope

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// 数据范围档位(挂在角色 SysAuthority.DataScope 上)
const (
	ScopeAll          = 1 // 全部
	ScopeDeptAndChild = 2 // 本部门及以下
	ScopeDept         = 3 // 本部门(不含子级)
	ScopeSelf         = 4 // 仅本人
	ScopeCustom       = 5 // 自定义部门集
)

// 审计事件类型
const (
	EventNoIdentity   = "no_identity"   // 无身份访问受控表(潜在旁路, 待补 ctx)
	EventBlockedWrite = "blocked_write" // 数据范围过滤后写操作影响 0 行(疑似越权尝试)
)

// Identity 一次请求的数据权限身份, 由 DataScope 中间件构建并注入 request.Context()
type Identity struct {
	UserID         uint   // 当前用户ID
	AuthorityID    uint   // 当前角色ID
	DeptID         uint   // 主部门(创建时盖章)
	DeptIDs        []uint // 直接归属部门(不含子级, 用于"本部门"档)
	VisibleDeptIDs []uint // 可见部门并集(所有归属部门的子树之并, 用于"本部门及以下"档)
	CustomDeptIDs  []uint // 自定义部门集(角色配置, 用于"自定义部门"档)
	Scope          int    // 数据权限档位("看全部"由 ScopeAll 表达, 不设超管旁路)
	IsSystem       bool   // 系统上下文(定时任务/CLI/初始化, 放行全部)
}

// AuditEvent 数据权限审计事件, 经 SetAuditHook 注入的写入器异步落表 sys_data_access_logs
type AuditEvent struct {
	EventType   string // 见 Event* 常量
	TargetTable string // 受控业务表名
	Operation   string // query / create / update / delete
	UserID      uint   // 事发用户(无身份事件为 0)
	AuthorityID uint
	Scope       int
	RequestID   string // 请求链路信息(来自 logger fields)
	Method      string
	Path        string
	Detail      string
}

// auditHook 由 initialize 注入(指向审计服务的异步入队), 解耦包依赖方向
var auditHook func(AuditEvent)

// SetAuditHook 注册审计事件写入器(应为非阻塞实现)
func SetAuditHook(h func(AuditEvent)) {
	auditHook = h
}

func emitAudit(db *gorm.DB, eventType, operation, detail string, id *Identity) {
	if auditHook == nil {
		return
	}
	evt := AuditEvent{
		EventType:   eventType,
		TargetTable: db.Statement.Table,
		Operation:   operation,
		Detail:      detail,
	}
	if id != nil {
		evt.UserID = id.UserID
		evt.AuthorityID = id.AuthorityID
		evt.Scope = id.Scope
	}
	if f := logger.FromCtx(db.Statement.Context); f != nil {
		evt.RequestID = f.RequestID
		evt.Method = f.HTTPMethod
		evt.Path = f.HTTPPath
	}
	auditHook(evt)
}

type identityKey struct{}

// WithIdentity 把数据权限身份写入 context
func WithIdentity(ctx context.Context, id *Identity) context.Context {
	return context.WithValue(ctx, identityKey{}, id)
}

// FromContext 从 context 取出数据权限身份
func FromContext(ctx context.Context) (*Identity, bool) {
	if ctx == nil {
		return nil, false
	}
	id, ok := ctx.Value(identityKey{}).(*Identity)
	return id, ok
}

// WithSystem 标记为系统上下文(定时任务/CLI/初始化), 数据权限回调一律放行。
// 用于无请求身份但确属系统行为的场景, 避免"无身份"被当成潜在越权而告警。
func WithSystem(ctx context.Context) context.Context {
	return WithIdentity(ctx, &Identity{IsSystem: true})
}

// RegisterCallbacks 在给定连接上注册数据权限全局回调。
// 全仓 green field(无既有自定义 callback), 命名回调保证幂等, 可安全重复调用。
//
// 约束: 不要在本函数之后再向同一连接注册 Before("gorm:delete") 的限制性回调(如租户过滤)——
// data_scope:stamp_delete 会当场 Build 定稿 SQL, 晚于它的回调对子句的修改会被静默丢弃。
func RegisterCallbacks(db *gorm.DB) {
	if db == nil {
		return
	}
	q := db.Callback().Query()
	if q.Get("data_scope:query") == nil {
		_ = q.Before("gorm:query").Register("data_scope:query", filterFor("query"))
	}
	u := db.Callback().Update()
	if u.Get("data_scope:update") == nil {
		_ = u.Before("gorm:update").Register("data_scope:update", filterFor("update"))
	}
	if u.Get("data_scope:stamp_update") == nil {
		_ = u.Before("gorm:update").Register("data_scope:stamp_update", stampUpdatedBy)
	}
	if u.Get("data_scope:audit_update") == nil {
		_ = u.After("gorm:update").Register("data_scope:audit_update", auditBlockedWrite("update"))
	}
	d := db.Callback().Delete()
	if d.Get("data_scope:delete") == nil {
		_ = d.Before("gorm:delete").Register("data_scope:delete", filterFor("delete"))
	}
	if d.Get("data_scope:stamp_delete") == nil {
		// 必须排在 data_scope:delete 之后: 本回调会当场 Build 定稿 SQL,
		// 数据范围 WHERE 必须已经在语句上, 否则过滤被旁路
		_ = d.Before("gorm:delete").After("data_scope:delete").Register("data_scope:stamp_delete", stampDeletedBy)
	}
	if d.Get("data_scope:audit_delete") == nil {
		_ = d.After("gorm:delete").Register("data_scope:audit_delete", auditBlockedWrite("delete"))
	}
	c := db.Callback().Create()
	if c.Get("data_scope:stamp") == nil {
		_ = c.Before("gorm:create").Register("data_scope:stamp", stampOwnership)
	}
}

func hasField(db *gorm.DB, name string) bool {
	return db.Statement.Schema != nil && db.Statement.Schema.LookUpField(name) != nil
}

// isSystemTable 系统管理表(sys_*)不参与行级数据权限:
// 其管理由 Casbin/菜单把关; 且 sys_users.dept_id 表示"主体的部门"而非"记录的归属部门"。
func isSystemTable(db *gorm.DB) bool {
	return strings.HasPrefix(db.Statement.Table, "sys_")
}

// filterFor 生成带操作类型的过滤回调(操作类型用于审计记录)
func filterFor(operation string) func(*gorm.DB) {
	return func(db *gorm.DB) { filterByScope(db, operation) }
}

// filterByScope 给 Query/Update/Delete 追加数据范围 WHERE。
// 护栏顺序很关键: 先按"表是否有归属列"过滤掉系统表, 再判身份, 保证系统表永不告警。
func filterByScope(db *gorm.DB, operation string) {
	if db.Statement.Schema == nil {
		return
	}
	// 同一 statement 只处理一次(Count 后 Find 复用 statement, 避免重复追加 WHERE)
	if _, done := db.Statement.Clauses["data_scope:applied"]; done {
		return
	}
	// 显式旁路: db.Set("data_scope:skip", true)
	if v, ok := db.Get("data_scope:skip"); ok {
		if b, _ := v.(bool); b {
			return
		}
	}
	// 系统管理表(sys_*)不做行级约束(sys_users 虽有 dept_id, 但那是主体的部门)
	if isSystemTable(db) {
		return
	}
	hasDept := hasField(db, "dept_id")
	hasCreatedBy := hasField(db, "created_by")
	// 无归属列的表 → 完全不碰
	if !hasDept && !hasCreatedBy {
		return
	}
	id, ok := FromContext(db.Statement.Context)
	if !ok || id == nil {
		// 无身份(定时任务/迁移/漏补ctx) → MVP 阶段响亮告警但放行, 并落审计表
		logger.WithCtx(db.Statement.Context).Mod("datascope").
			Warn("数据权限: 业务表[" + db.Statement.Table + "]访问无身份上下文, 已放行(待补 ctx / 或用 WithSystem)")
		emitAudit(db, EventNoIdentity, operation, "无身份上下文访问受控表, 已放行", nil)
		return
	}
	if id.IsSystem {
		return
	}
	// 漏写条件的 update/delete: 不注入范围 WHERE, 交回 gorm 的 ErrMissingWhereClause 护栏。
	// 否则引擎注入的范围条件会顶掉该护栏, 把"开发期报错"静默降级成"删/改光整个可见范围"。
	// AllowGlobalUpdate 属调用方显式声明的全量写, 照常注入(把全表收敛为范围内)。
	if (operation == "update" || operation == "delete") && !db.AllowGlobalUpdate && !hasWriteConditions(db) {
		return
	}

	db.Statement.Clauses["data_scope:applied"] = clause.Clause{}
	table := db.Statement.Table
	deptCol := table + ".dept_id"
	createdByCol := table + ".created_by"

	switch id.Scope {
	case ScopeAll:
		return
	case ScopeDeptAndChild:
		if hasDept {
			ids := id.VisibleDeptIDs
			if len(ids) == 0 {
				ids = []uint{0} // 无可见部门 → 匹配不到任何行(安全默认)
			}
			db.Where(deptCol+" IN ?", ids)
		}
	case ScopeDept:
		if hasDept {
			ids := id.DeptIDs
			if len(ids) == 0 {
				ids = []uint{0}
			}
			db.Where(deptCol+" IN ?", ids)
		}
	case ScopeSelf:
		if hasCreatedBy {
			db.Where(createdByCol+" = ?", id.UserID)
		} else if hasDept {
			ids := id.DeptIDs // 无 created_by 时降级为本部门
			if len(ids) == 0 {
				ids = []uint{0}
			}
			db.Where(deptCol+" IN ?", ids)
		}
	case ScopeCustom:
		if hasDept {
			ids := id.CustomDeptIDs
			if len(ids) == 0 {
				ids = []uint{0} // 未配置部门集 → 匹配不到任何行(安全默认)
			}
			db.Where(deptCol+" IN ?", ids)
		}
	}
}

// auditBlockedWrite 写操作(update/delete)被数据范围过滤后影响 0 行 → 疑似越权尝试, 落审计。
// 注意这是启发式信号: 目标行本就不存在也会命中, 用于排查而非定罪。
func auditBlockedWrite(operation string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if _, applied := db.Statement.Clauses["data_scope:applied"]; !applied {
			return
		}
		if db.Error != nil || db.RowsAffected != 0 {
			return
		}
		id, _ := FromContext(db.Statement.Context)
		if id == nil || id.Scope == ScopeAll {
			return
		}
		emitAudit(db, EventBlockedWrite, operation, "数据范围过滤后写操作影响 0 行(疑似越权尝试)", id)
	}
}

// stampOwnership 创建时自动盖 created_by / dept_id(仅当列存在)。
func stampOwnership(db *gorm.DB) {
	if db.Statement.Schema == nil || isSystemTable(db) {
		return
	}
	id, ok := FromContext(db.Statement.Context)
	if !ok || id == nil || id.IsSystem {
		return
	}
	if hasField(db, "created_by") && id.UserID != 0 {
		db.Statement.SetColumn("created_by", id.UserID, true)
	}
	if hasField(db, "dept_id") && id.DeptID != 0 {
		db.Statement.SetColumn("dept_id", id.DeptID, true)
	}
}

// hasWriteConditions 判断 update/delete 语句在引擎注入范围条件前是否已自带条件:
// 显式 WHERE, 或 Dest/Model 上的主键值(后者由 gorm 在后续流程转为主键 WHERE)。
func hasWriteConditions(db *gorm.DB) bool {
	stmt := db.Statement
	if c, ok := stmt.Clauses["WHERE"]; ok {
		if where, ok := c.Expression.(clause.Where); ok && len(where.Exprs) > 0 {
			return true
		}
	}
	if stmt.Schema == nil || !stmt.ReflectValue.IsValid() {
		return false
	}
	_, qv := schema.GetIdentityFieldValuesMap(stmt.Context, stmt.ReflectValue, stmt.Schema.PrimaryFields)
	if _, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, qv); len(values) > 0 {
		return true
	}
	if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
		_, qv = schema.GetIdentityFieldValuesMap(stmt.Context, reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
		if _, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, qv); len(values) > 0 {
			return true
		}
	}
	return false
}

// stampUpdatedBy 更新时自动盖 updated_by(仅当列存在)。
// SkipHooks(UpdateColumn/UpdateColumns 等)不盖, 与 gorm 对 AutoUpdateTime 的处理一致,
// 避免"updated_by 变了而 updated_at 没变"的撕裂审计。
func stampUpdatedBy(db *gorm.DB) {
	stmt := db.Statement
	if stmt.Schema == nil || stmt.SkipHooks || isSystemTable(db) {
		return
	}
	id, ok := FromContext(stmt.Context)
	if !ok || id == nil || id.IsSystem || id.UserID == 0 {
		return
	}
	if !hasField(db, "updated_by") {
		return
	}
	// Model(值)+Updates(结构体值) 这类不可寻址形态下 SetColumn 会 AddError(ErrInvalidValue),
	// 把原本合法的更新打挂; 该形态跳过盖章(与未装引擎时行为一致)
	if _, isMap := stmt.Dest.(map[string]interface{}); !isMap {
		if _, isMaps := stmt.Dest.([]map[string]interface{}); !isMaps {
			if stmt.ReflectValue.Kind() == reflect.Struct && !stmt.ReflectValue.CanAddr() {
				return
			}
		}
	}
	// 显式 Omit("updated_by") 尊重调用方意图
	for _, o := range stmt.Omits {
		if o == "updated_by" {
			return
		}
	}
	// Select 白名单模式下把 updated_by 补进白名单, 否则会被 ConvertToAssignments
	// 静默剔除(gorm 只豁免 AutoUpdateTime 字段), 出现"updated_at 变了而 updated_by 没变"
	if len(stmt.Selects) > 0 {
		found := false
		for _, s := range stmt.Selects {
			if s == "updated_by" || s == "*" {
				found = true
				break
			}
		}
		if !found {
			stmt.Selects = append(stmt.Selects, "updated_by")
		}
	}
	stmt.SetColumn("updated_by", id.UserID, true)
}

var deletedAtType = reflect.TypeOf(gorm.DeletedAt{})

// softDeleteField 找到 schema 上的 gorm.DeletedAt 软删除字段(含指针形式), 没有则为 nil(硬删除表)
func softDeleteField(s *schema.Schema) *schema.Field {
	for _, f := range s.Fields {
		if f.FieldType == deletedAtType || f.IndirectFieldType == deletedAtType {
			return f
		}
	}
	return nil
}

// stampDeletedBy 软删除时把 deleted_by 并进同一条 UPDATE(仅当表同时有 deleted_by 列
// 与 gorm.DeletedAt 软删除字段; 硬删除的行已不存在, 无处可盖)。
//
// 无法走"追加 SET 子句"的路线: gorm v1.31.1 中 clause.Set 的 MergeClause 是整体覆盖
// 而非合并, 且 SoftDeleteDeleteClause.ModifyStatement 会在子句应用当场 Build 定稿 SQL,
// 事后无处插手。因此本回调按同一流程(gorm@v1.31.1 soft_delete.go)抢先构建 SQL,
// 在 SET 里多并一个 deleted_by; 原软删除子句看到 SQL 非空即整体跳过。
// 升级 gorm 时需比对 soft_delete.go 的 SoftDeleteDeleteClause.ModifyStatement。
func stampDeletedBy(db *gorm.DB) {
	stmt := db.Statement
	if db.Error != nil || stmt.Schema == nil || stmt.SQL.Len() != 0 || stmt.Unscoped || isSystemTable(db) {
		return
	}
	if !hasField(db, "deleted_by") {
		return
	}
	sdField := softDeleteField(stmt.Schema)
	if sdField == nil {
		return
	}
	// 带 ZEROVALUE 标签的软删除字段语义不同(存活行 = 零值时间而非 NULL),
	// 本回调按 IS NULL 口径构建会让删除静默失效 → 不接管, 交回原生流程(不盖章)
	if _, hasZero := sdField.TagSettings["ZEROVALUE"]; hasZero {
		return
	}
	id, ok := FromContext(stmt.Context)
	if !ok || id == nil || id.IsSystem || id.UserID == 0 {
		return // 无身份/系统上下文: 不盖章, 交回标准软删除流程
	}

	// 先收集主键条件(与 gorm 软删除同源的两处提取), 和既有 WHERE 一起判定"是否有删除条件";
	// 都没有则不接管, 让 gorm 正常走到 ErrMissingWhereClause 护栏
	// (AllowGlobalUpdate 属显式全量写, gorm 会放行执行, 此时必须接管以免 deleted_by 漏盖)
	var pkExprs []clause.Expression
	_, queryValues := schema.GetIdentityFieldValuesMap(stmt.Context, stmt.ReflectValue, stmt.Schema.PrimaryFields)
	column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
	if len(values) > 0 {
		pkExprs = append(pkExprs, clause.IN{Column: column, Values: values})
	}
	if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
		_, queryValues = schema.GetIdentityFieldValuesMap(stmt.Context, reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
		column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
		if len(values) > 0 {
			pkExprs = append(pkExprs, clause.IN{Column: column, Values: values})
		}
	}
	if _, hasWhere := stmt.Clauses["WHERE"]; !hasWhere && len(pkExprs) == 0 && !db.AllowGlobalUpdate {
		return
	}

	curTime := db.NowFunc()
	stmt.AddClause(clause.Set{
		{Column: clause.Column{Name: sdField.DBName}, Value: curTime},
		{Column: clause.Column{Name: "deleted_by"}, Value: id.UserID},
	})
	stmt.SetColumn(sdField.DBName, curTime, true)
	stmt.SetColumn("deleted_by", id.UserID, true)
	if len(pkExprs) > 0 {
		stmt.AddClause(clause.Where{Exprs: pkExprs})
	}
	// 与原软删除流程一致: 追加 deleted_at IS NULL 条件并标记 soft_delete_enabled
	// (GVA 不使用 ZEROVALUE 标签, ZeroValue 固定为无效值 → 渲染成 IS NULL)
	gorm.SoftDeleteQueryClause{ZeroValue: sql.NullString{Valid: false}, Field: sdField}.ModifyStatement(stmt)
	stmt.AddClauseIfNotExists(clause.Update{})
	stmt.Build(stmt.DB.Callback().Update().Clauses...)
}
