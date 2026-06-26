# 基础组件库（reka-ui 底座）

本文件是项目「基础 UI 组件库」的唯一规范来源。任何 AI 在**新增、修改或使用**这套组件时，
都必须遵守这里的约定，确保组件统一接入主题系统、可换肤、暗色可用，并与构建管线保持一致。

## 适用范围（强制 / 可选）

本组件库**全局可用**（`g-` 前缀，kebab-case），但「是否必须用」按场景区分，不要在业务页面里一刀切地强推：

- **必须使用（项目配置 / 框架外壳）**：凡是**项目级配置与框架壳层**的 UI，必须用本组件库，
  以保证全站换肤、暗色、主题 token 始终一致。典型范围：
  - 顶栏 header、侧边/菜单 menu、整体布局 layout
  - 系统配置抽屉、主题 / 外观设置、预设管理等配置类面板
- **不强制（业务组件 / 业务页面）**：CRUD 列表、业务表单页、业务弹窗等**业务侧** UI，
  **以开发者使用习惯为主**，可继续用 Element Plus 等，不强制改用本库（确有需要时也可自行选用）。

一句话：**配置 / 外壳层强制 reka-ui，业务页面尊重习惯、不强制**。

## 选型结论

- 底座选用 **reka-ui**（无样式、可访问的 Vue 原语，前身 Radix Vue），对 UnoCSS **零桥接**。
- **不引入 shadcn-vue 的预制样式层**：本项目用 UnoCSS（preset-wind3），社区 `unocss-preset-shadcn`
  默认 presetWind4 与本项目错配、维护停滞，风险高。shadcn-vue 源码仅作为「设计蓝本」参考。
- 上色**只用项目自有的 UnoCSS 语义 token**，因此组件天然跟随 `themeStore` 换肤与暗色模式。

## 引入位置

- 组件库目录：`web/src/core/componentLibrary/`
- 统一出口（barrel）：`web/src/core/componentLibrary/index.js`
- 全局注册逻辑：`web/src/core/global.js` 的 `registerComponentLibrary()`
- 新增依赖：`reka-ui`、`class-variance-authority`(cva)、`clsx`、`tailwind-merge`

目录结构（一个组件一个目录，`Xxx.vue` 放实现、`index.js` 放出口 / cva 变体）：

```
core/componentLibrary/
├── index.js            # 总出口：re-export 各组件 + cn
├── utils.js            # cn()：clsx + tailwind-merge 合并 class
├── button/             # Button + buttonVariants(cva)
├── select/             # Select(:options 便捷模式) + Trigger/Content/Item 部件
├── switch/             # Switch
├── slider/             # Slider（单值 number 对外，内部包数组，支持 marks）
├── number-field/       # NumberField（数字步进输入）
└── color-picker/       # ColorPicker（Popover 内组合 reka 颜色原语，支持 alpha）
```

## 使用方式

### 1. 全局组件（推荐）

`core/global.js` 会把 barrel 导出的每个组件以 **`g-` 前缀（kebab-case）**注册为全局组件，
命名与项目里 `el-button` 等用法统一，全站直接用、无需 import：

| 组件        | 全局标签             |
| ----------- | -------------------- |
| Button      | `<g-button />`       |
| Select      | `<g-select />`       |
| Switch      | `<g-switch />`       |
| Slider      | `<g-slider />`       |
| NumberField | `<g-number-field />` |
| ColorPicker | `<g-color-picker />` |

注册是**自动遍历** barrel 导出实现的（`g-` + 导出名转 kebab-case），新增组件只要从 `index.js`
导出即自动获得全局标签，无需再改 `global.js`；`cn` 这类工具函数导出（非对象）会被跳过。
个别从 barrel re-export 的「非本库自有组件对象」（如 reka-ui 的 `SelectValue`，仅供 granular 模式按需 import）
不应获得全局标签，已由 `global.js` 的 `NON_GLOBAL_EXPORTS` 名单显式排除——新增此类 re-export 时同步登记。

