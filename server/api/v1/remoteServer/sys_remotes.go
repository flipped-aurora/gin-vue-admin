package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    remoteServerReq "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SysRemotesApi struct {
}

var sysRemotesService = service.ServiceGroupApp.RemoteServerServiceGroup.SysRemotesService


// CreateSysRemotes 创建远程服务器配置表
// @Tags SysRemotes
// @Summary 创建远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemotes true "创建远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysRemotes/createSysRemotes [post]
func (sysRemotesApi *SysRemotesApi) CreateSysRemotes(c *gin.Context) {
	var sysRemotes remoteServer.SysRemotes
	err := c.ShouldBindJSON(&sysRemotes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sysRemotes.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Code":{utils.NotEmpty()},
        "IsRemove":{utils.NotEmpty()},
        "Ip":{utils.NotEmpty()},
        "Port":{utils.NotEmpty()},
        "User":{utils.NotEmpty()},
        "Password":{utils.NotEmpty()},
    }
	if err := utils.Verify(sysRemotes, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := sysRemotesService.CreateSysRemotes(&sysRemotes); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysRemotes 删除远程服务器配置表
// @Tags SysRemotes
// @Summary 删除远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemotes true "删除远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRemotes/deleteSysRemotes [delete]
func (sysRemotesApi *SysRemotesApi) DeleteSysRemotes(c *gin.Context) {
	var sysRemotes remoteServer.SysRemotes
	err := c.ShouldBindJSON(&sysRemotes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sysRemotes.DeletedBy = utils.GetUserID(c)
	if err := sysRemotesService.DeleteSysRemotes(sysRemotes); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysRemotesByIds 批量删除远程服务器配置表
// @Tags SysRemotes
// @Summary 批量删除远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysRemotes/deleteSysRemotesByIds [delete]
func (sysRemotesApi *SysRemotesApi) DeleteSysRemotesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := sysRemotesService.DeleteSysRemotesByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysRemotes 更新远程服务器配置表
// @Tags SysRemotes
// @Summary 更新远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemotes true "更新远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysRemotes/updateSysRemotes [put]
func (sysRemotesApi *SysRemotesApi) UpdateSysRemotes(c *gin.Context) {
	var sysRemotes remoteServer.SysRemotes
	err := c.ShouldBindJSON(&sysRemotes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sysRemotes.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Code":{utils.NotEmpty()},
          "IsRemove":{utils.NotEmpty()},
          "Ip":{utils.NotEmpty()},
          "Port":{utils.NotEmpty()},
          "User":{utils.NotEmpty()},
          "Password":{utils.NotEmpty()},
      }
    if err := utils.Verify(sysRemotes, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := sysRemotesService.UpdateSysRemotes(sysRemotes); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysRemotes 用id查询远程服务器配置表
// @Tags SysRemotes
// @Summary 用id查询远程服务器配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query remoteServer.SysRemotes true "用id查询远程服务器配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysRemotes/findSysRemotes [get]
func (sysRemotesApi *SysRemotesApi) FindSysRemotes(c *gin.Context) {
	var sysRemotes remoteServer.SysRemotes
	err := c.ShouldBindQuery(&sysRemotes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysRemotes, err := sysRemotesService.GetSysRemotes(sysRemotes.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysRemotes": resysRemotes}, c)
	}
}

// GetSysRemotesList 分页获取远程服务器配置表列表
// @Tags SysRemotes
// @Summary 分页获取远程服务器配置表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query remoteServerReq.SysRemotesSearch true "分页获取远程服务器配置表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRemotes/getSysRemotesList [get]
func (sysRemotesApi *SysRemotesApi) GetSysRemotesList(c *gin.Context) {
	var pageInfo remoteServerReq.SysRemotesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysRemotesService.GetSysRemotesInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
