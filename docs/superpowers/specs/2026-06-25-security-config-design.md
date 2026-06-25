# 安全配置模块 设计文档

- 日期：2026-06-25
- 模块定位：系统本体（`server/.../system`）下的「安全配置」能力板块
- 状态：设计已确认，待写实现计划

## 1. 背景与目标

为系统本体提供一组可由管理员通过 UI 实时配置的安全能力，统一收敛到「安全配置」板块：

1. 验证码配置（防爆破阈值、超时、尺寸等）
2. 密码复杂度要求
3. 限流（登录/敏感接口）
4. 连续登录失败锁定（按账号）
5. 密码过期（如 90 天有效期，超期登录强制改密）

同时，借此机会把散落的缓存逻辑统一为一个「有 Redis 用 Redis、无 Redis 用内存」的通用缓存工具，彻底退役 `global.BlackCache`。

实现遵循项目既有分层 `Router -> API -> Service -> Model`、统一 `response`、`enter.go` 装配、`source/system` 种子机制，并需创建初始化的菜单与 API。

## 2. 范围（YAGNI）

- 只做上述 5 项能力 + 通用缓存工具。
- 限流仅作用于登录/敏感接口，不做全局限流。
- 不做 IP 黑白名单、2FA、额外审计（沿用既有登录日志）。

## 3. 总体架构

分两个阶段，写在同一份 spec，实现计划里按顺序落地：

- **Phase 0｜通用缓存工具**（地基）：新增统一 `Cache` 抽象 + 双后端，迁移所有现有缓存消费方，删除 `BlackCache`。
- **Phase 1｜安全配置功能**：新表 + Service + 5 项能力集成点 + API/Router/种子 + 前端页面。

## 4. Phase 0：通用缓存工具

### 4.1 位置与命名

- 包：`server/utils/gva_cache`
- 全局句柄：`global.GVA_CACHE`（类型为 `gva_cache.Cache` 接口）

### 4.2 接口

按现有消费方与限流需要定义（方法签名以实现时对齐 `local_cache` 现有 API 为准）：

```go
type Cache interface {
    Get(key string) (any, bool)
    Set(key string, value any, ttl time.Duration)
    SetDefault(key string, value any)            // 无过期 / 默认过期
    Increment(key string, n int64) (int64, error)
    IncrementWithExpire(key string, n int64, ttl time.Duration) (int64, error) // 限流用：首次设置过期
    Exists(key string) bool
    Delete(key string)
}
```

### 4.3 两套后端

- `memoryCache`：包装现有 `local_cache`（保持无 Redis 时的行为与语义）。
- `redisCache`：包装 `go-redis`。`Set`→`SET key val EX ttl`；`Increment`→`INCR`；`IncrementWithExpire`→`INCR` + 首次 `EXPIRE`；`Exists`→`EXISTS`；`SetDefault` 无 ttl。

### 4.4 初始化

- 在 Redis 初始化之后选择后端：`global.GVA_REDIS != nil` → `redisCache`，否则 `memoryCache`。
- 赋值给 `global.GVA_CACHE`，删除 `global.BlackCache` 及 `initialize/other.go` 中其初始化。

### 4.5 迁移消费方（全量）

- `middleware/jwt.go`：jwt 黑名单查询 → `GVA_CACHE`
- `service/system/jwt_black_list.go`：黑名单写入 → `GVA_CACHE`
- `api/v1/system/sys_captcha.go`、`sys_user.go`：验证码/登录计数 → `GVA_CACHE`
- `middleware/limit_ip.go`：限流改走 `GVA_CACHE`（使无 Redis 也能限流）

### 4.6 风险与降级

- jwt 黑名单语义保持与原内存实现一致（原为无过期 `SetDefault`，迁移后保持；Redis 下成为分布式，属增强）。
- 运行时 Redis 异常：`redisCache` 方法返回 error → 记录日志并 fail-open（限流/锁定计数失败时放行，避免误伤正常用户）。

## 5. Phase 1：安全配置功能

### 5.1 数据模型

新表 `system.SysSecurityConfig`（单行，固定 id=1），启动加载入内存缓存，保存即热更新：

| 分组 | 字段 | 说明 |
|---|---|---|
| 验证码 | `CaptchaOpen int` | 错误 N 次后出验证码，0=每次都要 |
| 验证码 | `CaptchaTimeout int` | 防爆破计数缓存超时（秒） |
| 验证码 | `KeyLong / ImgWidth / ImgHeight int` | 验证码长度/宽/高 |
| 密码复杂度 | `PwdMinLength int` | 最小长度 |
| 密码复杂度 | `PwdRequireUpper / Lower / Digit / Special bool` | 四类字符开关 |
| 限流 | `LimitEnable bool` | 是否开启 |
| 限流 | `LimitWindow int` | 窗口（秒） |
| 限流 | `LimitCount int` | 窗口内最大次数 |
| 失败锁定 | `LockEnable bool` | 是否开启 |
| 失败锁定 | `LockThreshold int` | 失败次数阈值 |
| 失败锁定 | `LockDuration int` | 锁定时长（分钟） |
| 密码过期 | `PwdExpireEnable bool` | 是否开启 |
| 密码过期 | `PwdExpireDays int` | 有效天数（如 90） |

SysUser 新增字段：`PasswordUpdatedAt *time.Time`（密码最后修改时间，AutoMigrate 加列）。

### 5.2 Service

`SecurityConfigService`：
- `Get()`：读取单行；不存在则按默认（含从 config.yaml captcha 取的初值）创建并返回。
- `Set(cfg)`：持久化 → 刷新内存缓存 → 若「密码过期」由关变开，则回填存量用户 `PasswordUpdatedAt = now`（仅对 NULL）。
- `Current()`：返回内存缓存的当前配置（被登录、中间件、密码校验器热读）。启动时和每次 `Set` 后刷新。

