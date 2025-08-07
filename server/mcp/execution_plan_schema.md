# ExecutionPlan 结构体格式说明

## 概述
ExecutionPlan 是用于自动化模块创建的执行计划结构体，包含了创建包和模块所需的所有信息。

## 完整结构体定义

```go
type ExecutionPlan struct {
    PackageName        string                           `json:"packageName"`        // 包名，如："user", "order", "product"
    PackageType        string                           `json:"packageType"`        // "plugin" 或 "package"
    NeedCreatedPackage bool                             `json:"needCreatedPackage"` // 是否需要创建包
    NeedCreatedModules bool                             `json:"needCreatedModules"` // 是否需要创建模块
    PackageInfo        *request.SysAutoCodePackageCreate `json:"packageInfo,omitempty"`  // 包信息（当NeedCreatedPackage=true时必需）
    ModulesInfo        []*request.AutoCode              `json:"modulesInfo,omitempty"` // 模块信息数组（当NeedCreatedModules=true时必需，支持批量创建）
    Paths              map[string]string                `json:"paths,omitempty"`       // 路径信息
}
```

## 子结构体详细说明

### 1. SysAutoCodePackageCreate 结构体

```go
type SysAutoCodePackageCreate struct {
    Desc        string `json:"desc"`        // 描述，如："用户管理模块"
    Label       string `json:"label"`       // 展示名，如："用户管理"
    Template    string `json:"template"`    // 模板类型："plugin" 或 "package"
    PackageName string `json:"packageName"` // 包名，如："user"
    Module      string `json:"-"`          // 模块名（自动填充，无需设置）
}
```

### 2. AutoCode 结构体（核心字段）

```go
type AutoCode struct {
    Package             string           `json:"package"`             // 包名
    TableName           string           `json:"tableName"`           // 数据库表名
    BusinessDB          string           `json:"businessDB"`          // 业务数据库名
    StructName          string           `json:"structName"`          // 结构体名称
    PackageName         string           `json:"packageName"`         // 文件名称
    Description         string           `json:"description"`         // 结构体中文名称
    Abbreviation        string           `json:"abbreviation"`        // 结构体简称
    HumpPackageName     string           `json:"humpPackageName"`     // 驼峰命名的包名
    GvaModel            bool             `json:"gvaModel"`            // 是否使用GVA默认Model
    AutoMigrate         bool             `json:"autoMigrate"`         // 是否自动迁移表结构
    AutoCreateResource  bool             `json:"autoCreateResource"`  // 是否自动创建资源标识
    AutoCreateApiToSql  bool             `json:"autoCreateApiToSql"`  // 是否自动创建API
    AutoCreateMenuToSql bool             `json:"autoCreateMenuToSql"` // 是否自动创建菜单
    AutoCreateBtnAuth   bool             `json:"autoCreateBtnAuth"`   // 是否自动创建按钮权限
    OnlyTemplate        bool             `json:"onlyTemplate"`        // 是否只生成模板
    IsTree              bool             `json:"isTree"`              // 是否树形结构
    TreeJson            string           `json:"treeJson"`            // 树形结构JSON字段
    IsAdd               bool             `json:"isAdd"`               // 是否新增
    Fields              []*AutoCodeField `json:"fields"`              // 字段列表
    GenerateWeb         bool             `json:"generateWeb"`         // 是否生成前端代码
    GenerateServer      bool             `json:"generateServer"`      // 是否生成后端代码
    Module              string           `json:"-"`                  // 模块（自动填充）
    DictTypes           []string         `json:"-"`                  // 字典类型（自动填充）
}
```

### 3. AutoCodeField 结构体（字段定义）

