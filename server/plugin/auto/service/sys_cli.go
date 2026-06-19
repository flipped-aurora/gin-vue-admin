package service

import (
	"errors"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model/request"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model/response"
	"gorm.io/gorm"
)

type cliService struct{}

func (s *cliService) CreateCli(req autoReq.CreateSysCliRequest) (autoModel.SysCli, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return autoModel.SysCli{}, errors.New("cli名称不能为空")
	}
	if exists, err := s.sysCliNameExists(name, 0); err != nil {
		return autoModel.SysCli{}, err
	} else if exists {
		return autoModel.SysCli{}, errors.New("存在同名CLI")
	}
	entity := autoModel.SysCli{
		Name:             name,
		Command:          strings.TrimSpace(req.Command),
		DisplayName:      strings.TrimSpace(req.DisplayName),
		Version:          strings.TrimSpace(req.Version),
		Description:      strings.TrimSpace(req.Description),
		Status:           strings.TrimSpace(req.Status),
		AuthMode:         "jwt",
		SkillName:        strings.TrimSpace(req.SkillName),
		SkillDescription: strings.TrimSpace(req.SkillDescription),
	}
	applySysCliDefaults(&entity)
	if err := global.GVA_DB.Create(&entity).Error; err != nil {
		return autoModel.SysCli{}, err
	}
	return entity, nil
}

func (s *cliService) UpdateCli(req autoReq.UpdateSysCliRequest) (autoModel.SysCli, error) {
	entity, err := s.getSysCliByID(req.ID)
	if err != nil {
		return autoModel.SysCli{}, err
	}
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return autoModel.SysCli{}, errors.New("cli名称不能为空")
	}
	if exists, err := s.sysCliNameExists(name, req.ID); err != nil {
		return autoModel.SysCli{}, err
	} else if exists {
		return autoModel.SysCli{}, errors.New("存在同名CLI")
	}
	entity.Name = name
	entity.Command = strings.TrimSpace(req.Command)
	entity.DisplayName = strings.TrimSpace(req.DisplayName)
	entity.Version = strings.TrimSpace(req.Version)
	entity.Description = strings.TrimSpace(req.Description)
	entity.Status = strings.TrimSpace(req.Status)
	entity.SkillName = strings.TrimSpace(req.SkillName)
	entity.SkillDescription = strings.TrimSpace(req.SkillDescription)
	applySysCliDefaults(&entity)
	if err := global.GVA_DB.Save(&entity).Error; err != nil {
		return autoModel.SysCli{}, err
	}
	return entity, nil
}

func (s *cliService) DeleteCli(req autoReq.DeleteSysCliRequest) error {
	if _, err := s.getSysCliByID(req.ID); err != nil {
		return err
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&autoModel.SysCliApi{}, "cli_id = ?", req.ID).Error; err != nil {
			return err
		}
		return tx.Delete(&autoModel.SysCli{}, "id = ?", req.ID).Error
	})
}

func (s *cliService) GetCliList(req autoReq.SysCliSearch) (autoRes.SysCliListResponse, error) {
	res := autoRes.SysCliListResponse{List: []autoRes.SysCliListItem{}, Page: req.Page, PageSize: req.PageSize}
	base := global.GVA_DB.Model(&autoModel.SysCli{})
	if name := strings.TrimSpace(req.Name); name != "" {
		base = base.Where("name LIKE ? OR display_name LIKE ?", "%"+name+"%", "%"+name+"%")
	}
	if status := strings.TrimSpace(req.Status); status != "" {
		base = base.Where("status = ?", status)
	}
	if err := base.Count(&res.Total).Error; err != nil {
		return res, err
	}
	type cliListRow struct {
		autoModel.SysCli
		ApiCount int64 `gorm:"column:api_count"`
	}
	query := base.Select("sys_clis.*, COUNT(sys_cli_apis.id) AS api_count").
		Joins("LEFT JOIN sys_cli_apis ON sys_cli_apis.cli_id = sys_clis.id AND sys_cli_apis.deleted_at IS NULL").
		Group("sys_clis.id").
		Order("sys_clis.id DESC")
	if req.PageSize > 0 {
		offset := req.PageSize * (req.Page - 1)
		query = query.Offset(offset).Limit(req.PageSize)
	}
	var rows []cliListRow
	if err := query.Find(&rows).Error; err != nil {
		return res, err
	}
	for _, row := range rows {
		res.List = append(res.List, autoRes.SysCliListItem{SysCli: row.SysCli, ApiCount: row.ApiCount})
	}
	return res, nil
}

