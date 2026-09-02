package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysVersionApi struct{}

// buildMenuTree 构建菜单树结构
func buildMenuTree(menus []system.SysBaseMenu) []system.SysBaseMenu {
	// 创建菜单映射
	menuMap := make(map[uint]*system.SysBaseMenu)
	for i := range menus {
		menuMap[menus[i].ID] = &menus[i]
	}

	// 构建树结构
	var rootMenus []system.SysBaseMenu
	for _, menu := range menus {
		if menu.ParentId == 0 {
			// 根菜单
			menuData := convertMenuToStruct(menu, menuMap)
			rootMenus = append(rootMenus, menuData)
		}
	}

	// 按sort排序根菜单
	sort.Slice(rootMenus, func(i, j int) bool {
		return rootMenus[i].Sort < rootMenus[j].Sort
	})

	return rootMenus
}

// convertMenuToStruct 将菜单转换为结构体并递归处理子菜单
func convertMenuToStruct(menu system.SysBaseMenu, menuMap map[uint]*system.SysBaseMenu) system.SysBaseMenu {
	result := system.SysBaseMenu{
		Path:      menu.Path,
		Name:      menu.Name,
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Meta:      menu.Meta,
	}

	// 清理并复制参数数据
	if len(menu.Parameters) > 0 {
		cleanParameters := make([]system.SysBaseMenuParameter, 0, len(menu.Parameters))
		for _, param := range menu.Parameters {
			cleanParam := system.SysBaseMenuParameter{
				Type:  param.Type,
				Key:   param.Key,
				Value: param.Value,
				// 不复制 ID, CreatedAt, UpdatedAt, SysBaseMenuID
			}
			cleanParameters = append(cleanParameters, cleanParam)
		}
		result.Parameters = cleanParameters
	}

	// 清理并复制菜单按钮数据
	if len(menu.MenuBtn) > 0 {
		cleanMenuBtns := make([]system.SysBaseMenuBtn, 0, len(menu.MenuBtn))
		for _, btn := range menu.MenuBtn {
			cleanBtn := system.SysBaseMenuBtn{
				Name: btn.Name,
				Desc: btn.Desc,
				// 不复制 ID, CreatedAt, UpdatedAt, SysBaseMenuID
			}
			cleanMenuBtns = append(cleanMenuBtns, cleanBtn)
		}
		result.MenuBtn = cleanMenuBtns
	}

	// 查找并处理子菜单
	var children []system.SysBaseMenu
	for _, childMenu := range menuMap {
		if childMenu.ParentId == menu.ID {
			childData := convertMenuToStruct(*childMenu, menuMap)
			children = append(children, childData)
		}
	}

	// 按sort排序子菜单
	if len(children) > 0 {
		sort.Slice(children, func(i, j int) bool {
			return children[i].Sort < children[j].Sort
		})
		result.Children = children
	}

	return result
}

// DeleteSysVersion 删除版本管理
// @Tags SysVersion
// @Summary 删除版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysVersion true "删除版本管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysVersion/deleteSysVersion [delete]
func (sysVersionApi *SysVersionApi) DeleteSysVersion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := sysVersionService.DeleteSysVersion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysVersionByIds 批量删除版本管理
// @Tags SysVersion
// @Summary 批量删除版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysVersion/deleteSysVersionByIds [delete]
func (sysVersionApi *SysVersionApi) DeleteSysVersionByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := sysVersionService.DeleteSysVersionByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// FindSysVersion 用id查询版本管理
// @Tags SysVersion
// @Summary 用id查询版本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询版本管理"
// @Success 200 {object} response.Response{data=system.SysVersion,msg=string} "查询成功"
// @Router /sysVersion/findSysVersion [get]
func (sysVersionApi *SysVersionApi) FindSysVersion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resysVersion, err := sysVersionService.GetSysVersion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysVersion, c)
}

// GetSysVersionList 分页获取版本管理列表
// @Tags SysVersion
// @Summary 分页获取版本管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysVersionSearch true "分页获取版本管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysVersion/getSysVersionList [get]
func (sysVersionApi *SysVersionApi) GetSysVersionList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.SysVersionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysVersionService.GetSysVersionInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSysVersionPublic 不需要鉴权的版本管理接口
// @Tags SysVersion
// @Summary 不需要鉴权的版本管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysVersion/getSysVersionPublic [get]
func (sysVersionApi *SysVersionApi) GetSysVersionPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysVersionService.GetSysVersionPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的版本管理接口信息",
	}, "获取成功", c)
}

