package system

import (
	"fmt"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliApi struct{}

// CreateCli 创建CLI
// @Tags Cli
// @Summary 创建CLI
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.CreateSysCliRequest true "CLI基础信息"
// @Success 200 {object} response.Response{data=object,msg=string} "创建成功"
// @Router /cli/createCli [post]
func (a *CliApi) CreateCli(c *gin.Context) {
	var req systemReq.CreateSysCliRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.CreateCli(req)
	if err != nil {
		global.GVA_LOG.Error("创建CLI失败!", zap.Error(err))
		response.FailWithMessage("创建CLI失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "创建成功", c)
}

// GetCliList 获取CLI列表
// @Tags Cli
// @Summary 获取CLI列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.SysCliSearch true "查询条件"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cli/getCliList [post]
func (a *CliApi) GetCliList(c *gin.Context) {
	var req systemReq.SysCliSearch
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.GetCliList(req)
	if err != nil {
		global.GVA_LOG.Error("获取CLI列表失败!", zap.Error(err))
		response.FailWithMessage("获取CLI列表失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: res.List, Total: res.Total, Page: res.Page, PageSize: res.PageSize}, "获取成功", c)
}

// GetCliDetail 获取CLI详情
// @Tags Cli
// @Summary 获取CLI详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.FindSysCliRequest true "CLI ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cli/getCliDetail [post]
func (a *CliApi) GetCliDetail(c *gin.Context) {
	var req systemReq.FindSysCliRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.GetCliDetail(req)
	if err != nil {
		global.GVA_LOG.Error("获取CLI详情失败!", zap.Error(err))
		response.FailWithMessage("获取CLI详情失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "获取成功", c)
}

// UpdateCli 更新CLI
// @Tags Cli
// @Summary 更新CLI
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.UpdateSysCliRequest true "CLI基础信息"
// @Success 200 {object} response.Response{data=object,msg=string} "更新成功"
// @Router /cli/updateCli [post]
func (a *CliApi) UpdateCli(c *gin.Context) {
	var req systemReq.UpdateSysCliRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.UpdateCli(req)
	if err != nil {
		global.GVA_LOG.Error("更新CLI失败!", zap.Error(err))
		response.FailWithMessage("更新CLI失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "更新成功", c)
}

// DeleteCli 删除CLI
// @Tags Cli
// @Summary 删除CLI
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.DeleteSysCliRequest true "CLI ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cli/deleteCli [post]
func (a *CliApi) DeleteCli(c *gin.Context) {
	var req systemReq.DeleteSysCliRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cliService.DeleteCli(req); err != nil {
		global.GVA_LOG.Error("删除CLI失败!", zap.Error(err))
		response.FailWithMessage("删除CLI失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// AddCliApis 增加CLI关联API
// @Tags Cli
// @Summary 增加CLI关联API
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.AddSysCliApisRequest true "CLI绑定API"
// @Success 200 {object} response.Response{data=object,msg=string} "保存成功"
// @Router /cli/addCliApis [post]
func (a *CliApi) AddCliApis(c *gin.Context) {
	var req systemReq.AddSysCliApisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.AddCliApis(req)
	if err != nil {
		global.GVA_LOG.Error("增加CLI关联API失败!", zap.Error(err))
		response.FailWithMessage("增加CLI关联API失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "保存成功", c)
}

// RemoveCliApis 减少CLI关联API
// @Tags Cli
// @Summary 减少CLI关联API
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.RemoveSysCliApisRequest true "CLI解绑API"
// @Success 200 {object} response.Response{data=object,msg=string} "移除成功"
// @Router /cli/removeCliApis [post]
func (a *CliApi) RemoveCliApis(c *gin.Context) {
	var req systemReq.RemoveSysCliApisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.RemoveCliApis(req)
	if err != nil {
		global.GVA_LOG.Error("减少CLI关联API失败!", zap.Error(err))
		response.FailWithMessage("减少CLI关联API失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "移除成功", c)
}

// PreviewManifest 预览CLI Manifest
// @Tags Cli
// @Summary 预览CLI Manifest
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.PreviewSysCliManifestRequest true "CLI ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cli/previewManifest [post]
func (a *CliApi) PreviewManifest(c *gin.Context) {
	var req systemReq.PreviewSysCliManifestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.PreviewManifest(req.CliID)
	if err != nil {
		global.GVA_LOG.Error("预览CLI Manifest失败!", zap.Error(err))
		response.FailWithMessage("预览CLI Manifest失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "获取成功", c)
}

// DownloadManifest 下载CLI Manifest
// @Tags Cli
// @Summary 下载CLI Manifest
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body systemReq.PreviewSysCliManifestRequest true "CLI ID"
// @Success 200 {file} file "manifest json"
// @Router /cli/downloadManifest [post]
func (a *CliApi) DownloadManifest(c *gin.Context) {
	var req systemReq.PreviewSysCliManifestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileName, payload, err := cliService.DownloadManifest(req.CliID)
	if err != nil {
		global.GVA_LOG.Error("下载CLI Manifest失败!", zap.Error(err))
		response.FailWithMessage("下载CLI Manifest失败: "+err.Error(), c)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	c.Data(http.StatusOK, "application/json", payload)
}

// BuildCliBinary 编译并下载可直接运行的 CLI 二进制
// @Tags Cli
// @Summary 编译并下载CLI二进制
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/octet-stream
// @Param data body systemReq.BuildSysCliBinaryRequest true "CLI ID 与目标平台"
// @Success 200 {file} file "cli binary"
// @Router /cli/buildCli [post]
func (a *CliApi) BuildCliBinary(c *gin.Context) {
	var req systemReq.BuildSysCliBinaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileName, payload, err := cliService.BuildCliBinary(req)
	if err != nil {
		global.GVA_LOG.Error("编译CLI二进制失败!", zap.Error(err))
		response.FailWithMessage("编译CLI二进制失败: "+err.Error(), c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	c.Data(http.StatusOK, "application/octet-stream", payload)
}

func (a *CliApi) FindCli(c *gin.Context) {
	var req commonReq.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := cliService.GetCliDetail(systemReq.FindSysCliRequest{ID: uint(req.ID)})
	if err != nil {
		global.GVA_LOG.Error("获取CLI详情失败!", zap.Error(err))
		response.FailWithMessage("获取CLI详情失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "获取成功", c)
}
