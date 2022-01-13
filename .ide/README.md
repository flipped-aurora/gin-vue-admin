# 使用SmartIDE开发和调试Gin-Vuew-Admin

SmartIDE是下一代的云原生IDE，可以帮助你一件启动开源项目的集成开发环境，直接进入编码调试，免除安装SDK，IDE和其他相关工具比如数据库管理工具的麻烦。

我们已经对Gin-Vue-Admin进行了SmartIDE适配，可以一键启动包含以下工具的 **标准化全栈开发环境(SFDE - Standard Fullstack Development Environment)**：

- 完整支持Vue的Node.js开发工具语言包（SDK）
- 完整支持Go语言Gin框架的开发工具语言包（SDK）
- 前端开发工具VSCode WebIDE
- 后端开发工具GoLand WebIDE
- 数据管理工具PHPMyAdmin用于管理Gva后台的MySQL数据库

本文档将对如何使用SmartIDE进行Gin-Vue-Admin项目的前后端联调进行描述。

## 完整操作视频

为了便于大家更直观的了解使用SmartIDE开发调试Gin-Vue-Admin的过程，我们在B站提上提供了视频供大家参考，视频地址如下：

https://www.bilibili.com/video/BV1eL4y1b7ep/

## 本地模式启动项目

使用SmartIDE启动Gin-Vue-Admin的开发调试非常简单，仅需要两个步骤

1. 按照 [安装手册](https://smartide.cn/zh/docs/install/) 完成 SmartIDE 本地命令行工具的安装
2. 使用以下命令一键启动以上所列的SFDE

```shell
## SmartIDE是一款跨平台开发工具，您可以在Windows或者MacOS上执行同样的指令
smartide start https://github.com/SmartIDE/gin-vue-admin.git
```

以上命令会在运行命令的当前目录自动完成代码克隆，拉取开发环境镜像，启动容器，自动开启浏览器打开VSCode WebIDE以及自动恢复vue前端项目的npm依赖包并启动前端项目等一系列动作。

以上动作完成后，看到类似如下的VSCode WebIDE窗口。

> VSCode WebIDE的地址是 https://localhost:6800

![](images/vscode-webide.png)

我们的环境中还内置了JetBrain GoLand WebIDE

> JetBrain WebIDE的地址是 https://localhost:8887

![](images/goland-webide.png)

## 远程主机模式启动项目

以上B站操作使用平所使用的是远程主机模式，远程主机模式允许你将SmartIDE的开发环境一键部署到一台安装了Docker环境的远程主机上，并使用WebIDE远程连接到这台主机进行开发，对于比较复杂的项目来说这样做可以让你扩展本地开发机的能力，实现云端本地一体化的开发体验。

使用远程模式也仅需要两个步骤

> 注意：远程主机模式下你不必在本地安装Docker环境，只需要安装好SmartIDE的命令行工具即可

1. 按照 [Docker & Docker-Compose 安装手册 (Linux服务器)](https://smartide.cn/zh/docs/install/docker-install-linux/) 准备好一台远程主机
2. 按照以下指令启动项目

```shell
# 将远程主机添加到SmartIDE中
smartide host add <IpAddress> --username <SSH-UserName> --password <SSH-Password> --port <SSH-Port默认为22>

# 获取主机ID
smartide host list

# 使用远程主机启动项目
smartide start --host <主机ID> https://github.com/SmartIDE/gin-vue-admin.git
```

## 前后端联调

使用SmartIDE启动环境后，前端应用已经自动启动；此时只需要启动后端调试模式即可开始调试Gin-Vue-Admin。进入联调模式的环境状态如下

![](images/gva-debug.png)

调试相关的入口如下：

- VSCode WebIDE http://localhost:6800
- 前端应用 http://localhost:8080
- JetBrain GoLand WebIDE http://localhost:8887
- 后端应用(Swagger-UI) http://localhost:8888/swagger-ui/index.html
- 数据库管理PHPMyAdmin http://localhost:8090

## 技术支持

**特别说明:** SmartIDE本身是开源产品，并且对独立开发者提供免费使用授权。

大家可以通过以下链接获取SmartIDE的技术支持

- 产品官网 https://SmartIDE.cn
  - 通过产品官网上的二维码可以加入 [Smart早鸟群] 与其他的 Smart Developer 一起交流
- 开源首页：SmartIDE采用GitHub和Gitee双通道开源模式，方便国内开发者访问
  - https://githbu.com/SmartIDE
  - https://gitee.com/SmartIDE
  
  大家可以通过以上地址提交Issue，产品组的小伙伴会及时给予反馈

  > 如果大家喜欢我们的产品，请给予 Star 支持

- B站频道：我们定制组织直播，为大家更新产品开发进展
  - https://space.bilibili.com/1001970523

  > 如果大家喜欢我们的产品和视频，一定要记得 “三连” 



