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
|   |   ├── layout  --  layout约束页面 
|   |   |   ├── aside -- 
|   |   |   ├── bottomInfo -- bottomInfo
|   |   |   ├── screenfull -- 全屏设置
|   |   |   └── index.vue -- base 约束
|   |   ├── login --结算单管理 
|   |   ├── person --结算单管理 
|   |   ├── superAdmin -- 超级管理员操作
|   |   └── home.vue -- page 入口页面 
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