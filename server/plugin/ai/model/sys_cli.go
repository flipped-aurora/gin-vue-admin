package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysCli struct {
	global.GVA_MODEL
	Name             string `json:"name" gorm:"column:name;size:128;not null;uniqueIndex;comment:CLI唯一标识"`
	Command          string `json:"command" gorm:"column:command;size:128;not null;default:'';comment:CLI主命令"`
	DisplayName      string `json:"displayName" gorm:"column:display_name;size:128;not null;comment:CLI展示名称"`
	Version          string `json:"version" gorm:"column:version;size:64;not null;default:v1;comment:CLI版本"`
	Description      string `json:"description" gorm:"column:description;type:text;comment:CLI描述"`
	Status           string `json:"status" gorm:"column:status;size:32;not null;default:enabled;comment:CLI状态"`
	AuthMode         string `json:"authMode" gorm:"column:auth_mode;size:32;not null;default:jwt;comment:认证方式"`
	SkillName        string `json:"skillName" gorm:"column:skill_name;size:128;comment:AI Skill名称"`
	SkillDescription string `json:"skillDescription" gorm:"column:skill_description;type:text;comment:AI Skill描述"`
	ScenariosJSON    string `json:"scenariosJson" gorm:"column:scenarios_json;type:text;comment:调用场景链路JSON"`
}

func (SysCli) TableName() string {
	return "sys_clis"
}

type SysCliApi struct {
	global.GVA_MODEL
	CliID            uint   `json:"cliId" gorm:"column:cli_id;not null;uniqueIndex:idx_cli_api;comment:CLI ID"`
	ApiID            uint   `json:"apiId" gorm:"column:api_id;not null;uniqueIndex:idx_cli_api;comment:API ID"`
	CommandName      string `json:"commandName" gorm:"column:command_name;size:128;comment:命令名覆盖"`
	CommandDesc      string `json:"commandDesc" gorm:"column:command_desc;type:text;comment:命令说明覆盖"`
	ParamsOverride   string `json:"paramsOverride" gorm:"column:params_override;type:text;comment:参数定义覆盖JSON"`
	ApiBrief         string `json:"apiBrief" gorm:"column:api_brief;size:255;comment:API简介覆盖"`
	ResponseOverride string `json:"responseOverride" gorm:"column:response_override;type:text;comment:返回字段定义覆盖JSON"`
	Enabled          bool   `json:"enabled" gorm:"column:enabled;not null;default:true;comment:是否启用"`
	Sort             int    `json:"sort" gorm:"column:sort;not null;default:0;comment:排序"`
}

func (SysCliApi) TableName() string {
	return "sys_cli_apis"
}

// CliScenario 表示一个命名调用场景，是一张由命令/判断节点与条件边组成的图。
type CliScenario struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Sort        int               `json:"sort"`
	Nodes       []CliScenarioNode `json:"nodes"`
	Edges       []CliScenarioEdge `json:"edges"`
}

// CliScenarioNode 节点：command 执行命令；decision 不执行命令，做汇聚+合并判断，靠出边带条件分支。
type CliScenarioNode struct {
	ID          string `json:"id"`                     // 节点 ID，边引用
	Type        string `json:"type"`                   // "command" | "decision"
	CommandName string `json:"commandName,omitempty"`  // command：执行的命令名
	Alias       string `json:"alias,omitempty"`        // 节点别名，场景内唯一；供 InputNote/Condition 用 别名.字段 引用上游出参
	InputNote   string `json:"inputNote,omitempty"`    // 入参来源，用 别名.字段 引用上游，可多源 "create.id + query.token"
	Note        string `json:"note,omitempty"`         // command:说明；decision:合并判断描述
}

// CliScenarioEdge 边：from→to，可带条件；条件非空表示满足该条件才走这条边（分支）。
type CliScenarioEdge struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Condition string `json:"condition,omitempty"` // 空表示默认流转
	Note      string `json:"note,omitempty"`
}
