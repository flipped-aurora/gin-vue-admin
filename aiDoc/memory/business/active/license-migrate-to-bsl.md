# 仓库许可证迁移到 BSL 1.1

## 基本信息

- 提出日期：2026-05-27
- 当前状态：`active`
- 需求类型：许可证与文档合规变更
- 优先级：高
- 需求文件：`aiDoc/memory/business/active/license-migrate-to-bsl.md`

## 用户原始意图摘要

将当前仓库许可证从 Apache 2.0 改为 BSL 协议，并同步更新正式公开文档中的对应表述。

## 影响范围

- 后端：无
- 前端：本次不改界面授权文案与展示位
- 文档：`LICENSE`、`README.md`
- 插件 / 模块：无

## 涉及对象

- 模块：仓库根部许可证与说明文档
- 接口：无
- 页面：无
- 配置：无

## 已确认约束

- 使用 BSL 1.1 标准文本
- `Change Date` 设为自本次迁移起 3 年后
- `Change License` 为 Apache License 2.0
- README 继续强调商用需授权
- 本次只做最小可审计迁移，不扩展到前端展示位和 CI

## 当前进展

- 已完成许可迁移设计确认
- 已写入设计文档：`docs/superpowers/specs/2026-05-27-license-bsl-design.md`
- 待按设计修改 `LICENSE` 与 `README.md`

## 后续待办

- 将 `LICENSE` 替换为 BSL 1.1 文本
- 更新 `README.md` 中 Apache 2.0 的正式许可证表述
- 复查正式文档中的授权语义一致性

## 更新规则

- 同一需求始终维护在同一个文件中
- 新信息优先补充到对应段落，不要另起一份重复记录
- 只有需求状态变化时，才在 `active/` 与 `done/` 之间移动文件
