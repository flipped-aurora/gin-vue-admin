export { default as ColorPicker } from './ColorPicker.vue'

// 有限枚举 format 的可选值集中导出，作为单一事实源，供 prop validator 复用（对齐 Button 的 BUTTON_VARIANTS）
export const COLOR_FORMATS = ['hex', 'rgb']
