package service

import (
	"context"
	"errors"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/request"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
	"gorm.io/gorm"
)

type mcService struct{}

func (s *mcService) CreateMcp(ctx context.Context, req autoReq.CreateSysMcpRequest) (autoModel.SysMcp, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return autoModel.SysMcp{}, errors.New("mcp名称不能为空")
	}
	if exists, err := s.sysMcpNameExists(ctx, name, 0); err != nil {
		return autoModel.SysMcp{}, err
	} else if exists {
		return autoModel.SysMcp{}, errors.New("存在同名MCP")
	}
	entity := autoModel.SysMcp{
		Name:          name,
		DisplayName:   strings.TrimSpace(req.DisplayName),
		Description:   strings.TrimSpace(req.Description),
		Status:        strings.TrimSpace(req.Status),
		Version:       strings.TrimSpace(req.Version),
		ScenariosJSON: req.ScenariosJSON,
	}
	applySysMcpDefaults(&entity)
	if err := global.GVA_DB.WithContext(ctx).Create(&entity).Error; err != nil {
		return autoModel.SysMcp{}, err
	}
	return entity, nil
}

func (s *mcService) UpdateMcp(ctx context.Context, req autoReq.UpdateSysMcpRequest) (autoModel.SysMcp, error) {
	entity, err := s.getSysMcpByID(ctx, req.ID)
	if err != nil {
		return autoModel.SysMcp{}, err
	}
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return autoModel.SysMcp{}, errors.New("mcp名称不能为空")
	}
	if exists, err := s.sysMcpNameExists(ctx, name, req.ID); err != nil {
		return autoModel.SysMcp{}, err
	} else if exists {
		return autoModel.SysMcp{}, errors.New("存在同名MCP")
	}
	entity.Name = name
	entity.DisplayName = strings.TrimSpace(req.DisplayName)
	entity.Description = strings.TrimSpace(req.Description)
	entity.Status = strings.TrimSpace(req.Status)
	entity.Version = strings.TrimSpace(req.Version)
	entity.ScenariosJSON = req.ScenariosJSON
	applySysMcpDefaults(&entity)
	if err := global.GVA_DB.WithContext(ctx).Save(&entity).Error; err != nil {
		return autoModel.SysMcp{}, err
	}
	return entity, nil
}

func (s *mcService) DeleteMcp(ctx context.Context, req autoReq.DeleteSysMcpRequest) error {
	if _, err := s.getSysMcpByID(ctx, req.ID); err != nil {
		return err
	}
	return global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&autoModel.SysMcpApi{}, "mcp_id = ?", req.ID).Error; err != nil {
			return err
		}
		return tx.Delete(&autoModel.SysMcp{}, "id = ?", req.ID).Error
	})
}

func (s *mcService) GetMcpList(ctx context.Context, req autoReq.SysMcpSearch) (autoRes.SysMcpListResponse, error) {
	res := autoRes.SysMcpListResponse{List: []autoRes.SysMcpListItem{}, Page: req.Page, PageSize: req.PageSize}
	base := global.GVA_DB.WithContext(ctx).Model(&autoModel.SysMcp{})
	if name := strings.TrimSpace(req.Name); name != "" {
		base = base.Where("name LIKE ? OR display_name LIKE ?", "%"+name+"%", "%"+name+"%")
	}
	if status := strings.TrimSpace(req.Status); status != "" {
		base = base.Where("status = ?", status)
	}
	if err := base.Count(&res.Total).Error; err != nil {
		return res, err
	}
	type mcpListRow struct {
		autoModel.SysMcp
		ApiCount int64 `gorm:"column:api_count"`
	}
	query := base.Select("sys_mcps.*, COUNT(sys_mcp_apis.id) AS api_count").
		Joins("LEFT JOIN sys_mcp_apis ON sys_mcp_apis.mcp_id = sys_mcps.id AND sys_mcp_apis.deleted_at IS NULL").
		Group("sys_mcps.id").
		Order("sys_mcps.id DESC")
	if req.PageSize > 0 {
		offset := req.PageSize * (req.Page - 1)
		query = query.Offset(offset).Limit(req.PageSize)
	}
	var rows []mcpListRow
	if err := query.Find(&rows).Error; err != nil {
		return res, err
	}
	for _, row := range rows {
		res.List = append(res.List, autoRes.SysMcpListItem{SysMcp: row.SysMcp, ApiCount: row.ApiCount})
	}
	return res, nil
}

func (s *mcService) GetMcpDetail(ctx context.Context, req autoReq.FindSysMcpRequest) (autoRes.SysMcpDetailResponse, error) {
	return s.getMcpDetailByID(ctx, req.ID)
}

