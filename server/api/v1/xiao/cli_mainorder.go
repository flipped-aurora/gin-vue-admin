package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliMainorderApi struct{}

// CreateCliMainorder 创建订单总表
// @Tags CliMainorder
// @Summary 创建订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainorder true "创建订单总表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /climainorder/createCliMainorder [post]
func (climainorderApi *CliMainorderApi) CreateCliMainorder(c *gin.Context) {
	var climainorder xiao.CliMainorder
	err := c.ShouldBindJSON(&climainorder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainorderService.CreateCliMainorder(&climainorder)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliMainorder 删除订单总表
// @Tags CliMainorder
// @Summary 删除订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainorder true "删除订单总表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /climainorder/deleteCliMainorder [delete]
func (climainorderApi *CliMainorderApi) DeleteCliMainorder(c *gin.Context) {
	ID := c.Query("ID")
	err := climainorderService.DeleteCliMainorder(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliMainorderByIds 批量删除订单总表
// @Tags CliMainorder
// @Summary 批量删除订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /climainorder/deleteCliMainorderByIds [delete]
func (climainorderApi *CliMainorderApi) DeleteCliMainorderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := climainorderService.DeleteCliMainorderByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliMainorder 更新订单总表
// @Tags CliMainorder
// @Summary 更新订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainorder true "更新订单总表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /climainorder/updateCliMainorder [put]
func (climainorderApi *CliMainorderApi) UpdateCliMainorder(c *gin.Context) {
	var climainorder xiao.CliMainorder
	err := c.ShouldBindJSON(&climainorder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainorderService.UpdateCliMainorder(climainorder)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliMainorder 用id查询订单总表
// @Tags CliMainorder
// @Summary 用id查询订单总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliMainorder true "用id查询订单总表"
// @Success 200 {object} response.Response{data=xiao.CliMainorder,msg=string} "查询成功"
// @Router /climainorder/findCliMainorder [get]
func (climainorderApi *CliMainorderApi) FindCliMainorder(c *gin.Context) {
	ID := c.Query("ID")
	reclimainorder, err := climainorderService.GetCliMainorder(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclimainorder, c)
}

// GetCliMainorderList 分页获取订单总表列表
// @Tags CliMainorder
// @Summary 分页获取订单总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainorderSearch true "分页获取订单总表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /climainorder/getCliMainorderList [get]
func (climainorderApi *CliMainorderApi) GetCliMainorderList(c *gin.Context) {
	var pageInfo xiaoReq.CliMainorderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := climainorderService.GetCliMainorderInfoList(pageInfo)
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

// GetCliMainorderPublic 不需要鉴权的订单总表接口
// @Tags CliMainorder
// @Summary 不需要鉴权的订单总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainorderSearch true "分页获取订单总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainorder/getCliMainorderPublic [get]
func (climainorderApi *CliMainorderApi) GetCliMainorderPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	climainorderService.GetCliMainorderPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订单总表接口信息",
	}, "获取成功", c)
}

// Buy 客户端购买方法
// @Tags CliMainorder
// @Summary 客户端购买方法
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliOrder true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /climainorder/buy [POST]
func (climainorderApi *CliMainorderApi) Buy(c *gin.Context) {
	// 请添加自己的业务逻辑
	var cliorder xiao.CliOrder
	err := c.ShouldBindJSON(&cliorder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainorderService.Buy(&cliorder)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}
