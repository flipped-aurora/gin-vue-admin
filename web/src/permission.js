import router from './router'
import {getToken} from '@/utils/auth'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import getPageTitle from '@/utils/page'
import { store } from '@/store/index'

let asyncRouterFlag = 0

const whiteList = ['/login']
router.beforeEach(async(to, from, next) => {
    NProgress.start()

    document.title = getPageTitle(to.meta.title)

    const token = await getToken()
        // 在白名单中的判断情况
    if (to.path === '/login') {
        if (token) {
            next({ path: '/layout/dashboard' })
            NProgress.done()
        } else {
            next()
            NProgress.done()
        }
    } else {
        // 不在白名单中并且已经登陆的时候
        if (token) {
            // 添加flag防止多次获取动态路由和栈溢出
            if (!asyncRouterFlag) {
                try{
                    asyncRouterFlag++
                    await store.dispatch('router/SetAsyncRouter')
                    const asyncRouters = store.getters['router/asyncRouters']
                    if(asyncRouters){
                        router.addRoutes(asyncRouters)
                        next({...to, replace: true })
                        NProgress.done()
                    }
                 
                  
                }catch(error){
                    console.log(error)
                }
              
            } else {
                next()
                NProgress.done()
            }
        }
        // 不在白名单中并且未登陆的时候
        if (!token) {
            next(`login?redirect=${to.path}`)
            NProgress.done()
        }
    }
})

router.afterEach(() => {
    NProgress.done()
  })
  