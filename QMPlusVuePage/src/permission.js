import router from './router'
import { asyncRouterHandle } from '@/utils/asyncRouter';

router.beforeEach((to, from, next) => {
    next()
        // asyncRouterHandle(asyncRouter)   // 等待动态使用 VUEX持久化 会将其放入 vuex并且动态生成左侧列表
        // router.addRoutes(asyncRouter)
})

const asyncRouter = [{
    path: '/layout',
    name: 'layout',
    component: 'view/layout/index.vue',
    meta: {
        title: '首页',
    },
    children: [{
        path: 'dashbord',
        name: 'dashbord',
        component: 'view/dashbord/index.vue'
    }, {
        path: "test",
        name: "test",
        component: "view/test/index.vue"
    }]
}]