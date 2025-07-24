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
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysVersionApi struct{}

// buildMenuTree 构建菜单树结构
func buildMenuTree(menus []system.SysBaseMenu) []map[string]interface{} {
	// 创建菜单映射
	menuMap := make(map[uint]*system.SysBaseMenu)
	for i := range menus {
		menuMap[menus[i].ID] = &menus[i]
	}

	// 构建树结构
	var rootMenus []map[string]interface{}
	for _, menu := range menus {
		if menu.ParentId == 0 {
			// 根菜单
			menuData := convertMenuToMap(menu, menuMap)
			rootMenus = append(rootMenus, menuData)
		}
	}

	// 按sort排序根菜单
	sort.Slice(rootMenus, func(i, j int) bool {
		sortI, _ := rootMenus[i]["sort"].(int)
		sortJ, _ := rootMenus[j]["sort"].(int)
		return sortI < sortJ
	})

	return rootMenus
}

// convertMenuToMap 将菜单转换为map并递归处理子菜单
func convertMenuToMap(menu system.SysBaseMenu, menuMap map[uint]*system.SysBaseMenu) map[string]interface{} {
	result := map[string]interface{}{
		"path":      menu.Path,
		"name":      menu.Name,
		"hidden":    menu.Hidden,
		"component": menu.Component,
		"sort":      menu.Sort,
		"meta":      menu.Meta,
	}

	// 处理参数
	if len(menu.Parameters) > 0 {
		params := make([]map[string]interface{}, 0)
		for _, param := range menu.Parameters {
			params = append(params, map[string]interface{}{
				"type":  param.Type,
				"key":   param.Key,
				"value": param.Value,
			})
		}
		result["parameters"] = params
	}

	// 处理菜单按钮
	if len(menu.MenuBtn) > 0 {
		btns := make([]map[string]interface{}, 0)
		for _, btn := range menu.MenuBtn {
			btns = append(btns, map[string]interface{}{
				"name": btn.Name,
				"desc": btn.Desc,
			})
		}
		result["menuBtn"] = btns
	}

	// 查找并处理子菜单
	var children []map[string]interface{}
	for _, childMenu := range menuMap {
		if childMenu.ParentId == menu.ID {
			childData := convertMenuToMap(*childMenu, menuMap)
			children = append(children, childData)
		}
	}

	// 按sort排序子菜单
	if len(children) > 0 {
		sort.Slice(children, func(i, j int) bool {
			sortI, _ := children[i]["sort"].(int)
			sortJ, _ := children[j]["sort"].(int)
			return sortI < sortJ
		})
		result["children"] = children
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
	userID := utils.GetUserID(c)
	err := sysVersionService.DeleteSysVersion(ctx, ID, userID)
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
	userID := utils.GetUserID(c)
	err := sysVersionService.DeleteSysVersionByIds(ctx, IDs, userID)
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

	// 处理菜单数据，构建递归的children结构
	processedMenus := buildMenuTree(menuData)

	// 处理API数据，移除ID字段
	processedApis := make([]map[string]interface{}, 0)
	for _, api := range apiData {
		apiMap := map[string]interface{}{
			"path":        api.Path,
			"description": api.Description,
			"apiGroup":    api.ApiGroup,
			"method":      api.Method,
		}
		processedApis = append(processedApis, apiMap)
	}

	// 构建导出数据
	exportData := map[string]interface{}{
		"version": map[string]interface{}{
			"name":        req.VersionName,
			"code":        req.VersionCode,
			"description": req.Description,
			"exportTime":  time.Now().Format("2006-01-02 15:04:05"),
		},
		"menus": processedMenus,
		"apis":  processedApis,
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
	version.CreatedBy = utils.GetUserID(c)

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
		basicData := map[string]interface{}{
			"version": map[string]interface{}{
				"name":        *version.VersionName,
				"code":        *version.VersionCode,
				"description": *version.Description,
				"exportTime":  version.CreatedAt.Format("2006-01-02 15:04:05"),
			},
			"menus": []interface{}{},
			"apis":  []interface{}{},
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
// @Param data body map[string]interface{} true "版本JSON数据"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /sysVersion/importVersion [post]
func (sysVersionApi *SysVersionApi) ImportVersion(c *gin.Context) {
	ctx := c.Request.Context()

	// 获取JSON数据
	var importData map[string]interface{}
	err := c.ShouldBindJSON(&importData)
	if err != nil {
		response.FailWithMessage("解析JSON数据失败:"+err.Error(), c)
		return
	}

	// 验证数据格式
	versionInfo, ok := importData["version"].(map[string]interface{})
	if !ok {
		response.FailWithMessage("版本信息格式错误", c)
		return
	}

	// 导入菜单数据
	if menusData, exists := importData["menus"]; exists {
		err = sysVersionService.ImportMenus(ctx, menusData)
		if err != nil {
			global.GVA_LOG.Error("导入菜单失败!", zap.Error(err))
			response.FailWithMessage("导入菜单失败:"+err.Error(), c)
			return
		}
	}

	// 导入API数据
	if apisData, exists := importData["apis"]; exists {
		err = sysVersionService.ImportApis(ctx, apisData)
		if err != nil {
			global.GVA_LOG.Error("导入API失败!", zap.Error(err))
			response.FailWithMessage("导入API失败:"+err.Error(), c)
			return
		}
	}

	// 创建导入记录
	jsonData, _ := json.Marshal(importData)
	version := system.SysVersion{
		VersionName: utils.Pointer(fmt.Sprintf("%v", versionInfo["name"])),
		VersionCode: utils.Pointer(fmt.Sprintf("%v_imported_%s", versionInfo["code"], time.Now().Format("20060102150405"))),
		Description: utils.Pointer(fmt.Sprintf("导入版本: %v", versionInfo["description"])),
		VersionData: utils.Pointer(string(jsonData)),
	}
	version.CreatedBy = utils.GetUserID(c)

	err = sysVersionService.CreateSysVersion(ctx, &version)
	if err != nil {
		global.GVA_LOG.Error("保存导入记录失败!", zap.Error(err))
		// 这里不返回错误，因为数据已经导入成功
	}

	response.OkWithMessage("导入成功", c)
}
