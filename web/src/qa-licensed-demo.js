import '@/style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'element-plus/dist/index.css'
import 'uno.css'

import * as ElIconModules from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import { createPinia } from 'pinia'
import { createApp, h } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'

import LogoHeader from '@/view/layout/header/index.vue'
import Dashboard from '@/view/dashboard/index.vue'
import { useAppStore } from '@/pinia/modules/app'
import { useUserStore } from '@/pinia/modules/user'

const pinia = createPinia()
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'qa',
      component: { render: () => h('div') },
      meta: { matched: [] }
    }
  ]
})

const appStore = useAppStore(pinia)
const userStore = useUserStore(pinia)

const updateDevice = () => {
  appStore.toggleDevice(window.innerWidth < 768 ? 'mobile' : 'desktop')
}

appStore.config.show_watermark = false
userStore.ResetUserInfo({
  nickName: '演示管理员',
  authority: { authorityName: '超级管理员' },
  authorities: []
})
updateDevice()
window.addEventListener('resize', updateDevice)

const app = createApp({
  render: () =>
    h('div', { class: 'h-screen bg-gray-50 text-slate-700 dark:bg-slate-800' }, [
      h(LogoHeader),
      h('main', { class: 'h-full pt-16' }, [h(Dashboard)])
    ])
})

for (const iconName in ElIconModules) {
  app.component(iconName, ElIconModules[iconName])
}

app.config.globalProperties.$GIN_VUE_ADMIN = {
  appName: 'Gin-Vue-Admin'
}

app.use(ElementPlus).use(pinia).use(router)
await router.isReady()
app.mount('#app')
