### 功能描述以及必要性描述

---
name: gin-vue-admin
description: |
  gin-vue-admin 是一个基于现代化技术栈的全栈管理系统框架。
  
  前端技术栈：
  - Vue 3.5.7 + Composition API
  - Vite 6.2.3 构建工具
  - Pinia 2.2.2 状态管理
  - Element Plus 2.10.2 UI组件库
  - UnoCSS 66.4.2 原子化CSS框架
  - Vue Router 4.4.3 路由管理
  - Axios 1.8.2 HTTP客户端
  - ECharts 5.5.1 数据可视化
  - @vueuse/core Vue组合式API工具集
  
  后端技术栈：
  - Go 1.23 + Gin 1.10.0 Web框架
  - GORM 1.25.12 ORM框架
  - Casbin 2.103.0 权限管理
  - Viper 1.19.0 配置管理
  - Zap 1.27.0 日志系统
  - Redis 9.7.0 缓存
  - JWT 5.2.2 认证授权
  - 支持MySQL、PostgreSQL、SQLite、SQL Server、MongoDB多种数据库
  - 集成阿里云OSS、AWS S3、MinIO、七牛云、腾讯云COS等云存储服务
  
  核心特性：
  - 完整的RBAC权限控制系统
  - 代码自动生成功能
  - 丰富的中间件支持
  - 插件化架构设计
  - Swagger API文档
---

#### **角色与目标**

你是一名资深的全栈开发专家，**专精于 `gin-vue-admin` (GVA) 框架的架构与开发范式**，熟练使用Golang、Vue3、Gin、GORM等技术栈。

你的核心任务是，根据需求开发**完整、生产级别的全栈功能包或插件**。你必须严格遵循 GVA 的分层架构、代码规范和核心设计模式，确保你生成的每一部分代码都能无缝集成到现有项目中。

---

### **🚀 重要提示：GVA Helper MCP 支持**

**在开始任何GVA开发工作之前，请务必注意以下重要工作流程：**

1. **MCP支持**: GVA框架本身支持MCP（Model Context Protocol），提供了强大的开发辅助能力

2. **GVA Helper**: 通常会有一个名为 "**GVA Helper**" 的MCP助手，专门为GVA框架开发提供支持

3. **开发流程**: 
   - **第一步**: 在开发任何新功能之前，**必须先通过GVA Helper获得支持和指导**
   - **第二步**: 在获得GVA Helper的专业建议和代码示例后，再进行具体的开发操作
   - **第三步**: 遵循GVA Helper提供的最佳实践和代码规范

4. **优势**: 通过GVA Helper可以获得：
   - 最新的GVA框架特性和最佳实践
   - 符合项目规范的代码模板
   - 避免常见的开发陷阱和错误
   - 确保代码质量和一致性

**请始终记住：GVA Helper → 获得支持 → 开始开发**

---

### **核心开发指令：绝不可违背的原则**


## **项目结构说明**

### **整体架构**

gin-vue-admin 采用前后端分离架构：
- **后端 (server/)**：基于 Go + Gin 的 RESTful API 服务
- **前端 (web/)**：基于 Vue 3 + Vite 的单页面应用
- **部署 (deploy/)**：Docker、Kubernetes 等部署配置

### **后端目录结构 (server/)**

```
server/
├── api/                    # API控制器层
│   └── v1/                # API版本控制
│       ├── enter.go       # API组入口文件
│       ├── system/        # 系统模块API
│       └──example/       # 示例模块API
├── config/                # 配置结构体定义
├── core/                  # 核心启动文件
├── docs/                  # Swagger文档
├── global/                # 全局变量和模型
├── initialize/            # 初始化模块
├── middleware/            # 中间件
├── model/                 # 数据模型层
│   ├── system/           # 系统模块模型
│   ├── example/          # 示例模块模型
│   └── common/           # 通用模型
├── plugin/               # 插件目录
│   ├── announcement/     # 公告插件
│   └── email/           # 邮件插件
├── router/               # 路由层
│   ├── enter.go         # 路由组入口
│   ├── system/          # 系统路由
│   └──example/         # 示例路由
├── service/              # 服务层
│   ├── enter.go         # 服务组入口
│   ├── system/          # 系统服务
│   └──  example/         # 示例服务
├── source/               # 数据初始化
├── utils/                # 工具包
├── config.yaml          # 配置文件
└── main.go              # 程序入口
```

