import { asyncRouterHandle } from '@/utils/asyncRouter';

import { asyncMenu } from '@/api/menu'

export const router = {
    namespaced: true,
    state: {
        asyncRouters: []
    },
    mutations: {
        // 设置动态路由
        setAsyncRouter(state, asyncRouters) {
            state.asyncRouters = asyncRouters
        }
    },
    actions: {
        // 从后台获取动态路由
        async SetAsyncRouter({ commit }) {
            const asyncRouterRes = await asyncMenu()
            const asyncRouter = asyncRouterRes.data.menus
            asyncRouter.push({
                path: '*',
                redirect: '/404'

            })
            asyncRouterHandle(asyncRouter) // 等待动态使用 VUEX持久化 会将其放入 vuex并且动态生成左侧列表
            commit('setAsyncRouter', asyncRouter)
        }
    },
    getters: {
        // 获取动态路由
        asyncRouters(state) {
            return state.asyncRouters
        }
    }
}