import { defineConfig } from '@unocss/vite';
import presetWind3 from '@unocss/preset-wind3';
import transformerDirectives from '@unocss/transformer-directives'
import { unoTheme } from './src/theme/adapters/uno'

export default defineConfig({
  // UnoCSS 默认扫描管线只覆盖 .vue/.jsx/.tsx 等模板文件，不扫描纯 .js/.ts。
  // 基础组件库（reka-ui 底座，位于 src/core/componentLibrary）刻意把类名维护在 .js 的 cva 里
  //（如 button 的 buttonVariants），不显式纳入扫描的话，只出现在这些 .js 里的原子类
  //（bg-primary / bg-error / hover:bg-primary-600 ...）不会被生成，按钮就会丢失主题色。
  // 这里在保留默认文件类型的前提下，额外把 core/componentLibrary 下的 .js/.ts 纳入扫描。
  // 详见 aiDoc/frontend-backend/component-library.md。
  content: {
    pipeline: {
      include: [
        /\.(vue|svelte|[jt]sx|mdx?|astro|elm|php|phtml|html)($|\?)/,
        /[\\/]core[\\/]componentLibrary[\\/].*\.[jt]s($|\?)/
      ]
    }
  },
  shortcuts: {
    // 业务卡片盒子的统一表面：容器底色（白/暗自适应）+ 主文本色 + 内边距，
    // 替代散落的 bg-white text-slate-700 dark:bg-slate-900 dark:text-slate-400 串
    'gva-surface': 'p-4 my-2 bg-container text-base-text',
    // header 工具图标点击区域：扁平 36px，与右侧用户胶囊（30px 头像 + 内边距≈36px）等高，视觉权重对齐；圆角跟随全局 --gva-radius；
    // hover 出现清晰的中性灰底——用半透明叠加（黑/白各 10%）而非 bg-muted：
    // muted 亮色值≈slate-50，与白色顶栏几乎同色导致 hover 底看不出来；叠加层与顶栏底色无关，深浅两态都清晰可见。
    'gva-tool-btn': 'inline-flex items-center justify-center w-9 h-9 rounded-[var(--gva-radius)] text-muted-foreground hover:text-base-text hover:bg-black/10 dark:hover:bg-white/10 cursor-pointer outline-none transition-colors',
    // header 用户菜单项：图标+文案，hover/键盘聚焦/子菜单展开高亮。
    // 高亮底与工具栏按钮一致——中性灰叠加（黑/白各 10%）而非 bg-muted：muted 亮色值≈slate-50 带冷蓝调，观感偏浅蓝且与白色面板对比过弱。
    'gva-menu-item': 'flex items-center gap-2 w-full px-2.5 py-2 rounded-lg text-[14px] text-base-text cursor-pointer outline-none transition-colors hover:bg-black/10 dark:hover:bg-white/10 data-[highlighted]:bg-black/10 dark:data-[highlighted]:bg-white/10 data-[state=open]:bg-black/10 dark:data-[state=open]:bg-white/10',
    // header 用户菜单项 - 危险态（退出登录，红色语义）
    'gva-menu-item-danger': 'text-error hover:bg-error/10 data-[highlighted]:bg-error/10',
    // header 用户菜单分隔线：极细 0.5px，组间留白（边框色 token 自适应深色）
    'gva-menu-sep': 'h-px my-1 mx-2 bg-border opacity-50',
    // header 用户身份信息块（不可交互 label）
    'gva-menu-label': 'block px-2.5 py-2 cursor-default outline-none'
  },
  theme: {
    ...unoTheme,
    backgroundColor: {
      main: 'rgb(var(--layout-bg-color))'
    },
    textColor: {
      active: 'rgb(var(--primary-color))'
    },
    boxShadowColor: {
      active: 'rgb(var(--primary-color))'
    },
    borderColor: {
      'table-border': 'var(--el-border-color-lighter)'
    }
  },
  presets: [
    presetWind3({ dark: 'class' })
  ],
  transformers: [
    transformerDirectives(),
  ],
})
