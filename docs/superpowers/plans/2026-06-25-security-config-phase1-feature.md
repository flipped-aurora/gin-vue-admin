# 安全配置功能 Phase 1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在 gin-vue-admin 系统本体下交付一个由管理员通过 UI 实时配置的「安全配置」板块，覆盖验证码、密码复杂度、限流、按账号失败锁定、密码过期强制改密 5 项能力，全部消费 Plan A 产出的 `global.GVA_CACHE`。

**Architecture:** 单行配置表 `SysSecurityConfig`（固定 id=1）作为唯一数据源，启动加载入 `SecurityConfigService` 进程内缓存（`Current()` 热读），保存即热更新。各能力集成点（登录、验证码、限流中间件、改密中间件）读 `Current()`，计数/锁定/限流走 `global.GVA_CACHE`（有 Redis 用 Redis，否则内存，fail-open）。严格遵循 `Router -> API -> Service -> Model` 分层，通过 `enter.go` 装配，`source/system` 种子机制初始化 API/菜单/默认配置。

**Tech Stack:** Go 1.x + Gin + GORM + golang-jwt/v5；Plan A 的 `server/utils/gva_cache`（`global.GVA_CACHE`）；Vue3 + Element Plus + Pinia + UnoCSS（presetWind3）。

## Global Constraints

- module 路径前缀：`github.com/flipped-aurora/gin-vue-admin/server`。
- 分层 `Router -> API -> Service -> Model`；Service 不依赖 `gin.Context`；通过 `service.ServiceGroupApp` / `api.ApiGroupApp` / `enter.go` 装配。
- 统一响应走 `model/common/response` 包；统一分页走 `response.PageResult`。
- Swagger 注释完整且与真实行为一致：`@Success` 落到具体类型（详情用 `response.Response{data=Model,msg=string}`，列表用 `response.PageResult{list=[]Model}`）；私有组才写 `@Security ApiKeyAuth`，公开组不写。细则见 `aiDoc/modules/backend-layer-rules.md`。
- 前端样式优先 UnoCSS 原子类（见 `aiDoc/frontend-backend/frontend-rules.md`）；仅在原子类难表达时才写 `<style scoped>`；避免内联 `style`。
- 不读 `node_modules/` 与 go `pkg/mod` 源码。
- 【跨计划缓存契约，不得改名改签名】包 `server/utils/gva_cache`，全局 `global.GVA_CACHE`（类型 `gva_cache.Cache`），接口：
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
  本计划只消费 `global.GVA_CACHE`，假定 Plan A 已完成且 `global.GVA_CACHE` 在 Redis 初始化之后被赋值（有 Redis 用 redis 后端，否则内存后端）。
- 验证码配置以 `SysSecurityConfig` 表为唯一来源；`config.yaml` 的 `captcha` 仅作首次默认值。
- 缓存运行时异常 fail-open（限流/锁定计数失败时放行，避免误伤）。
- 密码过期关闭时完全跳过该逻辑。

---

### Task 1: 数据模型 SysSecurityConfig + SysUser.PasswordUpdatedAt + AutoMigrate

**Files:**
- Create: `server/model/system/sys_security_config.go`
- Modify: `server/model/system/sys_user.go:33`（在 `OriginSetting` 字段后追加 `PasswordUpdatedAt`）
- Modify: `server/initialize/gorm.go:63`（在 `system.SysParams{},` 之后注册新表）
- Test: `server/model/system/sys_security_config_test.go`

**Interfaces:**
- Produces:
  - `type system.SysSecurityConfig struct{...}`，`func (SysSecurityConfig) TableName() string`（返回 `"sys_security_config"`）。
  - `func (c SysSecurityConfig) CaptchaTimeoutDuration() time.Duration`。
  - `func DefaultSecurityConfig(captcha config.Captcha) SysSecurityConfig`（返回带默认值的配置，id 留 0 由调用方设为 1）。
  - `system.SysUser.PasswordUpdatedAt *time.Time`（json `passwordUpdatedAt`）。

- [ ] **Step 1: 写失败测试**

创建 `server/model/system/sys_security_config_test.go`：

```go
package system

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
)

func TestSysSecurityConfigTableName(t *testing.T) {
	if got := (SysSecurityConfig{}).TableName(); got != "sys_security_config" {
		t.Fatalf("TableName() = %q, want %q", got, "sys_security_config")
	}
}

func TestCaptchaTimeoutDuration(t *testing.T) {
	c := SysSecurityConfig{CaptchaTimeout: 3600}
	if got := c.CaptchaTimeoutDuration(); got != time.Hour {
		t.Fatalf("CaptchaTimeoutDuration() = %v, want %v", got, time.Hour)
	}
}

func TestDefaultSecurityConfig(t *testing.T) {
	cfg := DefaultSecurityConfig(config.Captcha{
		KeyLong:            6,
		ImgWidth:           240,
		ImgHeight:          80,
		OpenCaptcha:        0,
		OpenCaptchaTimeOut: 3600,
	})
	if cfg.KeyLong != 6 || cfg.ImgWidth != 240 || cfg.ImgHeight != 80 {
		t.Fatalf("captcha fields not copied: %+v", cfg)
	}
	if cfg.CaptchaOpen != 0 || cfg.CaptchaTimeout != 3600 {
		t.Fatalf("captcha open/timeout not copied: %+v", cfg)
	}
	if cfg.PwdMinLength != 8 {
		t.Fatalf("PwdMinLength default = %d, want 8", cfg.PwdMinLength)
	}
	if cfg.LockThreshold != 5 || cfg.LockDuration != 30 {
		t.Fatalf("lock defaults wrong: %+v", cfg)
	}
	if cfg.PwdExpireDays != 90 {
		t.Fatalf("PwdExpireDays default = %d, want 90", cfg.PwdExpireDays)
	}
}
```

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./model/system/ -run 'TestSysSecurityConfig|TestCaptchaTimeoutDuration|TestDefaultSecurityConfig' -v`
Expected: 编译失败，`undefined: SysSecurityConfig` / `undefined: DefaultSecurityConfig`。

- [ ] **Step 3: 写最小实现 — 新模型文件**

创建 `server/model/system/sys_security_config.go`：

```go
package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysSecurityConfig 安全配置 单行表 固定 id=1 启动加载入内存缓存 保存即热更新
type SysSecurityConfig struct {
	global.GVA_MODEL
	// 验证码
	CaptchaOpen    int `json:"captchaOpen" gorm:"default:0;comment:错误N次后出验证码 0=每次都要"`
	CaptchaTimeout int `json:"captchaTimeout" gorm:"default:3600;comment:防爆破计数缓存超时(秒)"`
	KeyLong        int `json:"keyLong" gorm:"default:6;comment:验证码长度"`
	ImgWidth       int `json:"imgWidth" gorm:"default:240;comment:验证码宽度"`
	ImgHeight      int `json:"imgHeight" gorm:"default:80;comment:验证码高度"`
	// 密码复杂度
	PwdMinLength    int  `json:"pwdMinLength" gorm:"default:8;comment:密码最小长度"`
	PwdRequireUpper bool `json:"pwdRequireUpper" gorm:"default:false;comment:需大写字母"`
	PwdRequireLower bool `json:"pwdRequireLower" gorm:"default:false;comment:需小写字母"`
	PwdRequireDigit bool `json:"pwdRequireDigit" gorm:"default:false;comment:需数字"`
	PwdRequireSpecial bool `json:"pwdRequireSpecial" gorm:"default:false;comment:需特殊字符"`
	// 限流
	LimitEnable bool `json:"limitEnable" gorm:"default:false;comment:是否开启限流"`
	LimitWindow int  `json:"limitWindow" gorm:"default:60;comment:限流窗口(秒)"`
	LimitCount  int  `json:"limitCount" gorm:"default:30;comment:窗口内最大次数"`
	// 失败锁定
	LockEnable    bool `json:"lockEnable" gorm:"default:false;comment:是否开启失败锁定"`
	LockThreshold int  `json:"lockThreshold" gorm:"default:5;comment:失败次数阈值"`
	LockDuration  int  `json:"lockDuration" gorm:"default:30;comment:锁定时长(分钟)"`
	// 密码过期
	PwdExpireEnable bool `json:"pwdExpireEnable" gorm:"default:false;comment:是否开启密码过期"`
	PwdExpireDays   int  `json:"pwdExpireDays" gorm:"default:90;comment:密码有效天数"`
}

func (SysSecurityConfig) TableName() string {
	return "sys_security_config"
}

// CaptchaTimeoutDuration 防爆破计数缓存超时
func (c SysSecurityConfig) CaptchaTimeoutDuration() time.Duration {
	return time.Duration(c.CaptchaTimeout) * time.Second
}

// LockDurationTimeout 锁定时长
func (c SysSecurityConfig) LockDurationTimeout() time.Duration {
	return time.Duration(c.LockDuration) * time.Minute
}

// LimitWindowDuration 限流窗口
func (c SysSecurityConfig) LimitWindowDuration() time.Duration {
	return time.Duration(c.LimitWindow) * time.Second
}

