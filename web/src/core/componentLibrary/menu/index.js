// 只导出 Menu（PascalCase）→ 由 global.js 自动注册为 <g-menu>。
// MenuItem / MenuFlyout / Horizontal* 保持私有，避免泄漏为 g- 全局标签。
export { default as Menu } from './Menu.vue'
export { MENU_THEMES } from './shared'
