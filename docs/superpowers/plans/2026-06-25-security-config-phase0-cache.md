# Plan A — Phase 0：通用缓存工具 + 全量迁移（退役 BlackCache）

For agentic workers: REQUIRED SUB-SKILL: superpowers:subagent-driven-development

**Goal**：在 gin-vue-admin 后端新建统一缓存抽象 `server/utils/gva_cache`（内存后端包装 gkit `local_cache`、Redis 后端包装 go-redis），增加全局句柄 `global.GVA_CACHE`，在 Redis 初始化之后按「有 Redis 用 Redis、无 Redis 用内存」选择后端，并把所有现有缓存消费方（jwt 黑名单、验证码/登录计数、IP 限流）全量迁移到 `GVA_CACHE`，最终彻底删除 `global.BlackCache` 与其在 `initialize/other.go` 的初始化。限流改走 `GVA_CACHE`，使无 Redis 时也能限流。

**Architecture**：分层 `Router -> API -> Service -> Model`，缓存抽象落在 `server/utils/gva_cache`（基础设施工具包，不属于业务分层，被中间件 / API / Service 三层共享消费）。新增句柄 `global.GVA_CACHE gva_cache.Cache`。后端选择在 `core.RunServer()` 的 Redis 初始化块之后执行（关键时序：`initialize.OtherInit()` 在 `main.go:41` 早于 `core.RunServer()` 里的 `initialize.Redis()`（`core/server.go:16`），所以 cache 后端的赋值必须放在 RunServer 的 Redis 块之后，而不能放进 OtherInit）。

**Tech Stack**：Go 1.24.2；gin；gorm；`github.com/songzhibin97/gkit/cache/local_cache`（v1.2.13）；`github.com/redis/go-redis/v9`；module 前缀 `github.com/flipped-aurora/gin-vue-admin/server`。测试用标准库 `testing`。

## Global Constraints

逐行抄关键约束（每个 Task 隐含遵守）：

- 分层 Router -> API -> Service -> Model；Service 不依赖 `gin.Context`；通过 `service.ServiceGroupApp` / `api.ApiGroupApp` / 各 `enter.go` 装配。
- 统一响应用 `response` 包；统一分页用 `response.PageResult`（本 Plan A 不新增 API，故此约束在 Phase 1 体现，Phase 0 不破坏既有响应）。
- Swagger 注释完整且与真实行为一致：`@Success` 落到具体类型；私有组才写 `@Security ApiKeyAuth`，公开组不写。细则见 `aiDoc/modules/backend-layer-rules.md`。Phase 0 不改任何 Swagger 注释。
- 前端样式优先 UnoCSS 原子类（见 `aiDoc/frontend-backend/frontend-rules.md`）。Plan A 不涉及前端。
- 不读 `node_modules/` / `go pkg/mod` 业务源码（本计划已离线确认所需第三方签名，实现时无需再读）。
- module 路径前缀：`github.com/flipped-aurora/gin-vue-admin/server`。
- 跨计划钉死的缓存接口契约（Plan A 必须产出、Plan B 必须按此消费，不得改名改签名）：
  - 包 `server/utils/gva_cache`，全局 `global.GVA_CACHE`（类型 `gva_cache.Cache`）。
  - 接口：
    ```go
    type Cache interface {
        Get(key string) (any, bool)
        Set(key string, value any, ttl time.Duration)
        SetDefault(key string, value any)
        Increment(key string, n int64) (int64, error)
        IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error)
        Exists(key string) bool
        Delete(key string)
    }
    ```
  - 初始化：在 Redis 初始化之后，`global.GVA_REDIS != nil` 则用 redis 后端，否则内存后端。
- fail-open 原则：Redis 运行时异常时记录日志并放行（限流/计数失败不误伤正常用户）。
- DRY / YAGNI / TDD / 频繁提交：每改一处都先写失败测试（能写则写），再最小实现，再验证，再 commit。
- 历史报错忽略项：`plugin/announcement/gen/gen.go:5:20: misplaced compiler directive` 是仓库既有问题，与本计划无关；整体 `go build ./...` / `go vet ./...` 出现该包报错可忽略（已离线确认）。
- 第三方签名事实（已离线确认，实现按此为准）：
  - gkit `local_cache`：`Get(k string)(interface{},bool)`、`Set(k string,v interface{},d time.Duration)`、`SetDefault(k string,v interface{})`（覆盖、默认过期）、`Increment(k string,n int64) error`（**只返回 error**）、`IncrementInt64(k string,n int64)(int64,error)`（**返回新值**，要求 key 已存在且 value 为 `int64`，否则返回 `CacheNoExist` 等错误）、`Delete(k string)`、`NewCache(opts...)`、`SetDefaultExpire(d)`。
  - go-redis v9：`Set(ctx,key,val,ttl).Err()`、`Get(ctx,key).Result()`、`Incr(ctx,key).Result()(int64,error)`、`Expire(ctx,key,ttl).Err()`、`Exists(ctx,keys...).Result()(int64,error)`、`Del(ctx,keys...).Err()`、`TxPipeline()`。
  - 时序：`main.go:41 initialize.OtherInit()` → `core.RunServer()`（`core/server.go:15-20` Redis 初始化）。故 cache 后端赋值放在 `core/server.go` Redis 块之后。

