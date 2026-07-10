import config from './config'
import { h } from 'vue'

// 统一导入el-icon图标
import * as ElIconModules from '@element-plus/icons-vue'
import svgIcon from '@/components/svgIcon/svgIcon.vue'
// 基础组件库（reka-ui 底座，src/core/componentLibrary）：全局以 g- 前缀（kebab-case）注册，
// 全站可直接 <g-button /> / <g-select /> 使用，无需逐文件 import。
// 规范见 aiDoc/frontend-backend/component-library.md
import * as ComponentLibrary from '@/core/componentLibrary'
// 导入转换图标名称的函数

const createIconComponent = (name) => ({
  name: 'SvgIcon',
  render() {
    return h(svgIcon, {
      localIcon: name
    })
  }
})

const registerIcons = async (app) => {
  const iconModules = import.meta.glob('@/assets/icons/**/*.svg') // 系统目录 svg 图标
  const pluginIconModules = import.meta.glob(
    '@/plugin/**/assets/icons/**/*.svg'
  ) // 插件目录 svg 图标
  const mergedIconModules = Object.assign({}, iconModules, pluginIconModules) // 合并所有 svg 图标
  let allKeys = []
  for (const path in mergedIconModules) {
    let pluginName = ''
    if (path.startsWith('/src/plugin/')) {
      pluginName = `${path.split('/')[3]}-`
    }
    const iconName = path
      .split('/')
      .pop()
      .replace(/\.svg$/, '')
    // 如果iconName带空格则不加入到图标库中并且提示名称不合法
    if (iconName.indexOf(' ') !== -1) {
      console.error(`icon ${iconName}.svg includes whitespace in ${path}`)
      continue
    }
    const key = `${pluginName}${iconName}`
    const iconComponent = createIconComponent(key)
    config.logs.push({
      key: key,
      label: key
    })
    app.component(key, iconComponent)

    // 开发模式下列出所有 svg 图标，方便开发者直接查找复制使用
    allKeys.push(key)
  }

  import.meta.env.MODE == 'development' &&
    console.log(`所有可用的本地图标: ${allKeys.join(', ')}`)
}

// 把基础组件库的每个组件以 g- 前缀注册为全局组件（Button -> g-button、NumberField -> g-number-field ...）。
// 用短横线（kebab-case）命名，和项目里 el-button 等用法保持统一。
const toKebabCase = (name) => name.replace(/([a-z0-9])([A-Z])/g, '$1-$2').toLowerCase()

// barrel 里会混入「非本库自有组件」的 re-export（如 reka-ui 的 SelectValue，仅供 granular 模式按需 import），
// 它们是组件对象、会通过下面 typeof 判断，但不应获得全局 g- 标签。用显式名单挡掉，
// 既修掉 g-select-value 这类无人消费的命名空间泄漏，又保持「本库自有组件导出即自动注册」的设计不变。
// 枚举数组导出（BUTTON_VARIANTS / MENU_THEMES 等）typeof 也是 object，由 Array.isArray 统一跳过，无需登记名单。
const NON_GLOBAL_EXPORTS = new Set(['SelectValue'])

const registerComponentLibrary = (app) => {
  for (const [name, component] of Object.entries(ComponentLibrary)) {
    if (NON_GLOBAL_EXPORTS.has(name)) continue
    if (component && typeof component === 'object' && !Array.isArray(component)) {
      app.component(`g-${toKebabCase(name)}`, component)
    }
  }
}

export const register = (app) => {
  // 统一注册el-icon图标
  for (const iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
  }
  app.component('SvgIcon', svgIcon)
  registerIcons(app)
  // 基础组件库全局注册（g- 前缀）
  registerComponentLibrary(app)
  app.config.globalProperties.$GIN_VUE_ADMIN = config
}
