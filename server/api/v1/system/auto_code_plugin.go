package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
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
	webStr := "web插件安装成功"
	serverStr := "server插件安装成功"
	if web == -1 {
		webStr = "web端插件未成功安装，请按照文档自行解压安装，如果为纯后端插件请忽略此条提示"
	}
	if server == -1 {
		serverStr = "server端插件未成功安装，请按照文档自行解压安装，如果为纯前端插件请忽略此条提示"
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

// Packaged
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "打包插件"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/pubPlug [get]
func (a *AutoCodePluginApi) Packaged(c *gin.Context) {
	plugName := c.Query("plugName")
	zipPath, err := autoCodePluginService.PubPlug(plugName)
	if err != nil {
		global.GVA_LOG.Error("打包失败!", zap.Error(err))
		response.FailWithMessage("打包失败"+err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("打包成功,文件路径为:%s", zipPath), c)
}