func (s *mcService) AddMcpApis(ctx context.Context, req autoReq.AddSysMcpApisRequest) (autoRes.SysMcpDetailResponse, error) {
	if _, err := s.getSysMcpByID(ctx, req.McpID); err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	if len(req.Bindings) == 0 {
		return s.getMcpDetailByID(ctx, req.McpID)
	}
	apiIDs := make([]uint, 0, len(req.Bindings))
	seen := make(map[uint]struct{}, len(req.Bindings))
	for _, binding := range req.Bindings {
		if binding.ApiID == 0 {
			return autoRes.SysMcpDetailResponse{}, errors.New("apiId不能为空")
		}
		if _, ok := seen[binding.ApiID]; ok {
			continue
		}
		seen[binding.ApiID] = struct{}{}
		apiIDs = append(apiIDs, binding.ApiID)
	}
	if err := s.ensureApisExist(ctx, apiIDs); err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	if err := global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range req.Bindings {
			var existing autoModel.SysMcpApi
			err := tx.Unscoped().Where("mcp_id = ? AND api_id = ?", req.McpID, item.ApiID).First(&existing).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				entity := autoModel.SysMcpApi{McpID: req.McpID, ApiID: item.ApiID, CommandName: strings.TrimSpace(item.CommandName), CommandDesc: strings.TrimSpace(item.CommandDesc), ParamsOverride: strings.TrimSpace(item.ParamsOverride), ApiBrief: strings.TrimSpace(item.ApiBrief), ResponseOverride: strings.TrimSpace(item.ResponseOverride), Enabled: item.Enabled, Sort: item.Sort}
				if err := tx.Create(&entity).Error; err != nil {
					return err
				}
				continue
			}
			if err != nil {
				return err
			}
			if err := tx.Unscoped().Model(&existing).Updates(map[string]any{"command_name": strings.TrimSpace(item.CommandName), "command_desc": strings.TrimSpace(item.CommandDesc), "params_override": strings.TrimSpace(item.ParamsOverride), "api_brief": strings.TrimSpace(item.ApiBrief), "response_override": strings.TrimSpace(item.ResponseOverride), "enabled": item.Enabled, "sort": item.Sort, "deleted_at": nil}).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	return s.getMcpDetailByID(ctx, req.McpID)
}

func (s *mcService) RemoveMcpApis(ctx context.Context, req autoReq.RemoveSysMcpApisRequest) (autoRes.SysMcpDetailResponse, error) {
	if _, err := s.getSysMcpByID(ctx, req.McpID); err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	if len(req.ApiIDs) == 0 {
		return s.getMcpDetailByID(ctx, req.McpID)
	}
	if err := global.GVA_DB.WithContext(ctx).Unscoped().Delete(&autoModel.SysMcpApi{}, "mcp_id = ? AND api_id IN ?", req.McpID, req.ApiIDs).Error; err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	return s.getMcpDetailByID(ctx, req.McpID)
}

// BuildCommandsFromBindings 把启用的 MCP 绑定转成能力定义（command 列表），复用 CLI 的 manifest 管线。
// 供 MCP 动态注册：每个 command 对应一个 MCP tool。
func (s *mcService) BuildCommandsFromBindings(ctx context.Context) ([]autoRes.SysCliManifestCommand, error) {
	bindings, err := s.ListEnabledBindings(ctx)
	if err != nil {
		return nil, err
	}
	if len(bindings) == 0 {
		return []autoRes.SysCliManifestCommand{}, nil
	}
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return nil, err
	}
	commands := make([]autoRes.SysCliManifestCommand, 0, len(bindings))
	usedNames := make(map[string]int)
	// 复用 CLI 管线：传零值 SysCli（管线只用 binding.Api + 覆盖项，SysCli 仅影响示例前缀）
	zeroCli := autoModel.SysCli{}
	for _, b := range bindings {
		cliBinding := sysCliManifestBinding{Binding: b.Cli, Api: b.Api}
		command, err := buildSysCliManifestCommand(zeroCli, spec, cliBinding, usedNames)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

// McpPromptDef 一个 MCP 编排 prompt 的定义：name + 描述 + 渲染好的 markdown 指引。
type McpPromptDef struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Markdown    string `json:"markdown"`
}

