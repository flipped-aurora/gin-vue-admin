import config from './config'
import { h } from 'vue'

// 统一导入el-icon图标
import * as ElIconModules from '@element-plus/icons-vue'
// 导入转换图标名称的函数

const createIconComponent = (svgContent) => ({
  name: 'SvgIcon',
  props: {
    iconClass: {
      type: String,
      default: '',
    },
    className: {
      type: String,
      default: '',
    },
  },
  render() {
    const { className } = this
    return h('span', {
      class: className,
      ariaHidden: true,
      innerHTML: svgContent,
    })
  },
})

const registerIcons = async(app) => {
  const iconModules = import.meta.glob('@/assets/icons/**/*.svg')
  for (const path in iconModules) {
    const response = await fetch(path)
    const svgContent = await response.text()
    const iconName = path.split('/').pop().replace(/\.svg$/, '')
    // 如果iconName带空格则不加入到图标库中并且提示名称不合法
    if (iconName.indexOf(' ') !== -1) {
      console.error(`icon ${iconName}.svg includes whitespace`)
      continue
    }
    const iconComponent = createIconComponent(svgContent)
    config.logs.push({
      'key': iconName,
      'label': iconName,
    })
    app.component(iconName, iconComponent)
  }
}

export const register = (app) => {
  // 统一注册el-icon图标
  for (const iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
  }
  registerIcons(app)
  app.config.globalProperties.$GIN_VUE_ADMIN = config
}
