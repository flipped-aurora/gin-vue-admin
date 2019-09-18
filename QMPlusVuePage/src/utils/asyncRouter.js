const _import = require('./_import') //获取组件的方法
export const asyncRouterHandle = (asyncRouter) => {
    asyncRouter.map(item => {
        if (item.component) {
            item.component = _import(item.component)
        } else {
            delete item['component']
        }
        if (item.children) {
            asyncRouterHandle(item.children)
        }
    })
}