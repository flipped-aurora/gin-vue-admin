# 插件示例

这个目录用于说明插件级别的完整组织方式。

## 包含内容

- `full-plugin-example.md`: 一个插件从目录结构到初始化入口的大图示例

## 适用范围

当 AI 需要新增、修改、重构或生成以下内容时，必须在编辑前先读本目录：

- `server/plugin/<name>/`
- `web/src/plugin/<name>/`
- `server/resource/plugin/server/`
- `server/resource/plugin/web/`

同时必须阅读 `../../modules/plugin-development.md`、涉及层的当前生成模板和对应分层示例，并在完成必读清单后、文件编辑前的工作更新中列出已读内容。真实插件只可在示例明确限定的文件与职责范围内参考，不得整目录照抄。
