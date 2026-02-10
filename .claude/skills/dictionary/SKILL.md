---
name: dictionary
description: 字典相关工具说明
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# 字典工具

## 字典生成器 (generate_dictionary_options)

智能生成字典选项并自动创建字典记录。

### 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `dictType` | 是 | 字典类型，唯一标识 |
| `fieldDesc` | 是 | 字段描述，AI 据此生成选项 |
| `dictName` | 否 | 字典名称，不提供则自动生成 |
| `description` | 否 | 字典描述 |
| `options` | 否 | 选项列表 JSON |

### 智能生成示例

输入：`fieldDesc: "用户状态"`

自动生成：

```json
[
  {"label": "正常", "value": "1", "sort": 1},
  {"label": "禁用", "value": "2", "sort": 2}
]
```

### 手动指定选项

```json
{
  "dictType": "gender",
  "fieldDesc": "性别",
  "options": "[{\"label\":\"男\",\"value\":\"1\",\"sort\":1},{\"label\":\"女\",\"value\":\"2\",\"sort\":2}]"
}
```

---

## 字典查询器 (query_dictionaries)

查询系统中已有的字典信息。

### 参数

| 参数 | 说明 |
|------|------|
| `dictType` | 可选，指定字典类型精确查询 |
| `includeDisabled` | 是否包含禁用项，默认 false |
| `detailsOnly` | 是否只返回详情，默认 false |

### 返回格式

```json
{
  "dictionaries": [
    {
      "type": "gender",
      "name": "性别",
      "options": [
        {"label": "男", "value": "1"},
        {"label": "女", "value": "2"}
      ]
    }
  ]
}
```

---

## 常见字典场景

| 场景 | dictType 建议 | 典型选项 |
|------|--------------|---------|
| 状态 | `status` | 正常/禁用 |
| 性别 | `gender` | 男/女/未知 |
| 优先级 | `priority` | 高/中/低 |
| 订单状态 | `order_status` | 待支付/已支付/已发货/已完成 |
| 审批状态 | `approval_status` | 待审批/已通过/已拒绝 |

