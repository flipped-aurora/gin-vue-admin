package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// @Tags InitDB
// @Summary 初始化用户数据库
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /init/db [post]
func InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("非法访问")
		response.FailWithMessage("非法访问", c)
		return
	}
	var dbInfo model.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("自动创建数据库失败", zap.Any("err", err))
		response.FailWithMessage("自动创建数据库失败", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}