```go
type AutoCodeField struct {
    FieldName       string      `json:"fieldName"`       // 字段名
    FieldDesc       string      `json:"fieldDesc"`       // 字段中文描述
    FieldType       string      `json:"fieldType"`       // 字段类型：string, int, bool, time.Time等
    FieldJson       string      `json:"fieldJson"`       // JSON标签名
    DataTypeLong    string      `json:"dataTypeLong"`    // 数据库字段长度
    Comment         string      `json:"comment"`         // 数据库字段注释
    ColumnName      string      `json:"columnName"`      // 数据库列名
    FieldSearchType string      `json:"fieldSearchType"` // 搜索类型：EQ, LIKE, BETWEEN等
    FieldSearchHide bool        `json:"fieldSearchHide"` // 是否隐藏查询条件
    DictType        string      `json:"dictType"`        // 字典类型
    Form            bool        `json:"form"`            // 是否在表单中显示
    Table           bool        `json:"table"`           // 是否在表格中显示
    Desc            bool        `json:"desc"`            // 是否在详情中显示
    Excel           bool        `json:"excel"`           // 是否支持导入导出
    Require         bool        `json:"require"`         // 是否必填
    DefaultValue    string      `json:"defaultValue"`    // 默认值
    ErrorText       string      `json:"errorText"`       // 校验失败提示
    Clearable       bool        `json:"clearable"`       // 是否可清空
    Sort            bool        `json:"sort"`            // 是否支持排序
    PrimaryKey      bool        `json:"primaryKey"`      // 是否主键
    DataSource      *DataSource `json:"dataSource"`      // 数据源配置（用于关联其他表）
    CheckDataSource bool        `json:"checkDataSource"` // 是否检查数据源
    FieldIndexType  string      `json:"fieldIndexType"`  // 索引类型
}
```

### 4. DataSource 结构体（关联表配置）

```go
type DataSource struct {
    DBName       string `json:"dbName"`       // 关联的数据库名称
    Table        string `json:"table"`        // 关联的表名
    Label        string `json:"label"`        // 用于显示的字段名（如name、title等）
    Value        string `json:"value"`        // 用于存储的值字段名（通常是id）
    Association  int    `json:"association"`  // 关联关系：1=一对一，2=一对多
    HasDeletedAt bool   `json:"hasDeletedAt"` // 关联表是否有软删除字段
}
```

## 使用示例

### 示例1：创建新包和批量创建多个模块

```json
{
  "packageName": "user",
  "packageType": "package",
  "needCreatedPackage": true,
  "needCreatedModules": true,
  "packageInfo": {
    "desc": "用户管理模块",
    "label": "用户管理",
    "template": "package",
    "packageName": "user"
  },
  "modulesInfo": [
    {
      "package": "user",
      "tableName": "sys_users",
      "businessDB": "",
      "structName": "User",
      "packageName": "user",
      "description": "用户",
      "abbreviation": "user",
      "humpPackageName": "user",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "onlyTemplate": false,
      "isTree": false,
      "treeJson": "",
      "isAdd": true,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "Username",
          "fieldDesc": "用户名",
          "fieldType": "string",
          "fieldJson": "username",
          "dataTypeLong": "50",
          "comment": "用户名",
          "columnName": "username",
          "fieldSearchType": "LIKE",
          "fieldSearchHide": false,
          "dictType": "",
          "form": true,
          "table": true,
          "desc": true,
          "excel": true,
          "require": true,
          "defaultValue": "",
          "errorText": "请输入用户名",
          "clearable": true,
          "sort": false,
          "primaryKey": false,
          "dataSource": {
            "dbName": "gva",
            "table": "sys_users",
            "label": "username",
            "value": "id",
            "association": 2,
            "hasDeletedAt": true
          },
          "checkDataSource": true,
          "fieldIndexType": ""
        },
        {
          "fieldName": "Email",
          "fieldDesc": "邮箱",
          "fieldType": "string",
          "fieldJson": "email",
          "dataTypeLong": "100",
          "comment": "邮箱地址",
          "columnName": "email",
          "fieldSearchType": "EQ",
          "fieldSearchHide": false,
          "dictType": "",
          "form": true,
          "table": true,
          "desc": true,
          "excel": true,
          "require": true,
          "defaultValue": "",
          "errorText": "请输入邮箱",
          "clearable": true,
          "sort": false,
          "primaryKey": false,
          "dataSource": null,
          "checkDataSource": false,
          "fieldIndexType": "index"
        }
      ]
    },
    {
      "package": "user",
      "tableName": "user_profiles",
      "businessDB": "",
      "structName": "UserProfile",
      "packageName": "user",
      "description": "用户档案",
      "abbreviation": "userProfile",
      "humpPackageName": "user",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "onlyTemplate": false,
      "isTree": false,
      "treeJson": "",
      "isAdd": true,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "UserID",
          "fieldDesc": "用户ID",
          "fieldType": "int",
          "fieldJson": "userId",
          "dataTypeLong": "",
          "comment": "关联用户ID",
          "columnName": "user_id",
          "fieldSearchType": "EQ",
          "fieldSearchHide": false,
          "dictType": "",
          "form": true,
          "table": true,
          "desc": true,
          "excel": true,
          "require": true,
          "defaultValue": "",
          "errorText": "请选择用户",
          "clearable": true,
          "sort": false,
          "primaryKey": false,
          "dataSource": null,
          "checkDataSource": false,
          "fieldIndexType": "index"
        },
        {
          "fieldName": "Avatar",
          "fieldDesc": "头像",
          "fieldType": "string",
          "fieldJson": "avatar",
          "dataTypeLong": "255",
          "comment": "用户头像URL",
          "columnName": "avatar",
          "fieldSearchType": "",
          "fieldSearchHide": true,
          "dictType": "",
          "form": true,
          "table": true,
          "desc": true,
          "excel": false,
          "require": false,
          "defaultValue": "",
          "errorText": "",
          "clearable": true,
          "sort": false,
          "primaryKey": false,
          "dataSource": null,
          "checkDataSource": false,
          "fieldIndexType": ""
        }
      ]
    }
  ]
}
```

