import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import { run } from '@/core/gin-vue-admin.js'
import '@/permission'
import { store } from '@/store/index'

import { auth } from '@/directive/auth'

import App from './App.vue'
const app = createApp(App)
run(app)
auth(app)
app.config.productionTip = false
app.use(store).use(router).use(ElementPlus, { locale: zhCn }).mount('#app')

export default app