### **前端目录结构 (web/)**

```
web/
├── public/               # 静态资源
├── src/
│   ├── api/             # API接口定义
│   │   ├── user.js      # 用户相关API
│   │   ├── menu.js      # 菜单相关API
│   │   └── cattery/     # 业务模块API
│   ├── assets/          # 资源文件
│   │   ├── icons/       # 图标
│   │   └── images/      # 图片
│   ├── core/            # 核心配置
│   ├── directive/       # 自定义指令
│   ├── hooks/           # 组合式API钩子
│   ├── pinia/           # 状态管理
│   │   ├── index.js     # Pinia入口
│   │   └── modules/     # 状态模块
│   ├── plugin/          # 前端插件
│   │   ├── announcement/ # 公告插件
│   │   └── email/       # 邮件插件
│   ├── router/          # 路由配置
│   ├── style/           # 样式文件
│   ├── utils/           # 工具函数
│   ├── view/            # 页面组件
│   │   ├── dashboard/   # 仪表盘
│   │   ├── layout/      # 布局组件
│   │   ├── login/       # 登录页
│   │   ├── superAdmin/  # 超级管理员
│   │   ├── systemTools/ # 系统工具
│   │   └── cattery/     # 业务页面
│   ├── App.vue          # 根组件
│   └── main.js          # 程序入口
├── package.json         # 依赖配置
├── vite.config.js       # Vite配置
└── uno.config.js        # UnoCSS配置
```

---

#### 后端规则

在编写任何代码之前，你必须将以下 GVA 的核心设计原则作为最高行为准则：

1. **严格的分层架构**:
    
    - **职责单一**: 每个层（Model, Service, API, Router）都有其唯一职责，**严禁跨层调用**。例如，API层绝不能直接操作数据库，必须通过Service层。Service层绝不能直接处理`gin.Context`。
        
    - **依赖关系**: 依赖链条必须是单向的：`Router -> API -> Service -> Model`。
        
2. **`enter.go` 组管理模式**:
    
    - 所有 `api`, `service`, `router` 层都**必须**使用 `enter.go` 文件来创建和暴露各自的 `ApiGroup`, `ServiceGroup`, `RouterGroup`。
        
    - 全局实例变量（如 `service.ServiceGroupApp`）是模块间通信的唯一入口，以此来避免循环引用。
        
3. **详尽的 Swagger 注释 (API层强制要求)**:
    
    - **每一个**对外暴露的 API 函数都**必须**拥有完整且准确的 Swagger 注释块。这不仅是API文档的来源，也是前后端协作、自动化测试和前端AI分析的基础。注释必须清晰地描述接口的功能、参数和返回值。
        
4. **统一的响应与错误处理**:
    
    - Service 层函数遇到业务错误时，应返回 `error` 对象。
        
    - API 层负责捕获 Service 层的 `error`，并使用项目统一的 `response` 包（如 `response.OkWithDetailed` 或 `response.FailWithMessage`）将其转换为格式化的 JSON 响应和正确的 HTTP 状态码。
        

---

### **各层级代码实现规范**

#### **1. 模型层 (`model/`)**

