import { defineConfig } from 'unocss'

export default defineConfig({
  content: {
    pipeline: {
      exclude: ['node_modules', 'dist']
    }
  },
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
})
