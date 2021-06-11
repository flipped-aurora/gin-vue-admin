import Vue from 'vue'
import App from './App.vue'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import { store } from '@/store'
Vue.config.productionTip = false

import { auth } from '@/directive/auth'
// 按钮权限指令
auth(Vue)

export default new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')