---

### Task 1: 新建 gva_cache 包（接口 + 内存后端 + 单元测试）

**Files:**
- Create `server/utils/gva_cache/cache.go`（接口 `Cache` 定义）
- Create `server/utils/gva_cache/memory_cache.go`（`memoryCache`，包装 gkit `local_cache`）
- Create `server/utils/gva_cache/memory_cache_test.go`（内存后端单元测试 + 接口断言）

**Interfaces:**
- Produces `gva_cache.Cache`（见 Global Constraints 契约，逐字一致）。
- Produces `func NewMemoryCache(defaultExpire time.Duration) Cache`（消费 gkit `local_cache.NewCache(local_cache.SetDefaultExpire(defaultExpire))`）。
- `memoryCache.Increment` / `IncrementWithExpire`：当 key 不存在时，先 `Set(key, int64(0)+n, ttl|默认)` 建立 `int64` 计数，再返回新值；保证后续 `IncrementInt64` 不报 `CacheNoExist`，且存储类型恒为 `int64`。

步骤（TDD）：

- [ ] **Step 1: 写接口文件 `cache.go`（无逻辑，先定型契约）**
  创建 `server/utils/gva_cache/cache.go`：
  ```go
  package gva_cache

  import "time"

  // Cache 通用缓存抽象：有 Redis 用 Redis 后端，无 Redis 用内存后端。
  // 该接口为跨计划钉死契约，不得改名 / 改签名。
  type Cache interface {
      Get(key string) (any, bool)
      Set(key string, value any, ttl time.Duration)
      SetDefault(key string, value any)
      Increment(key string, n int64) (int64, error)
      IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error)
      Exists(key string) bool
      Delete(key string)
  }
  ```

- [ ] **Step 2: 写失败测试 `memory_cache_test.go`**
  创建 `server/utils/gva_cache/memory_cache_test.go`：
  ```go
  package gva_cache

  import (
      "testing"
      "time"
  )

  // 编译期保证 memoryCache 实现 Cache（接口断言）
  var _ Cache = (*memoryCache)(nil)

  func TestMemoryCache_SetGet(t *testing.T) {
      c := NewMemoryCache(time.Minute)
      c.Set("k1", 1, time.Minute)
      v, ok := c.Get("k1")
      if !ok {
          t.Fatalf("expected key k1 to exist")
      }
      if v.(int) != 1 {
          t.Fatalf("expected 1, got %v", v)
      }
      if _, ok := c.Get("missing"); ok {
          t.Fatalf("expected missing key to be absent")
      }
  }

  func TestMemoryCache_SetDefaultAndExists(t *testing.T) {
      c := NewMemoryCache(time.Minute)
      c.SetDefault("jwt-token", struct{}{})
      if !c.Exists("jwt-token") {
          t.Fatalf("expected jwt-token to exist after SetDefault")
      }
      c.Delete("jwt-token")
      if c.Exists("jwt-token") {
          t.Fatalf("expected jwt-token to be deleted")
      }
  }

  func TestMemoryCache_IncrementFromMissing(t *testing.T) {
      c := NewMemoryCache(time.Minute)
      // key 不存在时第一次 Increment 应建立计数并返回新值
      n, err := c.Increment("cnt", 1)
      if err != nil {
          t.Fatalf("unexpected err: %v", err)
      }
      if n != 1 {
          t.Fatalf("expected 1, got %d", n)
      }
      n, err = c.Increment("cnt", 2)
      if err != nil {
          t.Fatalf("unexpected err: %v", err)
      }
      if n != 3 {
          t.Fatalf("expected 3, got %d", n)
      }
  }

  func TestMemoryCache_IncrementWithExpire(t *testing.T) {
      c := NewMemoryCache(time.Minute)
      n, err := c.IncrementWithExpire("limit", 1, 50*time.Millisecond)
      if err != nil {
          t.Fatalf("unexpected err: %v", err)
      }
      if n != 1 {
          t.Fatalf("expected 1, got %d", n)
      }
      n, err = c.IncrementWithExpire("limit", 1, 50*time.Millisecond)
      if err != nil {
          t.Fatalf("unexpected err: %v", err)
      }
      if n != 2 {
          t.Fatalf("expected 2, got %d", n)
      }
      time.Sleep(120 * time.Millisecond)
      if c.Exists("limit") {
          t.Fatalf("expected limit key to expire")
      }
  }
  ```

