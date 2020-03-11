
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

# 开发者(贡献者)列表
    [@piexlmax](https://github.com/piexlmax) 项目发起者 蒋*兆
    [@krank666](https://github.com/krank666) 前端示例作者 尹*
    [@1319612909](https://github.com/1319612909) 前端css优化 杜*兰
    [@granty1](https://github.com/granty1) 代码积极贡献者  印*林
    
# gin-vue-admin gin+vue开源快速项目模板
    更新代码后如果遇到前端报错请执行 npm i  前端开发会不定期增加新的依赖 请注意更新
# 重构记录
    2019年12月12日 17:15 第一次重构完成 主要更新后台数据库结构 引入casbin进行鉴权 后台代码结构整理 前端页面结构变更
    注意！！！
    在此时间之前拉取项目的朋友，pull代码后必须重新导入数据库数据，之前数据库结构不可使用
本模板使用前端ui框架为 element-ui https://element.eleme.cn/#/zh-CN 前端组件可查看elementUi文档使用
## 写在前面
    本项目主要是小弟在接各种私活的时候发现频繁得书写CURD、权限管理、用户管理、列表、分页、上传下载、日志包封装、文档自动化等等功能
    为了方方便各位个人开发者快速搭建项目基础模型，于是开发本模板，也为了以后自己可以省下一些初期开发的功夫
    本身是前端出身，所以对于后端的项目目录或者架构可能偏向前端思维，如果有什么地方需要改动感谢大家iss
    强烈希望在各位的帮助下，此项目可以用于企业级项目的开发
    
## 一点建议
    各位在clone项目以后，把db文件导入自己创建的库后，最好前往七牛云申请自己的空间地址，
    替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
    
## 测试环境地址（脚本持续恶意破坏中，正在修复，请尊重开源作者，切勿恶意破坏）
    测试环境被脚本恶意攻击，利用开源中上传七牛云操作，脚本持续上传删除，从而导致服务负载过大，无法持续启动，针对方案正在实施。
    开源不易，请各位按照视频教学，本地搭建环境。
    http://qmplus.henrongyi.top/ （被脚本持续攻击中，正在解决）
    登陆以后为最高权限，动api权限或者菜单权限均有可能导致数据错乱，系统无法使用。请自己创建账号并设置自己角色后进行测试。
    为防止恶意操作 会定期恢复数据库 如发现系统无法使用 请联系开发者

## 环境搭建教学视频

腾讯视频：https://v.qq.com/x/page/e3008xjxqtu.html    (等待最新视频录制)
    
## 模板使用教学及展示视频

腾讯视频：https://v.qq.com/x/page/c3008y2ukba.html    (等待最新视频录制)

## 技术选型
    1.后端采用golang框架gin，快速搭建基础restful风格API
    2.前端项目采用VUE框架，构建基础页面
    3.数据库采用Mysql(5.6.44)版本不同可能会导致SQL导入失败
    4.使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制
    5.使用swagger构建自动化文档
    6.使用fsnotify和viper实现json格式配置文件
    7.使用logrus实现日志记录
    8.使用gorm实现对数据库的基本操作

## 项目说明
    golang项目存放于QMPlusServer文件夹下，
    QMPlusServer内部static/config存放mysql相关配置。可以根据自己的mysql数据库名 用户名 密码修改对应配置
    vue项目存放于QMPlusVuePage文件夹下
    开源不易，感谢各位支持，错误指出即刻改正，改写纠错，感谢star支持
## TODO
    1.基本用户注册登录 √
    2.用户等基础数据CURD √
    3.调用des实现数据加密 √
    4.实现基于jwt的权限管理 √
    5.实现基于七牛云的文件上传操作（需提前注册七牛云账号） √（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
    6.等装了分页方法，实现分页接口并且复制粘贴就可使用分页 √
    7.前端分页mixin封装 分页方法调用mixins即可 √
    8.图片上传前端下载功能 √ <后端下载功能开发中>
    9.富文本编辑器，MarkDown编辑器功能嵌入 √
    10.增加条件搜索示例 前端文件参考src\view\superAdmin\api\api.vue 后台文件参考 model\dnModel\api.go √
    11.增加了多点登录限制 体验需要再 static\config中 把 system中的useMultipoint 修改为 true(需要自行配置redis和config中的redis参数)(测试阶段，有bug请及时反馈)√
    12.增加了配置文件管理功能 √
    13.大文件分片上传功能示例 √
    14...看项目进度想到什么做什么,主要目的是方便各位快速接私活，完成项目基础功能
    15.各位开发者可以在issues提出自己工作中遇到的重复性大或者实现复杂的需求，我尽可能提供示例代码。
## 计划任务
    1.导入，导出Excel
    2.Echart图表支持
    3.object模式树
    4.工作流，任务交接功能开发
    5.结构体前端直接录入生成对应结构体及自动建表
    6.单独前端使用模式以及数据模拟
    7.前端定制化生成form，table代码以及对应的后端结构体
## 使用说明
    1.golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本
    2.支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包
    3.go server建议使用goland运行 减少出错可能性
    4.前端项目node建议高于V8.6.0
    5.到前端项目目录下运行 npm i 安装所需依赖
    6.依赖安装完成直接运行 npm run serve即可启动项目
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
## 团队博客
    https://blog.henrongyi.top，内有前端框架教学视频，GOLANG基础入门视频正在筹备中。
    如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994,欢迎您提出宝贵的需求。
## docker镜像
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

<div align=center>
<h3>qq交流群:622360840</h3>
<h3>微信交流群可以添加任意一位开发者备注"加入gin-vue-admin交流群"</h3>
</div>


## 更新日志

2020/01/07 角色增加数据资源功能 增加数据资源关联返回 演示环境代码已同步 开启了多点登录拦截 可能会被其他人挤掉
2020/01/13 增加了配置管理功能 此功能不发表至测试环境 待保护机制以及服务重启机制发开完成后才会发表值测试环境 请自行clone且导入sql体验

## golang基础教学视频录制中...
地址:https://space.bilibili.com/322210472/channel/detail?cid=108884
