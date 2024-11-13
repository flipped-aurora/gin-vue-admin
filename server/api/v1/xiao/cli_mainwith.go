package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliMainwithApi struct{}

// CreateCliMainwith 创建提币总表
// @Tags CliMainwith
// @Summary 创建提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainwith true "创建提币总表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /climainwith/createCliMainwith [post]
func (climainwithApi *CliMainwithApi) CreateCliMainwith(c *gin.Context) {
	var climainwith xiao.CliMainwith
	err := c.ShouldBindJSON(&climainwith)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainwithService.CreateCliMainwith(&climainwith)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliMainwith 删除提币总表
// @Tags CliMainwith
// @Summary 删除提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainwith true "删除提币总表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /climainwith/deleteCliMainwith [delete]
func (climainwithApi *CliMainwithApi) DeleteCliMainwith(c *gin.Context) {
	ID := c.Query("ID")
	err := climainwithService.DeleteCliMainwith(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliMainwithByIds 批量删除提币总表
// @Tags CliMainwith
// @Summary 批量删除提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /climainwith/deleteCliMainwithByIds [delete]
func (climainwithApi *CliMainwithApi) DeleteCliMainwithByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := climainwithService.DeleteCliMainwithByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliMainwith 更新提币总表
// @Tags CliMainwith
// @Summary 更新提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliMainwith true "更新提币总表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /climainwith/updateCliMainwith [put]
func (climainwithApi *CliMainwithApi) UpdateCliMainwith(c *gin.Context) {
	var climainwith xiao.CliMainwith
	err := c.ShouldBindJSON(&climainwith)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = climainwithService.UpdateCliMainwith(climainwith)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliMainwith 用id查询提币总表
// @Tags CliMainwith
// @Summary 用id查询提币总表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliMainwith true "用id查询提币总表"
// @Success 200 {object} response.Response{data=xiao.CliMainwith,msg=string} "查询成功"
// @Router /climainwith/findCliMainwith [get]
func (climainwithApi *CliMainwithApi) FindCliMainwith(c *gin.Context) {
	ID := c.Query("ID")
	reclimainwith, err := climainwithService.GetCliMainwith(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclimainwith, c)
}

// GetCliMainwithList 分页获取提币总表列表
// @Tags CliMainwith
// @Summary 分页获取提币总表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainwithSearch true "分页获取提币总表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /climainwith/getCliMainwithList [get]
func (climainwithApi *CliMainwithApi) GetCliMainwithList(c *gin.Context) {
	var pageInfo xiaoReq.CliMainwithSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := climainwithService.GetCliMainwithInfoList(pageInfo)
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

// GetCliMainwithPublic 不需要鉴权的提币总表接口
// @Tags CliMainwith
// @Summary 不需要鉴权的提币总表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliMainwithSearch true "分页获取提币总表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /climainwith/getCliMainwithPublic [get]
func (climainwithApi *CliMainwithApi) GetCliMainwithPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	err := climainwithService.GetCliMainwithPublic()
	if err != nil {

		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("静态结算成功", c)
}
