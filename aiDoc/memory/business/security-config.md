# 新增「安全配置」模块（系统本体安全能力）

## 基本信息

- 提出日期：2026-06-25
- 当前状态：`active`（设计已确认，待写实现计划）
- 需求类型：系统本体新增功能模块
- 设计文档：`docs/superpowers/specs/2026-06-25-security-config-design.md`

## 用户原始意图摘要

在 system 下新增「安全配置」功能板块，能力包含：验证码配置、密码复杂度要求、限流、连续登录失败锁定、密码过期（管理员可设有效期，如 90 天，超期登录强制改密）。需要创建初始化的菜单与 API。

## 影响范围

- 后端：`server/utils/gva_cache`（新，通用缓存）、`server/model/system`、`server/api/v1/system`、`server/service/system`、`server/router/system`、`server/source/system`、`server/initialize`、`server/middleware`（jwt/limit_ip）。
- 前端：`web/src/view/system/security/`、`web/src/api/securityConfig.js`、登录流程（强制改密跳转）。
- 数据：新表 `sys_security_config`（单行）；`sys_users` 加列 `password_updated_at`。

## 已确认约束

- 配置存储：专用 DB 单行表 `SysSecurityConfig`，启动入内存、保存即热更新。
- 通用缓存：新建「有 Redis 用 Redis、无则内存」的统一 `Cache`（`utils/gva_cache` + `global.GVA_CACHE`），**全量迁移并彻底退役 `global.BlackCache`**（jwt 黑名单 / 验证码计数 / 登录失败计数 / limit_ip 都改用它）。见 [[ai-plugin-migration]] 的同类「不改逻辑搬迁」风格。
- 限流：仅登录/敏感接口 + 可配置（enable/window/count）。
- 失败锁定：按账号；阈值 + 锁定时长可配；成功登录清零；与按 IP 的验证码防爆破独立并存。
- 密码过期：JWT 加 `MustChangePwd` claim + 中间件强拦（只放行改密/用户信息/登出），前端强制跳转改密；存量用户仅在开启时回填 `PasswordUpdatedAt=now`。
- 密码复杂度：可配最小长度 + 大小写/数字/特殊四类开关；在 管理员建用户 / 自助改密 / 重置密码 三处生效。
- 验证码配置：迁到新表作唯一来源，config.yaml 的 captcha 仅作首次默认值。
- 路由前缀 `/securityConfig`；Swagger 注释遵循 `aiDoc/modules/backend-layer-rules.md` 新规范。

## 当前进展

- ✅ 需求澄清与设计确认（10 项关键决策）
- ✅ 设计文档落盘
- ⬜ 实现计划（writing-plans）
- ⬜ Phase 0 通用缓存 + 迁移
- ⬜ Phase 1 安全配置功能（模型/服务/集成/API/种子/前端）

## 后续待办 / 注意

- 实现分两阶段：先做缓存地基再做功能。
- 注意迁移顺序：建表 + 种子 + `password_updated_at` 加列先于首次使用。
- 缓存运行时异常按 fail-open 降级，避免误伤正常用户。
