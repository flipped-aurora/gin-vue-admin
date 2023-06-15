package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/test"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysUsersApi struct {
}

var sysUsersService = service.ServiceGroupApp.TestServiceGroup.SysUsersService


// CreateSysUsers 创建SysUsers
// @Tags SysUsers
// @Summary 创建SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.SysUsers true "创建SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/createSysUsers [post]
func (sysUsersApi *SysUsersApi) CreateSysUsers(c *gin.Context) {
	var sysUsers test.SysUsers
	err := c.ShouldBindJSON(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysUsersService.CreateSysUsers(&sysUsers); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysUsers 删除SysUsers
// @Tags SysUsers
// @Summary 删除SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.SysUsers true "删除SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUsers/deleteSysUsers [delete]
func (sysUsersApi *SysUsersApi) DeleteSysUsers(c *gin.Context) {
	var sysUsers test.SysUsers
	err := c.ShouldBindJSON(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysUsersService.DeleteSysUsers(sysUsers); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysUsersByIds 批量删除SysUsers
// @Tags SysUsers
// @Summary 批量删除SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysUsers/deleteSysUsersByIds [delete]
func (sysUsersApi *SysUsersApi) DeleteSysUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysUsersService.DeleteSysUsersByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysUsers 更新SysUsers
// @Tags SysUsers
// @Summary 更新SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.SysUsers true "更新SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUsers/updateSysUsers [put]
func (sysUsersApi *SysUsersApi) UpdateSysUsers(c *gin.Context) {
	var sysUsers test.SysUsers
	err := c.ShouldBindJSON(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysUsersService.UpdateSysUsers(sysUsers); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysUsers 用id查询SysUsers
// @Tags SysUsers
// @Summary 用id查询SysUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.SysUsers true "用id查询SysUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUsers/findSysUsers [get]
func (sysUsersApi *SysUsersApi) FindSysUsers(c *gin.Context) {
	var sysUsers test.SysUsers
	err := c.ShouldBindQuery(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysUsers, err := sysUsersService.GetSysUsers(sysUsers.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysUsers": resysUsers}, c)
	}
}

// GetSysUsersList 分页获取SysUsers列表
// @Tags SysUsers
// @Summary 分页获取SysUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.SysUsersSearch true "分页获取SysUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/getSysUsersList [get]
func (sysUsersApi *SysUsersApi) GetSysUsersList(c *gin.Context) {
	var pageInfo testReq.SysUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysUsersService.GetSysUsersInfoList(pageInfo); err != nil {
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