// DefaultSecurityConfig 由 config.yaml 的 captcha 生成默认单行配置 调用方负责设 id=1
func DefaultSecurityConfig(captcha config.Captcha) SysSecurityConfig {
	return SysSecurityConfig{
		CaptchaOpen:     captcha.OpenCaptcha,
		CaptchaTimeout:  captcha.OpenCaptchaTimeOut,
		KeyLong:         captcha.KeyLong,
		ImgWidth:        captcha.ImgWidth,
		ImgHeight:       captcha.ImgHeight,
		PwdMinLength:    8,
		PwdRequireUpper: false,
		PwdRequireLower: false,
		PwdRequireDigit: false,
		PwdRequireSpecial: false,
		LimitEnable:     false,
		LimitWindow:     60,
		LimitCount:      30,
		LockEnable:      false,
		LockThreshold:   5,
		LockDuration:    30,
		PwdExpireEnable: false,
		PwdExpireDays:   90,
	}
}
```

- [ ] **Step 4: 写最小实现 — SysUser 加列**

在 `server/model/system/sys_user.go` 的 `OriginSetting` 字段（第 33 行）之后追加（注意 import 需加 `"time"`）。先在 import 块加 `"time"`：

```go
import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/google/uuid"
)
```

再追加字段（替换 `OriginSetting` 行后插入一行）：

```go
	OriginSetting common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	PasswordUpdatedAt *time.Time `json:"passwordUpdatedAt" gorm:"comment:密码最后修改时间"`                                                        //密码最后修改时间
```

- [ ] **Step 5: 写最小实现 — AutoMigrate 注册**

在 `server/initialize/gorm.go` 第 63 行 `system.SysParams{},` 之后追加：

```go
		system.SysParams{},
		system.SysSecurityConfig{},
```

- [ ] **Step 6: 跑测试看通过 + 编译**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./model/system/ -run 'TestSysSecurityConfig|TestCaptchaTimeoutDuration|TestDefaultSecurityConfig' -v && go build ./...`
Expected: 三个测试 PASS；`go build ./...` 无错误。

- [ ] **Step 7: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/model/system/sys_security_config.go server/model/system/sys_security_config_test.go server/model/system/sys_user.go server/initialize/gorm.go && git commit -m "feat(security): add SysSecurityConfig model and SysUser.PasswordUpdatedAt"
```

---

### Task 2: 密码复杂度校验器 utils.ValidatePasswordComplexity

**Files:**
- Create: `server/utils/password_complexity.go`
- Test: `server/utils/password_complexity_test.go`

**Interfaces:**
- Consumes: `system.SysSecurityConfig`（Task 1）。
- Produces: `func ValidatePasswordComplexity(pwd string, cfg system.SysSecurityConfig) error`（不满足返回可读中文 error，满足返回 nil）。

- [ ] **Step 1: 写失败测试（表驱动）**

创建 `server/utils/password_complexity_test.go`：

```go
package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestValidatePasswordComplexity(t *testing.T) {
	base := system.SysSecurityConfig{PwdMinLength: 8}
	all := system.SysSecurityConfig{
		PwdMinLength:      8,
		PwdRequireUpper:   true,
		PwdRequireLower:   true,
		PwdRequireDigit:   true,
		PwdRequireSpecial: true,
	}
	cases := []struct {
		name    string
		pwd     string
		cfg     system.SysSecurityConfig
		wantErr bool
	}{
		{"too short", "Ab1!", base, true},
		{"min length ok", "abcdefgh", base, false},
		{"all rules ok", "Abcdef1!", all, false},
		{"missing upper", "abcdef1!", all, true},
		{"missing lower", "ABCDEF1!", all, true},
		{"missing digit", "Abcdefg!", all, true},
		{"missing special", "Abcdefg1", all, true},
		{"empty pwd fails min", "", base, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePasswordComplexity(tc.pwd, tc.cfg)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("expected nil, got %v", err)
			}
		})
	}
}
```

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/ -run TestValidatePasswordComplexity -v`
Expected: 编译失败，`undefined: ValidatePasswordComplexity`。

- [ ] **Step 3: 写最小实现**

创建 `server/utils/password_complexity.go`：

```go
package utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// ValidatePasswordComplexity 按安全配置校验密码复杂度 不满足返回可读错误
func ValidatePasswordComplexity(pwd string, cfg system.SysSecurityConfig) error {
	if cfg.PwdMinLength > 0 && len(pwd) < cfg.PwdMinLength {
		return fmt.Errorf("密码长度不能少于 %d 位", cfg.PwdMinLength)
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range pwd {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}
	var missing []string
	if cfg.PwdRequireUpper && !hasUpper {
		missing = append(missing, "大写字母")
	}
	if cfg.PwdRequireLower && !hasLower {
		missing = append(missing, "小写字母")
	}
	if cfg.PwdRequireDigit && !hasDigit {
		missing = append(missing, "数字")
	}
	if cfg.PwdRequireSpecial && !hasSpecial {
		missing = append(missing, "特殊字符")
	}
	if len(missing) > 0 {
		return fmt.Errorf("密码必须包含%s", strings.Join(missing, "、"))
	}
	return nil
}
```

- [ ] **Step 4: 跑测试看通过**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/ -run TestValidatePasswordComplexity -v`
Expected: 所有子用例 PASS。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/utils/password_complexity.go server/utils/password_complexity_test.go && git commit -m "feat(security): add ValidatePasswordComplexity validator"
```

---

### Task 3: SecurityConfigService（Get/Set/Current + 回填 + 启动加载）

**Files:**
- Create: `server/service/system/sys_security_config.go`
- Modify: `server/service/system/enter.go:19`（在 `SysParamsService` 后注册 `SecurityConfigService`）
- Test: `server/service/system/sys_security_config_test.go`

**Interfaces:**
- Consumes: `system.SysSecurityConfig`、`system.DefaultSecurityConfig`（Task 1）。
- Produces:
  - `type SecurityConfigService struct{}`
  - `func (s *SecurityConfigService) Get() (system.SysSecurityConfig, error)`（读单行，不存在则按默认创建）。
  - `func (s *SecurityConfigService) Set(cfg system.SysSecurityConfig) error`（持久化 → 刷新内存缓存 → 密码过期由关变开时回填存量 NULL 用户）。
  - `func (s *SecurityConfigService) Current() system.SysSecurityConfig`（返回内存缓存当前配置；未加载则惰性 `Get`）。
  - `func (s *SecurityConfigService) LoadAll()`（启动时调用，加载入内存）。
  - 进程内缓存通过包级 `var securityConfigCache atomic.Value`（存 `system.SysSecurityConfig`）实现，被 `Current()`/`Set()`/`LoadAll()` 读写。

- [ ] **Step 1: 写失败测试**

创建 `server/service/system/sys_security_config_test.go`：

```go
package system

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 仅测试纯内存缓存读写 不触碰 DB
func TestSecurityConfigCurrentReadsCache(t *testing.T) {
	s := &SecurityConfigService{}
	want := system.SysSecurityConfig{PwdMinLength: 12, LockThreshold: 3}
	setSecurityConfigCache(want)
	got := s.Current()
	if got.PwdMinLength != 12 || got.LockThreshold != 3 {
		t.Fatalf("Current() = %+v, want PwdMinLength=12 LockThreshold=3", got)
	}
}

func TestSetSecurityConfigCacheRoundTrip(t *testing.T) {
	c := system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 45}
	setSecurityConfigCache(c)
	got := getSecurityConfigCache()
	if !got.PwdExpireEnable || got.PwdExpireDays != 45 {
		t.Fatalf("cache round trip failed: %+v", got)
	}
}
```

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./service/system/ -run 'TestSecurityConfigCurrentReadsCache|TestSetSecurityConfigCacheRoundTrip' -v`
Expected: 编译失败，`undefined: SecurityConfigService` / `undefined: setSecurityConfigCache`。

- [ ] **Step 3: 写最小实现 — service**

创建 `server/service/system/sys_security_config.go`：

```go
package system

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type SecurityConfigService struct{}

// securityConfigCache 进程内当前生效配置 热读
var securityConfigCache atomic.Value

func setSecurityConfigCache(cfg system.SysSecurityConfig) {
	securityConfigCache.Store(cfg)
}

func getSecurityConfigCache() system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	return system.SysSecurityConfig{}
}

// Get 读取单行配置 不存在则按 config.yaml 默认创建并返回
func (s *SecurityConfigService) Get() (system.SysSecurityConfig, error) {
	var cfg system.SysSecurityConfig
	err := global.GVA_DB.Where("id = ?", 1).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cfg = system.DefaultSecurityConfig(global.GVA_CONFIG.Captcha)
		cfg.ID = 1
		if err = global.GVA_DB.Create(&cfg).Error; err != nil {
			return cfg, err
		}
		return cfg, nil
	}
	return cfg, err
}

// Set 持久化配置 刷新内存缓存 密码过期由关变开时回填存量 NULL 用户
func (s *SecurityConfigService) Set(cfg system.SysSecurityConfig) error {
	prev, err := s.Get()
	if err != nil {
		return err
	}
	cfg.ID = 1
	if err = global.GVA_DB.Save(&cfg).Error; err != nil {
		return err
	}
	setSecurityConfigCache(cfg)
	// 密码过期由关变开 回填存量 PasswordUpdatedAt 为 NULL 的用户
	if cfg.PwdExpireEnable && !prev.PwdExpireEnable {
		now := time.Now()
		if err = global.GVA_DB.Model(&system.SysUser{}).
			Where("password_updated_at IS NULL").
			Update("password_updated_at", now).Error; err != nil {
			return err
		}
	}
	return nil
}

