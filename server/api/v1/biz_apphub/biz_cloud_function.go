package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
    biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BizCloudFunctionApi struct {}

// CreateBizCloudFunction 创建云函数
// @Tags BizCloudFunction
// @Summary 创建云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "创建云函数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bizCloudFunction/createBizCloudFunction [post]
func (bizCloudFunctionApi *BizCloudFunctionApi) CreateBizCloudFunction(c *gin.Context) {
	var bizCloudFunction biz_apphub.BizCloudFunction
	err := c.ShouldBindJSON(&bizCloudFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bizCloudFunctionService.CreateBizCloudFunction(&bizCloudFunction)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBizCloudFunction 删除云函数
// @Tags BizCloudFunction
// @Summary 删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "删除云函数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bizCloudFunction/deleteBizCloudFunction [delete]
func (bizCloudFunctionApi *BizCloudFunctionApi) DeleteBizCloudFunction(c *gin.Context) {
	ID := c.Query("ID")
	err := bizCloudFunctionService.DeleteBizCloudFunction(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizCloudFunctionByIds 批量删除云函数
// @Tags BizCloudFunction
// @Summary 批量删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bizCloudFunction/deleteBizCloudFunctionByIds [delete]
func (bizCloudFunctionApi *BizCloudFunctionApi) DeleteBizCloudFunctionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := bizCloudFunctionService.DeleteBizCloudFunctionByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizCloudFunction 更新云函数
// @Tags BizCloudFunction
// @Summary 更新云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "更新云函数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizCloudFunction/updateBizCloudFunction [put]
func (bizCloudFunctionApi *BizCloudFunctionApi) UpdateBizCloudFunction(c *gin.Context) {
	var bizCloudFunction biz_apphub.BizCloudFunction
	err := c.ShouldBindJSON(&bizCloudFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bizCloudFunctionService.UpdateBizCloudFunction(bizCloudFunction)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizCloudFunction 用id查询云函数
// @Tags BizCloudFunction
// @Summary 用id查询云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphub.BizCloudFunction true "用id查询云函数"
// @Success 200 {object} response.Response{data=biz_apphub.BizCloudFunction,msg=string} "查询成功"
// @Router /bizCloudFunction/findBizCloudFunction [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) FindBizCloudFunction(c *gin.Context) {
	ID := c.Query("ID")
	rebizCloudFunction, err := bizCloudFunctionService.GetBizCloudFunction(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rebizCloudFunction, c)
}

// GetBizCloudFunctionList 分页获取云函数列表
// @Tags BizCloudFunction
// @Summary 分页获取云函数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCloudFunctionSearch true "分页获取云函数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizCloudFunction/getBizCloudFunctionList [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) GetBizCloudFunctionList(c *gin.Context) {
	var pageInfo biz_apphubReq.BizCloudFunctionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bizCloudFunctionService.GetBizCloudFunctionInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetBizCloudFunctionPublic 不需要鉴权的云函数接口
// @Tags BizCloudFunction
// @Summary 不需要鉴权的云函数接口
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCloudFunctionSearch true "分页获取云函数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bizCloudFunction/getBizCloudFunctionPublic [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) GetBizCloudFunctionPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的云函数接口信息",
    }, "获取成功", c)
}
