// Package response 提供统一的API响应结构体和工具函数
//
// 功能说明：
// 本包定义了gin-vue-admin框架的标准API响应格式，所有API接口都应该使用
// 本包提供的函数来返回响应，确保响应格式的一致性。
//
// 核心特性：
// - 统一的响应结构体（Response）
// - 成功响应的快捷方法（Ok系列）
// - 失败响应的快捷方法（Fail系列）
// - 支持自定义消息和数据
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一的API响应结构体
//
// 功能说明：
// 所有API接口都应该使用此结构体返回响应，确保前端能够统一处理。
//
// 字段说明：
// - Code: 响应状态码，0表示成功，7表示失败
// - Data: 响应数据，可以是任意类型（对象、数组、字符串等）
// - Msg: 响应消息，用于提示用户操作结果
//
// 使用示例：
//
//	// 成功响应示例
//	response.OkWithData(gin.H{"id": 1, "name": "张三"}, c)
//	// 返回: {"code": 0, "data": {"id": 1, "name": "张三"}, "msg": "成功"}
//
//	// 失败响应示例
//	response.FailWithMessage("用户不存在", c)
//	// 返回: {"code": 7, "data": {}, "msg": "用户不存在"}
type Response struct {
	Code int         `json:"code"` // 状态码：0=成功，7=失败
	Data interface{} `json:"data"` // 响应数据，可以是任意类型
	Msg  string      `json:"msg"`  // 响应消息
}

const (
	ERROR   = 7 // 错误状态码
	SUCCESS = 0 // 成功状态码
)

// Result 构建并返回JSON响应
//
// 功能说明：
// 这是所有响应函数的底层实现，统一处理响应格式。
//
// 参数说明：
// - code: 状态码（SUCCESS或ERROR）
// - data: 响应数据
// - msg: 响应消息
// - c: Gin上下文对象
//
// 使用示例：
//
//	Result(SUCCESS, gin.H{"id": 1}, "操作成功", c)
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok 返回成功响应（无数据，默认消息）
//
// 功能说明：
// 最简单的成功响应，用于不需要返回数据的操作（如删除、更新等）。
//
// 使用示例：
//
//	// 删除成功后
//	response.Ok(c)
//	// 返回: {"code": 0, "data": {}, "msg": "操作成功"}
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// OkWithMessage 返回成功响应（无数据，自定义消息）
//
// 功能说明：
// 成功响应，但需要自定义提示消息。
//
// 参数说明：
// - message: 自定义的成功消息
//
// 使用示例：
//
//	response.OkWithMessage("用户创建成功", c)
//	// 返回: {"code": 0, "data": {}, "msg": "用户创建成功"}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// OkWithData 返回成功响应（有数据，默认消息）
//
// 功能说明：
// 成功响应并返回数据，消息使用默认值"成功"。
//
// 参数说明：
// - data: 要返回的数据，可以是对象、数组等
//
// 使用示例：
//
//	user := User{ID: 1, Name: "张三"}
//	response.OkWithData(user, c)
//	// 返回: {"code": 0, "data": {"id": 1, "name": "张三"}, "msg": "成功"}
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

// OkWithDetailed 返回成功响应（有数据，自定义消息）
//
// 功能说明：
// 最完整的成功响应，同时指定数据和消息。
//
// 参数说明：
// - data: 要返回的数据
// - message: 自定义的成功消息
//
// 使用示例：
//
//	response.OkWithDetailed(gin.H{"id": 1}, "查询成功", c)
//	// 返回: {"code": 0, "data": {"id": 1}, "msg": "查询成功"}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// Fail 返回失败响应（无数据，默认消息）
//
// 功能说明：
// 最简单的失败响应，用于操作失败但不需要详细错误信息的场景。
//
// 使用示例：
//
//	if err != nil {
//	    response.Fail(c)
//	    return
//	}
//	// 返回: {"code": 7, "data": {}, "msg": "操作失败"}
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

// FailWithMessage 返回失败响应（无数据，自定义消息）
//
// 功能说明：
// 失败响应，需要自定义错误消息。
//
// 参数说明：
// - message: 自定义的错误消息
//
// 使用示例：
//
//	if user == nil {
//	    response.FailWithMessage("用户不存在", c)
//	    return
//	}
//	// 返回: {"code": 7, "data": {}, "msg": "用户不存在"}
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// NoAuth 返回未授权响应
//
// 功能说明：
// 专门用于身份验证失败的情况，HTTP状态码为401。
//
// 参数说明：
// - message: 未授权的原因说明
//
// 使用示例：
//
//	if !isAuthenticated {
//	    response.NoAuth("请先登录", c)
//	    return
//	}
//	// 返回HTTP 401: {"code": 7, "data": null, "msg": "请先登录"}
func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

// FailWithDetailed 返回失败响应（有数据，自定义消息）
//
// 功能说明：
// 失败响应，同时返回错误相关的数据（如验证错误详情）。
//
// 参数说明：
// - data: 错误相关的数据（如字段验证错误列表）
// - message: 自定义的错误消息
//
// 使用示例：
//
//	errors := map[string]string{
//	    "email": "邮箱格式不正确",
//	    "phone": "手机号不能为空",
//	}
//	response.FailWithDetailed(errors, "数据验证失败", c)
//	// 返回: {"code": 7, "data": {"email": "邮箱格式不正确", ...}, "msg": "数据验证失败"}
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
