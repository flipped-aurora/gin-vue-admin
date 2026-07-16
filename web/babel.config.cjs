// Babel 配置
//
// 用 .cjs 而非 .js：本项目 package.json 为 "type": "module"，.js 会被当作 ESM，
// 而 @babel/eslint-parser 需要「同步」读取 babel 配置（ESLint 是同步解析的），
// ESM 配置无法同步加载。.cjs 强制 CommonJS，保证 eslint 能正常读取本文件。
//
// 目前主要供 @babel/eslint-parser 解析 .js / .vue<script> 使用；
// 如日后需要 babel 做实际转译，可直接复用这份 preset。
module.exports = {
  presets: [
    ['@babel/preset-env', { targets: 'defaults' }]
  ]
}
