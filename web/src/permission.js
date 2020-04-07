import router from './router'
import { store } from '@/store/index'

let asyncRouterFlag = 0

const whiteList = ['login', 'register']

router.beforeEach(async(to, from, next) => {
    const token = store.getters['user/token']
        // if (token) {
        //     const expiresAt = store.getters['user/expiresAt']
        //     const nowUnix = new Date().getTime()
        //     const hasExpires = (expiresAt - nowUnix) < 0
        //     if (hasExpires) {
        //         store.dispatch['user/claerAll']
        //     }
        // }
        // 在白名单中的判断情况
    if (whiteList.indexOf(to.name) > -1) {
        if (token) {
            next({ path: '/layout/dashboard' })
        } else {
            next()
        }
    } else {
        // 不在白名单中并且已经登陆的时候
        if (token) {
            // 添加flag防止多次获取动态路由和栈溢出
            if (!asyncRouterFlag) {
                asyncRouterFlag++
                await store.dispatch('router/SetAsyncRouter')
                const asyncRouters = store.getters['router/asyncRouters']
                router.addRoutes(asyncRouters)
                next({...to, replace: true })
            } else {
                next()
            }
        }
        // 不在白名单中并且未登陆的时候
        if (!token) {
            next({
                name: "login",
                query: {
                    redirect: document.location.hash
                }
            })
        }
    }
})