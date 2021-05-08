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
	var (
		paramGame        request.ParamGame
		createConnection request.CreateConnection
	)
	if err := c.ShouldBind(&createConnection); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&paramGame); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err, token := service.OpenConnection(&paramGame, &createConnection); err != nil {
		global.GVA_LOG.Error("创建连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "创建连接失败", c)
	} else {
		response.OkWithDetailed(token, "创建连接成功", c)
	}
}

func CloseConnection(c *gin.Context) {
	var (
		paramGame       request.ParamGame
		closeConnection request.CloseConnection
	)
	if err := c.ShouldBindJSON(&closeConnection); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&paramGame); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := service.CloseConnection(&paramGame, &closeConnection); err != nil {
		global.GVA_LOG.Error("关闭连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "关闭连接失败", c)
	} else {
		response.OkWithMessage("关闭连接成功", c)
	}
}

func GameRequest(c *gin.Context) {
	var (
		paramRequest request.ParamRequest
		gameRequest  request.GameRequest
	)
	if err := c.ShouldBindJSON(&gameRequest); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&paramRequest); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err, data := service.GameRequest(&paramRequest, &gameRequest); err != nil {
		global.GVA_LOG.Error("发送请求失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "发送请求失败", c)
	} else {
		response.OkWithDetailed(data, "发送请求成功", c)
	}
}
