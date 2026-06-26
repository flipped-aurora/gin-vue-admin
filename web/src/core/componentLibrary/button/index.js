import { cva } from 'class-variance-authority'
import { FOCUS_RING } from '../utils'

export { default as Button } from './Button.vue'

/**
 * 变体 / 尺寸的可选值集中导出：作为单一事实源，既驱动下面 cva 的语义，
 * 也供 Button.vue 的 prop `validator` 复用做白名单校验（JS 栈下等价于 TS 的字面量联合约束）。
 * 新增 / 改名 variant 或 size 时，这里与下方 cva 的 variants 同步增改。
 */
export const BUTTON_VARIANTS = [
  'default',
  'destructive',
  'outline',
  'outline-primary',
  'outline-success',
  'secondary',
  'ghost'
]
export const BUTTON_SIZES = ['default', 'sm', 'lg', 'icon']

/**
 * 按钮样式变体。
 *
 * 所有颜色都走项目既有的主题 token（--primary-color / --error-color / --container-bg-color /
 * --border-color / --muted-color ...），因此天然跟随 themeStore 的换肤与暗色模式，
 * 调用方无需再内联 `:style="{ backgroundColor: settings.themeColor }"`。
 *
 * 注意：刻意不使用 `bg-primary/10` 这类对 CSS 变量 token 取透明度的写法，
 * 改用语义 token（如 hover:bg-muted）保证在亮/暗两套配色下都成立。
 */
export const buttonVariants = cva(
  // 项目未启用「button 背景重置」式 preflight，原生 <button> 会带浏览器默认底色（#efefef）。
  // 基类先用 bg-transparent 兜底，各 variant 的显式底色会经 twMerge 覆盖它；
  // 唯独 ghost 没有底色，正好保持透明（hover 再上 muted）。
  `appearance-none inline-flex items-center justify-center gap-1.5 whitespace-nowrap rounded-md bg-transparent text-sm font-medium transition-all duration-150 ${FOCUS_RING} disabled:pointer-events-none disabled:opacity-50`,
  {
    variants: {
      variant: {
        default: 'bg-primary text-white shadow-sm hover:bg-primary-600',
        destructive: 'bg-error text-white shadow-sm hover:bg-error-600',
        outline: 'border border-border bg-container text-base-text hover:bg-muted',
        // 带色描边的次级按钮：常态描边 + 文字着色，hover 填充实色、文字转白；
        // 调用方用 variant 表达颜色，不再手写 border-/text- 覆盖。
        'outline-primary': 'border border-primary bg-container text-primary hover:bg-primary hover:text-white',
        'outline-success': 'border border-success bg-container text-success hover:bg-success hover:text-white',
        // hover 走语义 token 换色（与其余 variant 同机制），不再用 opacity
        secondary: 'bg-muted text-base-text hover:bg-control-track',
        ghost: 'text-base-text hover:bg-muted'
      },
      size: {
        default: 'h-9 px-4 py-2',
        sm: 'h-8 rounded-md px-3 text-xs',
        lg: 'h-10 rounded-md px-6 text-base',
        icon: 'h-8 w-8'
      }
    },
    defaultVariants: {
      variant: 'default',
      size: 'default'
    }
  }
)