// Current 返回内存缓存当前配置 未加载则惰性 Get
func (s *SecurityConfigService) Current() system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	cfg, err := s.Get()
	if err == nil {
		setSecurityConfigCache(cfg)
	}
	return cfg
}

// LoadAll 启动时加载配置入内存缓存
func (s *SecurityConfigService) LoadAll() {
	cfg, err := s.Get()
	if err != nil {
		global.GVA_LOG.Error("加载安全配置失败!")
		return
	}
	setSecurityConfigCache(cfg)
}
```

- [ ] **Step 4: 写最小实现 — enter 注册**

在 `server/service/system/enter.go` 第 19 行 `SysParamsService` 之后追加：

```go
	SysParamsService
	SecurityConfigService
```

- [ ] **Step 5: 跑测试看通过 + 编译**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./service/system/ -run 'TestSecurityConfigCurrentReadsCache|TestSetSecurityConfigCacheRoundTrip' -v && go build ./...`
Expected: 两个测试 PASS；`go build ./...` 无错误。

- [ ] **Step 6: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/service/system/sys_security_config.go server/service/system/sys_security_config_test.go server/service/system/enter.go && git commit -m "feat(security): add SecurityConfigService get/set/current with backfill"
```

---

### Task 4: 启动加载安全配置（OtherInit）

**Files:**
- Modify: `server/initialize/other.go:13-32`（`OtherInit` 末尾加载安全配置入内存）

**Interfaces:**
- Consumes: `service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.LoadAll()`（Task 3）。
- Produces: 启动后 `securityConfigService.Current()` 返回 DB 中持久化的配置（无行则自动创建默认行）。

- [ ] **Step 1: 写最小实现 — 在 OtherInit 末尾加载**

修改 `server/initialize/other.go`，import 增加 service 包，并在函数末尾加载（`BlackCache` 初始化由 Plan A 移除，这里只追加加载逻辑；若 Plan A 已删除 `local_cache` import，则不要再引用它）：

```go
package initialize

import (
	"bufio"
	"os"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
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

	// 加载安全配置入内存缓存
	service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.LoadAll()
}
```

> 注：若 Plan A 尚未移除 `global.BlackCache` 的初始化，保留原 `local_cache` 行不动，仅在函数末尾追加 `service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.LoadAll()` 与对应 `service` import。本步骤以 Plan A 已退役 `BlackCache` 为前提给出完整文件。

- [ ] **Step 2: 编译验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误。

- [ ] **Step 3: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/initialize/other.go && git commit -m "feat(security): load security config into memory cache on startup"
```

---

### Task 5: CustomClaims 加 MustChangePwd + LoginToken 支持

**Files:**
- Modify: `server/model/system/request/jwt.go:9-13`（`CustomClaims` 加字段）
- Modify: `server/utils/claims.go:137-148`（新增 `LoginTokenWithExpire`，保留原 `LoginToken`）
- Test: `server/utils/claims_security_test.go`

**Interfaces:**
- Consumes: `request.CustomClaims`、`system.Login`。
- Produces:
  - `request.CustomClaims.MustChangePwd bool`（json `mustChangePwd`）。
  - `func LoginTokenWithExpire(user system.Login, mustChangePwd bool) (token string, claims systemReq.CustomClaims, err error)`。

- [ ] **Step 1: 写失败测试**

创建 `server/utils/claims_security_test.go`：

```go
package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestLoginTokenWithExpireSetsMustChangePwd(t *testing.T) {
	global.GVA_CONFIG.JWT = config.JWT{
		SigningKey:  "test-signing-key",
		ExpiresTime: "7d",
		BufferTime:  "1d",
		Issuer:      "GVA",
	}
	user := &system.SysUser{Username: "tester"}
	token, claims, err := LoginTokenWithExpire(user, true)
	if err != nil {
		t.Fatalf("LoginTokenWithExpire err = %v", err)
	}
	if !claims.MustChangePwd {
		t.Fatalf("claims.MustChangePwd = false, want true")
	}
	parsed, err := NewJWT().ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken err = %v", err)
	}
	if !parsed.MustChangePwd {
		t.Fatalf("parsed.MustChangePwd = false, want true")
	}
}
```

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/ -run TestLoginTokenWithExpireSetsMustChangePwd -v`
Expected: 编译失败，`undefined: LoginTokenWithExpire` / `claims.MustChangePwd undefined`。

- [ ] **Step 3: 写最小实现 — CustomClaims 加字段**

修改 `server/model/system/request/jwt.go`：

```go
// CustomClaims structure
type CustomClaims struct {
	BaseClaims
	BufferTime    int64
	MustChangePwd bool `json:"mustChangePwd"`
	jwt.RegisteredClaims
}
```

- [ ] **Step 4: 写最小实现 — LoginTokenWithExpire**

在 `server/utils/claims.go` 末尾（`LoginToken` 之后）追加：

```go
// LoginTokenWithExpire 签发登录 token 可携带 MustChangePwd 强制改密标记
func LoginTokenWithExpire(user system.Login, mustChangePwd bool) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
	})
	claims.MustChangePwd = mustChangePwd
	token, err = j.CreateToken(claims)
	return
}
```

- [ ] **Step 5: 跑测试看通过**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/ -run TestLoginTokenWithExpireSetsMustChangePwd -v`
Expected: PASS。

- [ ] **Step 6: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/model/system/request/jwt.go server/utils/claims.go server/utils/claims_security_test.go && git commit -m "feat(security): add MustChangePwd claim and LoginTokenWithExpire"
```

---

### Task 6: 失败锁定与密码过期工具函数（纯函数 + cache 计数）

**Files:**
- Create: `server/service/system/security_lock.go`
- Test: `server/service/system/security_lock_test.go`

**Interfaces:**
- Consumes: `global.GVA_CACHE`（Plan A）、`system.SysSecurityConfig`（Task 1）。
- Produces（包级函数，供登录入口调用）：
  - `func loginFailKey(username string) string`（返回 `"login_fail:" + username`）。
  - `func loginLockKey(username string) string`（返回 `"login_lock:" + username`）。
  - `func IsAccountLocked(username string) bool`（查 `global.GVA_CACHE.Exists(loginLockKey)`，cache 异常时 fail-open 由 Exists 自身保证返回 false）。
  - `func RecordLoginFail(username string, cfg system.SysSecurityConfig)`（计数 +1，达阈值置锁；cache 报错时记录日志并放弃，不 panic）。
  - `func ClearLoginFail(username string)`（删除计数与锁）。
  - `func IsPasswordExpired(passwordUpdatedAt *time.Time, cfg system.SysSecurityConfig, now time.Time) bool`（纯函数：未开启或时间为 nil 返回 false；`now - passwordUpdatedAt > PwdExpireDays*24h` 返回 true）。

- [ ] **Step 1: 写失败测试（纯函数 + 内存 cache 计数）**

创建 `server/service/system/security_lock_test.go`：

```go
package system

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

func TestIsPasswordExpired(t *testing.T) {
	now := time.Date(2026, 6, 25, 0, 0, 0, 0, time.UTC)
	cfg := system.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 90}

	old := now.AddDate(0, 0, -100)
	recent := now.AddDate(0, 0, -10)

	if !IsPasswordExpired(&old, cfg, now) {
		t.Fatalf("expected expired for 100 days old")
	}
	if IsPasswordExpired(&recent, cfg, now) {
		t.Fatalf("expected not expired for 10 days old")
	}
	if IsPasswordExpired(nil, cfg, now) {
		t.Fatalf("nil time should be not expired")
	}
	off := system.SysSecurityConfig{PwdExpireEnable: false, PwdExpireDays: 90}
	if IsPasswordExpired(&old, off, now) {
		t.Fatalf("disabled config should be not expired")
	}
}

func TestLockCounting(t *testing.T) {
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Hour)
	cfg := system.SysSecurityConfig{LockEnable: true, LockThreshold: 3, LockDuration: 30}
	user := "locktester"

	ClearLoginFail(user)
	if IsAccountLocked(user) {
		t.Fatalf("should not be locked initially")
	}
	RecordLoginFail(user, cfg)
	RecordLoginFail(user, cfg)
	if IsAccountLocked(user) {
		t.Fatalf("should not be locked at 2 fails (threshold 3)")
	}
	RecordLoginFail(user, cfg)
	if !IsAccountLocked(user) {
		t.Fatalf("should be locked at 3 fails")
	}
	ClearLoginFail(user)
	if IsAccountLocked(user) {
		t.Fatalf("clear should remove lock")
	}
}
```

> 注：测试使用 Plan A 约定的内存后端构造器 `gva_cache.NewMemoryCache(defaultExpire time.Duration) gva_cache.Cache`。若 Plan A 的构造器命名不同，按 Plan A 实际导出名替换（接口方法签名已钉死，构造器名以 Plan A 为准）。

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./service/system/ -run 'TestIsPasswordExpired|TestLockCounting' -v`
Expected: 编译失败，`undefined: IsPasswordExpired` / `undefined: RecordLoginFail`。

- [ ] **Step 3: 写最小实现**

创建 `server/service/system/security_lock.go`：

