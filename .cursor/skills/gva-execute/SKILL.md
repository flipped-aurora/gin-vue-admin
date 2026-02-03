---
name: gva-execute
description: GVA代码生成器 - ExecutionPlan 完整参数说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# GVA 代码生成器 (gva_execute)

## ExecutionPlan 结构体

```json
{
  "packageName": "包名(string，小写开头)",
  "packageType": "package 或 plugin",
  "needCreatedPackage": "是否创建包(bool)",
  "needCreatedModules": "是否创建模块(bool)",
  "needCreatedDictionaries": "是否创建字典(bool)",
  "packageInfo": {
    "desc": "描述",
    "label": "展示名",
    "template": "package 或 plugin",
    "packageName": "包名"
  },
  "modulesInfo": [{
    "package": "包名",
    "tableName": "数据库表名(蛇形命名)",
    "businessDB": "业务数据库(可选)",
    "structName": "结构体名(大驼峰)",
    "packageName": "文件名称",
    "description": "中文描述",
    "abbreviation": "简称",
    "humpPackageName": "文件名称(小驼峰)",
    "gvaModel": true,
    "autoMigrate": true,
    "autoCreateApiToSql": true,
    "autoCreateMenuToSql": true,
    "generateWeb": true,
    "generateServer": true,
    "fields": [/* 见下方字段说明 */]
  }],
  "dictionariesInfo": [{
    "dictType": "字典类型",
    "dictName": "字典名称",
    "description": "字典描述",
    "status": true,
    "options": [
      {"label": "显示名", "value": "值", "sort": 1}
    ]
  }]
}
```

## 字段配置 (fields)

```json
{
  "fieldName": "字段名(大写开头)",
  "fieldDesc": "字段描述",
  "fieldType": "字段类型",
  "fieldJson": "JSON标签",
  "dataTypeLong": "数据长度",
  "comment": "注释",
  "columnName": "数据库列名",
  "fieldSearchType": "搜索类型",
  "fieldSearchHide": false,
  "dictType": "字典类型",
  "form": true,
  "table": true,
  "desc": true,
  "excel": false,
  "require": false,
  "defaultValue": "",
  "clearable": true,
  "sort": false,
  "primaryKey": false,
  "dataSource": {/* 关联配置，见下方 */}
}
```

## 字段类型 (fieldType)

| 类型 | 说明 |
|------|------|
| `string` | 字符串 |
| `richtext` | 富文本 |
| `int` | 整型 |
| `int64` | 长整型 |
| `float64` | 浮点型 |
| `bool` | 布尔值 |
| `time.Time` | 时间 |
| `enum` | 枚举 |
| `picture` | 单图片(字符串) |
| `pictures` | 多图片(JSON) |
| `video` | 视频 |
| `file` | 文件(JSON) |
| `json` | JSON |
| `array` | 数组 |

## 搜索类型 (fieldSearchType)

`=`, `!=`, `>`, `>=`, `<`, `<=`, `LIKE`, `BETWEEN`, `IN`, `NOT IN`

## 关联配置 (dataSource)

用于配置字段与其他表的关联关系：

```json
{
  "dbName": "数据库名(默认留空)",
  "table": "关联表名",
  "label": "显示字段(如 name)",
  "value": "值字段(通常 id)",
  "association": 1,
  "hasDeletedAt": true
}
```

### association 值
- `1` = 一对一：当前实体只能关联另一实体的一个记录
- `2` = 一对多：当前实体可关联另一实体的多个记录（fieldType 自动改为 array）

### 获取表名
在 `server/model` 或 `plugin/xxx/model` 查看 `TableName()` 实现：
- `SysUser` → `sys_users`
- `ExaFileUploadAndDownload` → `exa_file_upload_and_downloads`

## 重要规则

1. **needCreatedPackage=true** 时 `packageInfo` 必需
2. **needCreatedModules=true** 时 `modulesInfo` 必需
3. **needCreatedDictionaries=true** 时 `dictionariesInfo` 必需
4. **gvaModel=true** 时系统自动创建 ID 主键，字段不要设置 `primaryKey=true`
5. **gvaModel=false** 时必须有一个字段 `primaryKey=true`
6. **packageType** 只能是 `package` 或 `plugin`，默认用 `package`
7. 模块创建时会**自动生成 API 和菜单**，不要再调用 `api_creator` 和 `menu_creator`

## 工作流

```shell
requirement_analyzer → gva_analyze → gva_execute
```

1. `requirement_analyzer`: 分析需求，设计模块架构
2. `gva_analyze`: 获取现有包/字典信息
3. `gva_execute`: 执行代码生成

