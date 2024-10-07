package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type AutoCodePackageApi struct{}

// Create
// @Tags      AutoCodePackage
// @Summary   创建package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAutoCodePackageCreate                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/createPackage [post]
func (a *AutoCodePackageApi) Create(c *gin.Context) {
	var info request.SysAutoCodePackageCreate
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if strings.Contains(info.PackageName, "\\") || strings.Contains(info.PackageName, "/") || strings.Contains(info.PackageName, "..") {
		response.FailWithMessage(global.Translate("general.invalidPackageName"), c)
		return
	} // PackageName可能导致路径穿越的问题 / 和 \ 都要防止
	err := autoCodePackageService.Create(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccss"), c)
}

// Delete
// @Tags      AutoCode
// @Summary   删除package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.GetById                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "删除package成功"
// @Router    /autoCode/delPackage [post]
func (a *AutoCodePackageApi) Delete(c *gin.Context) {
	var info common.GetById
	_ = c.ShouldBindJSON(&info)
	err := autoCodePackageService.Delete(c.Request.Context(), info)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deletFailErr"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// All
// @Tags      AutoCodePackage
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getPackage [post]
func (a *AutoCodePackageApi) All(c *gin.Context) {
	data, err := autoCodePackageService.All(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFailErr"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
		return
	}
	response.OkWithDetailed(gin.H{"pkgs": data}, global.Translate("general.getDataSuccess"), c)
}

// Templates
// @Tags      AutoCodePackage
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getTemplates [get]
func (a *AutoCodePackageApi) Templates(c *gin.Context) {
	data, err := autoCodePackageService.Templates(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFailErr"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
		return
	}
	response.OkWithDetailed(data, global.Translate("general.getDataSuccess"), c)
}
