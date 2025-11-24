/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'

const greenText = (text) => `\x1b[32m${text}\x1b[0m`

export const config = {
  appName: 'Gin-Vue-Admin',
  showViteLogo: true,
  keepAliveTabs: false,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(
      greenText(
        `> 欢迎使用Gin-Vue-Admin，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(greenText(`> 当前版本:v${packageInfo.version}`))
    console.log(greenText(`> 加群方式:微信：shouzi_1994 QQ群：470239250`))
    console.log(
      greenText(`> 项目地址：https://github.com/flipped-aurora/gin-vue-admin`)
    )
    console.log(greenText(`> 插件市场:https://plugin.gin-vue-admin.com`))
    console.log(
      greenText(`> GVA讨论社区:https://support.qq.com/products/371961`)
    )
    console.log(
      greenText(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      greenText(`> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`)
    )
    console.log(
      greenText(
        `--------------------------------------版权声明--------------------------------------`
      )
    )
    console.log(greenText(`** 版权所有方：flipped-aurora开源团队 **`))
    console.log(greenText(`** 版权持有公司：北京翻转极光科技有限责任公司 **`))
    console.log(
      greenText(
        `** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **`
      )
    )
    console.log('\n')
  }
}

export default config
