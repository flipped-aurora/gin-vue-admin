package request

import (
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysMcpSearch struct {
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
	commonReq.PageInfo
}

type CreateSysMcpRequest struct {
	Name          string `json:"name" form:"name"`
	DisplayName   string `json:"displayName" form:"displayName"`
	Description   string `json:"description" form:"description"`
	Status        string `json:"status" form:"status"`
	Version       string `json:"version" form:"version"`
	ScenariosJSON string `json:"scenariosJson" form:"scenariosJson"`
}

type UpdateSysMcpRequest struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	DisplayName   string `json:"displayName" form:"displayName"`
	Description   string `json:"description" form:"description"`
	Status        string `json:"status" form:"status"`
	Version       string `json:"version" form:"version"`
	ScenariosJSON string `json:"scenariosJson" form:"scenariosJson"`
}

type FindSysMcpRequest struct {
	ID uint `json:"id" form:"id"`
}

type DeleteSysMcpRequest struct {
	ID uint `json:"id" form:"id"`
}

// SysMcpApiBindingItem 一个绑定项 = 一个 API + 可选覆盖（仿 SysCliApiBindingItem）
type SysMcpApiBindingItem struct {
	ApiID            uint   `json:"apiId" form:"apiId"`
	CommandName      string `json:"commandName" form:"commandName"`
	CommandDesc      string `json:"commandDesc" form:"commandDesc"`
	ParamsOverride   string `json:"paramsOverride" form:"paramsOverride"`
	ApiBrief         string `json:"apiBrief" form:"apiBrief"`
	ResponseOverride string `json:"responseOverride" form:"responseOverride"`
	Enabled          bool   `json:"enabled" form:"enabled"`
	Sort             int    `json:"sort" form:"sort"`
}

type AddSysMcpApisRequest struct {
	McpID    uint                  `json:"mcpId" form:"mcpId"`
	Bindings []SysMcpApiBindingItem `json:"bindings" form:"bindings"`
}

type RemoveSysMcpApisRequest struct {
	McpID  uint   `json:"mcpId" form:"mcpId"`
	ApiIDs []uint `json:"apiIds" form:"apiIds"`
}
