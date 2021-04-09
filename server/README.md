
整理代码结构
``` lua
web
├── api/v1 -- 主要API
|   ├── sys_initdb.go -- ico
|   └── sys_user.go --  
├── config -- 配置文件 设定操作的结构体
|   ├── auto_code.go -- ico captcha.go
|   ├── ... -- ico captcha.go
|   └── zap.go -- core
├── core -- 主要结构代码
|   ├── server_other.go -- ico captcha.go
|   ├── ... -- ico captcha.go
|   └── zap.go -- 
├── docs -- 文档系统
|   ├── docs.go -- ico captcha.go
|   ├── swagger.json -- json
|   └── swagger.yaml -- yaml  
├── global -- global
├── initialize -- initialize 
├── middleware -- 中间键
├── model -- global
│   ├── request  -- 所有请求model结构体
|   |   ├── common.go 
|   |   ├── ...
|   |   └── sys_user.go -- yaml  
|   ├── response  -- 返回数据
|   |   ├── common.go 
|   |   ├── ...
|   |   └── sys_user.go -- yaml  
├── packfile -- 文件写入
├── resource -- 资源文件
├── router -- 路由
├── service -- service层
├── source -- 文件目录操作 
├── utils
├── config.yaml  -- 
├── Dockerfile  -- docker配置
├── go.mod    -- mod 配置
├── go.sum -- sum
├── latest_log  -- vue-cli 配置
└── main.go  -- package.json
```