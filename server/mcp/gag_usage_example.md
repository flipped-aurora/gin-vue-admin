# GAG工具使用示例 - 带用户确认流程

## 新的工作流程

现在GAG工具支持三步工作流程：
1. `analyze` - 分析现有模块信息
2. `confirm` - 请求用户确认创建计划
3. `execute` - 执行创建操作（需要用户确认）

## 使用示例

### 第一步：分析
```json
{
  "action": "analyze",
  "requirement": "创建一个图书管理功能"
}
```

### 第二步：确认
```json
{
  "action": "confirm",
  "executionPlan": {
    "packageName": "library",
    "moduleName": "Book",
    "packageType": "package",
    "needCreatedPackage": true,
    "needCreatedModules": true,
    "packageInfo": {
      "desc": "图书管理包",
      "label": "图书管理",
      "template": "package",
      "packageName": "library"
    },
    "modulesInfo": {
      "package": "library",
      "tableName": "library_books",
      "businessDB": "",
      "structName": "Book",
      "packageName": "library",
      "description": "图书信息",
      "abbreviation": "book",
      "humpPackageName": "Library",
      "gvaModel": true,
      "autoMigrate": true,
      "autoCreateResource": true,
      "autoCreateApiToSql": true,
      "autoCreateMenuToSql": true,
      "autoCreateBtnAuth": true,
      "onlyTemplate": false,
      "isTree": false,
      "treeJson": "",
      "isAdd": false,
      "generateWeb": true,
      "generateServer": true,
      "fields": [
        {
          "fieldName": "title",
          "fieldDesc": "书名",
          "fieldType": "string",
          "fieldJson": "title",
          "dataTypeLong": "255",
          "comment": "书名",
          "columnName": "title",
          "fieldSearchType": "LIKE",
          "fieldSearchHide": false,
          "dictType": "",
          "form": true,
          "table": true,
          "desc": true,
          "excel": true,
          "require": true,
          "defaultValue": "",
          "errorText": "请输入书名",
          "clearable": true,
          "sort": false,
          "primaryKey": false,
          "dataSource": {},
          "checkDataSource": false,
          "fieldIndexType": ""
        }
      ]
    }
  }
}
```

### 第三步：执行（需要确认参数）
```json
{
  "action": "execute",
  "executionPlan": {
    // ... 同上面的executionPlan
  },
  "packageConfirm": "yes",  // 确认创建包
  "modulesConfirm": "yes"   // 确认创建模块
}
```

## 确认参数说明

- `packageConfirm`: 当`needCreatedPackage`为true时必需
  - "yes": 确认创建包
  - "no": 取消创建包（停止后续处理）

- `modulesConfirm`: 当`needCreatedModules`为true时必需
  - "yes": 确认创建模块
  - "no": 取消创建模块（停止后续处理）

## 取消操作的行为

1. 如果用户在`packageConfirm`中选择"no"，系统将停止所有后续处理
2. 如果用户在`modulesConfirm`中选择"no"，系统将停止模块创建
3. 任何取消操作都会返回相应的取消消息，不会执行任何创建操作

## 注意事项

1. 必须先调用`confirm`来获取确认信息
2. 在`execute`时必须提供相应的确认参数
3. 确认参数的值必须是"yes"或"no"
4. 如果不需要创建包或模块，则不需要提供对应的确认参数