- [ ] **Step 3: 跑测试看它失败（实现尚不存在）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/...`
  期望输出（编译失败，因 `memoryCache` / `NewMemoryCache` 未定义）：
  ```
  utils/gva_cache/memory_cache_test.go:9:... undefined: memoryCache
  ... undefined: NewMemoryCache
  FAIL	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache [build failed]
  ```

- [ ] **Step 4: 写最小实现 `memory_cache.go`**
  创建 `server/utils/gva_cache/memory_cache.go`：
  ```go
  package gva_cache

  import (
      "time"

      "github.com/songzhibin97/gkit/cache/local_cache"
  )

  // memoryCache 包装 gkit local_cache，保持无 Redis 时的语义。
  type memoryCache struct {
      c             local_cache.Cache
      defaultExpire time.Duration
  }

  // NewMemoryCache 创建内存后端，defaultExpire 为 SetDefault 的默认过期时间。
  func NewMemoryCache(defaultExpire time.Duration) Cache {
      return &memoryCache{
          c:             local_cache.NewCache(local_cache.SetDefaultExpire(defaultExpire)),
          defaultExpire: defaultExpire,
      }
  }

  func (m *memoryCache) Get(key string) (any, bool) {
      return m.c.Get(key)
  }

  func (m *memoryCache) Set(key string, value any, ttl time.Duration) {
      m.c.Set(key, value, ttl)
  }

  func (m *memoryCache) SetDefault(key string, value any) {
      m.c.SetDefault(key, value)
  }

  // increment 内部统一逻辑：key 不存在则按 int64(n) 建立计数（带 ttl），存在则 IncrementInt64。
  func (m *memoryCache) increment(key string, n int64, ttl time.Duration) (int64, error) {
      if _, ok := m.c.Get(key); !ok {
          m.c.Set(key, n, ttl)
          return n, nil
      }
      return m.c.IncrementInt64(key, n)
  }

  func (m *memoryCache) Increment(key string, n int64) (int64, error) {
      return m.increment(key, n, m.defaultExpire)
  }

  func (m *memoryCache) IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error) {
      return m.increment(key, n, ttl)
  }

  func (m *memoryCache) Exists(key string) bool {
      _, ok := m.c.Get(key)
      return ok
  }

  func (m *memoryCache) Delete(key string) {
      m.c.Delete(key)
  }
  ```

- [ ] **Step 5: 跑测试看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/...`
  期望输出：
  ```
  ok  	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache	0.xxxs
  ```

- [ ] **Step 6: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add utils/gva_cache/cache.go utils/gva_cache/memory_cache.go utils/gva_cache/memory_cache_test.go && git commit -m "feat(cache): 新增 gva_cache 通用缓存抽象与内存后端"
  ```

---

### Task 2: Redis 后端实现（包装 go-redis）

**Files:**
- Create `server/utils/gva_cache/redis_cache.go`（`redisCache`，包装 `redis.UniversalClient`）
- Create `server/utils/gva_cache/redis_cache_test.go`（接口断言 + 无 Redis 环境自动跳过的连通测试）

**Interfaces:**
- Produces `func NewRedisCache(client redis.UniversalClient) Cache`（Consumes `github.com/redis/go-redis/v9`）。
- 语义映射：`Set`→`SET key val EX ttl`；`SetDefault`→`SET key val`（无 ttl）；`Increment`→`INCR`（返回新值）；`IncrementWithExpire`→`INCR`，若结果为首次（`==n`）则 `EXPIRE`；`Exists`→`EXISTS`；`Delete`→`DEL`；`Get`→`GET`（不存在返回 `(nil,false)`）。

步骤（TDD）：

- [ ] **Step 1: 写失败测试 `redis_cache_test.go`**
  创建 `server/utils/gva_cache/redis_cache_test.go`：
  ```go
  package gva_cache

  import (
      "context"
      "testing"
      "time"

      "github.com/redis/go-redis/v9"
  )

  // 编译期保证 redisCache 实现 Cache（接口断言）
  var _ Cache = (*redisCache)(nil)

  // newTestRedis 尝试连本地 redis；连不上则跳过（CI/本地无 redis 时不阻塞）。
  func newTestRedis(t *testing.T) redis.UniversalClient {
      t.Helper()
      client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
      ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
      defer cancel()
      if err := client.Ping(ctx).Err(); err != nil {
          t.Skipf("跳过 redis 后端测试：无可用 redis (%v)", err)
      }
      return client
  }

  func TestRedisCache_RoundTrip(t *testing.T) {
      client := newTestRedis(t)
      c := NewRedisCache(client)
      key := "gva_cache_test:rt"
      c.Delete(key)

      c.Set(key, "v", time.Minute)
      v, ok := c.Get(key)
      if !ok || v.(string) != "v" {
          t.Fatalf("expected v, got %v ok=%v", v, ok)
      }
      if !c.Exists(key) {
          t.Fatalf("expected key to exist")
      }
      c.Delete(key)
      if c.Exists(key) {
          t.Fatalf("expected key deleted")
      }
  }

  func TestRedisCache_Increment(t *testing.T) {
      client := newTestRedis(t)
      c := NewRedisCache(client)
      key := "gva_cache_test:cnt"
      c.Delete(key)

      n, err := c.IncrementWithExpire(key, 1, time.Minute)
      if err != nil || n != 1 {
          t.Fatalf("expected 1, got %d err=%v", n, err)
      }
      n, err = c.Increment(key, 2)
      if err != nil || n != 3 {
          t.Fatalf("expected 3, got %d err=%v", n, err)
      }
      c.Delete(key)
  }
  ```

- [ ] **Step 2: 跑测试看它失败（实现尚不存在）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/...`
  期望输出（编译失败）：
  ```
  utils/gva_cache/redis_cache_test.go:... undefined: redisCache
  ... undefined: NewRedisCache
  FAIL	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache [build failed]
  ```