```go
package system

import (
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func loginFailKey(username string) string {
	return "login_fail:" + username
}

func loginLockKey(username string) string {
	return "login_lock:" + username
}

// IsAccountLocked 查询账号是否处于锁定状态
func IsAccountLocked(username string) bool {
	return global.GVA_CACHE.Exists(loginLockKey(username))
}

// RecordLoginFail 记录一次登录失败 计数滚动窗口 TTL=LockDuration 达阈值置锁
func RecordLoginFail(username string, cfg system.SysSecurityConfig) {
	if !cfg.LockEnable {
		return
	}
	n, err := global.GVA_CACHE.IncrementWithExpire(loginFailKey(username), 1, cfg.LockDurationTimeout())
	if err != nil {
		global.GVA_LOG.Error("登录失败计数失败", zap.Error(err))
		return
	}
	if int(n) >= cfg.LockThreshold {
		global.GVA_CACHE.Set(loginLockKey(username), 1, cfg.LockDurationTimeout())
	}
}

// ClearLoginFail 清除失败计数与锁 登录成功调用
func ClearLoginFail(username string) {
	global.GVA_CACHE.Delete(loginFailKey(username))
	global.GVA_CACHE.Delete(loginLockKey(username))
}

// IsPasswordExpired 纯函数 判定密码是否过期
func IsPasswordExpired(passwordUpdatedAt *time.Time, cfg system.SysSecurityConfig, now time.Time) bool {
	if !cfg.PwdExpireEnable || cfg.PwdExpireDays <= 0 {
		return false
	}
	if passwordUpdatedAt == nil {
		return false
	}
	deadline := passwordUpdatedAt.AddDate(0, 0, cfg.PwdExpireDays)
	return now.After(deadline)
}
```

- [ ] **Step 4: 跑测试看通过**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./service/system/ -run 'TestIsPasswordExpired|TestLockCounting' -v`
Expected: 两个测试 PASS。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/service/system/security_lock.go server/service/system/security_lock_test.go && git commit -m "feat(security): add account lock counting and password expiry helpers"
```

---

### Task 7: 验证码接口接入安全配置（替换 GVA_CONFIG.Captcha / GVA_CACHE 计数）

**Files:**
- Modify: `server/api/v1/system/sys_captcha.go:28-59`（`Captcha` 改读 `securityConfigService.Current()`，计数走 `global.GVA_CACHE`）
- Modify: `server/api/v1/system/enter.go:31-59`（注册 `securityConfigService` 句柄）

**Interfaces:**
- Consumes: `securityConfigService.Current()`（Task 3）、`global.GVA_CACHE`（Plan A）。
- Produces: `Captcha` handler 使用安全配置生成验证码与防爆破计数。

- [ ] **Step 1: 注册 service 句柄**

在 `server/api/v1/system/enter.go` 的 var 块（约第 58 行 `skillsService` 之后）追加：

```go
	skillsService            = service.ServiceGroupApp.SystemServiceGroup.SkillsService
	securityConfigService    = service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService
```

- [ ] **Step 2: 写最小实现 — 重写 Captcha handler**

将 `server/api/v1/system/sys_captcha.go` 的 `Captcha` 函数体改为：

```go
// Captcha
// @Tags      Base
// @Summary   生成验证码
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	cfg := securityConfigService.Current()
	// 判断验证码是否开启
	openCaptcha := cfg.CaptchaOpen               // 错误 N 次后出验证码
	openCaptchaTimeOut := cfg.CaptchaTimeout      // 计数缓存超时
	key := c.ClientIP()
	v, ok := global.GVA_CACHE.Get(key)
	if !ok {
		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut)) // int64 以匹配 GVA_CACHE.Increment 计数类型(内存后端 IncrementInt64)
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 生成默认数字 driver
	driver := base64Captcha.NewDriverDigit(cfg.ImgHeight, cfg.ImgWidth, cfg.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: cfg.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}
```

> 注：`/base/captcha` 原 Swagger 含 `@Security ApiKeyAuth`，但该接口在公开组注册，按 Global Constraints 公开组不写 `@Security`，本步骤已移除该行修正历史不一致。

- [ ] **Step 3: 编译验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误（`interfaceToInt` 仍在本文件定义，`time`/`zap`/`base64Captcha` import 不变）。

- [ ] **Step 4: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/api/v1/system/sys_captcha.go server/api/v1/system/enter.go && git commit -m "feat(security): captcha api reads security config and uses GVA_CACHE"
```

---

### Task 8: 登录接口集成（锁定检查 + 验证码 + 失败计数 + 过期强制改密）

**Files:**
- Modify: `server/api/v1/system/sys_user.go:27-161`（重写 `Login` 与 `TokenNext`）
- Modify: `server/model/system/response/sys_user.go:11-15`（`LoginResponse` 加 `NeedChangePassword`）

**Interfaces:**
- Consumes: `securityConfigService.Current()`、`IsAccountLocked`、`RecordLoginFail`、`ClearLoginFail`、`IsPasswordExpired`（Task 3/6）、`utils.LoginTokenWithExpire`（Task 5）、`global.GVA_CACHE`（Plan A）。
- Produces:
  - `systemRes.LoginResponse.NeedChangePassword bool`（json `needChangePassword`）。
  - `Login` handler：锁定检查 → 验证码检查（IP 计数）→ 凭证校验（失败计数/锁定）→ 成功清零 → 过期检查 → 调 `TokenNext(c, user, needChange)`。
  - `func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser, mustChangePwd bool)`（签名新增 `mustChangePwd`）。

- [ ] **Step 1: LoginResponse 加字段**

修改 `server/model/system/response/sys_user.go`：

```go
type LoginResponse struct {
	User               system.SysUser `json:"user"`
	Token              string         `json:"token"`
	ExpiresAt          int64          `json:"expiresAt"`
	NeedChangePassword bool           `json:"needChangePassword"`
}
```

- [ ] **Step 2: 写最小实现 — 重写 Login**

将 `server/api/v1/system/sys_user.go` 的 `Login` 函数体替换为：

```go
// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	cfg := securityConfigService.Current()

	// 1. 账号锁定检查
	if cfg.LockEnable && system.IsAccountLocked(l.Username) {
		response.FailWithMessage("账号已锁定，请 "+strconv.Itoa(cfg.LockDuration)+" 分钟后再试", c)
		loginLogService.CreateLoginLog(system.SysLoginLog{
			Username: l.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
			Status: false, ErrorMessage: "账号已锁定",
		})
		return
	}

	// 2. 验证码检查（按 IP 计数）
	key := c.ClientIP()
	openCaptcha := cfg.CaptchaOpen
	openCaptchaTimeOut := cfg.CaptchaTimeout
	v, ok := global.GVA_CACHE.Get(key)
	if !ok {
		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut)) // int64 以匹配 GVA_CACHE.Increment 计数类型(内存后端 IncrementInt64)
	}
	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc && (l.Captcha == "" || l.CaptchaId == "" || !store.Verify(l.CaptchaId, l.Captcha, true)) {
		global.GVA_CACHE.Increment(key, 1)
		response.FailWithMessage("验证码错误", c)
		loginLogService.CreateLoginLog(system.SysLoginLog{
			Username: l.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
			Status: false, ErrorMessage: "验证码错误",
		})
		return
	}

	// 3. 凭证校验
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		global.GVA_CACHE.Increment(key, 1)
		system.RecordLoginFail(l.Username, cfg)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		loginLogService.CreateLoginLog(system.SysLoginLog{
			Username: l.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
			Status: false, ErrorMessage: "用户名不存在或者密码错误",
		})
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		global.GVA_CACHE.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		loginLogService.CreateLoginLog(system.SysLoginLog{
			Username: l.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
			Status: false, ErrorMessage: "用户被禁止登录", UserID: user.ID,
		})
		return
	}

	// 4. 登录成功 清除失败计数与锁
	system.ClearLoginFail(l.Username)

	// 5. 密码过期检查
	needChange := system.IsPasswordExpired(user.PasswordUpdatedAt, cfg, time.Now())
	b.TokenNext(c, *user, needChange)
}
```

- [ ] **Step 3: 写最小实现 — 重写 TokenNext（签名加 mustChangePwd）**

将 `server/api/v1/system/sys_user.go` 的 `TokenNext` 替换为：

```go
// TokenNext 登录以后签发 jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser, mustChangePwd bool) {
	token, claims, err := utils.LoginTokenWithExpire(&user, mustChangePwd)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 记录登录成功日志
	loginLogService.CreateLoginLog(system.SysLoginLog{
		Username: user.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
		Status: true, UserID: user.ID, ErrorMessage: "登录成功",
	})
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:               user,
			Token:              token,
			ExpiresAt:          claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			NeedChangePassword: mustChangePwd,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := utils.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User: user, Token: token,
			ExpiresAt:          claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			NeedChangePassword: mustChangePwd,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := utils.SetRedisJWT(token, user.GetUsername()); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User: user, Token: token,
			ExpiresAt:          claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			NeedChangePassword: mustChangePwd,
		}, "登录成功", c)
	}
}
```

- [ ] **Step 4: 编译验证（其它 TokenNext 调用点）**

`TokenNext` 仅在本文件 `Login` 内被调用，签名变更已同步。验证编译。

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/api/v1/system/sys_user.go server/model/system/response/sys_user.go && git commit -m "feat(security): login integrates lock/captcha/fail-count/password-expiry"
```

---

### Task 9: 密码复杂度三处接入 + 改密回填 PasswordUpdatedAt