func (s *cliService) GetCliDetail(req autoReq.FindSysCliRequest) (autoRes.SysCliDetailResponse, error) {
	return s.getCliDetailByID(req.ID)
}

func (s *cliService) AddCliApis(req autoReq.AddSysCliApisRequest) (autoRes.SysCliDetailResponse, error) {
	if _, err := s.getSysCliByID(req.CliID); err != nil {
		return autoRes.SysCliDetailResponse{}, err
	}
	if len(req.Bindings) == 0 {
		return s.getCliDetailByID(req.CliID)
	}
	apiIDs := make([]uint, 0, len(req.Bindings))
	seen := make(map[uint]struct{}, len(req.Bindings))
	for _, binding := range req.Bindings {
		if binding.ApiID == 0 {
			return autoRes.SysCliDetailResponse{}, errors.New("apiId不能为空")
		}
		if _, ok := seen[binding.ApiID]; ok {
			continue
		}
		seen[binding.ApiID] = struct{}{}
		apiIDs = append(apiIDs, binding.ApiID)
	}
	if err := s.ensureApisExist(apiIDs); err != nil {
		return autoRes.SysCliDetailResponse{}, err
	}
	if err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range req.Bindings {
			var existing autoModel.SysCliApi
			err := tx.Unscoped().Where("cli_id = ? AND api_id = ?", req.CliID, item.ApiID).First(&existing).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				entity := autoModel.SysCliApi{CliID: req.CliID, ApiID: item.ApiID, CommandName: strings.TrimSpace(item.CommandName), CommandDesc: strings.TrimSpace(item.CommandDesc), ParamsOverride: strings.TrimSpace(item.ParamsOverride), ApiBrief: strings.TrimSpace(item.ApiBrief), ResponseOverride: strings.TrimSpace(item.ResponseOverride), Enabled: item.Enabled, Sort: item.Sort}
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
		return autoRes.SysCliDetailResponse{}, err
	}
	return s.getCliDetailByID(req.CliID)
}

func (s *cliService) RemoveCliApis(req autoReq.RemoveSysCliApisRequest) (autoRes.SysCliDetailResponse, error) {
	if _, err := s.getSysCliByID(req.CliID); err != nil {
		return autoRes.SysCliDetailResponse{}, err
	}
	if len(req.ApiIDs) == 0 {
		return s.getCliDetailByID(req.CliID)
	}
	if err := global.GVA_DB.Unscoped().Delete(&autoModel.SysCliApi{}, "cli_id = ? AND api_id IN ?", req.CliID, req.ApiIDs).Error; err != nil {
		return autoRes.SysCliDetailResponse{}, err
	}
	return s.getCliDetailByID(req.CliID)
}

func (s *cliService) PreviewManifest(cliID uint) (autoRes.SysCliManifestResponse, error) {
	cli, bindings, err := s.getCliAndBindings(cliID)
	if err != nil {
		return autoRes.SysCliManifestResponse{}, err
	}
	return buildSysCliManifest(cli, bindings)
}

func (s *cliService) PreviewApiCommand(req autoReq.PreviewApiCommandRequest) (autoRes.SysCliManifestCommand, error) {
	api, err := s.getSysApiByID(req.ApiID)
	if err != nil {
		return autoRes.SysCliManifestCommand{}, err
	}
	cli, err := s.getSysCliByID(req.CliID)
	if err != nil {
		return autoRes.SysCliManifestCommand{}, err
	}
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return autoRes.SysCliManifestCommand{}, err
	}
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{ApiID: req.ApiID, Enabled: true},
		Api:     api,
	}
	return buildSysCliManifestCommand(cli, spec, binding, map[string]int{})
}

func (s *cliService) getSysApiByID(id uint) (sysModel.SysApi, error) {
	var api sysModel.SysApi
	if id == 0 {
		return api, errors.New("apiId不能为空")
	}
	if err := global.GVA_DB.First(&api, "id = ?", id).Error; err != nil {
		return api, err
	}
	return api, nil
}

