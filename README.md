
<div align=center>
<img src="http://qmplusimg.henrongyi.top/logo.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.10-red"/>
</div>


# gin-vue-admin gin+vue开源快速项目模板
# 重构记录
    2019年12月12日 17:15 第一次重构完成 主要更新后台数据库结构 引入casbin进行鉴权 后台代码结构整理 前端页面结构变更
    注意！！！
    在此时间之前拉取项目的朋友，pull代码后必须重新导入数据库数据，之前数据库结构不可使用
本模板使用前端ui框架为 element-ui https://element.eleme.cn/#/zh-CN 前端组件可查看elementUi文档使用
## 写在前面
    本项目主要是小弟在接各种私活的时候发现频繁得书写CURD、权限管理、用户管理、列表、分页、上传下载、日志包封装、文档自动化等等功能
    为了方方便各位个人开发者快速搭建项目基础模型，于是开发本模板，也为了以后自己可以剩下一些初期开发的功夫
    本身是前端出身，所以对于后端的项目目录或者架构可能偏向前端思维，如果有什么地方需要改动感谢大家iss
    强烈希望在各位的帮助下，此项目可以用于企业级项目的开发
    
## 一点建议
    各位在clone项目以后，把db文件导入自己创建的库后，最好前往七牛云申请自己的空间地址，
    替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
    
## 环境搭建教学视频

腾讯视频：https://v.qq.com/x/page/e3008xjxqtu.html
    
## 模板使用教学及展示视频

腾讯视频：https://v.qq.com/x/page/c3008y2ukba.html

## 技术选型
    1.后端采用golang框架gin，快速搭建基础restful风格API
    2.前端项目采用VUE框架，构建基础页面
    3.数据库采用Mysql(5.6.44)版本不同可能会导致SQL导入失败，可能会引用redis作为缓存数据库使用（待定）
    4.使用swagger构建自动化文档
    5.使用fsnotify和viper实现json格式配置文件
    6.使用logrus实现日志记录
    7.使用gorm实现对数据库的基本操作

## 项目说明
    golang项目存放于QMPlusServer文件夹下，内部config存放mysql相关配置。可以根据自己的mysql数据库名 用户名 密码修改对应配置
    vue项目存放于QMPlusVuePage文件夹下
    开源不易，感谢各位支持，错误指出即刻改正，改写纠错，感谢star支持
## TODO
    1.基本用户注册登录 √
    2.用户等基础数据CURD √
    3.调用des实现数据加密 √
    4.实现基于jwt的权限管理 
    5.实现基于七牛云的文件上传操作（需提前注册七牛云账号） √（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
    6.等装了分页方法，实现分页接口并且复制粘贴就可使用分页 √
    7.前端分页mixin封装 分页方法调用mixins即可 √
    8.图片上传前端下载功能 √ <后端下载功能开发中>
    9.增加条件搜索示例 前端文件参考src\view\superAdmin\api\api.vue 后台文件参考 model\dnModel\api.go √
    10...看项目进度想到什么做什么,主要目的是方便各位快速接私活，完成项目基础功能
    11.各位开发者可以在issues提出自己工作中遇到的重复性大或者实现复杂的需求，我尽可能提供示例代码。
## 计划任务
    1.富文本编辑器，MarkDown编辑器功能嵌入
    2.导入，导出Excel
    3.Echart图表支持
    4.object模式树
    5.工作流，任务交接功能开发
    6.结构体前端直接录入生成对应结构体及自动建表
    7.单独前端使用模式以及数据模拟
## 使用说明
    1.golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本
    2.支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包
    3.go server建议使用goland运行 减少出错可能性
    4.前端项目node建议高于V8.6.0
    5.到前端项目目录下运行 npm i 安装所需依赖
    6.依赖安装完成直接运行 npm run dev即可启动项目
    7.如果要使用swagger自动化文档 首先需要安装 swagger
````
go get -u github.com/swaggo/swag/cmd/swag
````
由于国内没法安装到X包下面的东西 如果可以翻墙 上面的命令就可以让你安心使用swagger了
如果没有翻墙的办法那就先装一下 gopm

````
go get -v -u github.com/gpmgo/gopm
````
此时你就可以使用 gopm了
这时候执行
````
gopm get -g -v github.com/swaggo/swag/cmd/swag
````
等待安装完成以后
到我们GOPATH下面的/src/github.com/swaggo/swag/cmd/swag路径
执行
````
go install
````
安装完成过后在项目目录下运行
````
swag init
````
项目文件夹下面会有 doc文件夹出现
这时候登录 localhost:8888/swagger/index.html
就可以看到 swagger文档啦
## 个人博客
    http://www.henrongyi.top，内有前端框架教学视频，GOLANG基础入门视频正在筹备中。
    如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994,欢迎您提出宝贵的需求。
 
## 最后
    正在研发前端独立版本分支...可以不依赖go服务进行使用的后台模板
    感谢krank666协同开发
    可使用的初始mysql脚本正在制作中...
## 联系方式
<div align=center style="float: left">
<img src="http://qmplusimg.henrongyi.top/jjz.jpg" width="180"/>
<H3>Mr.奇淼</H3>
</div>
<div align=center style="float: left">
<img src="http://qmplusimg.henrongyi.top/yx.jpg" width="180"/>
<H3>krank666微信</H3>
</div>