**Files:**
- Modify: `server/api/v1/system/sys_user.go`（`Register` 第 170 行附近、`ChangePassword` 第 206 行附近、`ResetPassword` 第 502 行附近三处加复杂度校验）
- Modify: `server/service/system/sys_user.go:28-38`（`Register` 写入 `PasswordUpdatedAt`）、`:69-81`（`ChangePassword` 改密回填）、`:333-336`（`ResetPassword` 回填）
- Test: 复用 Task 2 校验器测试 + 本任务编译验证

**Interfaces:**
- Consumes: `utils.ValidatePasswordComplexity`（Task 2）、`securityConfigService.Current()`（Task 3）。
- Produces: 三处密码入口均校验复杂度；`Register`/`ChangePassword`/`ResetPassword` 均写入 `PasswordUpdatedAt=now`。

- [ ] **Step 1: 写最小实现 — service Register 写时间**

将 `server/service/system/sys_user.go` 的 `Register`（第 28-38 行）改为：

```go
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	now := time.Now()
	u.PasswordUpdatedAt = &now
	err = global.GVA_DB.Create(&u).Error
	return u, err
}
```

在该文件 import 块加入 `"time"`（若尚无）。

- [ ] **Step 2: 写最小实现 — service ChangePassword 回填**

将 `ChangePassword`（第 69-81 行）改为：

```go
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error) {
	var user system.SysUser
	err = global.GVA_DB.Select("id, password").Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("原密码错误")
	}
	pwd := utils.BcryptHash(newPassword)
	now := time.Now()
	err = global.GVA_DB.Model(&user).Updates(map[string]interface{}{
		"password":            pwd,
		"password_updated_at": now,
	}).Error
	return err
}
```

- [ ] **Step 3: 写最小实现 — service ResetPassword 回填**

将 `ResetPassword`（第 333-336 行）改为：

```go
func (userService *UserService) ResetPassword(ID uint, password string) (err error) {
	now := time.Now()
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"password":            utils.BcryptHash(password),
		"password_updated_at": now,
	}).Error
	return err
}
```

- [ ] **Step 4: 写最小实现 — API Register 加复杂度校验**

在 `server/api/v1/system/sys_user.go` 的 `Register` 中，`utils.Verify(r, utils.RegisterVerify)` 校验通过后、构造 `authorities` 之前插入：

```go
	if err = utils.ValidatePasswordComplexity(r.Password, securityConfigService.Current()); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
```

- [ ] **Step 5: 写最小实现 — API ChangePassword 加复杂度校验**

在 `ChangePassword` 中，`utils.Verify(req, utils.ChangePasswordVerify)` 通过后、取 `uid` 之前插入：

```go
	if err = utils.ValidatePasswordComplexity(req.NewPassword, securityConfigService.Current()); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
```

- [ ] **Step 6: 写最小实现 — API ResetPassword 加复杂度校验**

读取 `ResetPassword`（第 502 行起）确认其入参字段名，然后在调用 `userService.ResetPassword` 之前插入复杂度校验。先确认字段：

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && sed -n '502,540p' api/v1/system/sys_user.go`
Expected: 看到 handler 内构造 `system.SysUser` 与默认新密码（项目默认重置为 `123456` 之类）。在 `userService.ResetPassword(...)` 调用前插入（用实际传入的密码变量名，下例假定为默认常量 `"123456"` 时跳过校验，否则校验传入值）：

```go
	// 仅当重置为非空自定义密码时校验复杂度；默认重置密码沿用项目原行为
	if user.Password != "" {
		if err := utils.ValidatePasswordComplexity(user.Password, securityConfigService.Current()); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
```

> 注：`ResetPassword` 的真实变量名以第 502-540 行实际代码为准；若其重置为固定默认密码且不接收用户输入，则本步骤仅保留对自定义密码分支的校验，不改变默认重置行为。

- [ ] **Step 7: 跑校验器测试 + 编译**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./utils/ -run TestValidatePasswordComplexity -v && go build ./...`
Expected: 校验器测试 PASS；`go build ./...` 无错误。

- [ ] **Step 8: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/api/v1/system/sys_user.go server/service/system/sys_user.go && git commit -m "feat(security): enforce password complexity and record PasswordUpdatedAt on register/change/reset"
```

---

### Task 10: 强制改密中间件 MustChangePwdGuard

**Files:**
- Create: `server/middleware/must_change_pwd.go`
- Modify: `server/initialize/router.go:68`（在 `JWTAuth()` 之后挂 `MustChangePwdGuard()`）
- Test: `server/middleware/must_change_pwd_test.go`

**Interfaces:**
- Consumes: `c.Get("claims")` → `*systemReq.CustomClaims`（`JWTAuth` 设置）、`request.CustomClaims.MustChangePwd`（Task 5）。
- Produces: `func MustChangePwdGuard() gin.HandlerFunc`（`MustChangePwd=true` 时仅放行 改密/用户信息/登出，其余返回 403）。

- [ ] **Step 1: 写失败测试**

创建 `server/middleware/must_change_pwd_test.go`：

```go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

func setupRouterWithClaims(mustChange bool, path string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &systemReq.CustomClaims{MustChangePwd: mustChange})
		c.Next()
	})
	r.Use(MustChangePwdGuard())
	r.Any("/*any", func(c *gin.Context) { c.Status(http.StatusOK) })
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, path, nil)
	r.ServeHTTP(w, req)
	return w
}

func TestMustChangePwdGuard(t *testing.T) {
	// 未强制改密 任意路径放行
	if w := setupRouterWithClaims(false, "/user/getUserList"); w.Code != http.StatusOK {
		t.Fatalf("not-must-change should pass, got %d", w.Code)
	}
	// 强制改密 改密路径放行
	if w := setupRouterWithClaims(true, "/user/changePassword"); w.Code != http.StatusOK {
		t.Fatalf("changePassword should pass under must-change, got %d", w.Code)
	}
	// 强制改密 用户信息放行
	if w := setupRouterWithClaims(true, "/user/getUserInfo"); w.Code != http.StatusOK {
		t.Fatalf("getUserInfo should pass under must-change, got %d", w.Code)
	}
	// 强制改密 其它路径 403
	if w := setupRouterWithClaims(true, "/user/getUserList"); w.Code != http.StatusForbidden {
		t.Fatalf("other path should be 403 under must-change, got %d", w.Code)
	}
}
```

- [ ] **Step 2: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/ -run TestMustChangePwdGuard -v`
Expected: 编译失败，`undefined: MustChangePwdGuard`。

- [ ] **Step 3: 写最小实现**

创建 `server/middleware/must_change_pwd.go`：

```go
package middleware

import (
	"net/http"
	"strings"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// mustChangePwdAllowList 强制改密状态下允许访问的接口后缀
var mustChangePwdAllowList = []string{
	"/user/changePassword",
	"/user/getUserInfo",
	"/jwt/jsonInBlacklist",
}

// MustChangePwdGuard 当 jwt 携带 MustChangePwd=true 时 仅放行改密/用户信息/登出 其余 403
func MustChangePwdGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, exists := c.Get("claims")
		if !exists {
			c.Next()
			return
		}
		claims, ok := raw.(*systemReq.CustomClaims)
		if !ok || !claims.MustChangePwd {
			c.Next()
			return
		}
		path := c.Request.URL.Path
		for _, allow := range mustChangePwdAllowList {
			if strings.HasSuffix(path, allow) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"code": 7,
			"data": gin.H{"needChangePassword": true},
			"msg":  "密码已过期，请先修改密码",
		})
		c.Abort()
	}
}
```

- [ ] **Step 4: 注册中间件**

修改 `server/initialize/router.go` 第 68 行：

```go
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.MustChangePwdGuard()).Use(middleware.CasbinHandler())
```

- [ ] **Step 5: 跑测试看通过 + 编译**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/ -run TestMustChangePwdGuard -v && go build ./...`
Expected: 测试四个断言 PASS；`go build ./...` 无错误。

- [ ] **Step 6: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/middleware/must_change_pwd.go server/middleware/must_change_pwd_test.go server/initialize/router.go && git commit -m "feat(security): add MustChangePwdGuard middleware"
```

---

### Task 11: 限流接入安全配置（cache 化 limit_ip 挂登录/验证码）

**Files:**
- Modify: `server/middleware/limit_ip.go:44-92`（`DefaultCheckOrMark`/`SetLimitWithTime` 改走 `global.GVA_CACHE`，新增 `SecurityLimit()`）
- Modify: `server/router/system/sys_user.go`（在 `/base/login`、`/base/captcha` 挂 `middleware.SecurityLimit()`）
- Test: `server/middleware/limit_ip_test.go`

**Interfaces:**
- Consumes: `global.GVA_CACHE`（Plan A）、`securityConfigService.Current()`（Task 3，经包级函数读取避免 import 循环 — 见下）。
- Produces:
  - `func CacheCheckOrMark(key string, expire int, limit int) error`（基于 `global.GVA_CACHE.IncrementWithExpire`，超限返回错误，cache 异常 fail-open）。
  - `func SecurityLimit() gin.HandlerFunc`（读安全配置：未开启则直接放行；否则按 window/count 限流）。

- [ ] **Step 1: 确认基础路由文件**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && sed -n '1,40p' router/system/sys_user.go`
Expected: 看到 `InitBaseRouter`，内含 `baseRouter.POST("login", baseApi.Login)`、`baseRouter.POST("captcha", baseApi.Captcha)`（公开组）。记下分组变量名。

- [ ] **Step 2: 写失败测试**

创建 `server/middleware/limit_ip_test.go`：

```go
package middleware

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
)

