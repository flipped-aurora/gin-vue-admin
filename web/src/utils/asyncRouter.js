const viewModules = import.meta.glob('../view/**/*.vue')
const pluginModules = import.meta.glob('../plugin/**/*.vue')

export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.forEach((item) => {
    if (item.component && typeof item.component === 'string') {
      item.meta.path = '/src/' + item.component
      if (item.component.split('/')[0] === 'view') {
        item.component = dynamicImport(viewModules, item.component)
      } else if (item.component.split('/')[0] === 'plugin') {
        item.component = dynamicImport(pluginModules, item.component)
      }
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}

function dynamicImport(dynamicViewsModules, component) {
  const keys = Object.keys(dynamicViewsModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === component
  })
  const matchKey = matchKeys[0]
  const matched = dynamicViewsModules[matchKey]
  // 找不到组件时返回占位组件，避免 component 为 undefined 导致路由白屏
  // （常见于数据库菜单记录指向已删除的组件，如历史遗留的 aiWorkflow）
  if (!matched) {
    console.warn(`[asyncRouter] 未找到组件: ${component}，已使用占位组件代替`)
    return {
      name: 'MissingComponentPlaceholder',
      render: () => null
    }
  }
  return matched
}
