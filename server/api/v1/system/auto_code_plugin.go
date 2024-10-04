package system

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodePluginApi struct{}

// Install
// @Tags      AutoCodePlugin
// @Summary   安装插件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     plug  formData  file                                              true  "this is a test file"
// @Success   200   {object}  response.Response{data=[]interface{},msg=string}  "安装插件成功"
// @Router    /autoCode/installPlugin [post]
func (a *AutoCodePluginApi) Install(c *gin.Context) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := autoCodePluginService.Install(header)
	webStr := global.Translate("system.auto_code_plugin.webPluginInstalledSuccess")
	serverStr := global.Translate("system.auto_code_plugin.serverPluginInstalledSuccess")
	if web == -1 {
		webStr = global.Translate("system.auto_code_plugin.webPluginInstallFail")
	}
	if server == -1 {
		serverStr = global.Translate("system.auto_code_plugin.serverPluginInstallFail")
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData([]interface{}{
		gin.H{
			"code": web,
			"msg":  webStr,
		},
		gin.H{
			"code": server,
			"msg":  serverStr,
		}}, c)
}
xpe
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     plugName  query    string  true  "插件名称"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/pubPlug [post]
func (a *AutoCodePluginApi) Packaged(c *gin.Context) {
	plugName := c.Query("plugName")
	zipPath, err := autoCodePluginService.PubPlug(plugName)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.auto_code_plugin.packingFailed"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.auto_code_plugin.packingFailedError")+err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf(global.Translate("system.auto_code_plugin.packingSuccess")+"%s", zipPath), c)
}

// Packaged
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/initMenu [post]
func (a *AutoCodePluginApi) InitMenu(c *gin.Context) {
	var menuInfo request.InitMenu
	err := c.ShouldBindJSON(&menuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodePluginService.InitMenu(menuInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.auto_code_plugin.initMenuCreationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.auto_code_plugin.initMenuCreationFailError") + err.Error(), c)
		return
	}
	response.OkWithMessage(global.Translate("system.auto_code_plugin.fileChangedSuccessfully"), c)
}

// Packaged
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/initAPI [post]
func (a *AutoCodePluginApi) InitAPI(c *gin.Context) {
	var apiInfo request.InitApi
	err := c.ShouldBindJSON(&apiInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodePluginService.InitAPI(apiInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.auto_code_plugin.initApiCreationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.auto_code_plugin.initApiCreationFailError") + err.Error(), c)
		return
	}
	response.OkWithMessage(global.Translate("system.auto_code_plugin.fileChangedSuccessfully"), c)
}
