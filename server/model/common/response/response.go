package response

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
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

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

type ResponseV2 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ResultV2(code int, data interface{}, msg string, c *gin.Context) {
	responseData := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	// 反射遍历 data 并添加到 responseData 中
	if data != nil {
		val := reflect.ValueOf(data)
		typ := reflect.TypeOf(data)

		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Tag.Get("json")
			if fieldName == "" {
				fieldName = typ.Field(i).Name
			}
			responseData[fieldName] = val.Field(i).Interface()
		}
	}
	c.JSON(http.StatusOK, responseData)
}

func OkV2(c *gin.Context) {
	ResultV2(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessageV2(message string, c *gin.Context) {
	ResultV2(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithDataV2(data interface{}, c *gin.Context) {
	ResultV2(SUCCESS, data, "查询成功", c)
}

func OkWithDetailedV2(data interface{}, message string, c *gin.Context) {
	ResultV2(SUCCESS, data, message, c)
}

func FailV2(c *gin.Context) {
	ResultV2(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessageV2(message string, c *gin.Context) {
	ResultV2(ERROR, map[string]interface{}{}, message, c)
}

func NoAuthV2(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ResponseV2{
		Code: 7,
		Msg:  message,
	})
}

func FailWithDetailedV2(data interface{}, message string, c *gin.Context) {
	ResultV2(ERROR, data, message, c)
}