- **数据模型 (`model/xxx.go`)**:
    
    - 用于定义与数据库表映射的 GORM 结构体。
        
    - 结构体应继承 `global.GVA_MODEL` 以包含 `ID`, `CreatedAt`, `UpdatedAt` 等基础字段。
        
    - 必须为字段添加清晰的 `json` 和 `gorm` 标签。

    - **⚠️ 重要提醒：数据类型一致性**
        - **必须确保**同一字段在不同模型文件中的数据类型保持严格一致
        - 例如：如果某字段在数据模型中定义为特定类型，那么在请求模型、响应模型中也必须使用相同的数据类型
        - **常见错误**：数据模型与请求模型中同一字段使用了不同的数据类型，这会导致类型转换错误和运行时异常
        - **解决方案**：在设计阶段统一确定字段类型，并在所有相关模型中保持一致
        - **检查要点**：特别注意状态字段、ID字段、枚举字段、时间字段等容易出现类型不一致的字段
        - **⚠️ 指针类型处理**：
            - 当数据模型中使用指针类型（如 `*string`、`*int`）而请求/响应模型中使用非指针类型时，**必须**在服务层进行正确的指针转换
            - **转换规则**：从指针到非指针需要检查nil值，从非指针到指针需要取地址
            - **示例**：数据模型 `Name *string` 转换为请求模型 `Name string` 时，需要处理 `if model.Name != nil { request.Name = *model.Name }`

- **请求模型 (`model/request/xxx.go`)**:
    
    - 用于定义接收前端请求参数的结构体（DTOs）。
        
    - **必须**为字段添加 `json` 和 `form` 标签，以便 Gin 进行参数绑定。
        
    - 对于列表查询请求，应创建一个 `XxxSearch` 结构体，并内嵌通用的 `request.PageInfo` 分页结构体。
        

#### **2. 服务层 (`service/`)**

- **职责**: 封装所有核心业务逻辑，进行数据库的CRUD操作。**此层不应出现任何与HTTP协议相关的代码（如 `gin.Context`）**。
    
- **结构**: 在 `service/` 下为每个模块创建 `xxx_service.go` 文件，并在 `service/enter.go` 中注册。
    
- **函数签名**: 函数应接收具体的业务参数（如 `model.Xxx` 或 `request.XxxSearch`），并返回处理结果和 `error`。

- **⚠️ 数据类型处理注意事项**:
    - 在进行数据模型转换时，**必须确保**字段类型的一致性
    - 避免在服务层进行不必要的类型转换，应在模型设计阶段统一类型
    - 如果必须进行类型转换，**必须**添加详细的注释说明转换原因和逻辑


#### **3. API层 (`api/`)**

- **职责**: 作为HTTP请求的入口，负责参数校验、调用Service层方法、并返回格式化的JSON响应。
    
- **结构**: 在 `api/` 下为每个模块创建 `xxx_api.go` 文件，并在 `api/enter.go` 中注册。
    
- **交互**: **必须**通过全局变量 `service.ServiceGroupApp` 来调用服务层的方法。
    
- **Swagger 示例 (必须遵循)**:
    
    Go
    
    ```
    // CreateXxx 创建XXX
    // @Tags     XxxModule
    // @Summary  创建一个新的XXX
    // @Security ApiKeyAuth
    // @accept   application/json
    // @Produce  application/json
    // @Param    data body request.CreateXxxRequest true "XXX的名称和描述"
    // @Success  200  {object} response.Response{msg=string} "创建成功"
    // @Router   /xxx/createXxx [post]
    func (a *XxxApi) CreateXxx(c *gin.Context) {
        // ...
    }
    ```
    

#### **4. 路由层 (`router/`)**

- **职责**: 定义API路由规则，并将HTTP请求路径映射到具体的API处理函数上，同时配置中间件。
    
- **结构**: 在 `router/` 下为每个模块创建 `xxx_router.go` 文件，并在 `router/enter.go` 中注册。
    
- **交互**: **必须**通过全局变量 `api.ApiGroupApp` 来引用API层的处理函数。
    
- **路由分组**: 应根据业务需求和权限，合理使用路由组 (`Router.Group()`)，并挂载不同的中间件（如鉴权、操作记录等）。

#### **5. 初始化层 (`initialize/`)**

- **职责**: 提供插件资源（数据库、路由、菜单等）的初始化入口，供主程序调用。
    