> **命名三层关系**（刻意分层，勿混用）：组件内 `defineOptions({ name: 'UiXxx' })` = devtools 显示名 /
> barrel 导出 `Xxx`（PascalCase）= 显式 import 名 / 全局标签 `g-xxx`（kebab-case）= 模板里用。

### 2. 显式 import（仍受支持）

需要按需引入、或使用 Select 的 granular 部件（`SelectTrigger / SelectContent / SelectItem`）时：

```js
import { Button, Select } from '@/core/componentLibrary'
// 或细到单组件目录
import { Button } from '@/core/componentLibrary/button'
```

## 必须遵守的 UI / 主题规范（硬约束）

1. **只用项目语义 token 上色**，禁止写死颜色。可用 token 来自 `web/src/theme/vars.js`：
   - 色板：`primary / info / success / warning / error`（含 `-50`~`-950` 阶梯）
   - 表面：`container`(卡片/浮层底)、`layout`(布局底)、`inverted`、`base-text`(主文本)、`border`、
     `muted`(弱底)、`muted-foreground`(弱文本)、`control-track`(控件未激活轨道：开关关闭态 / 滑块未填充)
   - 禁止用 `bg-gray-300 dark:bg-gray-600` 这类裸色阶 + 手写 `dark:` 变体上色，
     暗色应交给语义 token 在 CSS 变量层自适应（单个 `bg-control-track` 即可，无需再写 `dark:`）。
   - 阴影：`shadow-header / shadow-sider / shadow-tab / shadow-card`
2. **禁止内联换肤色**：不要写 `:style="{ backgroundColor: settings.themeColor }"`，
   颜色一律走 token（`bg-primary` 自动跟随换肤）。
3. **禁止对 CSS 变量 token 取透明度**：不写 `bg-primary/10`、`text-base-text/60` 这类
   透明度后缀（CSS 变量 + alpha 在亮/暗下不可靠）。需要弱化时改用**语义 token**，
   如 `hover:bg-muted`、`text-muted-foreground`。
4. **class 一律用 `cn()` 合并**（`clsx` 处理条件类 + `tailwind-merge` 消解冲突原子类），
   并把对外可覆盖的 `class` prop 放在最后参与合并。
5. **变体用 cva 维护**，写在组件目录的 `index.js`（如 `buttonVariants`），variant/size 各成一档。
6. **浮层用 `z-[3000]`**：组件常被放进 `el-drawer`/对话框里，下拉、Popover 的 `Content` 需
   `z-[3000]` 才能盖过 Element Plus 浮层。
7. **焦点态统一**：集中在 `componentLibrary/utils.js`，全组件引用、禁止各处手写导致漂移：
   - 控件本体可聚焦 → `FOCUS_RING`（`focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-300`）。
   - 包裹型容器（内部 input 才真正聚焦，如 NumberField / ColorField）→ `FOCUS_RING_WITHIN`（`focus-within:ring-2 focus-within:ring-primary-300`）。
8. **原生控件兜底**：项目未启用 button 背景重置式 preflight（避免影响共存的 EP），
   自定义 `<button>` 基类需带 `appearance-none bg-transparent`，再由各 variant 显式底色覆盖。
9. **`v-model` 用 `defineModel()`**：组件双向绑定一律走 `defineModel`，不写 `modelValue` prop +
   `defineEmits(['update:modelValue'])` 老样板；需变换 / 防抖 / 类型还原时也基于 `defineModel` 的 ref
   （可写 `computed` 桥接，或在回调里 `modelValue.value = ...`）。详见 `frontend-rules.md`「组件写法规范」。
10. **有限枚举 prop 加 `validator`**：如 `variant` / `size` / `format`，可选值集中到组件 `index.js` 导出
    （如 `BUTTON_VARIANTS`），供 cva 与 prop `validator` 复用做单一事实源。
11. **图标走全局 `<svg-icon>`**：组件内图标用 `<svg-icon icon="lucide:xxx" />`（在线 Iconify，优先 `lucide` 集），
    不手写裸 `<svg><path/></svg>`、不用 `el-icon` / `ep:` 图标集。详见 `frontend-rules.md`「图标规范」。

