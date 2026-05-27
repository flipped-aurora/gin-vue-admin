# GVA 仓库 BSL 迁移设计

## 目标

将当前仓库根部公开许可证从 Apache License 2.0 迁移为 Business Source License 1.1（BSL 1.1），并保证正式公开文档中的授权表述一致。

本次设计基于以下已确认约束：

- 采用 **BSL 1.1 标准文本**，不自定义许可主体结构
- `Change Date` 设为 **3 年后**
- 到期后的 `Change License` 为 **Apache License 2.0**
- README 中继续**强调商用需授权**
- 本次仅做**最小可审计迁移**，不扩展到前端界面文案、Banner、控制台提示或 CI 规则

## 变更范围

### 需要修改的文件

1. [LICENSE](../../LICENSE)
2. [README.md](../../README.md)
3. `aiDoc/memory/business/active/license-migrate-to-bsl.md`
4. [aiDoc/memory/business/demand-index.md](../../aiDoc/memory/business/demand-index.md)

### 不在本次范围内的内容

以下内容即使与“购买授权”相关，也不在本次最小迁移范围内：

- `web/src/core/config.js`
- `web/src/core/gin-vue-admin.js`
- `web/src/view/dashboard/`
- `web/src/view/systemTools/`
- 其他前端展示位、购买按钮、Banner、控制台提示
- CI、发布流程、构建校验
- 贡献者协议、额外的贡献条款治理

## 文案与许可设计

### LICENSE 文件

`LICENSE` 由当前 Apache 2.0 全文替换为 **BSL 1.1 标准文本**。

应填入的仓库级参数：

- `Licensed Work`: gin-vue-admin 仓库当前发布内容
- `Change Date`: 自本次迁移时点起 3 年后
- `Change License`: Apache License, Version 2.0

设计要求：

- 使用标准 BSL 1.1 文本，不自行增删条款
- 保留仓库主体的版权归属信息
- 让 `LICENSE` 单独即可表达“当前受 BSL 约束、未来自动转 Apache 2.0”的完整语义

### README 正式文案

README 中当前仍引用 Apache 2.0 的正式表述需要统一迁移为 BSL 语义，重点位置包括：

- 商业使用提示段
- 末尾注意事项段

文案目标：

- 明确当前许可证是 **BSL 1.1**
- 明确**商业/生产使用需遵循 BSL 规则并按仓库说明获取商业授权**
- 保留现有授权购买入口链接
- 删除“当前许可证仍是 Apache 2.0”的冲突表达

### 授权语义边界

本次 README 文案不追求扩展新的商业限制口径，只做与 BSL 迁移一致的公开合规更新。

也就是说：

- 不新增复杂 FAQ
- 不定义额外的前端展示策略
- 不新增多版本授权对照说明
- 不把 README 扩写成完整法务文档

## 数据与流程

本次迁移不涉及业务数据、数据库结构或前后端接口。

流程上只包含：

1. 更新正式许可证文本
2. 更新 README 中与许可证直接相关的公开说明
3. 记录业务需求记忆，确保后续协作能沿用这次已确认的许可方向

## 错误处理与风险控制

### 风险

1. `LICENSE` 已切到 BSL，但 README 仍保留 Apache 2.0 口径，造成公开表述冲突
2. README 改写过度，意外扩展了超出本次确认范围的商业限制
3. 仅改一处 README，遗漏另一处正式提醒段，导致仓库内正式文案不一致

### 控制方式

- 只修改正式许可证与 README 中明确指向许可证的段落
- 不碰前端界面和其他展示位，避免范围膨胀
- 修改后复查仓库根部正式文档，确认不再把 Apache 2.0 表述为“当前许可证”

## 验收标准

本次设计完成后的实现应满足：

1. [LICENSE](../../LICENSE) 不再是 Apache 2.0 全文，而是 BSL 1.1 正式文本
2. [README.md](../../README.md) 中正式许可证说明不再把 Apache 2.0 描述为当前许可证
3. README 继续保留商业授权入口，并明确商用需授权
4. 本次改动不触达前端授权展示位与业务代码
5. `aiDoc/memory/business/` 中存在对应需求记录，并在需求索引中可检索

## 测试与验证

由于本次仅涉及许可证与文档层修改，不需要运行前后端测试或手工 UI 验证。

验证方式为：

- 逐项检查 `LICENSE` 与 `README` 正式文案是否一致
- 搜索仓库根部正式文档中是否仍存在把 Apache 2.0 作为当前许可证的描述
- 检查业务需求记忆文件与索引是否同步更新
