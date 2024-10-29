package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliMainprofitApi struct{}

// CreateCliMainprofit 创建结算总表
// @Tags CliMainprofit
// @Summary 创建结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainprofit true "创建结算总表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /climainprofit/createCliMainprofit [post]
func (climainprofitApi *CliMainprofitApi) CreateCliMainprofit(c *gin.Context) {
	var climainprofit xiao.CliMainprofit
	err := c.ShouldBindJSON(&climainprofit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainprofitService.CreateCliMainprofit(&climainprofit)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliMainprofit 删除结算总表
// @Tags CliMainprofit
// @Summary 删除结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainprofit true "删除结算总表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /climainprofit/deleteCliMainprofit [delete]
func (climainprofitApi *CliMainprofitApi) DeleteCliMainprofit(c *gin.Context) {
	ID := c.Query("ID")
	err := climainprofitService.DeleteCliMainprofit(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliMainprofitByIds 批量删除结算总表
// @Tags CliMainprofit
// @Summary 批量删除结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /climainprofit/deleteCliMainprofitByIds [delete]
func (climainprofitApi *CliMainprofitApi) DeleteCliMainprofitByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := climainprofitService.DeleteCliMainprofitByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliMainprofit 更新结算总表
// @Tags CliMainprofit
// @Summary 更新结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainprofit true "更新结算总表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /climainprofit/updateCliMainprofit [put]
func (climainprofitApi *CliMainprofitApi) UpdateCliMainprofit(c *gin.Context) {
	var climainprofit xiao.CliMainprofit
	err := c.ShouldBindJSON(&climainprofit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainprofitService.UpdateCliMainprofit(climainprofit)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliMainprofit 用id查询结算总表
// @Tags CliMainprofit
// @Summary 用id查询结算总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliMainprofit true "用id查询结算总表"
// @Success 200 {object} response.Response{data=xiao.CliMainprofit,msg=string} "查询成功"
// @Router /climainprofit/findCliMainprofit [get]
func (climainprofitApi *CliMainprofitApi) FindCliMainprofit(c *gin.Context) {
	ID := c.Query("ID")
	reclimainprofit, err := climainprofitService.GetCliMainprofit(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclimainprofit, c)
}

// GetCliMainprofitList 分页获取结算总表列表
// @Tags CliMainprofit
// @Summary 分页获取结算总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainprofitSearch true "分页获取结算总表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /climainprofit/getCliMainprofitList [get]
func (climainprofitApi *CliMainprofitApi) GetCliMainprofitList(c *gin.Context) {
	var pageInfo xiaoReq.CliMainprofitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := climainprofitService.GetCliMainprofitInfoList(pageInfo)
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

// GetCliMainprofitPublic 不需要鉴权的结算总表接口
// @Tags CliMainprofit
// @Summary 不需要鉴权的结算总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainprofitSearch true  "分页获取结算总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainprofit/getCliMainprofitPublic [get]
func (climainprofitApi *CliMainprofitApi) GetCliMainprofitPublic(c *gin.Context) {
	// 此接口不需要鉴权
	var pageInfo xiaoReq.CliMainprofitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := climainprofitService.GetCliMainprofitInfoList(pageInfo)
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
