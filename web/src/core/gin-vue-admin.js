/*
* gin-vue-admin web框架组
*
* */
import Vue from 'vue'
import './element_lazy' // 按需加载element
import uploader from 'vue-simple-uploader'
import APlayer from '@moefe/vue-aplayer'
// time line css
import '../../node_modules/timeline-vuejs/dist/timeline-vuejs.css'
// 路由守卫
import Bus from '@/utils/bus'
//加载网站配置文件夹
import config from './config'
Vue.prototype.$GIN_VUE_ADMIN = config
console.log(config)
Vue.use(Bus)
Vue.use(APlayer, {
  defaultCover: 'https://github.com/u3u.png',
  productionTip: true
})
Vue.use(uploader)

console.log(`
   欢迎使用 Gin-Vue-Admin
   当前版本:V2.4.2
   加群方式:微信：shouzi_1994 QQ群：622360840
   默认自动化文档地址:http://127.0.0.1:${process.env.VUE_APP_SERVER_PORT}/swagger/index.html
   默认前端文件运行地址:http://127.0.0.1:${process.env.VUE_APP_CLI_PORT}
   如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`)
