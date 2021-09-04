const modules = import.meta.glob('../view/**/*.vue')

export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.map(item => {
    if (item.component) {
      item.component = dynamicImport(modules, item.component)
    } else {
      delete item['component']
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}

function dynamicImport(
  dynamicViewsModules,
  component
) {
  const keys = Object.keys(dynamicViewsModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === component
  })
  const matchKey = matchKeys[0]

  return dynamicViewsModules[matchKey]
}
