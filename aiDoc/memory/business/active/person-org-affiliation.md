# person 页面「所属架构」字段

## 基本信息

- 提出日期:2026-07-14
- 当前状态:`active`
- 需求类型:页面改造 + 后端接口增强
- 优先级:中
- 需求文件:`aiDoc/memory/business/active/person-org-affiliation.md`
- 设计文档:`docs/superpowers/specs/2026-07-14-person-org-affiliation-design.md`

## 用户原始意图摘要

把 `/layout/person` 页面的「用户 ID」字段改为「所属架构」,值为当前登录用户所属**公司/部门**全路径,多部门用「｜」分隔;跨部门时兼顾多条。

## 影响范围

- 后端:`server/model/system/sys_department.go`(加视图字段 `NamePath`)、`server/service/system/sys_department.go`(加 `FillNamePaths`)、`server/service/system/sys_user.go`(`GetUserInfo` 调用填充)
- 前端:`web/src/view/person/person.vue`
- 文档:本记忆 + 设计 spec
- 插件 / 模块:无

## 涉及对象

- 模块:系统用户 / 部门(组织架构)
- 接口:`/user/getUserInfo`(增强返回,新增 `departments[].namePath` 字段)
- 页面:`/layout/person`
- 配置:无

## 已确认约束

- 只显示「公司/部门」全路径,**不含岗位**(部门与岗位在数据上是两个独立多对多,无法一一配对)
- 多部门用全角竖线 `｜` 分隔
- 无部门时显示「未定义部门」
- 路径从根部门一路拼到叶子,层级多深显示多深
- 采用方案 C(增强人人可调的 `getUserInfo`),避免新增接口在已初始化库上的 Casbin 授权问题
- 左卡「基本信息」改为「所属架构」,右卡「账号资料设置」的用户 ID 行删除
- 图标用 Element Plus `OfficeBuilding`

## 当前进展

- 需求已澄清、设计已确认,spec 已落档,待实现

## 后续待办

- 后端:`SysDepartment.NamePath` 字段 + `FillNamePaths` + `GetUserInfo` 调用 + 单测
- 前端:`person.vue` 左卡改字段、右卡删行、`orgText` 计算属性
- 浏览器点触验证
- 完成后把本文件从 `active/` 移入 `done/` 并更新 `demand-index.md`

## 更新规则

- 本文件只承载 person 页「所属架构」这一个功能点
- 同为 person 页的其他新功能点,新建独立文件(前缀 `person-`)并反向链接