// ExportVersion 创建发版数据
// @Tags SysVersion
// @Summary 创建发版数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.ExportVersionRequest true "创建发版数据"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysVersion/exportVersion [post]
func (sysVersionApi *SysVersionApi) ExportVersion(c *gin.Context) {
	ctx := c.Request.Context()

	var req systemReq.ExportVersionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取选中的菜单数据
	var menuData []system.SysBaseMenu
	if len(req.MenuIds) > 0 {
		menuData, err = sysVersionService.GetMenusByIds(ctx, req.MenuIds)
		if err != nil {
			global.GVA_LOG.Error("获取菜单数据失败!", zap.Error(err))
			response.FailWithMessage("获取菜单数据失败:"+err.Error(), c)
			return
		}
	}

	// 获取选中的API数据
	var apiData []system.SysApi
	if len(req.ApiIds) > 0 {
		apiData, err = sysVersionService.GetApisByIds(ctx, req.ApiIds)
		if err != nil {
			global.GVA_LOG.Error("获取API数据失败!", zap.Error(err))
			response.FailWithMessage("获取API数据失败:"+err.Error(), c)
			return
		}
	}

	// 获取选中的字典数据
	var dictData []system.SysDictionary
	if len(req.DictIds) > 0 {
		dictData, err = sysVersionService.GetDictionariesByIds(ctx, req.DictIds)
		if err != nil {
			global.GVA_LOG.Error("获取字典数据失败!", zap.Error(err))
			response.FailWithMessage("获取字典数据失败:"+err.Error(), c)
			return
		}
	}

	// 处理菜单数据，构建递归的children结构
	processedMenus := buildMenuTree(menuData)

	// 处理API数据，清除ID和时间戳字段
	processedApis := make([]system.SysApi, 0, len(apiData))
	for _, api := range apiData {
		cleanApi := system.SysApi{
			Path:        api.Path,
			Description: api.Description,
			ApiGroup:    api.ApiGroup,
			Method:      api.Method,
		}
		processedApis = append(processedApis, cleanApi)
	}

	// 处理字典数据，清除ID和时间戳字段，包含字典详情
	processedDicts := make([]system.SysDictionary, 0, len(dictData))
	for _, dict := range dictData {
		cleanDict := system.SysDictionary{
			Name:   dict.Name,
			Type:   dict.Type,
			Status: dict.Status,
			Desc:   dict.Desc,
		}
		
		// 处理字典详情数据，清除ID和时间戳字段
		cleanDetails := make([]system.SysDictionaryDetail, 0, len(dict.SysDictionaryDetails))
		for _, detail := range dict.SysDictionaryDetails {
			cleanDetail := system.SysDictionaryDetail{
				Label:  detail.Label,
				Value:  detail.Value,
				Extend: detail.Extend,
				Status: detail.Status,
				Sort:   detail.Sort,
				// 不复制 ID, CreatedAt, UpdatedAt, SysDictionaryID
			}
			cleanDetails = append(cleanDetails, cleanDetail)
		}
		cleanDict.SysDictionaryDetails = cleanDetails
		
		processedDicts = append(processedDicts, cleanDict)
	}

	// 构建导出数据
	exportData := systemRes.ExportVersionResponse{
		Version: systemReq.VersionInfo{
			Name:        req.VersionName,
			Code:        req.VersionCode,
			Description: req.Description,
			ExportTime:  time.Now().Format("2006-01-02 15:04:05"),
		},
		Menus:        processedMenus,
		Apis:         processedApis,
		Dictionaries: processedDicts,
	}

	// 转换为JSON
	jsonData, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		global.GVA_LOG.Error("JSON序列化失败!", zap.Error(err))
		response.FailWithMessage("JSON序列化失败:"+err.Error(), c)
		return
	}

	// 保存版本记录
	version := system.SysVersion{
		VersionName: utils.Pointer(req.VersionName),
		VersionCode: utils.Pointer(req.VersionCode),
		Description: utils.Pointer(req.Description),
		VersionData: utils.Pointer(string(jsonData)),
	}

	err = sysVersionService.CreateSysVersion(ctx, &version)
	if err != nil {
		global.GVA_LOG.Error("保存版本记录失败!", zap.Error(err))
		response.FailWithMessage("保存版本记录失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建发版成功", c)
}

