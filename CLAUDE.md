# CLAUDE.md

gin-vue-admin 全栈管理系统 | Go+Gin后端 | Vue3+Vite前端

## 开发命令

### 后端 (server/)
```bash
cd server
go generate && go mod tidy   # 安装依赖
go run .                      # 开发服务器
go build -o server            # 生产构建
swag init                     # 生成Swagger
go test ./...                 # 测试
```

### 前端 (web/)
```bash
cd web
npm install                   # 安装依赖
npm run serve                 # 开发服务器
npm run build                 # 生产构建
npm run preview               # 预览构建
```

### 全栈
```bash
make build      # 构建前后端
make doc        # 生成Swagger
make image      # Docker镜像
```

## 目录概览

```
server/
├── api/v1/       # API控制器
├── service/      # 业务逻辑
├── model/        # 数据模型
├── router/       # 路由定义
├── plugin/       # 插件目录
├── initialize/   # 初始化模块
└── config.yaml   # 配置文件

web/src/
├── api/          # API接口
├── view/         # 页面组件
├── pinia/        # 状态管理
├── plugin/       # 前端插件
└── router/       # 路由配置
```

## 核心原则

1. **分层架构**: `Router → API → Service → Model`，严禁跨层调用
2. **enter.go模式**: 所有api/service/router层使用enter.go管理
3. **统一响应**: 使用`response`包返回JSON

## Skills (按需加载)

开发时使用以下skills获取详细规范：

| Skill | 场景 |
|-------|------|
| `/backend-dev` | Go后端、API开发 |
| `/frontend-dev` | Vue页面、组件开发 |
| `/plugin-dev` | 插件开发 |
| `/model-spec` | 数据模型、类型问题 |
| `/swagger-spec` | API文档编写 |
| `/gva-helper/*` | GVA MCP工具详细参数说明 |

### GVA Helper Skills

GVA MCP 工具的详细参数说明，按需加载：

| Skill | 说明 |
|-------|------|
| `/gva-execute` | ExecutionPlan 参数结构 |
| `/requirement-analyzer` | 需求分析器说明 |
| `/gva-analyze` | 系统分析器说明 |
| `/gva-review` | 代码审查器说明 |
| `/api-creator` | API创建器说明 |
| `/menu-creator` | 菜单创建器说明 |
| `/dictionary` | 字典工具说明 |
| `/lister` | 查询工具说明 |

## 配置

- **后端**: `server/config.yaml` (数据库、Redis、存储)
- **前端开发**: `web/.env.development` (端口8080)
- **前端生产**: `web/.env.production`
- **默认端口**: 前端8080，后端8888

## 重要提醒

- 开发前检查 GVA Helper MCP 是否可用
- 参考插件: `server/plugin/announcement/`
- Swagger: `http://localhost:8888/swagger/index.html`
