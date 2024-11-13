package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliProfitApi struct{}

// CreateCliProfit 创建结算详情
// @Tags CliProfit
// @Summary 创建结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliProfit true "创建结算详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliprofit/createCliProfit [post]
func (cliprofitApi *CliProfitApi) CreateCliProfit(c *gin.Context) {
	var cliprofit xiao.CliProfit
	err := c.ShouldBindJSON(&cliprofit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliprofitService.CreateCliProfit(&cliprofit)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliProfit 删除结算详情
// @Tags CliProfit
// @Summary 删除结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliProfit true "删除结算详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliprofit/deleteCliProfit [delete]
func (cliprofitApi *CliProfitApi) DeleteCliProfit(c *gin.Context) {
	ID := c.Query("ID")
	err := cliprofitService.DeleteCliProfit(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliProfitByIds 批量删除结算详情
// @Tags CliProfit
// @Summary 批量删除结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliprofit/deleteCliProfitByIds [delete]
func (cliprofitApi *CliProfitApi) DeleteCliProfitByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliprofitService.DeleteCliProfitByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliProfit 更新结算详情
// @Tags CliProfit
// @Summary 更新结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliProfit true "更新结算详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliprofit/updateCliProfit [put]
func (cliprofitApi *CliProfitApi) UpdateCliProfit(c *gin.Context) {
	var cliprofit xiao.CliProfit
	err := c.ShouldBindJSON(&cliprofit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliprofitService.UpdateCliProfit(cliprofit)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliProfit 用id查询结算详情
// @Tags CliProfit
// @Summary 用id查询结算详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliProfit true "用id查询结算详情"
// @Success 200 {object} response.Response{data=xiao.CliProfit,msg=string} "查询成功"
// @Router /cliprofit/findCliProfit [get]
func (cliprofitApi *CliProfitApi) FindCliProfit(c *gin.Context) {
	ID := c.Query("ID")
	recliprofit, err := cliprofitService.GetCliProfit(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliprofit, c)
}

// GetCliProfitList 分页获取结算详情列表
// @Tags CliProfit
// @Summary 分页获取结算详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliProfitSearch true "分页获取结算详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliprofit/getCliProfitList [get]
func (cliprofitApi *CliProfitApi) GetCliProfitList(c *gin.Context) {
	var pageInfo xiaoReq.CliProfitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliprofitService.GetCliProfitInfoList(pageInfo)
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

// GetCliProfitPublic 不需要鉴权的结算详情接口
// @Tags CliProfit
// @Summary 不需要鉴权的结算详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliProfitSearch true "分页获取结算详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliprofit/getCliProfitPublic [get]
func (cliprofitApi *CliProfitApi) GetCliProfitPublic(c *gin.Context) {
	// 此接口不需要鉴权
	var pageInfo xiaoReq.CliProfitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliprofitService.GetCliProfitInfoList(pageInfo)
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
func (cliprofitApi *CliProfitApi) WithAsset(c *gin.Context) {
	var cliprofit xiao.CliProfit
	err := c.ShouldBindJSON(&cliprofit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliprofitService.WithAsset(&cliprofit)
	if err != nil {
		global.GVA_LOG.Error("提取本金失败!", zap.Error(err))
		response.FailWithMessage("提取本金失败:"+err.Error(), c)
	}
	response.OkWithMessage("提取本金成功", c)

}
