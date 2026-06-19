package response

import (
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysCliListItem struct {
	autoModel.SysCli
	ApiCount int64 `json:"apiCount"`
}

type SysCliListResponse struct {
	List     []SysCliListItem `json:"list"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
}

type SysCliBoundApi struct {
	autoModel.SysCliApi
	Api sysModel.SysApi `json:"api"`
}

type SysCliDetailResponse struct {
	Cli      autoModel.SysCli  `json:"cli"`
	Bindings []SysCliBoundApi `json:"bindings"`
}

type SysCliManifestResponse struct {
	Name     string                  `json:"name"`
	Version  string                  `json:"version"`
	Server   SysCliManifestServer    `json:"server"`
	Commands []SysCliManifestCommand `json:"commands"`
}

type SysCliManifestServer struct {
	BaseURL    string `json:"baseURL"`
	AuthHeader string `json:"authHeader"`
}

type SysCliManifestCommand struct {
	Name        string                        `json:"name"`
	Summary     string                        `json:"summary"`
	Description string                        `json:"description"`
	Method      string                        `json:"method"`
	Path        string                        `json:"path"`
	ContentType string                        `json:"contentType"`
	Source      SysCliManifestSource          `json:"source"`
	Parameters  []SysCliManifestParameter     `json:"parameters"`
	Response    []SysCliManifestResponseField `json:"response"`
	Examples    []string                      `json:"examples"`
	Warnings    []string                      `json:"warnings"`
}

type SysCliManifestSource struct {
	ApiID    uint   `json:"apiId"`
	ApiGroup string `json:"apiGroup"`
	From     string `json:"from"`
}

type SysCliManifestParameter struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Field       string `json:"field"`
}

type SysCliManifestResponseField struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
