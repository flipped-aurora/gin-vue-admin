# 前端约束

## 基础约束

- HTTP 请求统一走 `@/utils/request.js`
- 全局状态统一使用 Pinia
- 路由要带完整的元信息并考虑权限
- 需要动态路由时，沿用项目现有异步路由处理方式

## 命名规范

- 文件名使用 `kebab-case`
- 组件名使用 `PascalCase`
- 变量名使用 `camelCase`
- 常量名使用 `UPPER_SNAKE_CASE`

## 注释规范

- API 封装尽量补全 JSDoc
- 复杂组件要写清功能说明
- 关键业务逻辑可以补充必要的行内注释

## 组件写法规范

- **`v-model` 一律用 `defineModel()`**（项目 Vue ≥ 3.4），**禁止**再写
  `props.modelValue` + `defineEmits(['update:modelValue'])` + `@update:model-value` 这套老样板：
  - 纯透传：`const modelValue = defineModel({ type: X, default: ... })`，模板上对子组件直接 `v-model="modelValue"`。
  - 需要把值变换后再交给子组件（如对外单值、对内数组）：用**可写 `computed`** 作桥，
    `get` 返回变换后的值、`set` 里写回 `modelValue.value`。
  - 需要在写回前做防抖 / 拦截 / 类型还原（如取色器防抖、Select 原值映射）：仍用 `defineModel` 的 ref，
    在回调里 `modelValue.value = ...`，不要退回 `emit`。
  - 多个 v-model 用具名 `defineModel('xxx', { ... })`。
- 对外契约不变：`defineModel` 仍等价声明 `modelValue` prop 并 emit `update:modelValue`，
  调用方 `v-model` / `:model-value + @update:model-value` 都照常工作。
- 有限枚举 prop（如 `variant` / `size` / `format`）补 `validator` 做白名单校验；
  结构化 prop（`options` / `marks` 等）用 JSDoc `@typedef` 写清形状。
- 纯图形交互控件（开关、取色器、滑块等）要可传入 `ariaLabel` 等可访问名，调用方按语义补中文 `aria-label`。

## 样式规范

- 优先使用 UnoCSS
- 项目配置 / 框架外壳（header、menu、布局、系统设置抽屉等）的 UI 必须用组件库
  `@/core/componentLibrary`（全局 `<g-button />` 等），并遵守其主题 token 规范；
  详见 `frontend-backend/component-library.md`
- 业务组件 / 业务页面不强制用该组件库，按使用习惯可继续用 Element Plus 等，但也需要符合 gva-theme 主题 token 规范：
  - 避免手写 `text-/bg-/border-` 等原子类覆盖主题 token
  - 避免手写 `color: #xxxxxx` 等硬编码颜色
  - 避免手写 `font-size: xxpx` 等硬编码字号
- 避免内联样式
- 主题相关能力优先通过 CSS 变量控制
- 反例：用 `<div class="scenario-bar">` + scoped `.scenario-bar { display: flex; gap: 8px; align-items: center }`；正例：直接 `<div class="flex items-center gap-2">`

## 图标规范

- 图标统一用全局 `<svg-icon>`（`@/components/svgIcon/svgIcon.vue`），
  **不要用 `el-icon` + `@element-plus/icons-vue`**，也不要在模板里手写裸 `<svg><path/></svg>`。
- 在线图标传 Iconify 名：`<svg-icon icon="lucide:check" />`；
  本地 symbol 传 `local-icon`：`<svg-icon local-icon="close" />`（对应 `src/assets/icons/*.svg`）。
- 图标集优先 `lucide`（与基础组件库内置图标一致），**不要用 `ep:` 等 Element Plus 图标集**。

## 性能规范

- 路由优先懒加载
- 大列表考虑虚拟滚动或等价优化
- 合理复用缓存
- 图片上传与展示考虑压缩与加载成本