## 构建约束（关键，易踩坑）

- 把类名写进 **`.js/.ts`（cva 变体）** 时，UnoCSS 默认扫描管线**不扫 `.js/.ts`**，
  这些原子类不会被生成 → 组件丢色。
- 已在 `web/uno.config.js` 的 `content.pipeline.include` 加入规则覆盖本目录：
  `/[\\/]core[\\/]componentLibrary[\\/].*\.[jt]s($|\?)/`。
- 因此：新增组件的 cva 文件**必须放在 `core/componentLibrary/` 下**才会被扫描；
  若把含类名的 `.js/.ts` 放到别处，记得同步扩展该 include 规则。
- 改完 `uno.config.js` 需**重启 dev server** 才会重扫旧的 `.js`。

## 与 Element Plus 共存

- **在项目配置 / 框架外壳范围内**：表单控件优先用本库（button / select / switch / slider /
  number-field / color-picker），EP 在这一范围只保留非表单用途
  （`el-drawer` / `el-upload` / `ElMessage(Box)` 等）；图标走全局 `<svg-icon>`，不用 `el-icon`。
- **在业务页面范围内**：不强制改造，Element Plus 等既有用法按开发者习惯继续使用。
- 两者可在同一页面共存：本库浮层用 `z-[3000]` 盖过 EP 浮层即可。

## 组件速查

- **g-button**：`variant` = `default | destructive | outline | outline-primary | outline-success | secondary | ghost`；
  `size` = `default | sm | lg | icon`；`outline-primary` / `outline-success` 是带主色 / 成功色描边的次级按钮
  （hover 填充实色、文字转白），调用方用 variant 表达颜色、不要手写 `border-/text-` 覆盖；
  支持 `as` / `asChild`（reka `Primitive` 透传）；`loading` 异步提交时转圈并禁用点击，`disabled` 经原生属性禁用。
- **g-select**：默认走便捷 `:options` 模式（本地算当前文案，规避 reka SelectValue 首屏回填时机问题）；
  `option.value` 支持 `string | number | boolean`，回写**保留原值类型**（内部用 `String(value)` 映射桥接 reka）；
  便捷模式仅必填单选，需清空 / 多选时用 granular 部件（`g-select-trigger / g-select-content / g-select-item`）。
- **g-switch**：关=`control-track`、开=`primary`；纯图形控件，调用方按语义传 `aria-label`。
- **g-slider**：对外是单值 `number`（内部包成数组），支持 `marks`；未填充轨道走 `control-track`，可传 `aria-label`。
- **g-number-field**：数字步进输入；`+`/`-` 按 `step` 增减，手输的值也会吸附到 `step` 的倍数。
- **g-color-picker**：Popover 内组合 reka `ColorArea/ColorSlider/ColorField/ColorSwatchPicker`；
  `alpha` 开透明度通道，`format`=`hex|rgb`，`swatches` 传预设色卡；
  纯图形触发器可传 `title`（hover 提示，兼作可访问名兜底）/ `ariaLabel`；
  对外写回防抖 100ms，卸载时 flush 补发最终值。

## 新增 / 修改组件 checklist

1. 在 `core/componentLibrary/<name>/` 下建 `<Name>.vue` + `index.js`，并从总 `index.js` 导出。
2. 底座用 reka-ui 原语，颜色**只用语义 token**，逐条核对上面「硬约束」。
3. 变体写进 `index.js` 的 cva；class 用 `cn()` 合并、`class` prop 可覆盖；
   `v-model` 用 `defineModel()`（不写 `modelValue`+`emit` 老样板）；纯图形控件提供 `ariaLabel`。
4. 若类名出现在 `.js/.ts`，确认落在已被 `uno.config.js` include 覆盖的目录内。
5. 亮 / 暗两套配色、换主题色都自检一遍；放进抽屉的浮层确认 `z-[3000]` 不被遮挡。
6. 导出后会自动获得 `g-<name>`（kebab-case）全局标签，无需改 `global.js`。
