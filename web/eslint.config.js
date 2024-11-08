import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'

export default [
  {
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
  },

  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/build/*.js', '**/src/assets/**', '**/public/**'],
  },

  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],
]
