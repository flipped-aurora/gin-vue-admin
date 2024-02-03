package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/user"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserInfoApi struct {
}

var userDataService = service.ServiceGroupApp.UserServiceGroup.UserInfoService


// CreateUserInfo 创建用户信息
// @Tags UserInfo
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserInfo true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userData/createUserInfo [post]
func (userDataApi *UserInfoApi) CreateUserInfo(c *gin.Context) {
	var userData user.UserInfo
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userDataService.CreateUserInfo(&userData); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserInfo 删除用户信息
// @Tags UserInfo
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserInfo true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userData/deleteUserInfo [delete]
func (userDataApi *UserInfoApi) DeleteUserInfo(c *gin.Context) {
	id := c.Query("ID")
	if err := userDataService.DeleteUserInfo(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserInfoByIds 批量删除用户信息
// @Tags UserInfo
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userData/deleteUserInfoByIds [delete]
func (userDataApi *UserInfoApi) DeleteUserInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := userDataService.DeleteUserInfoByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserInfo 更新用户信息
// @Tags UserInfo
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserInfo true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userData/updateUserInfo [put]
func (userDataApi *UserInfoApi) UpdateUserInfo(c *gin.Context) {
	var userData user.UserInfo
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userDataService.UpdateUserInfo(userData); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserInfo 用id查询用户信息
// @Tags UserInfo
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserInfo true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userData/findUserInfo [get]
func (userDataApi *UserInfoApi) FindUserInfo(c *gin.Context) {
	id := c.Query("ID")
	if reuserData, err := userDataService.GetUserInfo(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserData": reuserData}, c)
	}
}

// GetUserInfoList 分页获取用户信息列表
// @Tags UserInfo
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserInfoSearch true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userData/getUserInfoList [get]
func (userDataApi *UserInfoApi) GetUserInfoList(c *gin.Context) {
	var pageInfo userReq.UserInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userDataService.GetUserInfoInfoList(pageInfo); err != nil {
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
