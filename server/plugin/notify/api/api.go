package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
}

func (s *Api) NotifyController(c *gin.Context) {
	if err := service.ServiceGroupApp.Send(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
