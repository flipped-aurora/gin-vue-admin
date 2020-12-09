import router from './router'
import { store } from '@/store/index'
import getPageTitle from '@/utils/page'

let asyncRouterFlag = 0

const whiteList = ['login']

router.beforeEach(async(to, from, next) => {
    const token = store.getters['user/token']
        // 在白名单中的判断情况
        //修改网页标签名称
    document.title = getPageTitle(to.meta.title)
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