唯一数据源：验证码配置以本表为准；config.yaml 的 captcha 仅作首次初始化默认值。

### 5.3 各能力集成点

1. **验证码配置**：登录/验证码接口改读 `securityConfigService.Current()`（替换 `global.GVA_CONFIG.Captcha`），计数走 `GVA_CACHE`（按 IP）。
2. **密码复杂度**：新增 `utils.ValidatePasswordComplexity(pwd, cfg) error`，在「管理员创建用户 / 用户自助改密 / 重置密码」三处统一调用，返回可读错误。
3. **限流**：cache 化的 `limit_ip` 按配置（enable/window/count）挂到 `/base/login`、`/base/captcha`。
4. **连续登录失败锁定（按账号）**：
   - 失败时 `login_fail:<username>` 计数（TTL=`LockDuration` 分钟，作为滚动统计窗口，避免额外配置项）；达 `LockThreshold` 时置 `login_lock:<username>`（TTL=`LockDuration`）。
   - 登录入口先查锁定，命中即拒：「账号已锁定，请 X 分钟后再试」。
   - 成功登录清除计数与锁。
   - 与按 IP 的验证码防爆破独立并存。
5. **密码过期 + 强制改密**：
   - SysUser 加 `PasswordUpdatedAt`；新建/改密/重置时写入。
   - 登录凭证校验通过后：若 `PwdExpireEnable` 且 `now - PasswordUpdatedAt > PwdExpireDays`：
     - 签发带 `MustChangePwd=true` claim 的 JWT；登录响应附 `needChangePassword=true`。
     - 中间件在该 claim 下只放行 改密 / 用户信息 / 登出，其余一律 403。
     - 前端据 flag 强制跳转改密页；改密成功后 `PasswordUpdatedAt = now` 并重发正常 token。
   - 关闭时完全跳过该逻辑。

### 5.4 API / Router / 种子

- API：`api/v1/system/sys_security_config.go`
  - `GetSecurityConfig` → `GET /securityConfig/getSecurityConfig`
  - `SetSecurityConfig` → `POST /securityConfig/setSecurityConfig`
  - 私有组（`@Security ApiKeyAuth`），按权限（superAdmin）控制。
  - Swagger 注释按项目规范：`@Success` 落到具体类型（`response.Response{data=system.SysSecurityConfig,msg=string}`）。
- Router：`router/system/sys_security_config.go` + `router/system/enter.go` + `initialize/router.go` 注册。
- enter.go：api / service / router 三层装配。
- 迁移：`initialize/gorm.go` 的 `RegisterTables` 注册 `SysSecurityConfig`；SysUser 加列由 AutoMigrate 处理。
- 种子：
  - `source/system/api.go`：新增 2 条 API（ApiGroup「安全配置」）。
  - `source/system/menu.go`：新增「安全配置」菜单（挂系统下，superAdmin 可见），指向前端组件。
  - 默认配置单行种子：新增一个 `source/system` 初始化器，按 config.yaml captcha 默认值写入单行。

### 5.5 前端

- 页面：`web/src/view/system/security/`（5 个分区/Tab：验证码、密码复杂度、限流、失败锁定、密码过期）。
- API 封装：`web/src/api/securityConfig.js`。
- 登录流程：处理 `needChangePassword` → 强制跳转改密页；改密后重新登录/刷新 token。
- 样式优先 UnoCSS 原子类（遵循 `aiDoc/frontend-backend/frontend-rules.md`）。

## 6. 数据流

- 启动：加载 `SysSecurityConfig` → 内存缓存（`Current()`）。
- 管理员保存：`SetSecurityConfig` → 持久化 + 刷新内存缓存（+ 必要时回填 `PasswordUpdatedAt`）→ 立即生效（热更新，无需重启）。
- 登录：锁定检查 → 验证码检查（按配置，IP 计数）→ 凭证校验（失败则计数/锁定）→ 成功 → 过期检查 → 签发 token（必要时附 `MustChangePwd`）。

## 7. 错误处理 / 边界

- 缓存后端运行时异常：记录日志 + fail-open（限流/锁定放行）。
- 配置行缺失：Service 自动创建默认行。
- 密码过期关闭：完全跳过。
- 迁移顺序：`SysSecurityConfig` 建表 + 种子先于首次使用；`PasswordUpdatedAt` 加列。
- 存量用户：仅在「密码过期」开启时，对 `PasswordUpdatedAt` 为 NULL 的用户回填为当前时间。

## 8. 测试

- 单元：密码复杂度校验器（表驱动）、过期计算、锁定计数（基于内存 cache）、`Cache` 双后端接口、`SecurityConfigService` get/set/回填。
- 集成（httptest）：登录失败锁定流程、过期强制改密流程（带 `MustChangePwd` 中间件拦截）。

## 9. 已确认的关键决策

1. 配置存储：专用 DB 单行表 `SysSecurityConfig`。
2. 缓存：全量迁移、彻底退役 `BlackCache`；包 `utils/gva_cache`，句柄 `global.GVA_CACHE`。
3. 限流范围：登录/敏感接口 + 可配置。
4. 失败锁定维度：按账号。
5. 密码过期 UX：JWT claim + 中间件强拦 + 前端强制跳转改密。
6. 密码复杂度：可配最小长度 + 四类字符开关，多点生效。
7. 验证码配置：迁到新表作唯一来源，config.yaml 作首次默认值。
8. 密码过期存量用户：仅开启时回填为当前时间。
9. 路由前缀：`/securityConfig`。
10. 交付：同一份 spec、分两阶段实施。