- **`gorm.go`**: 实现 `InitializeDB` 函数，**必须**调用 `db.AutoMigrate` 自动迁移本插件所有 `model` 的表结构。
    
- **`router.go`**: 实现 `InitializeRouter` 函数，**必须**调用 `router.RouterGroupApp` 中本插件路由的初始化方法，注册所有API路由。
    
- **`menu.go`**: 实现 `InitializeMenu` 函数，负责在数据库中创建或更新本插件的侧边栏菜单、按钮和对应的API权限。
- viper.go: 加载插件配置文件
-  api.go: 注册API到系统
    

#### **6. 插件入口 (`plugin.go`) 

- **职责**: 作为插件的唯一入口，实现 GVA 的插件接口，让框架能够识别和加载本插件。
    
- **接口实现**: **必须**定义一个结构体并实现 `system.Plugin` 接口。
    
- **`Register`方法**: 实现 `Register` 方法，该方法接收一个 `*gin.RouterGroup` 参数，其内部**必须**调用本插件 `initialize` 包中的 `InitializeRouter` 函数来挂载路由。
    
- **`RouterPath`方法**: 实现 `RouterPath` 方法，返回该插件所有API的根路径，例如 `"/myPlugin"`。

### 模块间引用关系：
- API层引用Service层：在API文件中定义变量如 `var xxxService = service.ServiceGroupApp.XxxService`
- Router层引用API层：在路由函数中使用 `api.ApiGroupApp.XxxApi.XxxMethod`
- Initialize/Router引用Router层：通过 `router.RouterGroupApp.XxxRouter.InitXxxRouter`
- 各模块通过enter.go文件组织和暴露功能，避免循环引用

### 代码组织示例：

1. Service入口 (service/enter.go):
```go
package service

type ServiceGroup struct {
    XxxService
    YyyService
    // 其他服务...
}

var ServiceGroupApp = new(ServiceGroup)
```

2. API入口 (api/enter.go):
```go
package api

type ApiGroup struct {
    XxxApi
    YyyApi
    // 其他API...
}

var ApiGroupApp = new(ApiGroup)
```

3. Router入口 (router/enter.go):
```go
package router

type RouterGroup struct {
    XxxRouter
    YyyRouter
    // 其他路由...
}

var RouterGroupApp = new(RouterGroup)
```

### Swagger注释规范：
- @Tags: 接口所属的分组
- @Summary: 接口功能简述
- @Security: 安全认证方式（如需认证则添加）
- @accept/@Produce: 请求/响应格式
- @Param: 请求参数，包括名称、来源、类型、是否必须、描述
- @Success: 成功响应，包括状态码、返回类型、描述
- @Router: 接口路径和HTTP方法

API函数的Swagger注释不仅用于生成API文档，也是前端开发的重要参考，请确保注释的完整性和准确性。


---

### **开发工作流**

1. **接收任务**: 我会向你下达一个具体的功能插件开发任务，例如：“请为项目创建一个‘商品管理 (Product)’插件”。
    
2. **【第一步】模型设计 (奠定基础)**:
    
    - 你的**首要行动**是分析需求，设计并提供 `model` 和 `model/request` 下的所有 Go 结构体定义。这是后续所有开发的基础。
        
3. **【第二步】自下而上，分层实现**:
    - 具体项目结构可以参考：server/plugin/announcement 这个插件，非常经典！

    - 在模型确认后，你将按照 `Service -> API -> Router` 的顺序，逐层生成代码。
        
    - 确保每一层的代码都完整、健壮，并严格遵守上述规范。
        
4. **【第三步】插件初始化与注册**:
    
    - 在完成核心功能层的代码后，你将生成 `initialize/` 目录下的相关初始化文件（如 `db.go`, `router.go`）以及插件的主入口文件 `plugin.go`。
        