// BuildMcpPrompts 遍历所有 MCP，把有 scenariosJson 的渲染成 prompt 定义列表。
// 供 MCP 进程启动时注册成 MCP prompt（AI 客户端调 prompts/list 可见）。
func (s *mcService) BuildMcpPrompts(ctx context.Context) ([]McpPromptDef, error) {
	var mcps []autoModel.SysMcp
	if err := global.GVA_DB.WithContext(ctx).Where("status = ?", "enabled").Order("id ASC").Find(&mcps).Error; err != nil {
		return nil, err
	}
	prompts := make([]McpPromptDef, 0, len(mcps))
	for _, mcp := range mcps {
		raw := strings.TrimSpace(mcp.ScenariosJSON)
		if raw == "" {
			continue
		}
		scenarios, err := parseCliScenarios(raw)
		if err != nil || len(scenarios) == 0 {
			continue
		}
		md := renderScenariosMarkdown(scenarios, false)
		if strings.TrimSpace(md) == "" {
			continue
		}
		name := "orchestration_" + mcp.Name
		desc := mcp.DisplayName + " 的 tool 编排场景"
		if strings.TrimSpace(mcp.Description) != "" {
			desc = mcp.Description
		}
		prompts = append(prompts, McpPromptDef{Name: name, Description: desc, Markdown: md})
	}
	return prompts, nil
}

// PreviewMcpPrompt 按单个 mcpId 渲染编排 prompt（markdown），供后台预览。
// 复用 parseCliScenarios + renderScenariosMarkdown。
func (s *mcService) PreviewMcpPrompt(ctx context.Context, mcpID uint) (McpPromptDef, error) {
	mcp, err := s.getSysMcpByID(ctx, mcpID)
	if err != nil {
		return McpPromptDef{}, err
	}
	raw := strings.TrimSpace(mcp.ScenariosJSON)
	if raw == "" {
		return McpPromptDef{Name: "orchestration_" + mcp.Name, Description: mcp.DisplayName, Markdown: ""}, nil
	}
	scenarios, err := parseCliScenarios(raw)
	if err != nil {
		return McpPromptDef{}, err
	}
	md := renderScenariosMarkdown(scenarios, false)
	name := "orchestration_" + mcp.Name
	desc := mcp.DisplayName + " 的 tool 编排场景"
	if strings.TrimSpace(mcp.Description) != "" {
		desc = mcp.Description
	}
	return McpPromptDef{Name: name, Description: desc, Markdown: md}, nil
}
func (s *mcService) PreviewMcpManifest(ctx context.Context, mcpID uint) ([]autoRes.SysCliManifestCommand, error) {
	mcp, bindings, err := s.getMcpAndBindings(ctx, mcpID)
	if err != nil {
		return nil, err
	}
	_ = mcp // mcp 本身不影响 command 生成（管线只用 binding.Api + 覆盖项）
	if len(bindings) == 0 {
		return []autoRes.SysCliManifestCommand{}, nil
	}
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return nil, err
	}
	commands := make([]autoRes.SysCliManifestCommand, 0, len(bindings))
	usedNames := make(map[string]int)
	zeroCli := autoModel.SysCli{}
	for _, b := range bindings {
		if !b.Cli.Enabled {
			continue
		}
		cliBinding := sysCliManifestBinding{Binding: b.Cli, Api: b.Api}
		command, err := buildSysCliManifestCommand(zeroCli, spec, cliBinding, usedNames)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

// PreviewApiCommand 按单个 apiId 实时生成能力定义（command），供"自动生成"按钮调用。
// 仿 CLI 的 PreviewApiCommand，但不需要 mcpId——只按 apiId 查 SysApi + 管线派生。
func (s *mcService) PreviewApiCommand(ctx context.Context, apiID uint) (autoRes.SysCliManifestCommand, error) {
	api, err := s.getSysApiByID(ctx, apiID)
	if err != nil {
		return autoRes.SysCliManifestCommand{}, err
	}
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return autoRes.SysCliManifestCommand{}, err
	}
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{ApiID: apiID, Enabled: true},
		Api:     api,
	}
	zeroCli := autoModel.SysCli{}
	return buildSysCliManifestCommand(zeroCli, spec, binding, map[string]int{})
}

func (s *mcService) ListEnabledBindings(ctx context.Context) ([]mcpBinding, error) {
	var bindings []autoModel.SysMcpApi
	if err := global.GVA_DB.WithContext(ctx).Where("enabled = ?", true).Order("mcp_id ASC, sort ASC, id ASC").Find(&bindings).Error; err != nil {
		return nil, err
	}
	if len(bindings) == 0 {
		return []mcpBinding{}, nil
	}
	apiIDs := make([]uint, 0, len(bindings))
	for _, b := range bindings {
		apiIDs = append(apiIDs, b.ApiID)
	}
	var apis []sysModel.SysApi
	if err := global.GVA_DB.WithContext(ctx).Where("id IN ?", apiIDs).Find(&apis).Error; err != nil {
		return nil, err
	}
	apiMap := make(map[uint]sysModel.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.ID] = api
	}
	result := make([]mcpBinding, 0, len(bindings))
	for _, b := range bindings {
		api, ok := apiMap[b.ApiID]
		if !ok {
			continue
		}
		result = append(result, mcpBinding{Raw: b, Cli: b.ToCliApi(), Api: api})
	}
	return result, nil
}

