package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body model.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func CreateWorkFlow(c *gin.Context) {
	var wk model.SysWorkflow
	_ = c.ShouldBindJSON(&wk)
	if err := utils.Verify(wk, utils.WorkFlowVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.Create(wk); err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithMessage("注册失败", c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}
