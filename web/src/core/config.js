/**
 * 网站配置文件
 */

const config = {
  appName: 'Gin-Vue-Admin',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Gin-Vue-Admin，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.1`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信：shouzi_1994 QQ群：622360840`
      )
    )
    console.log(
      chalk.green(
        `> GVA讨论社区：https://support.qq.com/products/371961`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> 如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.github.com/flipped-aurora/gin-vue-admin/server.com/docs/coffee`
      )
    )
    console.log('\n')
  }
}

export default config
