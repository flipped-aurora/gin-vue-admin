# person 页面「所属架构」字段 — 设计

- 日期:2026-07-14
- 分支:`feat/dev-293-newPage`
- 状态:设计已确认,待实现

## 1. 背景与目标

`/layout/person`(`web/src/view/person/person.vue`)当前展示「用户 ID」字段(左卡「基本信息」、右卡「账号资料设置」各一处)。用户 ID 对个人页无实际意义,改为展示当前登录用户的**所属组织架构**。

目标:

1. 左卡「基本信息」的 **用户 ID → 所属架构**(只读展示)。
2. 右卡「账号资料设置」的 **用户 ID 整行删除**。
3. 「所属架构」的值 = 当前用户各所属部门的 **公司/部门全路径**,多部门用全角竖线 `｜` 分隔。
4. **不显示岗位**(见下方决策记录)。

## 2. 关键数据事实(调研结论)

- 部门 `SysDepartment` 是**自关联树**,靠 `ParentId` + `Ancestors`(祖级 ID 链,顶级为 `"0"`、不含自身,如 `"0,1,5"`)维护层级。"公司/部门"即从根到该部门的名字路径。
- 岗位 `SysPosition` 是平表。
- 用户↔部门(`sys_user_departments`)、用户↔岗位(`sys_user_positions`)是**两个互相独立的多对多**,数据里**不记录"用户在某部门担任某岗位"**,因此部门与岗位无法一一配对。
- `/user/getUserInfo` 后端已 `Preload("Departments").Preload("Positions").Preload("Dept")`,前端 `userStore.userInfo` 已带 `departments`/`positions`;但用户自己的部门对象只含自身 `name` + 祖级 ID,**不含祖级名字**。
- `/department/getDepartmentList` 等管理端接口挂在 PrivateGroup,受 Casbin 限制,**普通用户无权限**;`CasbinHandler` 无白名单。
- `/user/getUserInfo` 被显式授予所有默认角色(888/8881/9528),**人人可调**。

## 3. 决策记录(与用户确认)

| 决策点 | 结论 | 理由 |
|--------|------|------|
| 拼接格式 | **只显示 公司/部门**,不含岗位 | 部门与岗位数据上无法配对;用户示例本身即"总公司/研发部" |
| 改动位置 | 左卡改为「所属架构」,右卡整行删除 | 所属架构是只读信息,不适合放在带「修改」按钮的设置区 |
| 实现方案 | **方案 C:增强 `getUserInfo`** | 复用人人可调的接口 → 零权限配置,老库/新库/所有用户立即生效;代价仅一次小查询 + 一个视图字段(新增接口方案 B 在已初始化库上需给角色手动分配 API 权限) |
| 无部门文案 | **「未定义部门」** | 用户指定 |
| 路径深度 | 从根部门一路拼到叶子,层级多深显示多深 | 用户确认 |
| 图标 | Element Plus `OfficeBuilding` | 组织/公司语义;EP 图标已在 `web/src/core/global.js` 全局注册 |

## 4. 数据流

```
getUserInfo(已 Preload Departments)
  → 后端 FillNamePaths 解析每个部门的「公司/部门」全路径 → dept.NamePath
  → 随 userInfo 下发 → userStore.userInfo.departments[i].namePath
  → person.vue: orgText = 各 namePath 用「｜」拼接(空则「未定义部门」)
```

## 5. 后端设计

### 5.1 模型 `server/model/system/sys_department.go`

给 `SysDepartment` 增加视图字段(与现有 `Children []SysDepartment gorm:"-"` 同款,不建列):

```go
NamePath string `json:"namePath" gorm:"-"` // 公司/部门全路径名(内存组装,不建列)
```

### 5.2 部门 service `server/service/system/sys_department.go`

新增填充方法:

```go
// FillNamePaths 依据各部门 Ancestors(祖级ID链,顶级"0"、不含自身)
// 解析出 "公司/部门" 全路径名,写入每个部门的 NamePath。解析失败返回 error,调用方可容错忽略。
func (s *SysDepartmentService) FillNamePaths(ctx context.Context, depts []system.SysDepartment) error
```

实现要点:

1. 遍历 `depts`,收集每个部门 `Ancestors` 按 `,` 切分、去掉 `"0"` 后的祖级 ID,汇成去重 ID 集合。
2. 若集合非空,一次 `SELECT id, name FROM sys_departments WHERE id IN (...)` 得到 `id→name` 映射。
3. 对每个部门:`NamePath = strings.Join(祖级名切片 + 自身 Name, "/")`;祖级名在映射中缺失则跳过该段(不报错)。
4. 顶级部门(`Ancestors == "0"`)→ `NamePath = Name`。

> 实现注意:`depts []SysDepartment` 是值切片,须用 `for i := range depts { depts[i].NamePath = ... }` **按索引**写回,不能用 `for _, d := range depts`(那样改的是副本,不会生效)。由于 `reqUser.Departments` 与入参共享底层数组,按索引写回即可透传到返回结果。

### 5.3 用户 service `server/service/system/sys_user.go` — `GetUserInfo`

在 `Preload(...).First(&reqUser, ...)` 成功之后、返回之前,追加(同包直接调):

```go
if err := SysDepartmentServiceApp.FillNamePaths(ctx, reqUser.Departments); err != nil {
    // 容错:路径解析失败不阻断用户信息返回,前端会降级为裸部门名
}
```

### 5.4 Swagger

`GetUserInfo` 的 `@Success` 为 `data=map[string]interface{}`,返回类型未写死,**新增字段无需改注释**;如需保持 docs 同步可选重新生成。

## 6. 前端设计 `web/src/view/person/person.vue`

1. **左卡「基本信息」用户 ID 项(约 58–66 行)** 改为:
   - 图标:`<OfficeBuilding />`
   - label:`所属架构`
   - value:`{{ orgText }}`
2. **右卡「账号资料设置」用户 ID 整行(约 159–170 行)删除**,并把上一行「账号密码」加上 `last-row` class(去掉底边框)。
3. `<script setup>` 新增计算属性:

```js
import { computed } from 'vue'

const orgText = computed(() => {
  const list = userStore.userInfo.departments || []
  const text = list.map((d) => d.namePath || d.name).filter(Boolean).join(' ｜ ')
  return text || '未定义部门'
})
```

> `d.namePath || d.name` 兜底:即使后端未给全路径也能显示裸部门名,不留空白。

## 7. 边界

| 场景 | 结果 |
|------|------|
| 无部门 | `未定义部门` |
| 跨多部门 | `总公司/研发部 ｜ 总公司/营销部` |
| 多级部门 | `总公司/研发中心/后端组` |
| 祖级名缺失 | 跳过该段,不报错 |

## 8. 测试

- **后端单测**(`server/service/system` 已有 `_test.go` 惯例):`FillNamePaths` 覆盖 顶级 / 多级 / 多部门 / 空祖级 / 祖级缺失 场景。
- **前端**:按 `aiDoc/frontend-backend/page-click-testing.md` 做浏览器点触验证(需登录 token):
  - 有多部门用户 → 显示 `｜` 分隔的全路径
  - 无部门用户 → 显示「未定义部门」
  - 右卡不再出现「用户 ID」行,左卡显示「所属架构」

## 9. 影响范围

- 后端:`sys_department.go`(model + service)、`sys_user.go`(service)
- 前端:`person.vue`
- 不改动:数据库表结构、getUserInfo 接口契约(仅新增字段)、其他调用 getUserInfo 的页面(向后兼容)