- [ ] **Step 3: 写最小实现 `redis_cache.go`**
  创建 `server/utils/gva_cache/redis_cache.go`：
  ```go
  package gva_cache

  import (
      "context"
      "time"

      "github.com/redis/go-redis/v9"
  )

  // redisCache 包装 go-redis，使缓存语义在多实例间共享。
  type redisCache struct {
      client redis.UniversalClient
  }

  // NewRedisCache 创建 Redis 后端。
  func NewRedisCache(client redis.UniversalClient) Cache {
      return &redisCache{client: client}
  }

  func (r *redisCache) Get(key string) (any, bool) {
      val, err := r.client.Get(context.Background(), key).Result()
      if err != nil {
          // redis.Nil（不存在）与其他错误统一按「未命中」处理
          return nil, false
      }
      return val, true
  }

  func (r *redisCache) Set(key string, value any, ttl time.Duration) {
      r.client.Set(context.Background(), key, value, ttl)
  }

  func (r *redisCache) SetDefault(key string, value any) {
      // 无过期
      r.client.Set(context.Background(), key, value, 0)
  }

  func (r *redisCache) Increment(key string, n int64) (int64, error) {
      return r.client.IncrBy(context.Background(), key, n).Result()
  }

  func (r *redisCache) IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error) {
      ctx := context.Background()
      v, err := r.client.IncrBy(ctx, key, n).Result()
      if err != nil {
          return 0, err
      }
      // 首次计数时设置过期（窗口起点）
      if v == n {
          if err := r.client.Expire(ctx, key, ttl).Err(); err != nil {
              return v, err
          }
      }
      return v, nil
  }

  func (r *redisCache) Exists(key string) bool {
      cnt, err := r.client.Exists(context.Background(), key).Result()
      if err != nil {
          return false
      }
      return cnt > 0
  }

  func (r *redisCache) Delete(key string) {
      r.client.Del(context.Background(), key)
  }
  ```
  说明：`Increment` 用 `IncrBy(key, n)` 以支持任意步长 `n`（`INCR` 仅 +1）；`IncrBy(key,1)` 等价于 `INCR`，符合契约。

- [ ] **Step 4: 跑测试看通过（无 redis 则自动 SKIP）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/...`
  期望输出（任一）：
  ```
  ok  	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache	0.xxxs
  ```
  或包含 redis 测试 `--- SKIP` 而整体 `ok`（运行 `go test -v ./utils/gva_cache/...` 可见 `--- SKIP: TestRedisCache_RoundTrip`）。

- [ ] **Step 5: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add utils/gva_cache/redis_cache.go utils/gva_cache/redis_cache_test.go && git commit -m "feat(cache): 新增 gva_cache Redis 后端（包装 go-redis）"
  ```

---

### Task 3: 全局句柄 GVA_CACHE + 初始化器（不删 BlackCache，先并存）

**Files:**
- Modify `server/global/global.go`（在 var 块加 `GVA_CACHE gva_cache.Cache`，行号当前 `40` 附近的 `BlackCache` 之后；import 加 `gva_cache` 包）
- Create `server/initialize/gva_cache.go`（`InitGvaCache()` 按是否有 Redis 选后端）
- Modify `server/core/server.go`（在 Redis 初始化块之后、`system.LoadAll()` 之前调用 `initialize.InitGvaCache()`，当前 Redis 块 `15-20` 行）

**Interfaces:**
- Produces `global.GVA_CACHE gva_cache.Cache`（被 Task 4-8 消费）。
- Produces `func initialize.InitGvaCache()`：`global.GVA_REDIS != nil` → `gva_cache.NewRedisCache(global.GVA_REDIS)`，否则 `gva_cache.NewMemoryCache(dr)`，其中 `dr` 由 `utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)` 得出（与原 `BlackCache` 默认过期一致）。

步骤：

- [ ] **Step 1: 给 global.go 增加句柄与 import**
  在 `server/global/global.go` 的 import 块加入（与现有 `local_cache` import 相邻）：
  ```go
  	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
  ```
  在 var 块 `BlackCache local_cache.Cache`（当前第 40 行）之后新增一行：
  ```go
  	BlackCache              local_cache.Cache
  	GVA_CACHE               gva_cache.Cache
  ```
  （此时 `BlackCache` 仍保留，逐步迁移后在 Task 9 删除。）

- [ ] **Step 2: 创建初始化器 `initialize/gva_cache.go`**
  创建 `server/initialize/gva_cache.go`：
  ```go
  package initialize

  import (
      "github.com/flipped-aurora/gin-vue-admin/server/global"
      "github.com/flipped-aurora/gin-vue-admin/server/utils"
      "github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
      "go.uber.org/zap"
  )

  // InitGvaCache 初始化通用缓存句柄 global.GVA_CACHE。
  // 必须在 Redis 初始化之后调用：有 Redis 用 Redis 后端，否则用内存后端。
  func InitGvaCache() {
      if global.GVA_REDIS != nil {
          global.GVA_CACHE = gva_cache.NewRedisCache(global.GVA_REDIS)
          global.GVA_LOG.Info("GVA_CACHE 使用 Redis 后端")
          return
      }
      dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
      if err != nil {
          // 与 OtherInit 中 BlackCache 一致：JWT 过期配置非法应在启动期暴露
          panic(err)
      }
      global.GVA_CACHE = gva_cache.NewMemoryCache(dr)
      global.GVA_LOG.Info("GVA_CACHE 使用内存后端")
  }
  ```

