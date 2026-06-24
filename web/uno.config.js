import { defineConfig } from '@unocss/vite';
import presetWind3 from '@unocss/preset-wind3';
import transformerDirectives from '@unocss/transformer-directives'
import { unoTheme } from './src/theme/adapters/uno'

export default defineConfig({
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
