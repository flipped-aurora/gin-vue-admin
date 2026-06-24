import { defineConfig } from '@unocss/vite';
import presetWind3 from '@unocss/preset-wind3';
import transformerDirectives from '@unocss/transformer-directives'
import { unoTheme } from './src/theme/adapters/uno'

export default defineConfig({
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