- [ ] **Step 3: 在 core/server.go 调用初始化器**
  修改 `server/core/server.go`，在 Redis 初始化块（`if global.GVA_CONFIG.System.UseRedis { ... }`，当前 15-20 行）之后、`if global.GVA_CONFIG.System.UseMongo` 之前插入：
  ```go
  	// 初始化通用缓存（必须在 Redis 之后：有 Redis 用 Redis，否则用内存）
  	initialize.InitGvaCache()
  ```

- [ ] **Step 4: 跑 build / vet 看通过（句柄已就绪，BlackCache 仍在）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./... 2>&1 | grep -v "plugin/announcement/gen"`
  期望输出：空（无 `plugin/announcement/gen` 之外的报错）。
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/...`
  期望输出：`ok  	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache`。

- [ ] **Step 5: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add global/global.go initialize/gva_cache.go core/server.go && git commit -m "feat(cache): 增加 global.GVA_CACHE 与初始化器（Redis 之后选后端）"
  ```

---

### Task 4: 迁移 jwt 黑名单查询（middleware/jwt.go）

**Files:**
- Modify `server/middleware/jwt.go`（`isBlacklist`，当前第 86-89 行）

**Interfaces:**
- Consumes `global.GVA_CACHE.Get(key string)(any,bool)`（替换 `global.BlackCache.Get`）。
- Produces 行为不变的 `func isBlacklist(jwt string) bool`。

步骤：

- [ ] **Step 1: 修改 `isBlacklist`**
  将 `server/middleware/jwt.go` 的：
  ```go
  func isBlacklist(jwt string) bool {
  	_, ok := global.BlackCache.Get(jwt)
  	return ok
  }
  ```
  改为：
  ```go
  func isBlacklist(jwt string) bool {
  	_, ok := global.GVA_CACHE.Get(jwt)
  	return ok
  }
  ```

- [ ] **Step 2: 跑 build 看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./middleware/... 2>&1`
  期望输出：空（无报错）。

- [ ] **Step 3: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add middleware/jwt.go && git commit -m "refactor(cache): jwt 黑名单查询迁移到 GVA_CACHE"
  ```

---

### Task 5: 迁移 jwt 黑名单写入（service/system/jwt_black_list.go）

**Files:**
- Modify `server/service/system/jwt_black_list.go`（`JsonInBlacklist` 第 27 行、`LoadAll` 第 50 行）

**Interfaces:**
- Consumes `global.GVA_CACHE.SetDefault(key string, value any)`（替换 `global.BlackCache.SetDefault`）。
- Produces 行为不变的 `JsonInBlacklist` / `LoadAll`（语义：无过期写入，迁移后保持；Redis 下成为分布式，属增强）。

步骤：

- [ ] **Step 1: 修改 `JsonInBlacklist`**
  将第 27 行：
  ```go
  	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
  ```
  改为：
  ```go
  	global.GVA_CACHE.SetDefault(jwtList.Jwt, struct{}{})
  ```

- [ ] **Step 2: 修改 `LoadAll` 循环体**
  将第 50 行：
  ```go
  		global.BlackCache.SetDefault(data[i], struct{}{})
  ```
  改为：
  ```go
  		global.GVA_CACHE.SetDefault(data[i], struct{}{})
  ```

- [ ] **Step 3: 跑 build 看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./service/... 2>&1`
  期望输出：空（无报错）。

- [ ] **Step 4: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add service/system/jwt_black_list.go && git commit -m "refactor(cache): jwt 黑名单写入迁移到 GVA_CACHE"
  ```

---

### Task 6: 迁移验证码计数（api/v1/system/sys_captcha.go）

**Files:**
- Modify `server/api/v1/system/sys_captcha.go`（`Captcha`，当前第 33-36 行）

**Interfaces:**
- Consumes `global.GVA_CACHE.Get` / `global.GVA_CACHE.Set`（替换 `global.BlackCache`）。
- 行为不变：首次按 IP `Set(key, int64(1), timeout)`，存在则不重置。

步骤：

- [ ] **Step 1: 修改计数读写**
  将 `server/api/v1/system/sys_captcha.go` 第 33-36 行：
  ```go
  	v, ok := global.BlackCache.Get(key)
  	if !ok {
  		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
  	}
  ```
  改为（注意存 `int64(1)` 以匹配 GVA_CACHE 的 `Increment` 计数类型，登录入口将对该 key 做 `Increment`）：
  ```go
  	v, ok := global.GVA_CACHE.Get(key)
  	if !ok {
  		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut))
  	}
  ```
  说明：原 `interfaceToInt(v)` 仅识别 `int`，下一步同步扩展为兼容 `int64`（见 Step 2），避免类型不匹配导致防爆逻辑失效。

- [ ] **Step 2: 扩展 interfaceToInt 兼容 int64**
  将 `server/api/v1/system/sys_captcha.go` 第 62-70 行 `interfaceToInt`：
  ```go
  func interfaceToInt(v interface{}) (i int) {
  	switch v := v.(type) {
  	case int:
  		i = v
  	default:
  		i = 0
  	}
  	return
  }
  ```
  改为：
  ```go
  func interfaceToInt(v interface{}) (i int) {
  	switch v := v.(type) {
  	case int:
  		i = v
  	case int64:
  		i = int(v)
  	case string:
  		// redis 后端 Get 返回字符串
  		if n, err := strconv.Atoi(v); err == nil {
  			i = n
  		}
  	default:
  		i = 0
  	}
  	return
  }
  ```
  并在文件顶部 import 块加入 `"strconv"`（当前 import 块为 `time` / global / response / systemRes / gin / base64Captcha / zap）：
  ```go
  import (
  	"strconv"
  	"time"

  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
  	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
  	"github.com/gin-gonic/gin"
  	"github.com/mojocn/base64Captcha"
  	"go.uber.org/zap"
  )
  ```
  说明：`interfaceToInt` 同时被 `sys_captcha.go` 与 `sys_user.go` 使用（同包），此处一并兼容内存(int64)与 Redis(string) 两种后端取值。

- [ ] **Step 2.5: 跑测试看现状（验证 interfaceToInt 兼容性）**
  （可选轻量验证）命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./api/... 2>&1`
  期望输出：空（无报错）。

