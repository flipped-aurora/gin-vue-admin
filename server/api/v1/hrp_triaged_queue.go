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

// @Tags HrpTriagedQueue
// @Summary 创建HrpTriagedQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpTriagedQueue true "创建HrpTriagedQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrpTriagedQueue/createHrpTriagedQueue [post]
func CreateHrpTriagedQueue(c *gin.Context) {
	var hrpTriagedQueue model.HrpTriagedQueue
	_ = c.ShouldBindJSON(&hrpTriagedQueue)
	err := service.CreateHrpTriagedQueue(hrpTriagedQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HrpTriagedQueue
// @Summary 删除HrpTriagedQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpTriagedQueue true "删除HrpTriagedQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrpTriagedQueue/deleteHrpTriagedQueue [delete]
func DeleteHrpTriagedQueue(c *gin.Context) {
	var hrpTriagedQueue model.HrpTriagedQueue
	_ = c.ShouldBindJSON(&hrpTriagedQueue)
	err := service.DeleteHrpTriagedQueue(hrpTriagedQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HrpTriagedQueue
// @Summary 批量删除HrpTriagedQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrpTriagedQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrpTriagedQueue/deleteHrpTriagedQueueByIds [delete]
func DeleteHrpTriagedQueueByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHrpTriagedQueueByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HrpTriagedQueue
// @Summary 更新HrpTriagedQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpTriagedQueue true "更新HrpTriagedQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrpTriagedQueue/updateHrpTriagedQueue [put]
func UpdateHrpTriagedQueue(c *gin.Context) {
	var hrpTriagedQueue model.HrpTriagedQueue
	_ = c.ShouldBindJSON(&hrpTriagedQueue)
	err := service.UpdateHrpTriagedQueue(&hrpTriagedQueue)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HrpTriagedQueue
// @Summary 用id查询HrpTriagedQueue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrpTriagedQueue true "用id查询HrpTriagedQueue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrpTriagedQueue/findHrpTriagedQueue [get]
func FindHrpTriagedQueue(c *gin.Context) {
	var hrpTriagedQueue model.HrpTriagedQueue
	_ = c.ShouldBindQuery(&hrpTriagedQueue)
	err, rehrpTriagedQueue := service.GetHrpTriagedQueue(hrpTriagedQueue.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehrpTriagedQueue": rehrpTriagedQueue}, c)
	}
}

// @Tags HrpTriagedQueue
// @Summary 分页获取HrpTriagedQueue列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HrpTriagedQueueSearch true "分页获取HrpTriagedQueue列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrpTriagedQueue/getHrpTriagedQueueList [get]
func GetHrpTriagedQueueList(c *gin.Context) {
	var pageInfo request.HrpTriagedQueueSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHrpTriagedQueueInfoList(pageInfo)
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