### 示例2：仅在现有包中批量创建多个模块

```json
{
  "packageName": "system",
  "packageType": "package",
  "needCreatedPackage": false,
  "needCreatedModules": true,
  "packageInfo": null,
  "modulesInfo": [
    {
      "package": "system",
      "tableName": "sys_roles",
      "businessDB": "",
      "structName": "Role",
      "packageName": "system",
      "description": "角色",
      "abbreviation": "role",
      "humpPackageName": "system",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "onlyTemplate": false,
      "isTree": false,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "RoleName",
          "fieldDesc": "角色名称",
          "fieldType": "string",
          "fieldJson": "roleName",
          "dataTypeLong": "50",
          "comment": "角色名称",
          "columnName": "role_name",
          "fieldSearchType": "LIKE",
          "form": true,
          "table": true,
          "desc": true,
          "require": true
        }
      ]
    },
    {
      "package": "system",
      "tableName": "sys_permissions",
      "businessDB": "",
      "structName": "Permission",
      "packageName": "system",
      "description": "权限",
      "abbreviation": "permission",
      "humpPackageName": "system",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "onlyTemplate": false,
      "isTree": false,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "PermissionName",
          "fieldDesc": "权限名称",
          "fieldType": "string",
          "fieldJson": "permissionName",
          "dataTypeLong": "100",
          "comment": "权限名称",
          "columnName": "permission_name",
          "fieldSearchType": "LIKE",
          "form": true,
          "table": true,
          "desc": true,
          "require": true
        },
        {
          "fieldName": "PermissionCode",
          "fieldDesc": "权限代码",
          "fieldType": "string",
          "fieldJson": "permissionCode",
          "dataTypeLong": "50",
          "comment": "权限代码",
          "columnName": "permission_code",
          "fieldSearchType": "=",
          "form": true,
          "table": true,
          "desc": true,
          "require": true
        }
      ]
    }
  ]
}
```

### 示例3：模块关联关系配置详解

以下示例展示了如何配置不同类型的关联关系：