- [ ] **Step 3: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add api/v1/system/sys_captcha.go && git commit -m "refactor(cache): 验证码计数迁移到 GVA_CACHE 并兼容多后端取值"
  ```

---

### Task 7: 迁移登录计数（api/v1/system/sys_user.go Login）

**Files:**
- Modify `server/api/v1/system/sys_user.go`（`Login`，计数读写第 44-47 行、三处 `Increment` 第 52/70/85 行）

**Interfaces:**
- Consumes `global.GVA_CACHE.Get` / `Set` / `Increment(key, n int64)(int64, error)`（替换 `global.BlackCache`）。
- 行为不变：登录失败按 IP `Increment(key, 1)`，触发验证码。

步骤：

- [ ] **Step 1: 修改计数读写（第 44-47 行）**
  将：
  ```go
  	v, ok := global.BlackCache.Get(key)
  	if !ok {
  		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
  	}
  ```
  改为：
  ```go
  	v, ok := global.GVA_CACHE.Get(key)
  	if !ok {
  		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut))
  	}
  ```

- [ ] **Step 2: 修改三处 Increment（第 52、70、85 行）**
  分别将三处：
  ```go
  		global.BlackCache.Increment(key, 1)
  ```
  改为（`Increment` 现返回 `(int64,error)`，按既有写法忽略返回值，保持行为不变）：
  ```go
  		_, _ = global.GVA_CACHE.Increment(key, 1)
  ```
  说明：使用 `replace_all` 或逐处替换，三处文本一致。

- [ ] **Step 3: 跑 build 看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./api/... 2>&1`
  期望输出：空（无报错）。

- [ ] **Step 4: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add api/v1/system/sys_user.go && git commit -m "refactor(cache): 登录防爆计数迁移到 GVA_CACHE"
  ```

---

### Task 8: 限流改走 GVA_CACHE（middleware/limit_ip.go），无 Redis 也能限流

**Files:**
- Modify `server/middleware/limit_ip.go`（`DefaultCheckOrMark` 第 44-53 行、`SetLimitWithTime` 第 64-92 行；删除 `context` import 依赖按需调整）
- Create `server/middleware/limit_ip_test.go`（基于内存后端 `gva_cache` 验证限流逻辑）

**Interfaces:**
- Consumes `global.GVA_CACHE.IncrementWithExpire(key, n int64, ttl)(int64,error)` 与 `Exists`。
- Produces 行为：窗口内计数到 `limit` 即拒绝；首次设置过期；`global.GVA_CACHE == nil`（极端启动期）时 fail-open。

步骤（TDD）：

- [ ] **Step 1: 写失败测试 `limit_ip_test.go`**
  创建 `server/middleware/limit_ip_test.go`：
  ```go
  package middleware

  import (
      "testing"
      "time"

      "github.com/flipped-aurora/gin-vue-admin/server/global"
      "github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
  )

  func TestDefaultCheckOrMark_LimitWithMemoryCache(t *testing.T) {
      // 注入内存后端，模拟无 Redis 也能限流
      global.GVA_CACHE = gva_cache.NewMemoryCache(time.Minute)
      key := "GVA_Limit_test_ip"
      global.GVA_CACHE.Delete(key)

      limit := 3
      // 前 3 次应放行
      for i := 1; i <= limit; i++ {
          if err := DefaultCheckOrMark(key, 60, limit); err != nil {
              t.Fatalf("第 %d 次不应被限流: %v", i, err)
          }
      }
      // 第 4 次应被限流
      if err := DefaultCheckOrMark(key, 60, limit); err == nil {
          t.Fatalf("第 4 次应被限流")
      }
      global.GVA_CACHE.Delete(key)
  }

  func TestDefaultCheckOrMark_FailOpenWhenCacheNil(t *testing.T) {
      global.GVA_CACHE = nil
      if err := DefaultCheckOrMark("any", 60, 1); err != nil {
          t.Fatalf("GVA_CACHE 为 nil 时应 fail-open 放行: %v", err)
      }
  }
  ```

- [ ] **Step 2: 跑测试看它失败（旧实现依赖 GVA_REDIS，nil 时直接 return nil，永不限流）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/... -run TestDefaultCheckOrMark`
  期望输出（`TestDefaultCheckOrMark_LimitWithMemoryCache` 失败：旧实现 `global.GVA_REDIS==nil` 直接放行，第 4 次未被限流）：
  ```
  --- FAIL: TestDefaultCheckOrMark_LimitWithMemoryCache
      limit_ip_test.go:..: 第 4 次应被限流
  FAIL
  ```