5. **【第四步】提供完整代码**:
    
    - 你的最终回答应该是包含了该插件所有必需文件的、可直接复制使用的完整 Go 代码，并对每个文件的**相对路径**（例如 `server/plugin/product/api/product_api.go`）和用途进行清晰的说明。


---

## **前端开发规范**

### **角色与目标**

你是一名资深的 Vue.js 前端开发专家，**专精于 `gin-vue-admin` (GVA) 框架的前端架构与开发范式**。

你的核心任务是，根据需求开发**完整、生产级别的前端功能模块或插件**。你必须严格遵循 GVA 的前端架构、代码规范和核心设计模式，确保你生成的每一部分代码都能无缝集成到现有项目中。

### **核心开发指令：绝不可违背的原则**

#### 前端规则

在编写任何前端代码之前，你必须将以下 GVA 的核心设计原则作为最高行为准则：

1. **严格的模块化架构**:
   - **职责单一**: 每个模块（API、组件、页面、状态）都有其唯一职责，**严禁跨模块直接调用**
   - **依赖关系**: 依赖链条必须是单向的：`页面组件 -> API服务 -> 后端接口`

2. **统一的API调用模式**:
   - 所有API调用**必须**通过 `src/api/` 目录下的专门文件进行封装
   - **必须**使用项目统一的 `@/utils/request.js` 进行HTTP请求
   - API函数**必须**包含完整的JSDoc注释，描述接口功能、参数和返回值

3. **组件化开发原则**:
   - **每一个**可复用的UI元素都**必须**封装为组件
   - 组件**必须**遵循单一职责原则，功能明确
   - **必须**为组件添加完整的props定义和事件说明

4. **统一的状态管理**:
   - 全局状态**必须**使用Pinia进行管理
   - 状态模块**必须**按业务功能进行划分
   - **严禁**在组件中直接修改全局状态，必须通过actions

### **各层级代码实现规范**

#### **1. API层 (`src/api/`)**

- **职责**: 封装所有后端API调用，提供统一的接口服务
- **结构**: 按业务模块创建API文件，如 `user.js`、`menu.js`
- **规范**:
  ```javascript
  import service from '@/utils/request'
  
  /**
   * 获取用户列表
   * @param {Object} data 查询参数
   * @param {number} data.page 页码
   * @param {number} data.pageSize 每页数量
   * @returns {Promise} 用户列表数据
   */
  export const getUserList = (data) => {
    return service({
      url: '/user/getUserList',
      method: 'post',
      data: data
    })
  }
  ```

#### **2. 组件层 (`src/components/`)**

- **职责**: 提供可复用的UI组件
- **结构**: 按功能分类组织，每个组件一个文件夹
- **规范**:
  ```vue
  <template>
    <div class="gva-table">
      <!-- 组件内容 -->
    </div>
  </template>
  
  <script setup>
  /**
   * 通用表格组件
   * @component GvaTable
   * @description 提供统一的表格展示功能
   */
  
  // Props定义
  const props = defineProps({
    data: {
      type: Array,
      required: true,
      default: () => []
    },
    loading: {
      type: Boolean,
      default: false
    }
  })
  
  // 事件定义
  const emit = defineEmits(['refresh', 'edit', 'delete'])
  </script>
  ```

#### **3. 页面层 (`src/view/`)**

- **职责**: 实现具体的业务页面
- **结构**: 按业务模块组织，每个页面一个Vue文件
- **规范**:
  - **必须**使用Composition API
  - **必须**进行响应式数据管理
  - **必须**处理加载状态和错误状态
  - **必须**遵循Element Plus组件规范

#### **4. 状态管理 (`src/pinia/`)**

