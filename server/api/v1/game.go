package v1

import (
	"github.com/eyotang/game-api-admin/server/global"
	"github.com/eyotang/game-api-admin/server/model/request"
	"github.com/eyotang/game-api-admin/server/model/response"
	"github.com/eyotang/game-api-admin/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func OpenConnection(c *gin.Context) {
	var (
		param request.ParamGame
		body  request.CreateConnection
	)
	if err := c.ShouldBind(&body); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err, token := service.OpenConnection(&param, &body); err != nil {
		global.GVA_LOG.Error("创建连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "创建连接失败", c)
	} else {
		response.OkWithDetailed(token, "创建连接成功", c)
	}
}

func CloseConnection(c *gin.Context) {
	var (
		param  request.ParamGame
		header request.HeaderRequest
	)
	if err := c.ShouldBindHeader(&header); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := service.CloseConnection(&param, &header); err != nil {
		global.GVA_LOG.Error("关闭连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "关闭连接失败", c)
	} else {
		response.OkWithMessage("关闭连接成功", c)
	}
}

func GameRequest(c *gin.Context) {
	var (
		param  request.ParamRequest
		body   request.GameRequest
		header request.HeaderRequest
	)
	if err := c.ShouldBindJSON(&body); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := c.ShouldBindHeader(&header); err != nil {
		global.GVA_LOG.Error("头部校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "头部校验错误", c)
		return
	}
	if err, data := service.GameRequest(&header, &param, &body); err != nil {
		global.GVA_LOG.Error("发送请求失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "发送请求失败", c)
	} else {
		response.OkWithDetailed(data, "发送请求成功", c)
	}
}
