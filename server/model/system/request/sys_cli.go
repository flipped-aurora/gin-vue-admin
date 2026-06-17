package request

import (
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysCliSearch struct {
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
	commonReq.PageInfo
}

type CreateSysCliRequest struct {
	Name             string `json:"name" form:"name"`
	Command          string `json:"command" form:"command"`
	DisplayName      string `json:"displayName" form:"displayName"`
	Version          string `json:"version" form:"version"`
	Description      string `json:"description" form:"description"`
	Status           string `json:"status" form:"status"`
	SkillName        string `json:"skillName" form:"skillName"`
	SkillDescription string `json:"skillDescription" form:"skillDescription"`
}

type UpdateSysCliRequest struct {
	ID               uint   `json:"id" form:"id"`
	Name             string `json:"name" form:"name"`
	Command          string `json:"command" form:"command"`
	DisplayName      string `json:"displayName" form:"displayName"`
	Version          string `json:"version" form:"version"`
	Description      string `json:"description" form:"description"`
	Status           string `json:"status" form:"status"`
	SkillName        string `json:"skillName" form:"skillName"`
	SkillDescription string `json:"skillDescription" form:"skillDescription"`
}

type FindSysCliRequest struct {
	ID uint `json:"id" form:"id"`
}

type DeleteSysCliRequest struct {
	ID uint `json:"id" form:"id"`
}

type SysCliApiBindingItem struct {
	ApiID       uint   `json:"apiId" form:"apiId"`
	CommandName string `json:"commandName" form:"commandName"`
	Enabled     bool   `json:"enabled" form:"enabled"`
	Sort        int    `json:"sort" form:"sort"`
}

type AddSysCliApisRequest struct {
	CliID    uint                   `json:"cliId" form:"cliId"`
	Bindings []SysCliApiBindingItem `json:"bindings" form:"bindings"`
}

type RemoveSysCliApisRequest struct {
	CliID  uint   `json:"cliId" form:"cliId"`
	ApiIDs []uint `json:"apiIds" form:"apiIds"`
}

type PreviewSysCliManifestRequest struct {
	CliID uint `json:"cliId" form:"cliId"`
}

type BuildSysCliBinaryRequest struct {
	CliID   uint   `json:"cliId" form:"cliId"`
	GOOS    string `json:"goos" form:"goos"`
	GOARCH  string `json:"goarch" form:"goarch"`
	BaseURL string `json:"baseUrl" form:"baseUrl"`
}
