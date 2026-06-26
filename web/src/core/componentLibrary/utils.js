import { clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

/**
 * 合并 class：clsx 处理条件类，tailwind-merge 消解冲突的原子类（presetWind3 与 Tailwind v3 同语法）。
 * 这是 reka-ui / shadcn-vue 体系约定俗成的 `cn()` 基础工具，后续基础组件统一用它拼接 class。
 * @param {...any} inputs
 * @returns {string}
 */
export function cn(...inputs) {
  return twMerge(clsx(inputs))
}

// 焦点态单一真源，全组件统一引用，避免各处手写导致漂移：
// - FOCUS_RING：控件本体可聚焦（focus-visible，仅键盘聚焦显形）。
// - FOCUS_RING_WITHIN：包裹型容器（内部 input 才真正聚焦）用 focus-within。
export const FOCUS_RING = 'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-300'
export const FOCUS_RING_WITHIN = 'focus-within:ring-2 focus-within:ring-primary-300'
