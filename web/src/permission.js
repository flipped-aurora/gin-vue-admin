import router from './router'
import { store } from '@/store/index'
import getPageTitle from '@/utils/page'
let asyncRouterFlag = 0

const whiteList = ['Login', 'Init']

const getRouter = async() => {
    await store.dispatch('router/SetAsyncRouter')
    await store.dispatch('user/GetUserInfo')
    const asyncRouters = store.getters['router/asyncRouters']
    asyncRouters.forEach(asyncRouter => {
        router.addRoute(asyncRouter)
    })
}

async function handleKeepAlive(to) {
    if (to.matched && to.matched.length > 2) {
        for (let i = 1; i < to.matched.length; i++) {
            const element = to.matched[i - 1]
            if (element.name === 'layout') {
                to.matched.splice(i, 1)
                await handleKeepAlive(to)
            }
            // 如果没有按需加载完成则等待加载
            if (typeof element.components.default === 'function') {
                await element.components.default()
                await handleKeepAlive(to)
            }
        }
    }
}

router.beforeEach(async(to, from, next) => {
    handleKeepAlive(to)
    const token = store.getters['user/token']
        // 在白名单中的判断情况
    document.title = getPageTitle(to.meta.title)
    if (whiteList.indexOf(to.name) > -1) {
        if (token) {
            if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
                asyncRouterFlag++
                await getRouter()
            }
            next({ name: store.getters['user/userInfo'].authority.defaultRouter })
        } else {
            next()
        }
    } else {
        // 不在白名单中并且已经登陆的时候
        if (token) {
            // 添加flag防止多次获取动态路由和栈溢出
            if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
                asyncRouterFlag++
                await getRouter()
                next({...to, replace: true })
            } else {
                if (to.matched.length) {
                    next()
                } else {
                    next({ path: '/layout/404' })
                }
            }
        }
        // 不在白名单中并且未登陆的时候
        if (!token) {
            next({
                name: 'Login',
                query: {
                    redirect: document.location.hash
                }
            })
        }
    }
})