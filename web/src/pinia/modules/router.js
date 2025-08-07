import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/utils/bus.js'
import { asyncMenu } from '@/api/menu'
import { defineStore } from 'pinia'
import { ref, watchEffect } from 'vue'
import i18n from '@/i18n' // added by mohamed hassan to multilangauge
import pathInfo from '@/pathInfo.json'
import {useRoute} from "vue-router";

const notLayoutRouterArr = []
const keepAliveRoutersArr = []
const nameMap = {}

const formatRouter = (routes, routeMap, parent) => {
  routes &&
    routes.forEach((item) => {
      item.parent = parent
      item.meta.btns = item.btns
      item.meta.hidden = item.hidden
      if (item.meta.defaultMenu === true) {
        if (!parent) {
          item = { ...item, path: `/${item.path}` }
          notLayoutRouterArr.push(item)
        }
      }
      routeMap[item.name] = item
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap, item)
      }
    })
}

const KeepAliveFilter = (routes) => {
  routes &&
    routes.forEach((item) => {
      // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
      if (
        (item.children && item.children.some((ch) => ch.meta.keepAlive)) ||
        item.meta.keepAlive
      ) {
        const path = item.meta.path
        keepAliveRoutersArr.push(pathInfo[path])
        nameMap[item.name] = pathInfo[path]
      }
      if (item.children && item.children.length > 0) {
        KeepAliveFilter(item.children)
      }
    })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([])
  const asyncRouterFlag = ref(0)
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    
    // 1. 首先添加原有的keepAlive配置
    keepArrTemp.push(...keepAliveRoutersArr)
    
    history.forEach((item) => {
      // 2. 为所有history中的路由强制启用keep-alive
      // 通过routeMap获取路由信息，然后通过pathInfo获取组件名
      const routeInfo = routeMap[item.name]
      if (routeInfo && routeInfo.meta && routeInfo.meta.path) {
        const componentName = pathInfo[routeInfo.meta.path]
        if (componentName) {
          keepArrTemp.push(componentName)
        }
      }
      
      // 3. 如果子路由在tabs中打开，父路由也需要keepAlive
      if (nameMap[item.name]) {
        keepArrTemp.push(nameMap[item.name])
      }
    })
    
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }


  const route = useRoute()

  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])

  const topMenu = ref([])

  const leftMenu = ref([])

  const menuMap = {}

  const topActive = ref('')

  const setLeftMenu = (name) => {
    sessionStorage.setItem('topActive', name)
    topActive.value = name
    leftMenu.value = []
    if (menuMap[name]?.children) {
      leftMenu.value = menuMap[name].children
    }
    return menuMap[name]?.children
  }

  const findTopActive = (menuMap, routeName) => {
    for (let topName in menuMap) {
      const topItem = menuMap[topName];
      if (topItem.children?.some(item => item.name === routeName)) {
        return topName;
      }
      const foundName = findTopActive(topItem.children || {}, routeName);
      if (foundName) {
        return topName;
      }
    }
    return null;
  };

  watchEffect(() => {
    let topActive = sessionStorage.getItem('topActive')
    // 初始化菜单内容，防止重复添加
    topMenu.value = [];
    asyncRouters.value[0]?.children.forEach((item) => {
      if (item.hidden) return
      menuMap[item.name] = item
      topMenu.value.push({ ...item, children: [] })
    })
    if (!topActive || topActive === 'undefined' || topActive === 'null') {
      topActive = findTopActive(menuMap, route.name);
    }
    setLeftMenu(topActive)
  })

  const routeMap = {}
  // 从后台获取动态路由
  const SetAsyncRouter = async () => {
    asyncRouterFlag.value++
    const baseRouter = [
      {
        path: '/layout',
        name: 'layout',
        component: 'view/layout/index.vue',
        meta: {
          title: i18n.global.t('pinia.modules.router.bottomLayerLayout')
        },
        children: []
      }
    ]
    const asyncRouterRes = await asyncMenu()
    const asyncRouter = asyncRouterRes.data.menus
    asyncRouter &&
      asyncRouter.push({
        path: 'reload',
        name: 'Reload',
        hidden: true,
        meta: {
          title: '',
          closeTab: true
        },
        component: 'view/error/reload.vue'
      })
    formatRouter(asyncRouter, routeMap)
    baseRouter[0].children = asyncRouter
    if (notLayoutRouterArr.length !== 0) {
      baseRouter.push(...notLayoutRouterArr)
    }
    asyncRouterHandle(baseRouter)
    KeepAliveFilter(asyncRouter)
    asyncRouters.value = baseRouter
    return true
  }

  return {
    topActive,
    setLeftMenu,
    topMenu,
    leftMenu,
    asyncRouters,
    keepAliveRouters,
    asyncRouterFlag,
    SetAsyncRouter,
    routeMap
  }
})
