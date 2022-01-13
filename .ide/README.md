# 使用SmartIDE开发和调试Gin-Vuew-Admin

SmartIDE是下一代的云原生IDE，可以帮助你一件启动开源项目的集成开发环境，直接进入编码调试，免除安装SDK，IDE和其他相关工具比如数据库管理工具的麻烦。

我们已经对Gin-Vue-Admin进行了SmartIDE适配，可以一键启动包含以下工具的 **标准化全栈开发环境(SFDE - Standard Fullstack Development Environment)**：

- 完整支持Vue的Node.js开发工具语言包（SDK）
- 完整支持Go语言Gin框架的开发工具语言包（SDK）
- 前端开发工具VSCode WebIDE
- 后端开发工具GoLand WebIDE
- 数据管理工具PHPMyAdmin用于管理Gva后台的MySQL数据库

本文档将对如何使用SmartIDE进行Gin-Vue-Admin项目的前后端联调进行描述。

## 本地启动项目

使用SmartIDE启动Gin-Vue-Admin的开发调试非常简单，仅需要两个步骤

1. 按照 安装手册 完成 SmartIDE 本地命令行工具的安装
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


