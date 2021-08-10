# qm-plus-vue-page

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

整理代码结构
``` lua
web
├── public -- public
|   ├── favicon.ico -- ico
|   └── index.html -- index
├── src   -- 源代码
│   ├── api  -- 所有请求
│   ├── assets  --  主题 字体等静态资源
|   ├── components -- components组件
|   ├── core -- gva抽离的一些前端资源
|   |   ├── config.js -- 配置文件
|   |   └── element_lazy.js -- elementt按需引入文件
|   |   └── gin-vue-admin.js -- gva前端控制库
|   ├── directive -- 公用方法 
|   ├── mixins -- 公用方法
|   ├── router -- 路由权限
|   ├── store -- store 
|   |   ├── modules -- modules 
|   |   |   ├── dictionary.js -- 动态路由
|   |   |   ├── router.js -- 路由
|   |   |   └── user.js -- 用户权限菜单过滤
|   |   ├── getters.js -- getters
|   |   └── index.js -- index
|   ├── styles -- css
|   ├── utils -- utils 组件
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
|   |   |   ├── aside -- 
|   |   |   ├── bottomInfo -- bottomInfo
|   |   |   ├── screenfull -- 全屏设置
|   |   |   ├── setting    -- 系统设置
|   |   |   └── index.vue -- base 约束
|   |   ├── login --登录 
|   |   ├── person --个人中心 
|   |   ├── superAdmin -- 超级管理员操作
|   |   ├── system -- 系统检测页面
|   |   ├── systemTools -- 系统配置相关页面
|   |   └── routerHolder.vue -- page 入口页面 
│   ├── App.vue  -- 入口页面
│   ├── main.js  -- 入口文件 加载组件 初始化等
│   └── permission.js  -- 跳转
├── build.config.js  -- 环境变量build配置
├── openDocument.js  -- 商用代码公司自用产品无需授权
├── .babelrc    -- babel-loader 配置
├── .travis.yml -- 自动化CI配置
├── vue.config.js  -- vue-cli 配置
└── package.json  -- package.json
```