func TestCacheCheckOrMark(t *testing.T) {
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Hour)
	key := "GVA_Limit_test_1.2.3.4"
	// limit=2 前两次放行 第三次拒绝
	if err := CacheCheckOrMark(key, 60, 2); err != nil {
		t.Fatalf("1st should pass, got %v", err)
	}
	if err := CacheCheckOrMark(key, 60, 2); err != nil {
		t.Fatalf("2nd should pass, got %v", err)
	}
	if err := CacheCheckOrMark(key, 60, 2); err == nil {
		t.Fatalf("3rd should be limited, got nil")
	}
}
```

> 构造器名 `gva_cache.NewMemoryCache` 以 Plan A 实际导出名为准。

- [ ] **Step 3: 跑测试看它失败**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/ -run TestCacheCheckOrMark -v`
Expected: 编译失败，`undefined: CacheCheckOrMark`。

- [ ] **Step 4: 写最小实现 — cache 化限流函数**

在 `server/middleware/limit_ip.go` 末尾追加（保留原 `DefaultCheckOrMark`/`SetLimitWithTime` 不动，新增 cache 版以免影响其它调用方）：

```go
// CacheCheckOrMark 基于 GVA_CACHE 的限流计数 超限返回错误 cache 异常 fail-open
func CacheCheckOrMark(key string, expire int, limit int) error {
	n, err := global.GVA_CACHE.IncrementWithExpire(key, 1, time.Duration(expire)*time.Second)
	if err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
		return nil // fail-open
	}
	if int(n) > limit {
		return errors.New("请求太过频繁，请稍后再试")
	}
	return nil
}
```

确认 `limit_ip.go` import 含 `"errors"`、`"time"`、`"go.uber.org/zap"`、`global`（已存在）。

- [ ] **Step 5: 写最小实现 — SecurityLimit 中间件**

为避免 middleware → service 的 import 循环，在 service 层暴露一个读当前限流配置的轻量取值器。在 `server/service/system/sys_security_config.go` 末尾追加：

```go
// CurrentLimit 供中间件读取限流配置 返回 enable/window/count
func (s *SecurityConfigService) CurrentLimit() (enable bool, window int, count int) {
	cfg := s.Current()
	return cfg.LimitEnable, cfg.LimitWindow, cfg.LimitCount
}
```

在 `server/middleware/limit_ip.go` 末尾追加 `SecurityLimit`（通过 `service.ServiceGroupApp` 读取，middleware 已可依赖 service 包，参考既有 middleware 用法）：

```go
// SecurityLimit 按安全配置对登录/敏感接口限流 未开启则放行
func SecurityLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		enable, window, count := service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.CurrentLimit()
		if !enable {
			c.Next()
			return
		}
		key := "GVA_SecLimit" + c.ClientIP() + c.FullPath()
		if err := CacheCheckOrMark(key, window, count); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
```

在 `limit_ip.go` import 块加入 `"github.com/flipped-aurora/gin-vue-admin/server/service"`。

> 验证 middleware 是否已 import service：若出现 import 循环（service → middleware），改为在 service 包内不引用 middleware；当前 middleware → service 为单向，安全。编译时若报循环，将 `CurrentLimit` 读取改为通过 `global` 暴露的函数指针，但默认按此实现，编译验证决定。

- [ ] **Step 6: 写最小实现 — 路由挂载**

在 `server/router/system/sys_user.go` 的 `InitBaseRouter` 中，将 login/captcha 两行改为挂限流中间件（用 Step 1 确认的分组变量名，下例假定为 `baseRouter`）：

```go
		baseRouter.POST("login", middleware.SecurityLimit(), baseApi.Login)
		baseRouter.POST("captcha", middleware.SecurityLimit(), baseApi.Captcha)
```

确认该文件已 import `"github.com/flipped-aurora/gin-vue-admin/server/middleware"`，无则补。

- [ ] **Step 7: 跑测试看通过 + 编译**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./middleware/ -run TestCacheCheckOrMark -v && go build ./...`
Expected: 测试 PASS；`go build ./...` 无错误（若报 import 循环，按 Step 5 注释处理）。

- [ ] **Step 8: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/middleware/limit_ip.go server/service/system/sys_security_config.go server/router/system/sys_user.go && git commit -m "feat(security): cache-based rate limit on login/captcha by config"
```

---

### Task 12: API GetSecurityConfig / SetSecurityConfig

**Files:**
- Create: `server/api/v1/system/sys_security_config.go`
- Modify: `server/api/v1/system/enter.go:5-31`（`ApiGroup` 加 `SecurityConfigApi`）
- Test: 编译验证（HTTP 集成测试见 Task 16）

**Interfaces:**
- Consumes: `securityConfigService.Get()`/`Set()`（Task 3）。
- Produces:
  - `type SecurityConfigApi struct{}`
  - `func (s *SecurityConfigApi) GetSecurityConfig(c *gin.Context)` → `GET /securityConfig/getSecurityConfig`
  - `func (s *SecurityConfigApi) SetSecurityConfig(c *gin.Context)` → `POST /securityConfig/setSecurityConfig`

- [ ] **Step 1: 写最小实现 — API**

创建 `server/api/v1/system/sys_security_config.go`：

```go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SecurityConfigApi struct{}

// GetSecurityConfig
// @Tags      SecurityConfig
// @Summary   获取安全配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=system.SysSecurityConfig,msg=string}  "获取安全配置"
// @Router    /securityConfig/getSecurityConfig [get]
func (s *SecurityConfigApi) GetSecurityConfig(c *gin.Context) {
	cfg, err := securityConfigService.Get()
	if err != nil {
		global.GVA_LOG.Error("获取安全配置失败!", zap.Error(err))
		response.FailWithMessage("获取安全配置失败", c)
		return
	}
	response.OkWithDetailed(cfg, "获取成功", c)
}

// SetSecurityConfig
// @Tags      SecurityConfig
// @Summary   设置安全配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysSecurityConfig                                     true  "安全配置"
// @Success   200   {object}  response.Response{data=system.SysSecurityConfig,msg=string}  "设置安全配置"
// @Router    /securityConfig/setSecurityConfig [post]
func (s *SecurityConfigApi) SetSecurityConfig(c *gin.Context) {
	var cfg system.SysSecurityConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := securityConfigService.Set(cfg); err != nil {
		global.GVA_LOG.Error("设置安全配置失败!", zap.Error(err))
		response.FailWithMessage("设置安全配置失败", c)
		return
	}
	saved, _ := securityConfigService.Get()
	response.OkWithDetailed(saved, "设置成功", c)
}
```

- [ ] **Step 2: 注册到 ApiGroup**

在 `server/api/v1/system/enter.go` 的 `ApiGroup` struct（约第 30 行 `AIWorkflowSessionApi` 之后）追加：

```go
	AIWorkflowSessionApi
	SecurityConfigApi
```

- [ ] **Step 3: 编译验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误。

- [ ] **Step 4: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/api/v1/system/sys_security_config.go server/api/v1/system/enter.go && git commit -m "feat(security): add GetSecurityConfig/SetSecurityConfig api"
```

---

### Task 13: Router 注册 securityConfig

**Files:**
- Create: `server/router/system/sys_security_config.go`
- Modify: `server/router/system/enter.go:5-53`（`RouterGroup` 加 `SecurityConfigRouter`，var 块加 `securityConfigApi`）
- Modify: `server/initialize/router.go:97`（注册 `InitSecurityConfigRouter`）

**Interfaces:**
- Consumes: `securityConfigApi`（Task 12）。
- Produces: `func (s *SecurityConfigRouter) InitSecurityConfigRouter(Router *gin.RouterGroup)`（私有组）。

- [ ] **Step 1: 写最小实现 — Router**

创建 `server/router/system/sys_security_config.go`：

```go
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SecurityConfigRouter struct{}

// InitSecurityConfigRouter 初始化 安全配置 路由信息
func (s *SecurityConfigRouter) InitSecurityConfigRouter(Router *gin.RouterGroup) {
	securityConfigRouter := Router.Group("securityConfig").Use(middleware.OperationRecord())
	securityConfigRouterWithoutRecord := Router.Group("securityConfig")
	{
		securityConfigRouter.POST("setSecurityConfig", securityConfigApi.SetSecurityConfig) // 设置安全配置
	}
	{
		securityConfigRouterWithoutRecord.GET("getSecurityConfig", securityConfigApi.GetSecurityConfig) // 获取安全配置
	}
}
```

- [ ] **Step 2: 注册到 RouterGroup + var**

在 `server/router/system/enter.go` 的 `RouterGroup` struct（约第 26 行 `SkillsRouter` 之后）追加：

```go
	SkillsRouter
	SecurityConfigRouter
```

并在 var 块（约第 52 行 `aiWorkflowSessionApi` 之后）追加：

```go
	aiWorkflowSessionApi = api.ApiGroupApp.SystemApiGroup.AIWorkflowSessionApi
	securityConfigApi    = api.ApiGroupApp.SystemApiGroup.SecurityConfigApi
