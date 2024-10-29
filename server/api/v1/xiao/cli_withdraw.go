package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliWithdrawApi struct{}

// CreateCliWithdraw 创建提币详情
// @Tags CliWithdraw
// @Summary 创建提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliWithdraw true "创建提币详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliwithdraw/createCliWithdraw [post]
func (cliwithdrawApi *CliWithdrawApi) CreateCliWithdraw(c *gin.Context) {
	var cliwithdraw xiao.CliWithdraw
	err := c.ShouldBindJSON(&cliwithdraw)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliwithdrawService.CreateCliWithdraw(&cliwithdraw)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliWithdraw 删除提币详情
// @Tags CliWithdraw
// @Summary 删除提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliWithdraw true "删除提币详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliwithdraw/deleteCliWithdraw [delete]
func (cliwithdrawApi *CliWithdrawApi) DeleteCliWithdraw(c *gin.Context) {
	ID := c.Query("ID")
	err := cliwithdrawService.DeleteCliWithdraw(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliWithdrawByIds 批量删除提币详情
// @Tags CliWithdraw
// @Summary 批量删除提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliwithdraw/deleteCliWithdrawByIds [delete]
func (cliwithdrawApi *CliWithdrawApi) DeleteCliWithdrawByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliwithdrawService.DeleteCliWithdrawByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliWithdraw 更新提币详情
// @Tags CliWithdraw
// @Summary 更新提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliWithdraw true "更新提币详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliwithdraw/updateCliWithdraw [put]
func (cliwithdrawApi *CliWithdrawApi) UpdateCliWithdraw(c *gin.Context) {
	var cliwithdraw xiao.CliWithdraw
	err := c.ShouldBindJSON(&cliwithdraw)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliwithdrawService.UpdateCliWithdraw(cliwithdraw)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliWithdraw 用id查询提币详情
// @Tags CliWithdraw
// @Summary 用id查询提币详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliWithdraw true "用id查询提币详情"
// @Success 200 {object} response.Response{data=xiao.CliWithdraw,msg=string} "查询成功"
// @Router /cliwithdraw/findCliWithdraw [get]
func (cliwithdrawApi *CliWithdrawApi) FindCliWithdraw(c *gin.Context) {
	ID := c.Query("ID")
	recliwithdraw, err := cliwithdrawService.GetCliWithdraw(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliwithdraw, c)
}

// GetCliWithdrawList 分页获取提币详情列表
// @Tags CliWithdraw
// @Summary 分页获取提币详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliWithdrawSearch true "分页获取提币详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliwithdraw/getCliWithdrawList [get]
func (cliwithdrawApi *CliWithdrawApi) GetCliWithdrawList(c *gin.Context) {
	var pageInfo xiaoReq.CliWithdrawSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliwithdrawService.GetCliWithdrawInfoList(pageInfo)
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

// GetCliWithdrawPublic 不需要鉴权的提币详情接口
// @Tags CliWithdraw
// @Summary 不需要鉴权的提币详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliWithdrawSearch true "分页获取提币详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliwithdraw/getCliWithdrawPublic [get]
func (cliwithdrawApi *CliWithdrawApi) GetCliWithdrawPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	var pageInfo xiaoReq.CliWithdrawSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	public, err := cliwithdrawService.GetCliWithdrawPublic(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(public, "获取成功", c)
}
