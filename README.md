
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/gorm-1.9.12-red"/>
</div>

[English](./README-en.md) | 简体中文

[gitee地址](https://gitee.com/pixelmax/gin-vue-admin)|
[github地址](https://github.com/flipped-aurora/gin-vue-admin)

# 项目文档
[在线文档](https://www.gin-vue-admin.com/) : https://www.gin-vue-admin.com/

[从环境到部署教学视频](https://www.bilibili.com/video/BV1fV411y7dT)

[开发教学](https://www.gin-vue-admin.com/docs/help) (贡献者:  <a href="https://github.com/LLemonGreen">LLemonGreen</a> And <a href="https://github.com/fkk0509">Fann</a>)
- 前端UI框架：[element-ui](https://github.com/ElemeFE/element) 
- 后台框架：[gin](https://github.com/gin-gonic/gin) 

## 1. 基本介绍

### 1.1 项目介绍

[在线预览](http://demo.gin-vue-admin.com/)

测试用户名：admin

测试密码：123456

> Gin-vue-admin是一个基于vue和gin开发的全栈前后端分离的后台管理系统，集成jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

### 1.2 贡献指南
Hi! 首先感谢你使用 gin-vue-admin。

Gin-vue-admin 是一套为后台管理平台准备的一整套前后端分离架构式的开源框架，旨在快速搭建后台管理系统。

Gin-vue-admin 的成长离不开大家的支持，如果你愿意为 gin-vue-admin 贡献代码或提供建议，请阅读以下内容。

#### 1.2.1 Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。如果你在使用时产生了疑问，请到 Slack 或 [Gitter](https://gitter.im/ElemeFE/element) 里咨询。

- 在提交 issue 之前，请搜索相关内容是否已被提出。

#### 1.2.2 Pull Request 规范
- 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- <font color=red>确保 PR 是提交到 `develop` 分支，而不是 `master` 分支。</font>

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

### 1.3 版本列表

- master: 2.0, 用于生产环境
- develop: 2.0, 用于测试环境
- [gin-vue-admin_v2_dev](https://github.com/flipped-aurora/gin-vue-admin/tree/gin-vue-admin_v2_dev) (v2.0 [GormV1版本](https://v1.gorm.io)稳定分支)
- [gva_gormv2_dev](https://github.com/flipped-aurora/gin-vue-admin/tree/gva_gormv2_dev) (v2.0 [GormV2版本](https://v2.gorm.io)开发分支)

## 2. 使用说明

```
- node版本 > v8.6.0
- golang版本 >= v1.14
- IDE推荐：Goland
- gormv2版本初始化数据库可以利用批量创建功能，这里已经写好初始化代码，需要在main.go内打开 initialize.Data() 的注释即可
- 替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
```

> 使用docker-compose体验本项目
- 安装 docker-compose [官方文档](https://docs.docker.com/compose/install/)
    - ```shell script
       # 在Linux安装
       # 1.1 运行此命令以下载Docker Compose的当前稳定版本
       sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
       # 1.2 将可执行权限应用于二进制文件
       sudo chmod +x /usr/local/bin/docker-compose 
      ```
    - ```shell script
       # 使用Python的pip安装 
       pip3 install docker-compose -i https://pypi.tuna.tsinghua.edu.cn/simple
      ```
    - 使用 Docker Desktop 
        - Windows: https://hub.docker.com/editions/community/docker-ce-desktop-windows
        - Mac: https://hub.docker.com/editions/community/docker-ce-desktop-mac/

- 使用git克隆本项目
    - ```git
        git clone https://github.com/flipped-aurora/gin-vue-admin.git
      ```
    
- 使用docker-compose up一键启动启动项目
    - ```shell script
      # 使用docker-compose启动四个容器
      docker-compose up
      # 如果您修改了某些配置选项,可以使用此命令重新打包镜像
      docker-compose up --build
      # 使用docker-compose 后台启动
      docker-compose up -d
      ```

    - web项目预览 [http://127.0.0.1:8000](http://127.0.0.1:8000)

    - swagger文档 [http://127.0.0.1:8888/swagger/index.html](http://127.0.0.1:8888/swagger/index.html)

- 如果server的177.7.0.12这个容器内部ip被占用了,需要修改地方为

	- [docker-compose.yaml](./docker-compose.yaml)的第39行的177.7.0.12更换为你想要的ip
	- [.docker-compose/nginx/conf.d/my.conf](./.docker-compose/nginx/conf.d/my.conf)的第20行的177.7.0.12更换为你想要的ip

- docker-compose使用自定义的一个docker网络

	- ```dockerfile
		networks:
		  network:
		    ipam:
		      driver: default
		      config:
		        - subnet: '177.7.0.0/16' 
		```

	- 子网地址, 默认网关是177.7.0.1(docker-compose V2需要写,V3则不需要),具体信息看[官方文档](https://docs.docker.com/compose/compose-file/#ipv4_address-ipv6_address)

	- 默认的network名为gin-vue-admin_network,默认是bridge模式

	- 如果修改了子网,对应的每个service的ipv4_address都需要修改,还有[.docker-compose/nginx/conf.d/my.conf](./.docker-compose/nginx/conf.d/my.conf)的第20行的server的ip也需要修改

> <font color=red>**使用docker-compose进行部署本项目需注意的问题**</font>

- mysql数据库请使用装在服务器磁盘的本地数据库.
	- 避免使用docker容器内的mysql,可能会出现写入的问题, io比宿主机低  docker的持久化机制问题
- [init.sql](.docker-compose/docker-entrypoint-initdb.d/init.sql)是给docker-compose进行<font color=red>体验本项目</font>的, 禁止[init.sql](.docker-compose/docker-entrypoint-initdb.d/init.sql)使用进行项目数据的初始化, 数据库初始化[请使用此方法](https://www.gin-vue-admin.com/docs/help#step1%EF%BC%9A%E6%95%B0%E6%8D%AE%E5%BA%93%E5%88%9D%E5%A7%8B%E5%8C%96)
	- 使用[init.sql](.docker-compose/docker-entrypoint-initdb.d/init.sql)进行初始化出现的所有问题,请自行承担,与本项目无关
- 使用本项目的docker-compose进行部署时,请修改[docker-compose.yaml](./docker-compose.yaml)对应的[nginx配置](.docker-compose/nginx/conf.d/my.conf),mysql配置,networks配置,redis配置,按需自行更改.

### 2.1 web端

```bash
# clone the project
git clone https://github.com/piexlmax/gin-vue-admin.git

# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 2.2 server端

```bash
# 使用 go.mod

# 安装go依赖包
go list (go mod tidy)

# 编译
go build
```

> Zap日志库使用指南&&配置指南

Zap日志库的配置选择在[config.yaml](./server/config.yaml)下的zap

```yaml
# zap logger configuration
zap:
  level: 'debug'
  format: 'console'
  prefix: '[GIN-VUE-ADMIN]'
  director: 'log'
  link_name: 'latest_log'
  show_line: true
  encode_level: 'LowercaseColorLevelEncoder'
  stacktrace_key: 'stacktrace'
  log_in_console: true
```

| 配置名         | 配置的类型 | 说明                                                         |
| -------------- | ---------- | ------------------------------------------------------------ |
| level          | string     | level的模式的详细说明,请看[zap官方文档](https://pkg.go.dev/go.uber.org/zap?tab=doc#pkg-constants) <br />info: info模式,无错误的堆栈信息,只输出信息<br />debug:debug模式,有错误的堆栈详细信息<br />warn:warn模式<br />error: error模式,有错误的堆栈详细信息<br />dpanic: dpanic模式<br />panic: panic模式<br />fatal: fatal模式<br /> |
| format         | string     | console: 控制台形式输出日志<br />json: json格式输出日志      |
| prefix         | string     | 日志的前缀                                                   |
| director       | string     | 存放日志的文件夹,修改即可,不需要手动创建                     |
| link_name      | string     | 在server目录下会生成一个link_name的[软连接文件](https://baike.baidu.com/item/%E8%BD%AF%E9%93%BE%E6%8E%A5),链接的是director配置项的最新日志文件 |
| show_line      | bool       | 显示行号, 默认为true,不建议修改                              |
| encode_level   | string     | LowercaseLevelEncoder:小写<br /> LowercaseColorLevelEncoder:小写带颜色<br />CapitalLevelEncoder: 大写<br />CapitalColorLevelEncoder: 大写带颜色 |
| stacktrace_key | string     | 堆栈的名称,即在json格式输出日志时的josn的key                 |
| log_in_console | bool       | 是否输出到控制台,默认为true                                  |

- 开发环境 || 调试环境配置建议
	- `level:debug`
	- `format:console`
	- `encode_level:LowercaseColorLevelEncoder`或者`encode_leve:CapitalColorLevelEncoder`
- 部署环境配置建议
	- `level:error`
	- `format:json` 
	- `encode_level: LowercaseLevelEncoder `或者 `encode_level:CapitalLevelEncoder`
	- `log_in_console: false` 
- <font color=red>建议只是建议,按照自己的需求进行即可,给出建议仅供参考</font>

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

##### （1）可以翻墙
````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法翻墙

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.io](https://goproxy.io/zh/)

```bash
如果您使用的 Go 版本是 1.13 及以上(推荐)
# 启用 Go Modules 功能
go env -w GO111MODULE=on 
# 配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.io,direct

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.3.2 生成API文档

````
cd server
swag init
````
执行上面的命令后，server目录下会出现docs文件夹，登录http://localhost:8888/swagger/index.html，即可查看swagger文档


## 3. 技术选型

- 前端：用基于`vue`的`Element-UI`构建基础页面。
- 后端：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 数据库：采用`MySql`(5.6.44)版本，使用`gorm`实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用`fsnotify`和`viper`实现`yaml`格式的配置文件。
- 日志：使用`go-logging`实现日志记录。


## 4. 项目架构
### 4.1 系统架构图

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 前端详细设计图 （提供者:<a href="https://github.com/baobeisuper">baobeisuper</a>）

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 目录结构

```
    ├─server  	     （后端文件夹）
    │  ├─api            （API）
    │  ├─config         （配置包）
    │  ├─core  	        （內核）
    │  ├─docs  	        （swagger文档目录）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─resource       （资源）
    │  ├─router         （路由）
    │  ├─service         (服务)
    │  └─utils	        （公共功能）
    └─web            （前端文件）
        ├─public        （发布模板）
        └─src           （源码包）
            ├─api       （向后台发送ajax的封装层）
            ├─assets	（静态文件）
            ├─components（组件）
            ├─router	（前端路由）
            ├─store     （vuex 状态管理仓）
            ├─style     （通用样式文件）
            ├─utils     （前端工具库）
            └─view      （前端页面）

```

## 5. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理 
- 文件上传下载：实现基于七牛云的文件上传操作（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
- 分页封装：前端使用mixins封装分页，分页方法调用mixins即可 
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改（测试环境不开放此功能）。
- 富文本编辑器：MarkDown编辑器功能嵌入。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。 
```
前端文件参考: src\view\superAdmin\api\api.vue 
后台文件参考: model\dnModel\api.go 
```
- 多点登录限制：需要在`config.yaml`中把`system`中的`useMultipoint`修改为true(需要自行配置Redis和Config中的Redis参数，测试阶段，有bug请及时反馈)。
- 分片长传：提供文件分片上传和大文件分片上传功能示例。
- 表单生成器：表单生成器借助 [@form-generator](https://github.com/JakHuang/form-generator)。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。 

## 6. 计划任务

- [ ] 导入，导出Excel
- [ ] Echart图表支持
- [ ] 工作流，任务交接功能开发
- [ ] 单独前端使用模式以及数据模拟

## 7. 知识库 
## 7.1 团队博客

> https://www.yuque.com/flipped-aurora
>
>内有前端框架教学视频。如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994，欢迎您提出宝贵的需求。

## 7.2 教学视频

（1）环境搭建
> Bilibili：https://www.bilibili.com/video/BV1Fg4y187Bw/ (v1.0版本视频，v2.0操作相同目录不同)

（2）模板使用
> Bilibili：https://www.bilibili.com/video/BV16K4y1r7BD/ (v1.0版本视频，v2.0操作相同目录不同)

（3）2.0目录以及开发体验
> Bilibili：https://www.bilibili.com/video/BV1aV411d7Gm#reply2831798461

（4）golang基础教学视频录制中...
> https://space.bilibili.com/322210472/channel/detail?cid=108884

## 8. 联系方式

### 8.1 技术群

### QQ交流群：622360840
| QQ 群 |
|  :---:  |
| <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

### 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="http://qmplusimg.henrongyi.top/qrjjz.png"> 

添加微信，备注"加入gin-vue-admin交流群"

### [关于我们](https://www.gin-vue-admin.com/about/)

## 9. 捐赠

如果你觉得这个项目对你有帮助，你可以请作者喝饮料 :tropical_drink: [点我](https://www.gin-vue-admin.com/docs/coffee)

## 10. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。
