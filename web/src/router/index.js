import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
