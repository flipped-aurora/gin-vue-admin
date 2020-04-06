
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.10-red"/>
</div>

# 开发文档
[在线文档](http://doc.henrongyi.top/) 

本模板使用前端ui框架为 element-ui https://element.eleme.cn/#/zh-CN 前端组件可查看elementUi文档使用

## 版本列表
master为1.0版代码

[gin-vue-adminv 1.0稳定版](https://github.com/piexlmax/gin-vue-admin/tree/gin-vue-admin_v1_stable) v1.0保持更新和维护

[gin-vue-adminv 2.0测试版](https://github.com/piexlmax/gin-vue-admin) v2.0不再兼容v1.0



## 1. 基本介绍
>GIN-VUE-ADMIN是一个基于vue和gin开发的全栈前后端分离的后台管理系统，拥有jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

## 2. 技术选型
- 后端采用golang框架gin，快速搭建基础restful风格API

- 前端项目采用VUE框架，构建基础页面

- 数据库采用Mysql(5.6.44)版本不同可能会导致SQL导入失败

- 使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制

- 使用swagger构建自动化文档

- 使用fsnotify和viper实现json格式配置文件

- 使用logrus实现日志记录

- 使用gorm实现对数据库的基本操作

## 3. 项目目录

```
    ├─erver  	    （后端文件夹）
    │  ├─api            （API）
    │  ├─config         （配置包）
    │  ├─core  	        （內核）
    │  ├─db             （数据库脚本）
    │  ├─docs  	        （swagger文档目录）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─resource       （资源）
    │  ├─router         （路层）
    │  └─urtils		    （公共功能）
    └─web	         （前端文件）
        ├─public	    （发布模板）
        └─src           （源码包）
            ├─api	    （向后台发送ajax的封装层）
            ├─assets	（静态文件）
            ├─components（组件）
            ├─router	（前端路由）
            ├─store	    （vuex 状态管理仓）
            ├─style	    （通用样式文件）
            ├─utils	    （前端工具库）
            └─view	    （前端页面）

```

## 4. 主要功能
- 权限管理：基于jwt和casbin实现的权限管理 

- 文件上传下载：实现基于七牛云的文件上传操作（需提前注册七牛云账号） （为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）

- 分页封装：等装了分页方法，实现分页接口并且复制粘贴就可使用分页，前端分页mixin封装 分页方法调用mixins即可 

- 用户管理：系统管理员分配用户角色和角色权限。

- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。

- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。

- api管理：不同用户可调用的api接口的权限不同。

- 配置管理：配置文件可前台修改（测试环境不开放此功能）

- 富文本编辑器：富文本编辑器，MarkDown编辑器功能嵌入 

- 条件搜索：增加条件搜索示例 

```
前端文件参考: src\view\superAdmin\api\api.vue 
后台文件参考: model\sys_api.go 
```

- 多点登录限制：
体验需要在 config中 把 system中的useMultipoint 修改为 true(需要自行配置redis和config中的redis参数)(测试阶段，有bug请及时反馈)

- 分片上传：提供文件分片上传和大文件分片上传功能示例 

- 表单生成器：表单生成器借助 [@form-generator](https://github.com/JakHuang/form-generator)

- 代码生成器：后台基础逻辑以及简单curd的代码生成器 

## 5.  计划任务
-  导入，导出Excel

-  Echart图表支持

- 工作流，任务交接功能开发

- 单独前端使用模式以及数据模拟
## 6. 使用说明
- golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本

- 支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包

- go server建议使用goland运行 减少出错可能性

- 前端项目node建议高于V8.6.0

- 到前端项目目录下运行 npm i 安装所需依赖

- 依赖安装完成直接运行 npm run serve即可启动项目

### 6.1 生成swagger自动化API文档

#### 6.1.1 安装 swagger

- （1）可以翻墙

````
go get -u github.com/swaggo/swag/cmd/swag
````
- （2）无法翻墙

由于国内没法安装 go.org/x 包下面的东西，需要先安装gopm

````
go get -v -u github.com/gpmgo/gopm
````

- 再执行

````
gopm get -g -v github.com/swaggo/swag/cmd/swag
````

- 到我们GOPATH下的/src/github.com/swaggo/swag/cmd/swag路径
执行

````
go install
````

#### 6.1.2 生成API文档
- 安装完成过后在项目目录下运行

````
swag init
````

- 项目文件夹下面会有 doc文件夹出现，这时候登录 localhost:8888/swagger/index.html
就可以看到 swagger文档啦

## 7. 团队博客
    https://blog.henrongyi.top，内有前端框架教学视频，GoLang基础入门视频正在筹备中。
    如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994,欢迎您提出宝贵的需求。
    
## 8. docker镜像
   感谢 [@chenlinzhong](https://github.com/chenlinzhong)提供docker镜像
   
      #启动容器
      docker run -itd --net=host --name=go_container shareclz/go_node /bin/bash;
      
      #进入容器
      docker exec -it go_container /bin/bash;
      git clone https://github.com/piexlmax/gin-vue-admin.git /data1/www/htdocs/go/admin;
      
      #启动前端
      cd /data1/www/htdocs/go/admin/QMPlusVuePage;
      cnpm i ;
      npm run serve;
      
      #修改数据库配置
      vi /data1/www/htdocs/go/admin/QMPlusServer/static/dbconfig/config.json;
      
      #启动后端
      cd /data1/www/htdocs/go/admin/QMPlusServer;
      go run main.go;
      
## 9. 一点建议
    各位在clone项目以后，把db文件导入自己创建的库后，最好前往七牛云申请自己的空间地址，
    替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
    
## 10. 测试环境地址

测试环境:[http://qmplus.henrongyi.top/](http://qmplus.henrongyi.top/)
 
账号/密码: admin/123456

## 11. 环境搭建教学视频

bilibili：https://www.bilibili.com/video/BV1Fg4y187Bw/   (v1.0版本视频，v2.0操作相同目录不同)
    
## 12. 使用教学及展示视频

bilibili：https://www.bilibili.com/video/BV16K4y1r7BD/    (v1.0版本视频，v2.0操作相同目录不同)

## 13. 联系方式

|  奇淼   | krank666  |qq群|
|  :---:  |  :---: | :---: |
|  <img src="http://qmplusimg.henrongyi.top/jjz.jpg" width="180"/>  |  <img src="http://qmplusimg.henrongyi.top/yx.jpg" width="180"/> | <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

<div align=center>
<h3>qq交流群:622360840</h3>
<h3>微信交流群可以添加任意一位开发者备注"加入gin-vue-admin交流群"</h3>
</div>

## 14. 开发者(贡献者)列表

|  开发者   | 功能  | 姓名  |
|  ----  | ----  | ----  |
| [@piexlmax](https://github.com/piexlmax)  | 项目发起者 | 蒋 |
| [@krank666](https://github.com/krank666)  | 前端开发 | 尹 |
| [@1319612909](https://github.com/1319612909)  | 前端UI开发 |  杜 |
| [@granty1](https://github.com/granty1)  | 后台开发 | 印 |
| [@Ruio9244](https://github.com/Ruio9244)  | 全栈开发 | 严 |
| [@chen-chen-up](https://github.com/chen-chen-up)  | 新手开发 | 宋 |

## 15. 更新日志

|  日期   | 日志  |
|  :---:  | --- |
|2020/01/07| 角色增加数据资源功能 增加数据资源关联返回 演示环境代码已同步 开启了多点登录拦截 可能会被其他人挤掉 |
|2020/01/13| 增加了配置管理功能 此功能不发表至测试环境 待保护机制以及服务重启机制发开完成后才会发表值测试环境 请自行clone且导入sql体验 |
|2020/02/21| 修改了casbin的自定义鉴权方法，使其完全支持RESTFUL的/:params以及?query= 的接口模式 |
|2020/03/17| 增加了验证码功能 使用了 [@dchest/captcha](https://github.com/dchest/captcha)库 |
|2020/03/30| 代码生成器开发完成 表单生成器开发完成 使用了[@form-generator](https://github.com/JakHuang/form-generator) 库 |
|2020/04/01| 增加前端历史页签功能，增加（修改）条件查询示例，前端背景色调修改为白色 如不需要此功能可以在 view\/layout\/index\/   屏蔽HistoryComponent 背景色调 为本页260行 &.el-main 的background |
|2020/04/04| 启动2.x版本，项目文档规范化，日志功能改造，方法增加英文注释|

## 16. golang基础教学视频录制中...
---
地址:https://space.bilibili.com/322210472/channel/detail?cid=108884


## 17. 捐赠 :tropical_drink:
如果您想请团队喝可乐

|  支付宝   | 微信  |
|  :---:  | :---: |
| ![markdown](http://qmplusimg.henrongyi.top/zfb.png "支付宝") |  ![markdown](http://qmplusimg.henrongyi.top/wxzf.png "微信") |
