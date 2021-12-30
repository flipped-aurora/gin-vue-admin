import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import ElementPlus from 'element-plus'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { createI18n } from './i18n' // added by mohamed hassan to multilangauge

import '@/permission'
import { store } from '@/store/index'

import App from './App.vue'
const app = createApp(App)
app.config.productionTip = false

const i18n = createI18n() // added by mohamed hassan to multilangauge

app.use(i18n)

app.use(run)
  .use(auth)
  .use(store)
  .use(router)
  .use(ElementPlus, { i18n: (key, value) => i18n.t(key, value) }).mount('#app')

export default app
