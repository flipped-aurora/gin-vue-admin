export { default as Input } from './Input.vue'

/**
 * 尺寸 / 变体可选值集中导出：作为单一事实源，供 Input.vue 的 prop `validator` 复用做白名单校验。
 * - size：default(h-9) / lg(h-11)
 * - variant：default(盒式) / underline(下划线式)
 */
export const INPUT_SIZES = ['default', 'lg']
export const INPUT_VARIANTS = ['default', 'underline']
