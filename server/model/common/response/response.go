package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResQueryOrder struct {
	Code        int    `json:"code"`
	TradeStatus string `json:"tradeStatus"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Responsewx struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	ServerTime string      `json:"ServerTime"`
}

type ResponseWx struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultWxFail(code string, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusBadRequest, ResponseWx{
		code,
		msg,
	})
}

func ResultWx(code int, data interface{}, msg, ServerTime string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Responsewx{
		code,
		data,
		msg,
		ServerTime,
	})
}

func ResultWxOK(c *gin.Context) {
	c.Status(http.StatusOK)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWxCall(cede string, message string, c *gin.Context) {
	ResultWxFail(cede, message, c)
}

func OkWxCall(c *gin.Context) {
	ResultWxOK(c)
}

func WxQrCode(data interface{}, message, serverTime string, c *gin.Context) {
	ResultWx(SUCCESS, data, message, serverTime, c)
}

func OkMhtQueryOrder(code int, tradeStatus, status, time string, c *gin.Context) {
	//{"code":0,"tradeStatus":"支付成功","status":"S","timestamp":"1699067811355"}
	c.JSON(http.StatusOK, ResQueryOrder{
		Code:        code,
		TradeStatus: tradeStatus,
		Status:      status,
		Timestamp:   time,
	})
}

func FailMhtQueryOrder(code int, tradeStatus, status, time string, c *gin.Context) {
	//{"code":0,"tradeStatus":"支付成功","status":"S","timestamp":"1699067811355"}
	c.JSON(http.StatusOK, ResQueryOrder{
		Code:        code,
		TradeStatus: tradeStatus,
		Status:      status,
		Timestamp:   time,
	})
}
