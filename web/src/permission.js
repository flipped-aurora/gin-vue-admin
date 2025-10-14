import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'

// 配置 NProgress
Nprogress.configure({
  showSpinner: false,
  ease: 'ease',
  speed: 500
})

// 白名单路由
const WHITE_LIST = ['Login', 'Init']

function isExternalUrl(val) {
  return typeof val === 'string' && /^(https?:)?\/\//.test(val)
}

// 工具函数：统一路径归一化
function normalizeAbsolutePath(p) {
  const s = '/' + String(p || '')
  return s.replace(/\/+/g, '/')
}

function normalizeRelativePath(p) {
  return String(p || '').replace(/^\/+/, '')
}

// 安全注册：仅在路由名未存在时注册顶级路由
function addTopLevelIfAbsent(r) {
  if (!router.hasRoute(r.name)) {
    router.addRoute(r)
  }
}

// 将 n 级菜单扁平化为：
// - 常规：一级 layout + 二级页面组件
// - 若某节点 meta.defaultMenu === true：该节点为顶级（不包裹在 layout 下），其子节点作为该顶级的二级页面组件
function addRouteByChildren(route, segments = [], parentName = null) {
  // 跳过外链根节点
  if (isExternalUrl(route?.path) || isExternalUrl(route?.name) || isExternalUrl(route?.component)) {
    return
  }

  // 顶层 layout 仅用于承载，不参与路径拼接
  if (route?.name === 'layout') {
    route.children?.forEach((child) => addRouteByChildren(child, [], null))
    return
  }

  // 如果标记为 defaultMenu，则该路由应作为顶级路由（不包裹在 layout 下）
  if (route?.meta?.defaultMenu === true && parentName === null) {
    const fullPath = [...segments, route.path].filter(Boolean).join('/')
    const children = route.children ? [...route.children] : []
    const newRoute = { ...route, path: fullPath }
    delete newRoute.children
    delete newRoute.parent
    // 顶级路由使用绝对路径
    newRoute.path = normalizeAbsolutePath(newRoute.path)

    // 若已存在同名路由则整体跳过（之前应已处理过其子节点）
    if (router.hasRoute(newRoute.name)) return
    addTopLevelIfAbsent(newRoute)

    // 若该 defaultMenu 节点仍有子节点，继续递归处理其子节点（挂载到该顶级路由下）
    if (children.length) {
      // 重置片段，使其成为顶级下的二级相对路径
      children.forEach((child) => addRouteByChildren(child, [], newRoute.name))
    }
    return
  }

  // 还有子节点，继续向下收集路径片段（忽略外链片段）
  if (route?.children && route.children.length) {
    const nextSegments = isExternalUrl(route.path) ? segments : [...segments, route.path]
    route.children.forEach((child) => addRouteByChildren(child, nextSegments, parentName))
    return
  }

  // 叶子节点：注册为其父（defaultMenu 顶级或 layout）的二级子路由
  const fullPath = [...segments, route.path].filter(Boolean).join('/')
  const newRoute = { ...route, path: fullPath }
  delete newRoute.children
  delete newRoute.parent
  // 子路由使用相对路径，避免 /layout/layout/... 的问题
  newRoute.path = normalizeRelativePath(newRoute.path)

  if (parentName) {
    // 挂载到 defaultMenu 顶级路由下
    router.addRoute(parentName, newRoute)
  } else {
    // 常规：挂载到 layout 下
    router.addRoute('layout', newRoute)
  }
}

// 处理路由加载
const setupRouter = async (userStore) => {
  try {
    const routerStore = useRouterStore()
    await Promise.all([routerStore.SetAsyncRouter(), userStore.GetUserInfo()])

    // 确保先注册父级 layout
    const baseRouters = routerStore.asyncRouters || []
    const layoutRoute = baseRouters[0]
    if (layoutRoute?.name === 'layout' && !router.hasRoute('layout')) {
      const bareLayout = { ...layoutRoute, children: [] }
      router.addRoute(bareLayout)
    }

    // 扁平化：将 layout.children 与其余顶层异步路由一并作为二级子路由注册到 layout 下
    const toRegister = []
    if (layoutRoute?.children?.length) {
      toRegister.push(...layoutRoute.children)
    }
    if (baseRouters.length > 1) {
      baseRouters.slice(1).forEach((r) => {
        if (r?.name !== 'layout') toRegister.push(r)
      })
    }
  toRegister.forEach((r) => addRouteByChildren(r, [], null))
    return true
  } catch (error) {
    console.error('Setup router failed:', error)
    return false
  }
}

// 移除加载动画
const removeLoading = () => {
  const element = document.getElementById('gva-loading-box')
  element?.remove()
}

// 处理组件缓存
const handleKeepAlive = async (to) => {
  if (!to.matched.some((item) => item.meta.keepAlive)) return

  if (to.matched?.length > 2) {
    for (let i = 1; i < to.matched.length; i++) {
      const element = to.matched[i - 1]

      if (element.name === 'layout') {
        to.matched.splice(i, 1)
        await handleKeepAlive(to)
        continue
      }

      if (typeof element.components.default === 'function') {
        await element.components.default()
        await handleKeepAlive(to)
      }
    }
  }
}

// 处理路由重定向
const handleRedirect = (to, userStore) => {
  if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
    return { ...to, replace: true }
  }
  return { path: '/layout/404' }
}

// 路由守卫
router.beforeEach(async (to, from) => {
  const userStore = useUserStore()
  const routerStore = useRouterStore()
  const token = userStore.token

  Nprogress.start()

  // 处理元数据和缓存
  to.meta.matched = [...to.matched]
  await handleKeepAlive(to)

  // 设置页面标题
  document.title = getPageTitle(to.meta.title, to)

  if (to.meta.client) {
    return true
  }

  // 白名单路由处理
  if (WHITE_LIST.includes(to.name)) {
    if (token) {
      if(!routerStore.asyncRouterFlag){
        await setupRouter(userStore)
      }
      if(userStore.userInfo.authority.defaultRouter){
        return { name: userStore.userInfo.authority.defaultRouter }
      }
    }
    return  true
  }

  // 需要登录的路由处理
  if (token) {
    // 处理需要跳转到首页的情况
    if (sessionStorage.getItem('needToHome') === 'true') {
      sessionStorage.removeItem('needToHome')
      return { path: '/' }
    }

    // 处理异步路由
    if (!routerStore.asyncRouterFlag && !WHITE_LIST.includes(from.name)) {
      const setupSuccess = await setupRouter(userStore)

      if (setupSuccess && userStore.token) {
        return handleRedirect(to, userStore)
      }

      return {
        name: 'Login',
        query: { redirect: to.fullPath }
      }
    }

    return to.matched.length ? true : { path: '/layout/404' }
  }

  // 未登录跳转登录页
  return {
    name: 'Login',
    query: {
      redirect: to.fullPath
    }
  }
})

// 路由加载完成
router.afterEach(() => {
  document.querySelector('.main-cont.main-right')?.scrollTo(0, 0)
  Nprogress.done()
})

// 路由错误处理
router.onError((error) => {
  console.error('Router error:', error)
  Nprogress.remove()
})

// 移除初始加载动画
removeLoading()
