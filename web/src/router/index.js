import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},

{
  path: '/func_hub/detail',
  name: 'func_detail',
  component: () => import('@/view/biz_apphub/funcHub/func_detail.vue')
},

{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/:catchAll(.*)',
  meta: {
    closeTab: true,
  },
  component: () => import('@/view/error/index.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