// DownloadVersionJson 下载版本JSON数据
// @Tags SysVersion
// @Summary 下载版本JSON数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query string true "版本ID"
// @Success 200 {object} response.Response{data=object,msg=string} "下载成功"
// @Router /sysVersion/downloadVersionJson [get]
func (sysVersionApi *SysVersionApi) DownloadVersionJson(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("版本ID不能为空", c)
		return
	}

	// 获取版本记录
	version, err := sysVersionService.GetSysVersion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("获取版本记录失败!", zap.Error(err))
		response.FailWithMessage("获取版本记录失败:"+err.Error(), c)
		return
	}

	// 构建JSON数据
	var jsonData []byte
	if version.VersionData != nil && *version.VersionData != "" {
		jsonData = []byte(*version.VersionData)
	} else {
		// 如果没有存储的JSON数据，构建一个基本的结构
		basicData := systemRes.ExportVersionResponse{
			Version: systemReq.VersionInfo{
				Name:        *version.VersionName,
				Code:        *version.VersionCode,
				Description: *version.Description,
				ExportTime:  version.CreatedAt.Format("2006-01-02 15:04:05"),
			},
			Menus: []system.SysBaseMenu{},
			Apis:  []system.SysApi{},
		}
		jsonData, _ = json.MarshalIndent(basicData, "", "  ")
	}

	// 设置下载响应头
	filename := fmt.Sprintf("version_%s_%s.json", *version.VersionCode, time.Now().Format("20060102150405"))
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Length", strconv.Itoa(len(jsonData)))

	c.Data(http.StatusOK, "application/json", jsonData)
}

// ImportVersion 导入版本数据
// @Tags SysVersion
// @Summary 导入版本数据
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.ImportVersionRequest true "版本JSON数据"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /sysVersion/importVersion [post]
func (sysVersionApi *SysVersionApi) ImportVersion(c *gin.Context) {
	ctx := c.Request.Context()

	// 获取JSON数据
	var importData systemReq.ImportVersionRequest
	err := c.ShouldBindJSON(&importData)
	if err != nil {
		response.FailWithMessage("解析JSON数据失败:"+err.Error(), c)
		return
	}

	// 验证数据格式
	if importData.VersionInfo.Name == "" || importData.VersionInfo.Code == "" {
		response.FailWithMessage("版本信息格式错误", c)
		return
	}

	// 导入菜单数据
	if len(importData.ExportMenu) > 0 {
		if err := sysVersionService.ImportMenus(ctx, importData.ExportMenu); err != nil {
			global.GVA_LOG.Error("导入菜单失败!", zap.Error(err))
			response.FailWithMessage("导入菜单失败: "+err.Error(), c)
			return
		}
	}

	// 导入API数据
	if len(importData.ExportApi) > 0 {
		if err := sysVersionService.ImportApis(importData.ExportApi); err != nil {
			global.GVA_LOG.Error("导入API失败!", zap.Error(err))
			response.FailWithMessage("导入API失败: "+err.Error(), c)
			return
		}
	}

	// 导入字典数据
	if len(importData.ExportDictionary) > 0 {
		if err := sysVersionService.ImportDictionaries(importData.ExportDictionary); err != nil {
			global.GVA_LOG.Error("导入字典失败!", zap.Error(err))
			response.FailWithMessage("导入字典失败: "+err.Error(), c)
			return
		}
	}

	// 创建导入记录
	jsonData, _ := json.Marshal(importData)
	version := system.SysVersion{
		VersionName: utils.Pointer(importData.VersionInfo.Name),
		VersionCode: utils.Pointer(fmt.Sprintf("%s_imported_%s", importData.VersionInfo.Code, time.Now().Format("20060102150405"))),
		Description: utils.Pointer(fmt.Sprintf("导入版本: %s", importData.VersionInfo.Description)),
		VersionData: utils.Pointer(string(jsonData)),
	}

	err = sysVersionService.CreateSysVersion(ctx, &version)
	if err != nil {
		global.GVA_LOG.Error("保存导入记录失败!", zap.Error(err))
		// 这里不返回错误，因为数据已经导入成功
	}

	response.OkWithMessage("导入成功", c)
}
