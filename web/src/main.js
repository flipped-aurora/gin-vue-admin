import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
ElementUI.Dialog.props.closeOnClickModal.default = false
Vue.use(ElementUI);
// 引入封装的router
import router from '@/router/index'

// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

import '@/permission'
import { store } from '@/store/index'
Vue.config.productionTip = false

// 路由守卫
import Bus from '@/utils/bus.js'
Vue.use(Bus)

import APlayer from '@moefe/vue-aplayer';

Vue.use(APlayer, {
    defaultCover: 'https://github.com/u3u.png',
    productionTip: true,
});


import { auth } from '@/directive/auth'
// 按钮权限指令
auth(Vue)

import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

//引入echarts
import echarts from 'echarts'
import buildConfig from '../build.config';
Vue.prototype.$echarts = echarts;

console.log(`
       欢迎使用 Gin-Vue-Admin
       当前版本:V2.4.0
       加群方式:微信：shouzi_1994 QQ群：622360840
       默认自动化文档地址:http://127.0.0.1:` + buildConfig.goServerPort + `/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:`+ buildConfig.vueClientPort + `
       如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`)