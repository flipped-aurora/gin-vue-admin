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

## 样式规范

- **优先使用 UnoCSS 原子类**（`web/uno.config.js`，presetWind3）：布局、间距、排版、尺寸、颜色等能用原子类表达的，不再新写自定义 class
- 仅在原子类难以表达时才写 `<style scoped>`，典型场景：
  - `:deep()` 覆盖 Element Plus 内部样式
  - 伪类 / 伪元素、复杂选择器
  - 需要复用的成组复杂样式
- 避免内联 `style`；动态样式用绑定的原子 class 或 CSS 变量
- 遵循 Element Plus 现有设计风格
- 主题相关能力优先通过 CSS 变量控制
- 反例：用 `<div class="scenario-bar">` + scoped `.scenario-bar { display: flex; gap: 8px; align-items: center }`；正例：直接 `<div class="flex items-center gap-2">`

## 性能规范

- 路由优先懒加载
- 大列表考虑虚拟滚动或等价优化
- 合理复用缓存
- 图片上传与展示考虑压缩与加载成本