```

- [ ] **Step 3: 在初始化总路由注册**

在 `server/initialize/router.go` 第 97 行 `InitLoginLogRouter` 之后追加：

```go
		systemRouter.InitLoginLogRouter(PrivateGroup)                       // 登录日志
		systemRouter.InitSecurityConfigRouter(PrivateGroup)                 // 安全配置
```

- [ ] **Step 4: 编译验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/router/system/sys_security_config.go server/router/system/enter.go server/initialize/router.go && git commit -m "feat(security): register securityConfig router"
```

---

### Task 14: 种子 — API + 菜单 + 默认配置单行

**Files:**
- Modify: `server/source/system/api.go`（`entities` 切片加 2 条「安全配置」API）
- Modify: `server/source/system/menu.go`（`childMenus` 加「安全配置」菜单）
- Create: `server/source/system/security_config.go`（默认配置单行初始化器）

**Interfaces:**
- Consumes: `system.SysSecurityConfig`、`system.DefaultSecurityConfig`（Task 1）、`config.yaml` captcha。
- Produces: 启动初始化时写入 2 条 API、1 条菜单、1 行默认配置。

- [ ] **Step 1: 写最小实现 — API 种子**

在 `server/source/system/api.go` 的 `entities := []sysModel.SysApi{` 切片内追加（紧邻「API Token」分组块之后）：

```go
		{ApiGroup: "安全配置", Method: "GET", Path: "/securityConfig/getSecurityConfig", Description: "获取安全配置"},
		{ApiGroup: "安全配置", Method: "POST", Path: "/securityConfig/setSecurityConfig", Description: "设置安全配置"},
```

- [ ] **Step 2: 写最小实现 — 菜单种子**

在 `server/source/system/menu.go` 的 `childMenus` 切片内（superAdmin 子菜单块、`sysError` 行之后）追加：

```go
			{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "security", Name: "security", Component: "view/system/security/index.vue", Sort: 13, Meta: Meta{Title: "安全配置", Icon: "lock"}},
```

- [ ] **Step 3: 写最小实现 — 默认配置初始化器**

参考 `server/source/system/api.go` 的初始化器结构，创建 `server/source/system/security_config.go`：

```go
package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initSecurityConfig struct{}

const initOrderSecurityConfig = system.InitOrderSystem + 100

// auto run
func init() {
	system.RegisterInit(initOrderSecurityConfig, &initSecurityConfig{})
}

func (i *initSecurityConfig) InitializerName() string {
	return sysModel.SysSecurityConfig{}.TableName()
}

func (i *initSecurityConfig) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysSecurityConfig{})
}

func (i *initSecurityConfig) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysSecurityConfig{})
}

func (i *initSecurityConfig) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	cfg := sysModel.DefaultSecurityConfig(global.GVA_CONFIG.Captcha)
	cfg.ID = 1
	if err := db.Create(&cfg).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysSecurityConfig{}.TableName()+"默认配置初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), cfg)
	return next, nil
}

func (i *initSecurityConfig) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", 1).First(&sysModel.SysSecurityConfig{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
```

- [ ] **Step 4: 编译验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./...`
Expected: 无错误。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/source/system/api.go server/source/system/menu.go server/source/system/security_config.go && git commit -m "feat(security): seed security api, menu and default config row"
```

---

### Task 15: 前端 — API 封装 + 配置页 + 登录强制改密流程

**Files:**
- Create: `web/src/api/securityConfig.js`
- Create: `web/src/view/system/security/index.vue`
- Create: `web/src/view/system/security/forceChangePassword.vue`
- Modify: `web/src/pinia/modules/user.js:62-111`（`LoginIn` 处理 `needChangePassword`）

**Interfaces:**
- Consumes: 后端 `GET /securityConfig/getSecurityConfig`、`POST /securityConfig/setSecurityConfig`、`LoginResponse.needChangePassword`（Task 8/12）。
- Produces: 配置页（5 分区）、强制改密页、登录后跳转逻辑。

- [ ] **Step 1: 写最小实现 — API 封装**

创建 `web/src/api/securityConfig.js`：

```js
import service from '@/utils/request'

// @Tags SecurityConfig
// @Summary 获取安全配置
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=system.SysSecurityConfig,msg=string} "获取安全配置"
// @Router /securityConfig/getSecurityConfig [get]
export const getSecurityConfig = () => {
  return service({
    url: '/securityConfig/getSecurityConfig',
    method: 'get'
  })
}

// @Tags SecurityConfig
// @Summary 设置安全配置
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysSecurityConfig true "安全配置"
// @Success 200 {object} response.Response{data=system.SysSecurityConfig,msg=string} "设置安全配置"
// @Router /securityConfig/setSecurityConfig [post]
export const setSecurityConfig = (data) => {
  return service({
    url: '/securityConfig/setSecurityConfig',
    method: 'post',
    data
  })
}
```

- [ ] **Step 2: 写最小实现 — 配置页（5 Tab，UnoCSS）**

创建 `web/src/view/system/security/index.vue`：

```vue
<template>
  <div class="p-4">
    <el-card>
      <template #header>
        <div class="flex items-center justify-between">
          <span class="text-base font-medium">安全配置</span>
          <el-button type="primary" :loading="saving" @click="onSave">保存</el-button>
        </div>
      </template>

      <el-tabs v-model="activeTab">
        <el-tab-pane label="验证码" name="captcha">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="错误N次后出验证码">
              <el-input-number v-model="form.captchaOpen" :min="0" />
              <span class="ml-2 text-gray-400 text-sm">0 = 每次都需要验证码</span>
            </el-form-item>
            <el-form-item label="计数缓存超时(秒)">
              <el-input-number v-model="form.captchaTimeout" :min="1" />
            </el-form-item>
            <el-form-item label="验证码长度">
              <el-input-number v-model="form.keyLong" :min="1" />
            </el-form-item>
            <el-form-item label="验证码宽度">
              <el-input-number v-model="form.imgWidth" :min="1" />
            </el-form-item>
            <el-form-item label="验证码高度">
              <el-input-number v-model="form.imgHeight" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="密码复杂度" name="password">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="最小长度">
              <el-input-number v-model="form.pwdMinLength" :min="1" />
            </el-form-item>
            <el-form-item label="需大写字母">
              <el-switch v-model="form.pwdRequireUpper" />
            </el-form-item>
            <el-form-item label="需小写字母">
              <el-switch v-model="form.pwdRequireLower" />
            </el-form-item>
            <el-form-item label="需数字">
              <el-switch v-model="form.pwdRequireDigit" />
            </el-form-item>
            <el-form-item label="需特殊字符">
              <el-switch v-model="form.pwdRequireSpecial" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="限流" name="limit">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启限流">
              <el-switch v-model="form.limitEnable" />
            </el-form-item>
            <el-form-item label="窗口(秒)">
              <el-input-number v-model="form.limitWindow" :min="1" />
            </el-form-item>
            <el-form-item label="窗口内最大次数">
              <el-input-number v-model="form.limitCount" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="失败锁定" name="lock">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启失败锁定">
              <el-switch v-model="form.lockEnable" />
            </el-form-item>
            <el-form-item label="失败次数阈值">
              <el-input-number v-model="form.lockThreshold" :min="1" />
            </el-form-item>
            <el-form-item label="锁定时长(分钟)">
              <el-input-number v-model="form.lockDuration" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="密码过期" name="expire">
          <el-form :model="form" label-width="160px" class="max-w-2xl">
            <el-form-item label="开启密码过期">
              <el-switch v-model="form.pwdExpireEnable" />
            </el-form-item>
            <el-form-item label="有效天数">
              <el-input-number v-model="form.pwdExpireDays" :min="1" />
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getSecurityConfig, setSecurityConfig } from '@/api/securityConfig'

defineOptions({ name: 'SecurityConfig' })

const activeTab = ref('captcha')
const saving = ref(false)
const form = ref({
  captchaOpen: 0,
  captchaTimeout: 3600,
  keyLong: 6,
  imgWidth: 240,
  imgHeight: 80,
  pwdMinLength: 8,
  pwdRequireUpper: false,
  pwdRequireLower: false,
  pwdRequireDigit: false,
  pwdRequireSpecial: false,
  limitEnable: false,
  limitWindow: 60,
  limitCount: 30,
  lockEnable: false,
  lockThreshold: 5,
  lockDuration: 30,
  pwdExpireEnable: false,
  pwdExpireDays: 90
})

const load = async () => {
  const res = await getSecurityConfig()
  if (res.code === 0) {
    form.value = { ...form.value, ...res.data }
  }
}

const onSave = async () => {
  saving.value = true
  try {
    const res = await setSecurityConfig(form.value)
    if (res.code === 0) {
      form.value = { ...form.value, ...res.data }
      ElMessage.success('保存成功')
    }
  } finally {
    saving.value = false
  }
}

onMounted(load)
</script>
```

- [ ] **Step 3: 写最小实现 — 强制改密页**

创建 `web/src/view/system/security/forceChangePassword.vue`：

