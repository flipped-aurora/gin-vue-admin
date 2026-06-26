# 顶栏边框改为可调尺寸阴影（box-shadow）

## 基本信息

- 提出日期：2026-06-23
- 当前状态：`done`
- 需求类型：前端 / 主题系统增强
- 优先级：中
- 需求文件：`aiDoc/memory/business/done/theme-header-shadow.md`

## 用户原始意图摘要

从主题到页面全局替换掉 `--gva-header-border` 相关内容，改用 `box-shadow`；阴影设置项支持调整尺寸（含「无」），其中 `azir-fresh-blue` 预设固定为「无」。

## 影响范围

- 后端：无
- 前端：主题系统配置、落地层、样式默认值、顶栏页面、设置面板、内置预设
- 文档：本记忆
- 插件 / 模块：`gva-theme` 主题系统

## 涉及对象

- 模块：`web/src/theme`、`web/src/utils/theme.js`、`web/src/style/theme.scss`
- 接口：无
- 页面：
  - `web/src/view/layout/header/index.vue`（顶栏由 `borderBottom` 改为 `boxShadow: var(--gva-header-shadow)`）
  - `web/src/view/layout/setting/modules/layout/index.vue`（「顶栏边框」颜色选择器 → 「顶栏阴影」`el-select`：无/小/中/大）
- 配置：
  - 配置字段 `header_border`（颜色）→ `header_shadow`（尺寸档位）
  - CSS 变量 `--gva-header-border` → `--gva-header-shadow`

## 已确认约束

- `header_shadow` 档位取值：`none` / `sm` / `md` / `lg`，对应「无 / 小 / 中 / 大」
- `azir-fresh-blue` 预设固定为 `none`；`gva-classic` 预设为 `sm`；`dark-night` 不声明该字段，回退 `BASE_CONFIG` 的 `sm`
- 阴影按明暗各一套（`HEADER_SHADOWS.light` / `.dark`），不走 `autoDarkColor`（阴影非颜色）
- 采用「1px 硬发丝线 + 渐增模糊」模型：各档位保留一条清晰分隔线（贴近旧 1px 边框观感），模糊随档位增强；`none` 为完全无阴影
- `BASE_CONFIG.header_shadow` 默认 `sm`
- `theme.scss` 中 `:root` / `html.dark` 的 `--gva-header-shadow` 兜底默认值必须与 `HEADER_SHADOWS` 的 `sm` 保持同步

## 当前进展

- 已完成全部 7 处改动并落地工作区
- ESLint 通过；全仓 `header_border` / `--gva-header-border` 残留为 0
- 三视角（正确性 / 完整性 / 视觉 UX）对抗式评审结论均为 ship
- 配置无 localStorage 持久化，旧自定义预设携带的 `header_border` 旧键会被 `applyPreset` 静默忽略并回退默认，无需迁移

## 后续待办

- （可选，越界未改）`azir-fresh-blue.json` 的 `version: 1` 与 `PRESET_VERSION = 2` 不一致；该字段不参与内置预设迁移逻辑，纯属元数据一致性，可在后续顺手对齐为 2

## 更新规则

- 本文件只承载「顶栏阴影」这一个功能点
- 同属主题系统的其他功能点请用 `theme-` 前缀新建独立文件，并与本文件互相反向链接