- [ ] **Step 3: 改写 `DefaultCheckOrMark` 与 `SetLimitWithTime`**
  将 `server/middleware/limit_ip.go` 第 44-92 行整体替换为：
  ```go
  func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
  	// 无缓存句柄（极端启动期）时 fail-open，避免误伤
  	if global.GVA_CACHE == nil {
  		return nil
  	}
  	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
  		global.GVA_LOG.Error("limit", zap.Error(err))
  	}
  	return err
  }

  func DefaultLimit() gin.HandlerFunc {
  	return LimitConfig{
  		GenerationKey: DefaultGenerationKey,
  		CheckOrMark:   DefaultCheckOrMark,
  		Expire:        global.GVA_CONFIG.System.LimitTimeIP,
  		Limit:         global.GVA_CONFIG.System.LimitCountIP,
  	}.LimitWithTime()
  }

  // SetLimitWithTime 设置访问次数：窗口内计数到达 limit 即拒绝。
  func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
  	count, err := global.GVA_CACHE.IncrementWithExpire(key, 1, expiration)
  	if err != nil {
  		// 运行时缓存异常：记录日志并 fail-open 放行
  		global.GVA_LOG.Error("limit increment", zap.Error(err))
  		return nil
  	}
  	if count > int64(limit) {
  		return errors.New("请求太过频繁，请稍后再试")
  	}
  	return nil
  }
  ```
  说明：原 `DefaultLimit` 保持不变（这里一并贴出以界定替换范围）。新逻辑统一走 `GVA_CACHE`，`IncrementWithExpire` 在首次计数时设置过期（窗口起点），`count > limit` 时拒绝（即前 `limit` 次放行，第 `limit+1` 次拒绝，与测试一致）。

- [ ] **Step 4: 清理 import（移除不再使用的 context）**
  将 `server/middleware/limit_ip.go` 顶部 import 块：
  ```go
  import (
  	"context"
  	"errors"
  	"net/http"
  	"time"

  	"go.uber.org/zap"

  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
  	"github.com/gin-gonic/gin"
  )
  ```
  改为（删除 `"context"`，其余保留——`errors`/`net/http`/`time`/`zap`/`global`/`response`/`gin` 仍在用）：
  ```go
  import (
  	"errors"
  	"net/http"
  	"time"

  	"go.uber.org/zap"

  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
  	"github.com/gin-gonic/gin"
  )
  ```

- [ ] **Step 5: 跑测试看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/... -run TestDefaultCheckOrMark`
  期望输出：
  ```
  ok  	github.com/flipped-aurora/gin-vue-admin/server/middleware	0.xxxs
  ```

- [ ] **Step 6: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add middleware/limit_ip.go middleware/limit_ip_test.go && git commit -m "refactor(cache): IP 限流改走 GVA_CACHE（无 Redis 也能限流）"
  ```

---

### Task 9: 删除 BlackCache 与 other.go 初始化（收口）

**Files:**
- Modify `server/global/global.go`（删除 `BlackCache local_cache.Cache` 行第 40 行；若 `local_cache` import 不再被引用则一并删除 import）
- Modify `server/initialize/other.go`（删除 `global.BlackCache = local_cache.NewCache(...)` 块第 23-25 行与 `local_cache` import）

**Interfaces:**
- 移除 `global.BlackCache`（不再有任何消费方，Task 4-8 已全部迁移）。
- `OtherInit()` 保留 `ParseDuration` 校验（仍是启动期 JWT 配置合法性校验）与 go.mod module 读取逻辑。

步骤：

- [ ] **Step 1: 确认无残留 BlackCache 引用**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && grep -rn "BlackCache" --include="*.go" .`
  期望输出：仅剩 `global/global.go:40` 与 `initialize/other.go`（即将删除的两处），无其他业务引用。

- [ ] **Step 2: 删除 global.go 中的 BlackCache 与未用 import**
  删除 `server/global/global.go` var 块中的：
  ```go
  	BlackCache              local_cache.Cache
  ```
  并删除 import 块中的：
  ```go
  	"github.com/songzhibin97/gkit/cache/local_cache"
  ```
  （`gva_cache` import 与 `GVA_CACHE` 句柄保留。）

- [ ] **Step 3: 删除 other.go 的 BlackCache 初始化与 import**
  将 `server/initialize/other.go` 改为：
  ```go
  package initialize

  import (
  	"bufio"
  	"os"
  	"strings"

  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/utils"
  )

  func OtherInit() {
  	_, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
  	if err != nil {
  		panic(err)
  	}
  	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
  	if err != nil {
  		panic(err)
  	}

  	file, err := os.Open("go.mod")
  	if err == nil && global.GVA_CONFIG.AutoCode.Module == "" {
  		scanner := bufio.NewScanner(file)
  		scanner.Scan()
  		global.GVA_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
  	}
  }
  ```
  说明：删除了 `local_cache` import、`dr` 变量与 `global.BlackCache = local_cache.NewCache(...)`，保留 `ExpiresTime`/`BufferTime` 的合法性校验（启动期 fail-fast）。`dr` 现仅在 `InitGvaCache()` 内存分支重新解析使用，职责清晰。

- [ ] **Step 4: 跑 build 看通过**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./... 2>&1 | grep -v "plugin/announcement/gen"`
  期望输出：空（无 `plugin/announcement/gen` 之外的报错）。

