import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'uno.css';
import { createApp } from 'vue'
import ElementPlus from 'element-plus'

import 'element-plus/dist/index.css'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import i18n from './i18n' // added by mohamed hassan to multilingual

const app = createApp(App)
app.config.productionTip = false

app
  .use(run)
  .use(ElementPlus)
  .use(store)
  .use(auth)
  .use(router)
  .use(i18n)
  .mount('#app')
export default app
