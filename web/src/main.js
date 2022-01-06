import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'

import { store } from '@/store/index'
import { createPinia } from 'pinia'
import { setPermission } from '@/permission'
import App from './App.vue'
const app = createApp(App)
app.config.productionTip = false
const pinia = createPinia()

app
  .use(run)
  .use(pinia)
  .use(auth)
  .use(store)
  .use(router)
  .use(ElementPlus, { locale: zhCn }).mount('#app')

setPermission(router)

export default app