func (s *mcService) getMcpDetailByID(ctx context.Context, id uint) (autoRes.SysMcpDetailResponse, error) {
	mcp, err := s.getSysMcpByID(ctx, id)
	if err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	bindings, err := s.listMcpBindings(ctx, id)
	if err != nil {
		return autoRes.SysMcpDetailResponse{}, err
	}
	res := autoRes.SysMcpDetailResponse{Mcp: mcp, Bindings: make([]autoRes.SysMcpBoundApi, 0, len(bindings))}
	for _, b := range bindings {
		res.Bindings = append(res.Bindings, autoRes.SysMcpBoundApi{SysMcpApi: b.Raw, Api: b.Api})
	}
	return res, nil
}

// mcpBinding 同时持有原始 SysMcpApi（供详情返回）和转成 CLI 格式的 manifest binding（供管线复用）
type mcpBinding struct {
	Raw    autoModel.SysMcpApi
	Cli    autoModel.SysCliApi
	Api    sysModel.SysApi
}

func (s *mcService) getMcpAndBindings(ctx context.Context, id uint) (autoModel.SysMcp, []mcpBinding, error) {
	mcp, err := s.getSysMcpByID(ctx, id)
	if err != nil {
		return autoModel.SysMcp{}, nil, err
	}
	bindings, err := s.listMcpBindings(ctx, id)
	if err != nil {
		return autoModel.SysMcp{}, nil, err
	}
	return mcp, bindings, nil
}

func (s *mcService) getSysMcpByID(ctx context.Context, id uint) (autoModel.SysMcp, error) {
	var mcp autoModel.SysMcp
	if id == 0 {
		return mcp, errors.New("mcpId不能为空")
	}
	if err := global.GVA_DB.WithContext(ctx).First(&mcp, "id = ?", id).Error; err != nil {
		return mcp, err
	}
	return mcp, nil
}

func (s *mcService) sysMcpNameExists(ctx context.Context, name string, excludeID uint) (bool, error) {
	var existing autoModel.SysMcp
	query := global.GVA_DB.WithContext(ctx).Where("name = ?", name)
	if excludeID > 0 {
		query = query.Where("id <> ?", excludeID)
	}
	if err := query.First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *mcService) ensureApisExist(ctx context.Context, apiIDs []uint) error {
	var count int64
	if err := global.GVA_DB.WithContext(ctx).Model(&sysModel.SysApi{}).Where("id IN ?", apiIDs).Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(apiIDs)) {
		return errors.New("存在无效API")
	}
	return nil
}

func (s *mcService) getSysApiByID(ctx context.Context, id uint) (sysModel.SysApi, error) {
	var api sysModel.SysApi
	if id == 0 {
		return api, errors.New("apiId不能为空")
	}
	if err := global.GVA_DB.WithContext(ctx).First(&api, "id = ?", id).Error; err != nil {
		return api, err
	}
	return api, nil
}

func (s *mcService) listMcpBindings(ctx context.Context, mcpID uint) ([]mcpBinding, error) {
	var bindings []autoModel.SysMcpApi
	if err := global.GVA_DB.WithContext(ctx).Where("mcp_id = ?", mcpID).Order("sort ASC, id ASC, api_id ASC").Find(&bindings).Error; err != nil {
		return nil, err
	}
	if len(bindings) == 0 {
		return []mcpBinding{}, nil
	}
	apiIDs := make([]uint, 0, len(bindings))
	for _, b := range bindings {
		apiIDs = append(apiIDs, b.ApiID)
	}
	var apis []sysModel.SysApi
	if err := global.GVA_DB.WithContext(ctx).Where("id IN ?", apiIDs).Find(&apis).Error; err != nil {
		return nil, err
	}
	apiMap := make(map[uint]sysModel.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.ID] = api
	}
	result := make([]mcpBinding, 0, len(bindings))
	for _, b := range bindings {
		api, ok := apiMap[b.ApiID]
		if !ok {
			continue
		}
		result = append(result, mcpBinding{Raw: b, Cli: b.ToCliApi(), Api: api})
	}
	return result, nil
}

func applySysMcpDefaults(mcp *autoModel.SysMcp) {
	if mcp == nil {
		return
	}
	if strings.TrimSpace(mcp.DisplayName) == "" {
		mcp.DisplayName = mcp.Name
	}
	if strings.TrimSpace(mcp.Version) == "" {
		mcp.Version = "v1"
	}
	if strings.TrimSpace(mcp.Status) == "" {
		mcp.Status = "enabled"
	}
}