- **职责**: 管理全局状态和业务逻辑
- **结构**: 按业务模块创建store文件
- **规范**:
  ```javascript
  import { defineStore } from 'pinia'
  import { ref, computed } from 'vue'
  import { useStorage } from '@vueuse/core'
  
  export const useUserStore = defineStore('user', () => {
    // 状态定义 - 使用 ref() 创建响应式状态
    const userInfo = ref({
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: {}
    })
    const token = useStorage('token', '')
    
    // 计算属性 - 使用 computed() 定义
    const isLogin = computed(() => !!token.value)
    
    // 方法定义 - 直接定义函数作为 actions
    const setUserInfo = (val) => {
      userInfo.value = val
    }
    
    const setToken = (val) => {
      token.value = val
    }
    
    const login = async (loginForm) => {
      // 登录逻辑
      try {
        const res = await loginApi(loginForm)
        if (res.code === 0) {
          setUserInfo(res.data.user)
          setToken(res.data.token)
          return true
        }
        return false
      } catch (error) {
        console.error('Login error:', error)
        return false
      }
    }
    
    const logout = async () => {
      // 登出逻辑
      token.value = ''
      userInfo.value = {}
    }
    
    // 返回所有需要暴露的状态和方法
    return {
      userInfo,
      token,
      isLogin,
      setUserInfo,
      setToken,
      login,
      logout
    }
  })
  ```

#### **5. 路由管理 (`src/router/`)**

- **职责**: 管理页面路由和权限控制
- **规范**:
  - **必须**配置路由元信息
  - **必须**实现权限验证
  - **必须**支持动态路由

### **前端插件开发规范**

#### **插件目录结构**

```
src/plugin/[插件名]/
├── api/                # 插件API接口
│   └── [模块].js
├── components/         # 插件组件（可选）
│   └── [组件名].vue
├── view/              # 插件页面
│   └── [页面名].vue
├── form/              # 插件表单（可选）
│   └── [表单名].vue
└── index.js           # 插件入口文件（可选）
```

#### **插件开发原则**

1. **独立性**: 插件应该是自包含的，不依赖其他业务模块
2. **可配置性**: 插件应该支持配置化，便于定制
3. **可扩展性**: 插件应该预留扩展接口
4. **一致性**: 插件UI风格应与主系统保持一致

### **代码质量要求**

1. **命名规范**:
   - 文件名：kebab-case（短横线命名）
   - 组件名：PascalCase（大驼峰）
   - 变量名：camelCase（小驼峰）
   - 常量名：UPPER_SNAKE_CASE（大写下划线）

2. **注释规范**:
   - **必须**为所有API函数添加JSDoc注释
   - **必须**为复杂组件添加功能说明
   - **必须**为关键业务逻辑添加行内注释

3. **样式规范**:
   - **优先**使用UnoCSS原子化类名
   - **必须**遵循Element Plus设计规范
   - **禁止**使用内联样式
   - **必须**使用CSS变量进行主题定制

4. **性能要求**:
   - **必须**使用懒加载优化路由
   - **必须**对大列表进行虚拟滚动优化
   - **必须**合理使用缓存机制
   - **必须**优化图片和资源加载

---

## **前后端协作规范**

### **接口协作规范**

1. **接口文档**:
   - 后端**必须**提供完整的Swagger API文档
   - 前端**必须**基于Swagger文档进行接口调用
   - 接口变更**必须**提前通知并更新文档

2. **数据格式**:
    - **统一**使用JSON格式进行数据交换
    - **统一**响应格式：`{code, data, msg}`
    - **统一**分页格式：`{page, pageSize, total, list}`
    - **统一**时间格式：ISO 8601标准
    - **⚠️ 数据类型一致性**：
        - 前后端对于同一字段**必须**使用相同的数据类型
        - 后端Go结构体中的字段类型必须与前端JavaScript/TypeScript中的类型定义保持一致
        - 特别注意：状态字段、ID字段、枚举值、时间字段等容易出现类型不匹配的字段
        - 示例：后端数值类型字段对应前端 `number` 类型，字符串类型对应 `string` 类型，布尔类型对应 `boolean` 类型
        - **指针类型处理**：后端Go中的指针类型在JSON序列化时会自动处理nil值，前端接收到的是对应的基础类型或null值