```vue
<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <el-card class="w-96">
      <template #header>
        <span class="text-base font-medium">密码已过期，请修改密码</span>
      </template>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="原密码" prop="password">
          <el-input v-model="form.password" type="password" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="form.newPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirm">
          <el-input v-model="form.confirm" type="password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" @click="onSubmit">提交并重新登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { changePassword } from '@/api/user'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({ name: 'ForceChangePassword' })

const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const form = ref({ password: '', newPassword: '', confirm: '' })

const rules = {
  password: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
  newPassword: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
  confirm: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    {
      validator: (rule, value, cb) =>
        value === form.value.newPassword ? cb() : cb(new Error('两次密码不一致')),
      trigger: 'blur'
    }
  ]
}

const onSubmit = async () => {
  await formRef.value.validate()
  loading.value = true
  try {
    const res = await changePassword({
      password: form.value.password,
      newPassword: form.value.newPassword
    })
    if (res.code === 0) {
      ElMessage.success('修改成功，请重新登录')
      await userStore.LoginOut()
    }
  } finally {
    loading.value = false
  }
}
</script>
```

> 注：`changePassword` 入参字段名以 `web/src/api/user.js` 既有定义为准；若现有签名为 `{ password, newPassword }` 则直接复用，否则按实际字段调整。

- [ ] **Step 4: 写最小实现 — 登录流程处理 needChangePassword**

在 `web/src/pinia/modules/user.js` 的 `LoginIn` 中，`setToken(res.data.token)` 之后、初始化路由之前插入：

```js
      setUserInfo(res.data.user)
      setToken(res.data.token)

      // 密码过期 强制跳转改密页
      if (res.data.needChangePassword) {
        await router.push({ name: 'ForceChangePassword' })
        return true
      }
```

并确保前端路由表注册了 `ForceChangePassword`（在 `web/src/router/index.js` 的公开/基础路由数组加入；用项目既有静态路由写法）：

Run: `cd /Users/jiangjizhao/gin-vue-admin && sed -n '1,80p' web/src/router/index.js`
Expected: 看到静态路由数组（含 `login`/`init` 等）。在该数组追加：

```js
  {
    path: '/forceChangePassword',
    name: 'ForceChangePassword',
    component: () => import('@/view/system/security/forceChangePassword.vue'),
    meta: { title: '修改密码' }
  },
```

- [ ] **Step 5: 前端构建验证**

Run: `cd /Users/jiangjizhao/gin-vue-admin/web && npm run build`
Expected: 构建成功，无报错（首次需 `npm install`）。

- [ ] **Step 6: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add web/src/api/securityConfig.js web/src/view/system/security/ web/src/pinia/modules/user.js web/src/router/index.js && git commit -m "feat(security): security config page, force change password flow, login handling"
```

---

### Task 16: 集成测试 — 锁定流程 & 过期强制改密流程

**Files:**
- Create: `server/api/v1/system/security_integration_test.go`
- Test: 同文件

**Interfaces:**
- Consumes: `system.IsAccountLocked`/`RecordLoginFail`/`ClearLoginFail`/`IsPasswordExpired`（Task 6）、`MustChangePwdGuard`（Task 10）、`global.GVA_CACHE`（Plan A）、`request.CustomClaims.MustChangePwd`（Task 5）。
- Produces: 端到端验证锁定与强制改密拦截。

- [ ] **Step 1: 写失败测试（锁定流程 + 中间件拦截）**

创建 `server/api/v1/system/security_integration_test.go`：

```go
package system

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
	"github.com/gin-gonic/gin"
)

func TestAccountLockFlow(t *testing.T) {
	global.GVA_CACHE = gva_cache.NewMemoryCache(time.Hour)
	cfg := sysModel.SysSecurityConfig{LockEnable: true, LockThreshold: 2, LockDuration: 30}
	user := "integration_lock_user"

	systemService.ClearLoginFail(user)
	systemService.RecordLoginFail(user, cfg)
	if systemService.IsAccountLocked(user) {
		t.Fatalf("should not be locked after 1 fail")
	}
	systemService.RecordLoginFail(user, cfg)
	if !systemService.IsAccountLocked(user) {
		t.Fatalf("should be locked after threshold reached")
	}
	systemService.ClearLoginFail(user)
	if systemService.IsAccountLocked(user) {
		t.Fatalf("should be unlocked after clear")
	}
}

func TestMustChangePwdMiddlewareBlocks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &systemReq.CustomClaims{MustChangePwd: true})
		c.Next()
	})
	r.Use(middleware.MustChangePwdGuard())
	r.POST("/user/getUserList", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.POST("/user/changePassword", func(c *gin.Context) { c.Status(http.StatusOK) })

	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, httptest.NewRequest(http.MethodPost, "/user/getUserList", nil))
	if w1.Code != http.StatusForbidden {
		t.Fatalf("getUserList under must-change should be 403, got %d", w1.Code)
	}

	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, httptest.NewRequest(http.MethodPost, "/user/changePassword", nil))
	if w2.Code != http.StatusOK {
		t.Fatalf("changePassword under must-change should pass, got %d", w2.Code)
	}
}

func TestPasswordExpiryFlow(t *testing.T) {
	now := time.Now()
	cfg := sysModel.SysSecurityConfig{PwdExpireEnable: true, PwdExpireDays: 30}
	old := now.AddDate(0, 0, -40)
	if !systemService.IsPasswordExpired(&old, cfg, now) {
		t.Fatalf("40 days old with 30d expiry should be expired")
	}
	fresh := now.AddDate(0, 0, -5)
	if systemService.IsPasswordExpired(&fresh, cfg, now) {
		t.Fatalf("5 days old should not be expired")
	}
}
```

- [ ] **Step 2: 跑测试看它失败（实现存在时应直接通过；若早于依赖任务则编译失败）**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./api/v1/system/ -run 'TestAccountLockFlow|TestMustChangePwdMiddlewareBlocks|TestPasswordExpiryFlow' -v`
Expected: 在 Task 5/6/10 完成前为编译失败（`undefined`）；本任务在它们之后执行，预期 PASS。

- [ ] **Step 3: 跑测试看通过**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go test ./api/v1/system/ -run 'TestAccountLockFlow|TestMustChangePwdMiddlewareBlocks|TestPasswordExpiryFlow' -v`
Expected: 三个测试 PASS。

- [ ] **Step 4: 全量回归**

Run: `cd /Users/jiangjizhao/gin-vue-admin/server && go build ./... && go test ./...`
Expected: 编译无错误；新增测试包全部 PASS（既有依赖外部 DB/Redis 的测试若原本即跳过/失败，不在本计划改动范围）。

- [ ] **Step 5: Commit**

```bash
cd /Users/jiangjizhao/gin-vue-admin && git add server/api/v1/system/security_integration_test.go && git commit -m "test(security): integration tests for lock flow and must-change-pwd guard"
```

---

## Self-Review Coverage

spec 覆盖点 → 对应 Task 编号：

- 数据模型 `SysSecurityConfig` 单行表（spec 5.1，字段表）→ Task 1
- `SysUser.PasswordUpdatedAt *time.Time` 加列 + AutoMigrate 注册（spec 5.1 / 5.4）→ Task 1
- `SecurityConfigService` Get/Set/Current + 密码过期回填（spec 5.2）→ Task 3
- 启动加载配置入内存（spec 6 数据流「启动」）→ Task 4
- 密码复杂度校验器 `ValidatePasswordComplexity`（spec 5.3.2）→ Task 2；三处接入（建用户/自助改密/重置）→ Task 9
- 验证码配置迁到新表、登录/验证码接口读 `Current()`、计数走 `GVA_CACHE`（spec 5.3.1 / 4.5）→ Task 7（验证码接口）、Task 8（登录接口）
- 连续登录失败锁定（按账号，`login_fail:`/`login_lock:`，阈值/TTL/成功清零，入口先查锁）（spec 5.3.4）→ Task 6（计数/锁定/过期工具）、Task 8（登录入口集成）
- 限流 cache 化 `limit_ip` 按配置挂 `/base/login`、`/base/captcha`（spec 5.3.3）→ Task 11
- 密码过期 + 强制改密：`MustChangePwd` claim（spec 5.3.5）→ Task 5；过期判定 → Task 6；登录签发带 claim + `needChangePassword` → Task 8；中间件强拦 → Task 10；改密成功回填 `PasswordUpdatedAt` → Task 9
- API `GetSecurityConfig`/`SetSecurityConfig`（私有组，Swagger `response.Response{data=system.SysSecurityConfig,msg=string}`）（spec 5.4）→ Task 12
- Router `securityConfig` + enter + initialize/router 注册（spec 5.4）→ Task 13
- 种子：2 条 API（ApiGroup「安全配置」）、菜单（挂系统下）、默认配置单行（取 config.yaml captcha）（spec 5.4 种子）→ Task 14
- enter.go 三层装配（api / service / router）→ Task 3（service）、Task 7/12（api）、Task 13（router）
- 前端 `web/src/api/securityConfig.js`、`web/src/view/system/security/` 5 分区配置页、登录处理 `needChangePassword` 跳转强制改密页（spec 5.5）→ Task 15
- fail-open 缓存异常降级（spec 4.6 / 7）→ Task 6（`RecordLoginFail` 计数失败放弃）、Task 11（`CacheCheckOrMark` fail-open）
- 配置行缺失自动创建默认行（spec 7）→ Task 3（`Get` 惰性创建）、Task 14（种子）
- 密码过期关闭完全跳过（spec 7）→ Task 6（`IsPasswordExpired` 未开启返回 false）、Task 8（仅在 needChange 时签 claim）
- 测试：复杂度校验器表驱动 → Task 2；过期计算 → Task 6；锁定计数（内存 cache）→ Task 6；Service get/set/回填缓存 → Task 3；集成测试 锁定流程 + 过期强制改密流程 → Task 16
