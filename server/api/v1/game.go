package v1

import (
	"github.com/eyotang/game-proxy/server/global"
	"github.com/eyotang/game-proxy/server/model/request"
	"github.com/eyotang/game-proxy/server/model/response"
	"github.com/eyotang/game-proxy/server/service"
	"github.com/eyotang/game-proxy/server/service/shared"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VServedPlugin(c *gin.Context) {
	var (
		param request.ParamGame
	)
	if err := c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}

	if err, _ := service.ServedPlugin(&param); err != nil {
		global.GVA_LOG.Error("服务插件错误!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "服务插件错误", c)
	}
}

func VOwnedToken(c *gin.Context) {
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

	if err := service.OwnedToken(&param, &header); err != nil {
		global.GVA_LOG.Error("所属token错误!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "所属token错误", c)
	}
}

func Destroy(c *gin.Context) {
	var (
		param request.ParamGame
	)
	if err := c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err, result := service.Destroy(&param); err != nil {
		global.GVA_LOG.Error("销毁失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "销毁失败", c)
	} else {
		response.OkWithDetailed(result, "销毁成功", c)
	}
}

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
		header request.HeaderRequest
	)
	if err := c.ShouldBindHeader(&header); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err := service.CloseConnection(&header); err != nil {
		global.GVA_LOG.Error("关闭连接失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "关闭连接失败", c)
	} else {
		response.OkWithMessage("关闭连接成功", c)
	}
}

func GameRequest(c *gin.Context) {
	var (
		err     error
		param   request.ParamRequest
		header  request.HeaderRequest
		request shared.GameRequest
	)

	if err = c.ShouldBindUri(&param); err != nil {
		global.GVA_LOG.Error("参数校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "参数校验错误", c)
		return
	}
	if err = c.ShouldBindHeader(&header); err != nil {
		global.GVA_LOG.Error("头部校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "头部校验错误", c)
		return
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		global.GVA_LOG.Error("Body校验错误！", zap.String("err", err.Error()))
		response.FailWithDetailed(err.Error(), "Body校验错误", c)
		return
	}

	if err, data := service.GameRequest(&header, &param, &request); err != nil {
		global.GVA_LOG.Error("发送请求失败!", zap.Any("err", err))
		response.FailWithDetailed(err.Error(), "发送请求失败", c)
	} else {
		response.OkWithDetailed(data, "发送请求成功", c)
	}
}
