import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

//获取原型对象上的push函数
const originalPush = Router.prototype.push
//修改原型对象中的push方法
Router.prototype.push = function push(location) {
   return originalPush.call(this, location).catch(err => err)
}

const baseRouters = [{
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'login',
        component: () =>
            import ('@/view/login/login.vue')
    }
]

// 需要通过后台数据来生成的组件

const createRouter = () => new Router({
    routes: baseRouters
})

const router = createRouter()

export default router
