export { default as DropdownMenu } from './DropdownMenu.vue'
export { default as DropdownMenuContent } from './DropdownMenuContent.vue'
export { default as DropdownMenuItem } from './DropdownMenuItem.vue'
export { default as DropdownMenuLabel } from './DropdownMenuLabel.vue'
export { default as DropdownMenuSeparator } from './DropdownMenuSeparator.vue'

/**
 * 触发方式的可选值集中导出：作为单一事实源，供 DropdownMenu.vue 的 prop `validator`
 * 复用做白名单校验，也便于调用方按需引用。
 * - click：点击触发器展开 / 收起（reka 默认交互）
 * - hover：悬停触发器展开、移出（触发器与面板都离开）后延迟收起，对齐顶栏工具区的悬停习惯
 */
export const DROPDOWN_MENU_TRIGGERS = ['click', 'hover']
