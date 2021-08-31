/*
* gin-vue-admin web框架组
*
* */
// 加载网站配置文件夹
import config from './config'

export const run = function(app) {
  app.config.globalProperties.$GIN_VUE_ADMIN = config

  // app.use(uploader)

  console.log(`
     欢迎使用 Gin-Vue-Admin
     当前版本:V2.4.5 alpha
     加群方式:微信：shouzi_1994 QQ群：622360840
     默认自动化文档地址:http://127.0.0.1:${process.env.VUE_APP_SERVER_PORT}/swagger/index.html
     默认前端文件运行地址:http://127.0.0.1:${process.env.VUE_APP_CLI_PORT}
     如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
  `)
}

