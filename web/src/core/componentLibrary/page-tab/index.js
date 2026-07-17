// 只导出 PageTab（PascalCase）→ 由 global.js 自动注册为 <g-page-tab>。
// button / chrome / slider 三个模式子组件保持私有（不从这里导出），避免泄漏为 g- 全局标签。
export { default as PageTab } from './PageTab.vue'
export { PAGE_TAB_MODES } from './shared'
