import { cva } from 'class-variance-authority'

// 菜单项配色变体：三种菜单风格 × 选中态 × 角色（普通项 / 分组标题）。
// 深浅由侧栏容器的语义 token 决定（见 useSidebarTheme 的 .gva-sider-dark 作用域）。
// 配色只用语义 token + 主色弱底 rgb(var(--primary-color)/α)。
export const menuItemVariants = cva(
  'group/mi relative flex w-full items-center gap-2 appearance-none bg-transparent cursor-pointer select-none whitespace-nowrap text-[14px] leading-none transition-colors outline-none',
  {
    variants: {
      theme: {
        design: 'rounded-none',
        light: 'rounded-[var(--gva-radius)]',
        group: 'rounded-lg'
      },
      active: { true: '', false: '' },
      role: { item: '', header: '' }
    },
    compoundVariants: [
      // design：全宽直角 + 左侧主色竖条 + 主色弱底（保留原风格观感）
      {
        theme: 'design',
        active: true,
        class:
          "text-active font-medium bg-[rgb(var(--primary-color)/0.12)] before:absolute before:inset-y-0 before:left-0 before:w-[3px] before:rounded-r before:bg-primary before:content-['']"
      },
      {
        theme: 'design',
        active: false,
        class: 'text-base-text hover:bg-[rgb(var(--primary-color)/0.08)]'
      },
      // light：白底 + 主色实底选中
      { theme: 'light', active: true, class: 'bg-primary text-white' },
      { theme: 'light', active: false, class: 'text-base-text hover:bg-muted' },
      // group：分组标题 = 加粗基础色；子项 = 弱色；选中 = 浅灰药丸
      { theme: 'group', role: 'header', class: 'text-base-text font-semibold' },
      {
        theme: 'group',
        role: 'item',
        active: true,
        class: 'bg-muted text-base-text font-medium'
      },
      {
        theme: 'group',
        role: 'item',
        active: false,
        class: 'text-muted-foreground hover:bg-muted hover:text-base-text'
      }
    ],
    defaultVariants: { theme: 'design', active: false, role: 'item' }
  }
)

// 侧栏容器底色：统一跟随 container（随全局明暗 / .gva-sider-dark 作用域自适应）。
export const SIDEBAR_SURFACE = 'bg-container text-base-text'

// 折叠图标栏 / 飞出触发的基础按钮样式（图标居中、hover 底；focus ring 由调用方叠加）。
export const MENU_ICON_BUTTON =
  'relative flex w-full appearance-none items-center justify-center rounded-lg bg-transparent text-base-text transition-colors hover:bg-muted'