func (s *cliService) DownloadManifest(cliID uint) (string, []byte, error) {
	cli, bindings, err := s.getCliAndBindings(cliID)
	if err != nil {
		return "", nil, err
	}
	manifest, err := buildSysCliManifest(cli, bindings)
	if err != nil {
		return "", nil, err
	}
	payload, err := marshalSysCliManifest(manifest)
	if err != nil {
		return "", nil, err
	}
	return sysCliManifestFileName(cli), payload, nil
}

func (s *cliService) getCliDetailByID(id uint) (autoRes.SysCliDetailResponse, error) {
	cli, bindings, err := s.getCliAndBindings(id)
	if err != nil {
		return autoRes.SysCliDetailResponse{}, err
	}
	res := autoRes.SysCliDetailResponse{Cli: cli, Bindings: make([]autoRes.SysCliBoundApi, 0, len(bindings))}
	for _, binding := range bindings {
		res.Bindings = append(res.Bindings, autoRes.SysCliBoundApi{SysCliApi: binding.Binding, Api: binding.Api})
	}
	return res, nil
}

func (s *cliService) getCliAndBindings(id uint) (autoModel.SysCli, []sysCliManifestBinding, error) {
	cli, err := s.getSysCliByID(id)
	if err != nil {
		return autoModel.SysCli{}, nil, err
	}
	bindings, err := s.listCliBindings(id)
	if err != nil {
		return autoModel.SysCli{}, nil, err
	}
	return cli, bindings, nil
}

func (s *cliService) getSysCliByID(id uint) (autoModel.SysCli, error) {
	var cli autoModel.SysCli
	if id == 0 {
		return cli, errors.New("cliId不能为空")
	}
	if err := global.GVA_DB.First(&cli, "id = ?", id).Error; err != nil {
		return cli, err
	}
	return cli, nil
}

func (s *cliService) sysCliNameExists(name string, excludeID uint) (bool, error) {
	var existing autoModel.SysCli
	query := global.GVA_DB.Where("name = ?", name)
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

func (s *cliService) ensureApisExist(apiIDs []uint) error {
	var count int64
	if err := global.GVA_DB.Model(&sysModel.SysApi{}).Where("id IN ?", apiIDs).Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(apiIDs)) {
		return errors.New("存在无效API")
	}
	return nil
}

func (s *cliService) listCliBindings(cliID uint) ([]sysCliManifestBinding, error) {
	var bindings []autoModel.SysCliApi
	if err := global.GVA_DB.Where("cli_id = ?", cliID).Order("sort ASC, id ASC, api_id ASC").Find(&bindings).Error; err != nil {
		return nil, err
	}
	if len(bindings) == 0 {
		return []sysCliManifestBinding{}, nil
	}
	apiIDs := make([]uint, 0, len(bindings))
	for _, binding := range bindings {
		apiIDs = append(apiIDs, binding.ApiID)
	}
	var apis []sysModel.SysApi
	if err := global.GVA_DB.Where("id IN ?", apiIDs).Find(&apis).Error; err != nil {
		return nil, err
	}
	apiMap := make(map[uint]sysModel.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.ID] = api
	}
	result := make([]sysCliManifestBinding, 0, len(bindings))
	for _, binding := range bindings {
		api, ok := apiMap[binding.ApiID]
		if !ok {
			continue
		}
		result = append(result, sysCliManifestBinding{Binding: binding, Api: api})
	}
	return result, nil
}

func applySysCliDefaults(cli *autoModel.SysCli) {
	if cli == nil {
		return
	}
	if strings.TrimSpace(cli.Command) == "" {
		cli.Command = cli.Name
	}
	if strings.TrimSpace(cli.DisplayName) == "" {
		cli.DisplayName = cli.Name
	}
	if strings.TrimSpace(cli.Version) == "" {
		cli.Version = "v1"
	}
	if strings.TrimSpace(cli.Status) == "" {
		cli.Status = "enabled"
	}
	if strings.TrimSpace(cli.AuthMode) == "" {
		cli.AuthMode = "jwt"
	}
}