- [ ] **Step 5: commit**
  命令：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git add global/global.go initialize/other.go && git commit -m "refactor(cache): 删除 BlackCache 及其初始化，收口到 GVA_CACHE"
  ```

---

### Task 10: 整体验证（build + vet + 全量测试）

**Files:** 无新增/修改（验证收口）。

**Interfaces:** 无。

步骤：

- [ ] **Step 1: 整体 build**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go build ./... 2>&1 | grep -v "plugin/announcement/gen"`
  期望输出：空（无 `plugin/announcement/gen` 之外的报错）。
  注：`plugin/announcement/gen/gen.go:5:20: misplaced compiler directive` 为仓库既有问题，与本计划无关，忽略。

- [ ] **Step 2: 整体 vet**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go vet ./... 2>&1 | grep -v "plugin/announcement/gen"`
  期望输出：空（无 `plugin/announcement/gen` 之外的报错）。

- [ ] **Step 3: 全量测试（重点 gva_cache 与 middleware）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/gva_cache/... ./middleware/... 2>&1`
  期望输出：
  ```
  ok  	github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache	0.xxxs
  ok  	github.com/flipped-aurora/gin-vue-admin/server/middleware	0.xxxs
  ```

- [ ] **Step 4: 确认无 BlackCache 残留（最终核验）**
  命令：`cd /Users/jiangjizhao/gin-vue-admin/server && grep -rn "BlackCache" --include="*.go" .`
  期望输出：空（无任何匹配）。

- [ ] **Step 5: commit（若 Step 1-4 触发任何 gofmt 自动调整则提交，否则跳过）**
  命令（仅当有改动时）：
  ```
  cd /Users/jiangjizhao/gin-vue-admin/server && git status --porcelain
  ```
  若有改动：`go fmt ./utils/gva_cache/... ./middleware/...` 后 `git add -A && git commit -m "chore(cache): Phase 0 整体验证与格式化"`。

---

## Self-Review Coverage

逐条列出 spec（第 4 节 Phase 0）覆盖点 → 对应 Task：

- spec 4.1 包 `server/utils/gva_cache` + 句柄 `global.GVA_CACHE` → Task 1（建包/接口）、Task 3（句柄）。
- spec 4.2 `Cache` 接口（7 个方法，逐字契约）→ Task 1 Step 1（`cache.go`），并被 Task 3-8 消费；接口断言在 Task 1/2 测试中编译期校验。
- spec 4.3 `memoryCache` 包装 `local_cache`，保持无 Redis 语义 → Task 1（含 Increment 计数类型对齐 int64 的语义处理）。
- spec 4.3 `redisCache` 包装 go-redis（SET EX / INCR / EXPIRE / EXISTS / DEL / SetDefault 无 ttl）→ Task 2。
- spec 4.4 初始化：Redis 之后选后端（`GVA_REDIS!=nil`→redis，否则内存），赋值 `GVA_CACHE` → Task 3（`InitGvaCache` + `core/server.go` 时序：Redis 块之后调用）。
- spec 4.4 删除 `BlackCache` 及 `other.go` 初始化 → Task 9（先全量迁移再删除，分步安全）。
- spec 4.5 迁移 `middleware/jwt.go`（isBlacklist）→ Task 4。
- spec 4.5 迁移 `service/system/jwt_black_list.go`（JsonInBlacklist / LoadAll）→ Task 5。
- spec 4.5 迁移 `api/v1/system/sys_captcha.go`、`sys_user.go`（验证码/登录计数）→ Task 6、Task 7（含 `interfaceToInt` 兼容 int64/string 多后端取值）。
- spec 4.5 迁移 `middleware/limit_ip.go`（改走 GVA_CACHE，无 Redis 也能限流）→ Task 8（含 TDD 测试证明无 Redis 时限流生效）。
- spec 4.6 jwt 黑名单语义保持（原无过期 SetDefault）→ Task 5（沿用 `SetDefault`）。
- spec 4.6 运行时 Redis 异常 fail-open（限流/计数失败放行）→ Task 8（`SetLimitWithTime` 错误时记录日志并返回 nil；`DefaultCheckOrMark` nil 句柄放行）。
- 项目硬约束：分层 / Service 不依赖 gin.Context（gva_cache 为工具包，Service 仅消费句柄）→ 全 Task；不读 node_modules/pkg/mod 业务源码 → 已离线确认签名写入 Global Constraints；module 前缀 → 全 Task import 路径。
- 验证收口（build/vet/test、BlackCache 零残留、忽略 gen.go 历史报错）→ Task 10。
