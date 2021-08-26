import { asyncRouterHandle } from '@/utils/asyncRouter'

import { asyncMenu } from '@/api/menu'

const routerList = []

const formatRouter = (routes) => {
  routes && routes.map(item => {
    if ((!item.children || item.children.every(ch => ch.hidden)) && item.name !== '404' && !item.hidden) {
      routerList.push({ label: item.meta.title, value: item.name })
    }
    item.meta.hidden = item.hidden
    if (item.children && item.children.length > 0) {
      formatRouter(item.children)
    }
  })
}

export const router = {
  namespaced: true,
  state: {
    asyncRouters: [],
    routerList: routerList,
  },
  mutations: {
    setRouterList(state, routerList) {
      state.routerList = routerList
    },
    // 设置动态路由
    setAsyncRouter(state, asyncRouters) {
      state.asyncRouters = asyncRouters
    }
  },
  actions: {
    // 从后台获取动态路由
    async SetAsyncRouter({ commit }) {
      const baseRouter = [{
        path: '/layout',
        name: 'layout',
        component: 'view/layout/index.vue',
        meta: {
          title: '底层layout'
        },
        children: []
      }]
      const asyncRouterRes = await asyncMenu()
      const asyncRouter = asyncRouterRes.data.menus
      asyncRouter.push({
        path: '404',
        name: '404',
        hidden: true,
        meta: {
          title: '迷路了*。*',
        },
        component: 'view/error/index.vue'
      })
      formatRouter(asyncRouter)
      baseRouter[0].children = asyncRouter
      baseRouter.push({
        path: '/:catchAll(.*)',
        redirect: '/layout/404'

      })
      asyncRouterHandle(baseRouter)
      commit('setAsyncRouter', baseRouter)
      commit('setRouterList', routerList)
      return true
    }
  },
  getters: {
    // 获取动态路由
    asyncRouters(state) {
      return state.asyncRouters
    },
    routerList(state) {
      return state.routerList
    }
  }
}
