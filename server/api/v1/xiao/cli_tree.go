package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CliTreeApi struct{}

// CreateCliTree 创建用户关系表
// @Tags CliTree
// @Summary 创建用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliTree true "创建用户关系表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliTree/createCliTree [post]
func (cliTreeApi *CliTreeApi) CreateCliTree(c *gin.Context) {
	var cliTree xiao.CliTree
	err := c.ShouldBindJSON(&cliTree)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliTreeService.CreateCliTree(&cliTree)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliTree 删除用户关系表
// @Tags CliTree
// @Summary 删除用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliTree true "删除用户关系表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliTree/deleteCliTree [delete]
func (cliTreeApi *CliTreeApi) DeleteCliTree(c *gin.Context) {
	ID := c.Query("ID")
	err := cliTreeService.DeleteCliTree(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliTreeByIds 批量删除用户关系表
// @Tags CliTree
// @Summary 批量删除用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliTree/deleteCliTreeByIds [delete]
func (cliTreeApi *CliTreeApi) DeleteCliTreeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliTreeService.DeleteCliTreeByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliTree 更新用户关系表
// @Tags CliTree
// @Summary 更新用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliTree true "更新用户关系表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliTree/updateCliTree [put]
func (cliTreeApi *CliTreeApi) UpdateCliTree(c *gin.Context) {
	var cliTree xiao.CliTree
	err := c.ShouldBindJSON(&cliTree)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliTreeService.UpdateCliTree(cliTree)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliTree 用id查询用户关系表
// @Tags CliTree
// @Summary 用id查询用户关系表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliTree true "用id查询用户关系表"
// @Success 200 {object} response.Response{data=xiao.CliTree,msg=string} "查询成功"
// @Router /cliTree/findCliTree [get]
func (cliTreeApi *CliTreeApi) FindCliTree(c *gin.Context) {
	ID := c.Query("ID")
	recliTree, err := cliTreeService.GetCliTree(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliTree, c)
}

// GetCliTreeList 分页获取用户关系表列表
// @Tags CliTree
// @Summary 分页获取用户关系表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliTreeSearch true "分页获取用户关系表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliTree/getCliTreeList [get]
func (cliTreeApi *CliTreeApi) GetCliTreeList(c *gin.Context) {
	var pageInfo xiaoReq.CliTreeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliTreeService.GetCliTreeInfoList(pageInfo)
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

// GetCliTreePublic 不需要鉴权的用户关系表接口
// @Tags CliTree
// @Summary 不需要鉴权的用户关系表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliTreeSearch true "分页获取用户关系表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliTree/getCliTreePublic [get]
func (cliTreeApi *CliTreeApi) GetCliTreePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cliTreeService.GetCliTreePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户关系表接口信息",
	}, "获取成功", c)
}
