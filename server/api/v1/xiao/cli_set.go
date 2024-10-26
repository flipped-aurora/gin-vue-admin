package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliSetApi struct{}

// CreateCliSet 创建结算设置
// @Tags CliSet
// @Summary 创建结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSet true "创建结算设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliset/createCliSet [post]
func (clisetApi *CliSetApi) CreateCliSet(c *gin.Context) {
	var cliset xiao.CliSet
	err := c.ShouldBindJSON(&cliset)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = clisetService.CreateCliSet(&cliset)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliSet 删除结算设置
// @Tags CliSet
// @Summary 删除结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSet true "删除结算设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliset/deleteCliSet [delete]
func (clisetApi *CliSetApi) DeleteCliSet(c *gin.Context) {
	ID := c.Query("ID")
	err := clisetService.DeleteCliSet(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliSetByIds 批量删除结算设置
// @Tags CliSet
// @Summary 批量删除结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliset/deleteCliSetByIds [delete]
func (clisetApi *CliSetApi) DeleteCliSetByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := clisetService.DeleteCliSetByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliSet 更新结算设置
// @Tags CliSet
// @Summary 更新结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSet true "更新结算设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliset/updateCliSet [put]
func (clisetApi *CliSetApi) UpdateCliSet(c *gin.Context) {
	var cliset xiao.CliSet
	err := c.ShouldBindJSON(&cliset)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = clisetService.UpdateCliSet(cliset)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliSet 用id查询结算设置
// @Tags CliSet
// @Summary 用id查询结算设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliSet true "用id查询结算设置"
// @Success 200 {object} response.Response{data=xiao.CliSet,msg=string} "查询成功"
// @Router /cliset/findCliSet [get]
func (clisetApi *CliSetApi) FindCliSet(c *gin.Context) {
	ID := c.Query("ID")
	recliset, err := clisetService.GetCliSet(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliset, c)
}

// GetCliSetList 分页获取结算设置列表
// @Tags CliSet
// @Summary 分页获取结算设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetSearch true "分页获取结算设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliset/getCliSetList [get]
func (clisetApi *CliSetApi) GetCliSetList(c *gin.Context) {
	var pageInfo xiaoReq.CliSetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := clisetService.GetCliSetInfoList(pageInfo)
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

// GetCliSetPublic 不需要鉴权的结算设置接口
// @Tags CliSet
// @Summary 不需要鉴权的结算设置接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetSearch true "分页获取结算设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliset/getCliSetPublic [get]
func (clisetApi *CliSetApi) GetCliSetPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	clisetService.GetCliSetPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的结算设置接口信息",
	}, "获取成功", c)
}
