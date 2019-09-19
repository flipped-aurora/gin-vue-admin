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
            const baseRouter = [{
                path: '/layout',
                name: 'layout',
                component: "view/layout/index.vue",
                meta: {
                    title: "底层layout"
                },
                children: []
            }]
            const asyncRouterRes = await asyncMenu()
            const asyncRouter = asyncRouterRes.data.menus
            baseRouter[0].children = asyncRouter
            baseRouter.push({
                path: '*',
                redirect: '/404'

            })
            asyncRouterHandle(baseRouter)
            commit('setAsyncRouter', baseRouter)
        }
    },
    getters: {
        // 获取动态路由
        asyncRouters(state) {
            return state.asyncRouters
        }
    }
}