package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags HrpQueue
// @Summary 创建HrpQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpQueue true "创建HrpQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrpQueue/createHrpQueue [post]
func CreateHrpQueue(c *gin.Context) {
	var hrpQueue model.HrpQueue
	_ = c.ShouldBindJSON(&hrpQueue)
	err := service.CreateHrpQueue(hrpQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HrpQueue
// @Summary 删除HrpQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpQueue true "删除HrpQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrpQueue/deleteHrpQueue [delete]
func DeleteHrpQueue(c *gin.Context) {
	var hrpQueue model.HrpQueue
	_ = c.ShouldBindJSON(&hrpQueue)
	err := service.DeleteHrpQueue(hrpQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HrpQueue
// @Summary 批量删除HrpQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrpQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrpQueue/deleteHrpQueueByIds [delete]
func DeleteHrpQueueByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHrpQueueByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HrpQueue
// @Summary 更新HrpQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpQueue true "更新HrpQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrpQueue/updateHrpQueue [put]
func UpdateHrpQueue(c *gin.Context) {
	var hrpQueue model.HrpQueue
	_ = c.ShouldBindJSON(&hrpQueue)
	err := service.UpdateHrpQueue(&hrpQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HrpQueue
// @Summary 用id查询HrpQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpQueue true "用id查询HrpQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrpQueue/findHrpQueue [get]
func FindHrpQueue(c *gin.Context) {
	var hrpQueue model.HrpQueue
	_ = c.ShouldBindQuery(&hrpQueue)
	err, rehrpQueue := service.GetHrpQueue(hrpQueue.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehrpQueue": rehrpQueue}, c)
	}
}

// @Tags HrpQueue
// @Summary 分页获取HrpQueue列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HrpQueueSearch true "分页获取HrpQueue列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrpQueue/getHrpQueueList [get]
func GetHrpQueueList(c *gin.Context) {
	var pageInfo request.HrpQueueSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHrpQueueInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
