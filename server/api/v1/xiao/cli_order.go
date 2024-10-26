package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliOrderApi struct{}

// CreateCliOrder 创建订单详情
// @Tags CliOrder
// @Summary 创建订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliOrder true "创建订单详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliOrder/createCliOrder [post]
func (cliOrderApi *CliOrderApi) CreateCliOrder(c *gin.Context) {
	var cliOrder xiao.CliOrder
	err := c.ShouldBindJSON(&cliOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliOrderService.CreateCliOrder(&cliOrder)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliOrder 删除订单详情
// @Tags CliOrder
// @Summary 删除订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliOrder true "删除订单详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliOrder/deleteCliOrder [delete]
func (cliOrderApi *CliOrderApi) DeleteCliOrder(c *gin.Context) {
	ID := c.Query("ID")
	err := cliOrderService.DeleteCliOrder(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliOrderByIds 批量删除订单详情
// @Tags CliOrder
// @Summary 批量删除订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliOrder/deleteCliOrderByIds [delete]
func (cliOrderApi *CliOrderApi) DeleteCliOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliOrderService.DeleteCliOrderByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliOrder 更新订单详情
// @Tags CliOrder
// @Summary 更新订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliOrder true "更新订单详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliOrder/updateCliOrder [put]
func (cliOrderApi *CliOrderApi) UpdateCliOrder(c *gin.Context) {
	var cliOrder xiao.CliOrder
	err := c.ShouldBindJSON(&cliOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliOrderService.UpdateCliOrder(cliOrder)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliOrder 用id查询订单详情
// @Tags CliOrder
// @Summary 用id查询订单详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliOrder true "用id查询订单详情"
// @Success 200 {object} response.Response{data=xiao.CliOrder,msg=string} "查询成功"
// @Router /cliOrder/findCliOrder [get]
func (cliOrderApi *CliOrderApi) FindCliOrder(c *gin.Context) {
	ID := c.Query("ID")
	recliOrder, err := cliOrderService.GetCliOrder(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliOrder, c)
}

// GetCliOrderList 分页获取订单详情列表
// @Tags CliOrder
// @Summary 分页获取订单详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliOrderSearch true "分页获取订单详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliOrder/getCliOrderList [get]
func (cliOrderApi *CliOrderApi) GetCliOrderList(c *gin.Context) {
	var pageInfo xiaoReq.CliOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliOrderService.GetCliOrderInfoList(pageInfo)
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

// GetCliOrderPublic 不需要鉴权的订单详情接口
// @Tags CliOrder
// @Summary 不需要鉴权的订单详情接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliOrderSearch true "分页获取订单详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliOrder/getCliOrderPublic [get]
func (cliOrderApi *CliOrderApi) GetCliOrderPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cliOrderService.GetCliOrderPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订单详情接口信息",
	}, "获取成功", c)
}
