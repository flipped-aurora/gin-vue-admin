import { defineConfig } from 'unocss'
import presetWind3 from '@unocss/preset-wind3';

export default defineConfig({
  theme: {
    backgroundColor: {
      main: '#F5F5F5'
    },
    textColor: {
      active: 'var(--el-color-primary)'
    },
    boxShadowColor: {
      active: 'var(--el-color-primary)'
    },
    borderColor: {
      'table-border': 'var(--el-border-color-lighter)'
    }
  },
  presets: [
    presetWind3({ dark: 'class' })
  ],
})
