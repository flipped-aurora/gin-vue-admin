import * as path from 'path'
export default function fullImportPlugin() {
  let config
  return {
    name: 'fullImportElementPlus',
    async configResolved(conf) {
      config = conf
    },
    transform(code, id) {
      const sourcePath = path.join(config.root, 'src/main.js').split(path.sep).join('/')
      const targetPath = id.split(path.sep).join('/')
      if (sourcePath === targetPath) {
        const name = 'ElementPlus'
        // 引入 ElementPlus 和 样式
        code = code.replace(`import { createApp } from 'vue'`, ($1) => $1 + `\nimport ${name} from 'element-plus'`)
        code = code.replace(`import './style/element_visiable.scss'`, ($1) => $1 + `\nimport 'element-plus/theme-chalk/src/index.scss'`)
        code = code.replace('.mount(', ($1) => `.use(${name})` + $1)
        return code
      }
      return code
    }
  }
}
