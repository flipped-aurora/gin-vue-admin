import { createRouter, createWebHistory } from 'vue-router'

const routes = [{
    path: '/',
    name: 'Root',
    redirect: '/home',
    component: () => import('@/view/layout/index.vue'),
    children:[{
        path: 'home',
        name: 'Home',
        component: () => import('@/view/home/index.vue'),
    },{
        path: 'chat',
        name: 'Chat',
        component: () => import('@/view/chat/index.vue'),
    }]
}]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
