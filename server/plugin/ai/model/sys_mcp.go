package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// SysMcp 一组 MCP 配置（一个 MCP 实例 + 它对外暴露的 tool 集合）
type SysMcp struct {
	global.GVA_MODEL
	Name          string `json:"name" gorm:"column:name;size:128;not null;uniqueIndex;comment:MCP唯一标识"`
	DisplayName   string `json:"displayName" gorm:"column:display_name;size:128;not null;comment:MCP展示名称"`
	Description   string `json:"description" gorm:"column:description;type:text;comment:MCP描述"`
	Status        string `json:"status" gorm:"column:status;size:32;not null;default:enabled;comment:MCP状态"`
	Version       string `json:"version" gorm:"column:version;size:64;not null;default:v1;comment:MCP版本"`
	ScenariosJSON string `json:"scenariosJson" gorm:"column:scenarios_json;type:text;comment:tool间编排场景JSON"`
}

func (SysMcp) TableName() string {
	return "sys_mcps"
}

// SysMcpApi MCP 与 API 的绑定关系。每个绑定的 API 会被动态注册为一个 MCP tool。
type SysMcpApi struct {
	global.GVA_MODEL
	McpID            uint   `json:"mcpId" gorm:"column:mcp_id;not null;uniqueIndex:idx_mcp_api;comment:MCP ID"`
	ApiID            uint   `json:"apiId" gorm:"column:api_id;not null;uniqueIndex:idx_mcp_api;comment:API ID"`
	CommandName      string `json:"commandName" gorm:"column:command_name;size:128;comment:工具名覆盖"`
	CommandDesc      string `json:"commandDesc" gorm:"column:command_desc;type:text;comment:工具说明覆盖"`
	ParamsOverride   string `json:"paramsOverride" gorm:"column:params_override;type:text;comment:参数定义覆盖JSON"`
	ApiBrief         string `json:"apiBrief" gorm:"column:api_brief;size:255;comment:API简介覆盖"`
	ResponseOverride string `json:"responseOverride" gorm:"column:response_override;type:text;comment:返回字段定义覆盖JSON"`
	Enabled          bool   `json:"enabled" gorm:"column:enabled;not null;default:true;comment:是否启用"`
	Sort             int    `json:"sort" gorm:"column:sort;not null;default:0;comment:排序"`
}

func (SysMcpApi) TableName() string {
	return "sys_mcp_apis"
}

// ToCliApi 把 MCP 绑定转成 CLI 绑定格式，复用 CLI 的 manifest 生成管线。
// 字段一一对应（SysMcpApi 与 SysCliApi 的覆盖项结构完全一致），仅表/CliID/McpID 不同。
func (m SysMcpApi) ToCliApi() SysCliApi {
	return SysCliApi{
		CliID:            m.McpID, // 复用 CliID 字段占位（管线只用 ApiID + 覆盖项，不用 CliID 的语义）
		ApiID:            m.ApiID,
		CommandName:      m.CommandName,
		CommandDesc:      m.CommandDesc,
		ParamsOverride:   m.ParamsOverride,
		ApiBrief:         m.ApiBrief,
		ResponseOverride: m.ResponseOverride,
		Enabled:          m.Enabled,
		Sort:             m.Sort,
	}
}
