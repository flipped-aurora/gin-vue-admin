
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

[English](./README-en.md) | 简体中文

## 不同版本

我们会对以下四个版本持续维护，请选择适合自己的版本使用。最新技术栈为组合式api版本，已支持多语言（I18N）

[组合式API版（主）](https://github.com/flipped-aurora/gin-vue-admin) | 
[组合式API多语言(i18n)版](https://github.com/flipped-aurora/gin-vue-admin/tree/i18n-dev-new) | 
[声明式API版](https://github.com/flipped-aurora/gin-vue-admin/tree/v2.4.x) | 
[声明式API多语言(i18n)版](https://github.com/flipped-aurora/gin-vue-admin/tree/i18n-dev)

# 项目文档
[在线文档](https://www.gin-vue-admin.com) : https://www.gin-vue-admin.com

[初始化](https://www.gin-vue-admin.com/guide/start-quickly/initialization.html)
						       
[从环境到部署教学视频](https://www.bilibili.com/video/BV1Rg411u7xH)

[开发教学](https://www.gin-vue-admin.com/guide/start-quickly/env.html) (贡献者:  <a href="https://github.com/LLemonGreen">LLemonGreen</a> And <a href="https://github.com/fkk0509">Fann</a>)

[交流社区](https://support.qq.com/products/371961)

[插件市场](https://plugin.gin-vue-admin.com/)

# 重要提示

1.本项目从起步到开发到部署均有文档和详细视频教程

2.本项目需要您有一定的golang和vue基础

3.您完全可以通过我们的教程和文档完成一切操作，因此我们不再提供免费的技术服务，如需服务请进行[付费支持](https://www.gin-vue-admin.com/coffee/payment.html)

4.如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。您需保留如下版权声明信息，其余信息功能不做任何限制。如需剔除请[购买授权](https://www.gin-vue-admin.com/empower/index.html)

<img src="https://qmplusimg.henrongyi.top/%E6%8E%88%E6%9D%83.png" width="1000">

## 1. 基本介绍

### 1.1 项目介绍

> Gin-vue-admin是一个基于 [vue](https://vuejs.org) 和 [gin](https://gin-gonic.com) 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

[在线预览](http://demo.gin-vue-admin.com): http://demo.gin-vue-admin.com

测试用户名：admin

测试密码：123456

### 1.2 贡献指南
Hi! 首先感谢你使用 gin-vue-admin。

Gin-vue-admin 是一套为快速研发准备的一整套前后端分离架构式的开源框架，旨在快速搭建中小型项目。

Gin-vue-admin 的成长离不开大家的支持，如果你愿意为 gin-vue-admin 贡献代码或提供建议，请阅读以下内容。

#### 1.2.1 Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。如果你在使用时产生了疑问，请到 Slack 或 [Gitter](https://gitter.im/ElemeFE/element) 里咨询。

- 在提交 issue 之前，请搜索相关内容是否已被提出。

#### 1.2.2 Pull Request 规范
- 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

## 2. 使用说明

```
- node版本 > v16.8.3
- golang版本 >= v1.16
- IDE推荐：Goland
```

### 2.1 server项目

使用 `Goland` 等编辑工具，打开server目录，不可以打开 gin-vue-admin 根目录

```bash

# 克隆项目
git clone https://github.com/flipped-aurora/gin-vue-admin.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

##### （1）可以访问外国网站

````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn) 或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 如果嫌弃麻烦,可以使用go generate 编译前自动执行代码, 不过这个不能使用 `Goland` 或者 `Vscode` 的 命令行终端
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.3.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档

### 2.4 VSCode工作区

#### 2.4.1 开发

使用`VSCode`打开根目录下的工作区文件`gin-vue-admin.code-workspace`，在边栏可以看到三个虚拟目录：`backend`、`frontend`、`root`。

#### 2.4.2 运行/调试

在运行和调试中也可以看到三个task：`Backend`、`Frontend`、`Both (Backend & Frontend)`。运行`Both (Backend & Frontend)`可以同时启动前后端项目。

#### 2.4.3 settings

在工作区配置文件中有`go.toolsEnvVars`字段，是用于`VSCode`自身的go工具环境变量。此外在多go版本的系统中，可以通过`gopath`、`go.goroot`指定运行版本。

```json
    "go.gopath": null,
    "go.goroot": null,
```

## 3. 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql` > (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 4. 项目架构

### 4.1 系统架构图

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 前端详细设计图 （提供者:<a href="https://github.com/baobeisuper">baobeisuper</a>）

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 目录结构

```
    ├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── model           (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   ├── excel       (excel导入导出默认路径)                        
        │   ├── page        (表单生成器)                        
        │   └── template    (模板)                            
        ├── router          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
    
            web
        ├── babel.config.js
        ├── Dockerfile
        ├── favicon.ico
        ├── index.html                 -- 主页面
        ├── limit.js                   -- 助手代码
        ├── package.json               -- 包管理器代码
        ├── src                        -- 源代码
        │   ├── api                    -- api 组
        │   ├── App.vue                -- 主页面
        │   ├── assets                 -- 静态资源
        │   ├── components             -- 全局组件
        │   ├── core                   -- gva 组件包
        │   │   ├── config.js          -- gva网站配置文件
        │   │   ├── gin-vue-admin.js   -- 注册欢迎文件
        │   │   └── global.js          -- 统一导入文件
        │   ├── directive              -- v-auth 注册文件
        │   ├── main.js                -- 主文件
        │   ├── permission.js          -- 路由中间件
        │   ├── pinia                  -- pinia 状态管理器，取代vuex
        │   │   ├── index.js           -- 入口文件
        │   │   └── modules            -- modules
        │   │       ├── dictionary.js
        │   │       ├── router.js
        │   │       └── user.js
        │   ├── router                 -- 路由声明文件
        │   │   └── index.js
        │   ├── style                  -- 全局样式
        │   │   ├── base.scss
        │   │   ├── basics.scss
        │   │   ├── element_visiable.scss  -- 此处可以全局覆盖 element-plus 样式
        │   │   ├── iconfont.css           -- 顶部几个icon的样式文件
        │   │   ├── main.scss
        │   │   ├── mobile.scss
        │   │   └── newLogin.scss
        │   ├── utils                  -- 方法包库
        │   │   ├── asyncRouter.js     -- 动态路由相关
        │   │   ├── btnAuth.js         -- 动态权限按钮相关
        │   │   ├── bus.js             -- 全局mitt声明文件
        │   │   ├── date.js            -- 日期相关
        │   │   ├── dictionary.js      -- 获取字典方法 
        │   │   ├── downloadImg.js     -- 下载图片方法
        │   │   ├── format.js          -- 格式整理相关
        │   │   ├── image.js           -- 图片相关方法
        │   │   ├── page.js            -- 设置页面标题
        │   │   ├── request.js         -- 请求
        │   │   └── stringFun.js       -- 字符串文件
        |   ├── view -- 主要view代码
        |   |   ├── about -- 关于我们
        |   |   ├── dashboard -- 面板
        |   |   ├── error -- 错误
        |   |   ├── example --上传案例
        |   |   ├── iconList -- icon列表
        |   |   ├── init -- 初始化数据  
        |   |   |   ├── index -- 新版本
        |   |   |   ├── init -- 旧版本
        |   |   ├── layout  --  layout约束页面 
        |   |   |   ├── aside 
        |   |   |   ├── bottomInfo     -- bottomInfo
        |   |   |   ├── screenfull     -- 全屏设置
        |   |   |   ├── setting        -- 系统设置
        |   |   |   └── index.vue      -- base 约束
        |   |   ├── login              --登录 
        |   |   ├── person             --个人中心 
        |   |   ├── superAdmin         -- 超级管理员操作
        |   |   ├── system             -- 系统检测页面
        |   |   ├── systemTools        -- 系统配置相关页面
        |   |   └── routerHolder.vue   -- page 入口页面 
        ├── vite.config.js             -- vite 配置文件
        └── yarn.lock

```

## 5. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理。
- 文件上传下载：实现基于`七牛云`, `阿里云`, `腾讯云` 的文件上传操作(请开发自己去各个平台的申请对应 `token` 或者对应`key`)。
- 分页封装：前端使用 `mixins` 封装分页，分页方法调用 `mixins` 即可。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改(在线体验站点不开放此功能)。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。
	- 前端文件参考: [web/src/view/superAdmin/api/api.vue](https://github.com/flipped-aurora/gin-vue-admin/blob/master/web/src/view/superAdmin/api/api.vue)
    - 后台文件参考: [server/router/sys_api.go](https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/router/sys_api.go)
- 多点登录限制：需要在`config.yaml`中把`system`中的`use-multipoint`修改为true(需要自行配置Redis和Config中的Redis参数，测试阶段，有bug请及时反馈)。
- 分片长传：提供文件分片上传和大文件分片上传功能示例。
- 表单生成器：表单生成器借助 [@form-generator](https://github.com/JakHuang/form-generator) 。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。

## 6. 知识库 

## 6.1 团队博客

> https://www.yuque.com/flipped-aurora
>
>内有前端框架教学视频。如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994，欢迎您提出宝贵的需求。

## 6.2 教学视频

（1）手把手教学视频

> https://www.bilibili.com/video/BV1Rg411u7xH/

（2）后端目录结构调整介绍以及使用方法

> https://www.bilibili.com/video/BV1x44y117TT/

（3）golang基础教学视频

> bilibili：https://space.bilibili.com/322210472/channel/detail?cid=108884

（4）gin框架基础教学

> bilibili：https://space.bilibili.com/322210472/channel/detail?cid=126418&ctype=0

（5）gin-vue-admin 版本更新介绍视频

> bilibili：https://www.bilibili.com/video/BV1kv4y1g7nT

## 7. 联系方式

### 7.1 技术群

### QQ交流群：622360840
| QQ 群 |
|  :---:  |
| <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

### 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="http://qmplusimg.henrongyi.top/qrjjz.png"> 

添加微信，备注"加入gin-vue-admin交流群"

### [关于我们](https://www.gin-vue-admin.com/about/join.html)

## 8. 贡献者

感谢您对gin-vue-admin的贡献!

<a href="https://github.com/flipped-aurora/gin-vue-admin/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=flipped-aurora/gin-vue-admin" />
</a>

## 9. 捐赠

如果你觉得这个项目对你有帮助，你可以请作者喝饮料 :tropical_drink: [点我](https://www.gin-vue-admin.com/coffee/index.html)

## 10. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。