```json
{
  "packageName": "order",
  "packageType": "package",
  "needCreatedPackage": true,
  "needCreatedModules": true,
  "packageInfo": {
    "desc": "订单管理模块",
    "label": "订单管理",
    "template": "package",
    "packageName": "order"
  },
  "modulesInfo": [
    {
      "package": "order",
      "tableName": "orders",
      "structName": "Order",
      "packageName": "order",
      "description": "订单",
      "abbreviation": "order",
      "humpPackageName": "order",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "UserID",
          "fieldDesc": "下单用户",
          "fieldType": "uint",
          "fieldJson": "userId",
          "columnName": "user_id",
          "fieldSearchType": "EQ",
          "form": true,
          "table": true,
          "desc": true,
          "require": true,
          "dataSource": {
            "dbName": "gva",
            "table": "sys_users",
            "label": "username",
            "value": "id",
            "association": 2,
            "hasDeletedAt": true
          },
          "checkDataSource": true
        },
        {
          "fieldName": "ProductID",
          "fieldDesc": "商品",
          "fieldType": "uint",
          "fieldJson": "productId",
          "columnName": "product_id",
          "fieldSearchType": "EQ",
          "form": true,
          "table": true,
          "desc": true,
          "require": true,
          "dataSource": {
            "dbName": "gva",
            "table": "products",
            "label": "name",
            "value": "id",
            "association": 2,
            "hasDeletedAt": false
          },
          "checkDataSource": true
        },
        {
          "fieldName": "Status",
          "fieldDesc": "订单状态",
          "fieldType": "int",
          "fieldJson": "status",
          "columnName": "status",
          "fieldSearchType": "EQ",
          "form": true,
          "table": true,
          "desc": true,
          "require": true,
          "dictType": "order_status"
        }
      ]
    }
  ]
}
```

## DataSource 配置说明

### 关联关系类型
- **association: 1** - 一对一关联（如用户与用户档案）
- **association: 2** - 一对多关联（如用户与订单）

### 配置要点
1. **dbName**: 通常为 "gva"（默认数据库）
2. **table**: 关联表的实际表名
3. **label**: 用于前端显示的字段（如用户名、商品名称）
4. **value**: 用于存储关联ID的字段（通常是 "id"）
5. **hasDeletedAt**: 关联表是否支持软删除
6. **checkDataSource**: 建议设为true，会验证关联表是否存在

### 常见关联场景
- 用户关联：`{"table": "sys_users", "label": "username", "value": "id"}`
- 角色关联：`{"table": "sys_authorities", "label": "authorityName", "value": "authorityId"}`
- 部门关联：`{"table": "sys_departments", "label": "name", "value": "id"}`
- 分类关联：`{"table": "categories", "label": "name", "value": "id"}`

## 重要注意事项

1. **PackageType**: 只能是 "plugin" 或 "package"
2. **NeedCreatedPackage**: 当为true时，PackageInfo必须提供
3. **NeedCreatedModules**: 当为true时，ModulesInfo必须提供
4. **字段类型**: FieldType支持的类型包括：
   - string（字符串）
   - richtext（富文本）
   - int（整型）
   - bool（布尔值）
   - float64（浮点型）
   - time.Time（时间）
   - enum（枚举）
   - picture（单图片，字符串）
   - pictures（多图片，json字符串）
   - video（视频，字符串）
   - file（文件，json字符串）
   - json（JSON）
   - array（数组）
5. **搜索类型**: FieldSearchType支持：EQ, NE, GT, GE, LT, LE, LIKE, BETWEEN等
6. **索引类型**: FieldIndexType支持：index, unique等
7. **GvaModel**: 设置为true时会自动包含ID、CreatedAt、UpdatedAt、DeletedAt字段
8. **关联配置**: 使用dataSource时，确保关联表已存在，建议开启checkDataSource验证

## 常见错误避免

1. 确保PackageName和ModuleName符合Go语言命名规范
2. 字段名使用大写开头的驼峰命名
3. JSON标签使用小写开头的驼峰命名
4. 数据库列名使用下划线分隔的小写命名
5. 必填字段不要遗漏
6. 字段类型要与实际需求匹配