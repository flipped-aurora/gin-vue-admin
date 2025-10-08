/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'
import packageInfo from '../../package.json'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 Gin-Vue-Admin
       当前版本:v${packageInfo.version}
       加群方式:微信：shouzi_1994 QQ群：622360840
       项目地址：https://github.com/flipped-aurora/gin-vue-admin
       插件市场:https://plugin.gin-vue-admin.com
       GVA讨论社区:https://support.qq.com/products/371961
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
       如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/coffee/index.html
       --------------------------------------版权声明--------------------------------------
       ** 版权所有方：flipped-aurora开源团队 **
       ** 版权持有公司：北京翻转极光科技有限责任公司 **
       ** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **
       ** 感谢您对Gin-Vue-Admin的支持与关注 合法授权使用更有利于项目的长久发展**
    `)
  }
}
