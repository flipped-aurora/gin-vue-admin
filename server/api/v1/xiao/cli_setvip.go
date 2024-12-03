package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliSetvipApi struct{}

// CreateCliSetvip 创建团队设置
// @Tags CliSetvip
// @Summary 创建团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSetvip true "创建团队设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /clisetvip/createCliSetvip [post]
func (clisetvipApi *CliSetvipApi) CreateCliSetvip(c *gin.Context) {
	var clisetvip xiao.CliSetvip
	err := c.ShouldBindJSON(&clisetvip)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = clisetvipService.CreateCliSetvip(&clisetvip)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliSetvip 删除团队设置
// @Tags CliSetvip
// @Summary 删除团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSetvip true "删除团队设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /clisetvip/deleteCliSetvip [delete]
func (clisetvipApi *CliSetvipApi) DeleteCliSetvip(c *gin.Context) {
	ID := c.Query("ID")
	err := clisetvipService.DeleteCliSetvip(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliSetvipByIds 批量删除团队设置
// @Tags CliSetvip
// @Summary 批量删除团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /clisetvip/deleteCliSetvipByIds [delete]
func (clisetvipApi *CliSetvipApi) DeleteCliSetvipByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := clisetvipService.DeleteCliSetvipByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliSetvip 更新团队设置
// @Tags CliSetvip
// @Summary 更新团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliSetvip true "更新团队设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /clisetvip/updateCliSetvip [put]
func (clisetvipApi *CliSetvipApi) UpdateCliSetvip(c *gin.Context) {
	var clisetvip xiao.CliSetvip
	err := c.ShouldBindJSON(&clisetvip)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = clisetvipService.UpdateCliSetvip(clisetvip)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliSetvip 用id查询团队设置
// @Tags CliSetvip
// @Summary 用id查询团队设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliSetvip true "用id查询团队设置"
// @Success 200 {object} response.Response{data=xiao.CliSetvip,msg=string} "查询成功"
// @Router /clisetvip/findCliSetvip [get]
func (clisetvipApi *CliSetvipApi) FindCliSetvip(c *gin.Context) {
	ID := c.Query("ID")
	reclisetvip, err := clisetvipService.GetCliSetvip(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclisetvip, c)
}

// GetCliSetvipList 分页获取团队设置列表
// @Tags CliSetvip
// @Summary 分页获取团队设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetvipSearch true "分页获取团队设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /clisetvip/getCliSetvipList [get]
func (clisetvipApi *CliSetvipApi) GetCliSetvipList(c *gin.Context) {
	var pageInfo xiaoReq.CliSetvipSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := clisetvipService.GetCliSetvipInfoList(pageInfo)
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

// GetCliSetvipPublic 不需要鉴权的团队设置接口
// @Tags CliSetvip
// @Summary 不需要鉴权的团队设置接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliSetvipSearch true "分页获取团队设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /clisetvip/getCliSetvipPublic [get]
func (clisetvipApi *CliSetvipApi) GetCliSetvipPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	err := clisetvipService.GetCliSetvipPublic()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		return
	}
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的团队设置接口信息",
	}, "获取成功", c)
}
