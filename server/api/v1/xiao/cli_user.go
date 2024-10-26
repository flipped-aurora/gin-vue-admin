package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliUserApi struct{}

// CreateCliUser 创建cliUser表
// @Tags CliUser
// @Summary 创建cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliUser true "创建cliUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliUser/createCliUser [post]
func (cliUserApi *CliUserApi) CreateCliUser(c *gin.Context) {
	var cliUser xiao.CliUser
	err := c.ShouldBindJSON(&cliUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliUserService.CreateCliUser(&cliUser)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliUser 删除cliUser表
// @Tags CliUser
// @Summary 删除cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliUser true "删除cliUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliUser/deleteCliUser [delete]
func (cliUserApi *CliUserApi) DeleteCliUser(c *gin.Context) {
	ID := c.Query("ID")
	err := cliUserService.DeleteCliUser(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliUserByIds 批量删除cliUser表
// @Tags CliUser
// @Summary 批量删除cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliUser/deleteCliUserByIds [delete]
func (cliUserApi *CliUserApi) DeleteCliUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliUserService.DeleteCliUserByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliUser 更新cliUser表
// @Tags CliUser
// @Summary 更新cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliUser true "更新cliUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliUser/updateCliUser [put]
func (cliUserApi *CliUserApi) UpdateCliUser(c *gin.Context) {
	var cliUser xiao.CliUser
	err := c.ShouldBindJSON(&cliUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliUserService.UpdateCliUser(cliUser)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliUser 用id查询cliUser表
// @Tags CliUser
// @Summary 用id查询cliUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliUser true "用id查询cliUser表"
// @Success 200 {object} response.Response{data=xiao.CliUser,msg=string} "查询成功"
// @Router /cliUser/findCliUser [get]
func (cliUserApi *CliUserApi) FindCliUser(c *gin.Context) {
	ID := c.Query("ID")
	recliUser, err := cliUserService.GetCliUser(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliUser, c)
}

// GetCliUserList 分页获取cliUser表列表
// @Tags CliUser
// @Summary 分页获取cliUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliUserSearch true "分页获取cliUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliUser/getCliUserList [get]
func (cliUserApi *CliUserApi) GetCliUserList(c *gin.Context) {
	var pageInfo xiaoReq.CliUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliUserService.GetCliUserInfoList(pageInfo)
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

// GetCliUserPublic 不需要鉴权的cliUser表接口
// @Tags CliUser
// @Summary 不需要鉴权的cliUser表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliUserSearch true "分页获取cliUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliUser/getCliUserPublic [get]
func (cliUserApi *CliUserApi) GetCliUserPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cliUserService.GetCliUserPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的cliUser表接口信息",
	}, "获取成功", c)
}

// Rigester 客户端注册方法
// @Tags CliUser
// @Summary 客户端注册方法
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliUser true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cliUser/rigester [POST]
func (cliUserApi *CliUserApi) Rigester(c *gin.Context) {
	// 请添加自己的业务逻辑
	var cliUser xiao.CliUser
	err := c.ShouldBindJSON(&cliUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliUserService.Rigester(&cliUser)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("注册成功", c)
}