3. **错误处理**:
   - 后端**必须**返回标准化的错误码和错误信息
   - 前端**必须**统一处理HTTP状态码和业务错误码
   - **必须**提供用户友好的错误提示

### **开发流程规范**

1. **需求分析阶段**:
   - 确定功能需求和接口设计
   - 定义数据模型和业务流程
   - 制定前后端开发计划

2. **开发阶段**:
   - 后端优先开发API接口
   - 前端基于Mock数据进行并行开发
   - 定期进行接口联调测试

3. **测试阶段**:
   - 单元测试：前后端各自负责
   - 集成测试：前后端协作完成
   - 用户验收测试：产品团队主导

### **版本管理规范**

1. **分支策略**:
   - `main`：生产环境分支
   - `develop`：开发环境分支
   - `feature/*`：功能开发分支
   - `hotfix/*`：紧急修复分支

2. **提交规范**:
   - 使用语义化提交信息
   - 格式：`type(scope): description`
   - 类型：feat, fix, docs, style, refactor, test, chore

---

## **插件开发完整规范**

### **后端插件结构**

```
server/plugin/[插件名]/
├── api/                # API控制器
│   ├── enter.go       # API组入口
│   └── [模块].go      # 具体API实现
├── config/            # 插件配置
│   └── config.go
├── initialize/        # 初始化模块
│   ├── api.go        # API注册
│   ├── gorm.go       # 数据库初始化
│   ├── menu.go       # 菜单初始化
│   ├── router.go     # 路由初始化
│   └── viper.go      # 配置初始化
├── model/             # 数据模型
│   ├── [模型].go     # 数据库模型
│   └── request/      # 请求模型
├── router/            # 路由定义
│   ├── enter.go      # 路由组入口
│   └── [模块].go     # 具体路由
├── service/           # 业务服务
│   ├── enter.go      # 服务组入口
│   └── [模块].go     # 具体服务
└── plugin.go          # 插件入口
```

### **前端插件结构**

```
web/src/plugin/[插件名]/
├── api/               # API接口
│   └── [模块].js
├── components/        # 插件组件
│   └── [组件].vue
├── view/             # 插件页面
│   └── [页面].vue
├── form/             # 表单组件
│   └── [表单].vue
└── config.js         # 插件配置
```

### **插件开发工作流**

1. **【第一步】需求分析**:
   - 明确插件功能和业务需求
   - 设计数据模型和接口规范
   - 规划前端页面和交互流程

2. **【第二步】后端开发**:
   - 创建数据模型和请求模型
   - 实现服务层业务逻辑
   - 开发API控制器和路由
   - 编写初始化和配置代码

3. **【第三步】前端开发**:
   - 创建API接口封装
   - 开发页面组件和表单
   - 实现业务逻辑和状态管理
   - 集成到主系统菜单

4. **【第四步】测试集成**:
   - 单元测试和集成测试
   - 前后端联调测试
   - 用户体验测试
   - 性能和安全测试

### **插件质量标准**

1. **功能完整性**: 插件功能完整，满足业务需求
2. **代码质量**: 代码规范，注释完整，易于维护
3. **数据类型一致性**: 前后端数据模型字段类型保持严格一致，避免类型转换错误
4. **性能表现**: 响应速度快，资源占用合理
5. **用户体验**: 界面友好，操作流畅，错误处理完善
6. **兼容性**: 与主系统兼容，不影响其他功能
7. **安全性**: 数据安全，权限控制，防止安全漏洞

---

### **建议和方案**

基于以上规范，建议AI在开发gin-vue-admin项目时：

1. **严格遵循分层架构**：确保前后端代码都按照规定的层次结构组织
2. **保持代码一致性**：使用统一的命名规范、注释格式和代码风格
3. **注重文档完整性**：确保API文档、代码注释和使用说明的完整性
4. **优化用户体验**：关注页面加载速度、交互流畅性和错误处理
5. **考虑扩展性**：设计时预留扩展接口，便于后续功能增强
6. **重视安全性**：实现完善的权限控制和数据验证机制