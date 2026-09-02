import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'
import globals from 'globals'

export default [
  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],
  {
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
    languageOptions: {
      ecmaVersion: 'latest',
      sourceType: 'module',
      globals: globals.node
    },
    rules: {
      'vue/max-attributes-per-line': 0,
      'vue/no-v-model-argument': 0,
      'vue/multi-word-component-names': 'off',
      'no-lone-blocks': 'off',
      'no-extend-native': 'off',
      'no-unused-vars': ['error', { argsIgnorePattern: '^_' }]
    }
  },
  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/build/*.js', '**/src/assets/**', '**/public/**']
  }
]
