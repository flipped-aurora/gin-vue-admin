package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func OpenConnection(c *gin.Context) {
	var createConnection request.CreateConnection
	_ = c.ShouldBind(&createConnection)
	_ = c.ShouldBindUri(&createConnection)
	if err, token := service.OpenConnection(createConnection); err != nil {
		global.GVA_LOG.Error("创建连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "创建连接失败", c)
	} else {
		response.OkWithDetailed(token, "创建连接成功", c)
	}
}

func CloseConnection(c *gin.Context) {
	var closeConnection request.CloseConnection
	_ = c.ShouldBindJSON(&closeConnection)
	_ = c.ShouldBindUri(&closeConnection)
	if err := service.CloseConnection(closeConnection); err != nil {
		global.GVA_LOG.Error("关闭连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "关闭连接失败", c)
	} else {
		response.OkWithMessage("关闭连接成功", c)
	}
}

func GameRequest(c *gin.Context) {
	var gameRequest request.GameRequest
	_ = c.ShouldBindJSON(&gameRequest)
	if err, data := service.GameRequest(gameRequest); err != nil {
		global.GVA_LOG.Error("发送请求失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "发送请求失败", c)
	} else {
		response.OkWithDetailed(data, "发送请求成功", c)
	}
}
