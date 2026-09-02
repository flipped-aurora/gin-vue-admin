package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	_ = ctx
	entities := []model.SysApi{
		// 代码生成器
		{Path: "/autoCode/getDB", Description: "获取数据库列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/getTables", Description: "获取数据表列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/getColumn", Description: "获取字段列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/preview", Description: "预览自动代码", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/createTemp", Description: "生成自动代码", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcp", Description: "生成 MCP 工具", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStatus", Description: "获取 MCP 状态", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStart", Description: "启动 MCP", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStop", Description: "停止 MCP", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpList", Description: "获取 MCP 工具列表", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpRoutes", Description: "获取 MCP 路由", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpTest", Description: "测试 MCP 调用", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/pubPlug", Description: "打包插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/installPlugin", Description: "安装插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/removePlugin", Description: "移除插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getPluginList", Description: "获取插件列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/saveAIWorkflowSession", Description: "保存 AI 工作流会话", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getAIWorkflowSessionList", Description: "获取 AI 工作流列表", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getAIWorkflowSessionDetail", Description: "获取 AI 工作流详情", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/deleteAIWorkflowSession", Description: "删除 AI 工作流会话", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/dumpAIWorkflowMarkdown", Description: "导出 AI 工作流 Markdown", ApiGroup: "代码生成器", Method: "POST"},

		// 模板配置
		{Path: "/autoCode/getPackage", Description: "获取自动化包列表", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/delPackage", Description: "删除自动化包", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/createPackage", Description: "创建自动化包", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/getTemplates", Description: "获取模板列表", ApiGroup: "模板配置", Method: "GET"},

		// 代码生成器历史
		{Path: "/autoCode/getMeta", Description: "获取自动代码历史元数据", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/rollback", Description: "回滚自动代码历史", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/delSysHistory", Description: "删除自动代码历史", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/getSysHistory", Description: "获取自动代码历史列表", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/addFunc", Description: "追加自动代码方法", ApiGroup: "代码生成器历史", Method: "POST"},

		// skills
		{Path: "/skills/getTools", Description: "获取 AI 工具列表", ApiGroup: "skills", Method: "GET"},
		{Path: "/skills/getSkillList", Description: "获取技能列表", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getSkillDetail", Description: "获取技能详情", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveSkill", Description: "保存技能", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/deleteSkill", Description: "删除技能", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/createScript", Description: "创建脚本", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getScript", Description: "获取脚本", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveScript", Description: "保存脚本", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/createResource", Description: "创建资源", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getResource", Description: "获取资源", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveResource", Description: "保存资源", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/createReference", Description: "创建参考资料", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getReference", Description: "获取参考资料", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveReference", Description: "保存参考资料", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/createTemplate", Description: "创建模板", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getTemplate", Description: "获取模板", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveTemplate", Description: "保存模板", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/getGlobalConstraint", Description: "获取全局约束", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/saveGlobalConstraint", Description: "保存全局约束", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/packageSkill", Description: "打包技能", ApiGroup: "skills", Method: "POST"},
		{Path: "/skills/downloadOnlineSkill", Description: "下载在线技能", ApiGroup: "skills", Method: "POST"},
	}
	utils.RegisterApis(entities...)
}
