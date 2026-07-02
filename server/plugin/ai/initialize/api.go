package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	_ = ctx
	entities := []model.SysApi{
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

		// MCP Tools（处理逻辑仍在 core，路径前缀沿用 autoCode）
		{Path: "/autoCode/mcp", Description: "生成 MCP 工具", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStatus", Description: "获取 MCP 状态", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStart", Description: "启动 MCP", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpStop", Description: "停止 MCP", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpList", Description: "获取 MCP 工具列表", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpRoutes", Description: "获取 MCP 路由", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/mcpTest", Description: "测试 MCP 调用", ApiGroup: "代码生成器", Method: "POST"},

		// AI 工作流（处理逻辑仍在 core，路径前缀沿用 autoCode）
		{Path: "/autoCode/saveAIWorkflowSession", Description: "保存 AI 工作流会话", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getAIWorkflowSessionList", Description: "获取 AI 工作流列表", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getAIWorkflowSessionDetail", Description: "获取 AI 工作流详情", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/deleteAIWorkflowSession", Description: "删除 AI 工作流会话", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/dumpAIWorkflowMarkdown", Description: "导出 AI 工作流 Markdown", ApiGroup: "代码生成器", Method: "POST"},

		// cli
		{Path: "/cli/createCli", Description: "创建CLI", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/getCliList", Description: "获取CLI列表", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/getCliDetail", Description: "获取CLI详情", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/updateCli", Description: "更新CLI", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/deleteCli", Description: "删除CLI", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/addCliApis", Description: "增加CLI关联API", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/removeCliApis", Description: "减少CLI关联API", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/previewManifest", Description: "预览CLI Manifest", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/downloadManifest", Description: "下载CLI Manifest", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/buildCli", Description: "编译并下载CLI二进制", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/downloadSkill", Description: "下载CLI的AI Skill", ApiGroup: "CLI管理", Method: "POST"},
		{Path: "/cli/previewApiCommand", Description: "填充API命令", ApiGroup: "CLI管理", Method: "POST"},

		// mcpApi（动态 tool 注册）
		{Path: "/mcpApi/createMcp", Description: "创建MCP", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/getMcpList", Description: "获取MCP列表", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/getMcpDetail", Description: "获取MCP详情", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/updateMcp", Description: "更新MCP", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/deleteMcp", Description: "删除MCP", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/addMcpApis", Description: "增加MCP关联API", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/removeMcpApis", Description: "减少MCP关联API", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/previewManifest", Description: "预览MCP能力定义", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/previewPrompt", Description: "预览MCP编排prompt", ApiGroup: "MCP管理", Method: "POST"},
		{Path: "/mcpApi/previewApiCommand", Description: "按API生成能力定义", ApiGroup: "MCP管理", Method: "POST"},
	}
	utils.RegisterApis(entities...)
}
