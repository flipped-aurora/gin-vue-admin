/**
 * ESLint 扁平配置（flat config）
 *
 * 解析链：.js/.jsx 交给 @babel/eslint-parser（读 babel.config.cjs）；.vue 顶层由
 * vue-eslint-parser 解析 <template>，其 <script> 再委托给 babel。统一走 babel 而非默认
 * espree，是为了和项目 babel 工具链保持一致，日后引入 espree 尚不支持的语法也无需换解析器。
 */
import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'
import globals from 'globals'
import babelParser from '@babel/eslint-parser'
import vueParser from 'vue-eslint-parser'

// .js 与 .vue<script> 复用的一份 babel 解析选项，改一处即两处生效，避免漂移
const BABEL_PARSER_OPTIONS = {
  requireConfigFile: true,     // 必须读到 babel.config.cjs（缺失即报错，杜绝悄悄回退到无配置）
  ecmaFeatures: { jsx: true }  // 放开 JSX，渲染函数 / 个别组件偶有用到
}

// 不参与检查的构建产物与静态资源
const IGNORE_GLOBS = ['**/dist/**', '**/build/*.js', '**/src/assets/**', '**/public/**']

export default [
  // 基线规则集：ESLint 官方推荐 + Vue3 基础必备，后面的块只在此之上做增量覆盖
  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],

  // 纯脚本文件：整段交给 babel 解析
  {
    name: 'app/js-babel',
    files: ['**/*.{js,mjs,jsx}'],
    languageOptions: {
      ecmaVersion: 'latest',
      sourceType: 'module',
      parser: babelParser,
      parserOptions: BABEL_PARSER_OPTIONS,
      globals: globals.node
    }
  },

  // 单文件组件：顶层用 vue 解析器吃 <template>，<script> 再转交 babel
  {
    name: 'app/vue-babel',
    files: ['**/*.vue'],
    languageOptions: {
      ecmaVersion: 'latest',
      sourceType: 'module',
      parser: vueParser,
      parserOptions: { parser: babelParser, ...BABEL_PARSER_OPTIONS },
      globals: globals.node
    }
  },

  // CommonJS 配置文件（如 babel.config.cjs）：走 Node 环境，别把 module/require 判成未定义
  {
    name: 'app/commonjs',
    files: ['**/*.cjs'],
    languageOptions: {
      sourceType: 'commonjs',
      globals: globals.node
    }
  },

  // 规则微调：只列与本项目习惯相冲、需显式放宽/收紧的项，其余沿用基线
  {
    name: 'app/rules',
    files: ['**/*.{js,mjs,jsx,vue}'],
    rules: {
      'vue/max-attributes-per-line': 0,             // 每行属性数不限，排版交给人工
      'vue/no-v-model-argument': 0,                 // 允许 v-model:arg（EP 等组件写法）
      'vue/multi-word-component-names': 'off',      // 允许单词组件名（大量既有页面如此）
      'no-lone-blocks': 'off',                      // 允许独立块级作用域
      'no-extend-native': 'off',                    // 允许扩展原生原型（个别工具函数依赖）
      'no-unused-vars': ['error', { argsIgnorePattern: '^_' }] // 未用变量报错；下划线前缀参数放行
    }
  },

  // 忽略目录
  {
    name: 'app/files-to-ignore',
    ignores: IGNORE_GLOBS
  }
]
