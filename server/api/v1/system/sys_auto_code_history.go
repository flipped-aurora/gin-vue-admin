package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeHistoryApi struct{}

// First
// @Tags AutoCode
// @Summary 获取meta信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "请求参数"
// @Success 200 {object} response.Response{} "获取成功!"
// @Router /autoCode/getMeta [post]
func (a *AutoCodeHistoryApi) First(c *gin.Context) {
	var info request.GetById
	_ = c.ShouldBindJSON(&info)
	data, err := autoCodeHistoryService.First(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"meta": data}, "获取成功", c)
}

// Delete
// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /autoCode/delSysHistory [post]
func (a *AutoCodeHistoryApi) Delete(c *gin.Context) {
	var info request.GetById
	_ = c.ShouldBindJSON(&info)
	err := autoCodeHistoryService.Delete(&info)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// RollBack
// @Tags AutoCode
// @Summary 回滚自动生成代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "请求参数"
// @Success 200 {object} response.Response{} "回滚成功!"
// @Router /autoCode/rollback [post]
func (a *AutoCodeHistoryApi) RollBack(c *gin.Context) {
	var info request.GetById
	_ = c.ShouldBindJSON(&info)
	if err := autoCodeHistoryService.RollBack(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回滚成功", c)
}

// GetList
// @Tags AutoCode
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SysAutoHistory true "请求参数"
// @Success 200 {object} response.Response{} "获取成功!"
// @Router /autoCode/getSysHistory [post]
func (a *AutoCodeHistoryApi) GetList(c *gin.Context) {
	var search systemReq.SysAutoHistory
	_ = c.ShouldBindJSON(&search)
	list, total, err := autoCodeHistoryService.GetList(search.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "获取成功", c)
}
