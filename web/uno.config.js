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
    'gva-surface': 'p-4 my-2 bg-container text-base-text'
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
