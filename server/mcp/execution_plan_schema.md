# ExecutionPlan 结构体格式说明

## 概述
ExecutionPlan 是用于自动化模块创建的执行计划结构体，包含了创建包和模块所需的所有信息。

## 完整结构体定义

```go
type ExecutionPlan struct {
    PackageName        string                           `json:"packageName"`        // 包名，如："user", "order", "product"
    ModuleName         string                           `json:"moduleName"`         // 模块名，通常与结构体名相同
    PackageType        string                           `json:"packageType"`        // "plugin" 或 "package"
    NeedCreatedPackage bool                             `json:"needCreatedPackage"` // 是否需要创建包
    NeedCreatedModules bool                             `json:"needCreatedModules"` // 是否需要创建模块
    PackageInfo        *request.SysAutoCodePackageCreate `json:"packageInfo,omitempty"`  // 包信息（当NeedCreatedPackage=true时必需）
    ModulesInfo        *request.AutoCode                `json:"modulesInfo,omitempty"` // 模块信息（当NeedCreatedModules=true时必需）
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
    DataSource      *DataSource `json:"dataSource"`      // 数据源
    CheckDataSource bool        `json:"checkDataSource"` // 是否检查数据源
    FieldIndexType  string      `json:"fieldIndexType"`  // 索引类型
}
```

## 使用示例

### 示例1：创建新包和模块

```json
{
  "packageName": "user",
  "moduleName": "User",
  "packageType": "package",
  "needCreatedPackage": true,
  "needCreatedModules": true,
  "packageInfo": {
    "desc": "用户管理模块",
    "label": "用户管理",
    "template": "package",
    "packageName": "user"
  },
  "modulesInfo": {
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
        "dataSource": null,
        "checkDataSource": false,
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
  }
}
```

### 示例2：仅在现有包中创建模块

```json
{
  "packageName": "system",
  "moduleName": "Role",
  "packageType": "package",
  "needCreatedPackage": false,
  "needCreatedModules": true,
  "packageInfo": null,
  "modulesInfo": {
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
  }
}
```

## 重要注意事项

1. **PackageType**: 只能是 "plugin" 或 "package"
2. **NeedCreatedPackage**: 当为true时，PackageInfo必须提供
3. **NeedCreatedModules**: 当为true时，ModulesInfo必须提供
4. **字段类型**: FieldType支持的类型包括：string, int, int64, float64, bool, time.Time, enum, picture, video, file, pictures, array, richtext, json等
5. **搜索类型**: FieldSearchType支持：EQ, NE, GT, GE, LT, LE, LIKE, BETWEEN等
6. **索引类型**: FieldIndexType支持：index, unique等
7. **GvaModel**: 设置为true时会自动包含ID、CreatedAt、UpdatedAt、DeletedAt字段

## 常见错误避免

1. 确保PackageName和ModuleName符合Go语言命名规范
2. 字段名使用大写开头的驼峰命名
3. JSON标签使用小写开头的驼峰命名
4. 数据库列名使用下划线分隔的小写命名
5. 必填字段不要遗漏
6. 字段类型要与实际需求匹配