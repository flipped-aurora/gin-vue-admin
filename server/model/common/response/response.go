package response

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Data    *JsonStruct
	Context *gin.Context
}

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success              = JsonStruct{Code: 0, Msg: "成功"}
	ParameterError       = JsonStruct{Code: 1000, Msg: "参数错误"}
	AccountErr           = JsonStruct{Code: 1001, Msg: "账号或密码错误"}
	AccountNotExit       = JsonStruct{Code: 1002, Msg: "账号不存在"}
	AccountExit          = JsonStruct{Code: 1003, Msg: "账号已存在"}
	SendSMSOk            = JsonStruct{Code: 1004, Msg: "验证码发送成功"}
	ExceedLimitRequestIP = JsonStruct{Code: 1005, Msg: "请求次数超限"}
	SMSMismatch          = JsonStruct{Code: 1006, Msg: "验证码错误"}
	AccountOffSiteLogin  = JsonStruct{Code: 1007, Msg: "异地登录，如非本人操作请尽快修改密码！！！"}
	AccountDisable       = JsonStruct{Code: 1008, Msg: "账号已被禁用，如有疑问请联系客服。"}

	TokenNotExit     = JsonStruct{Code: 2000, Msg: "令牌不存在"}
	TokenExpired     = JsonStruct{Code: 2001, Msg: "登录已过期，请重新登录。"}
	TokenInvalid     = JsonStruct{Code: 2002, Msg: "令牌不合法解析失败"}
	TokenNotValidYet = JsonStruct{Code: 2003, Msg: "令牌尚未激活"}

	TokenCreateError   = JsonStruct{Code: 2004, Msg: "令牌创建失败"}
	ServerError        = JsonStruct{Code: 9999, Msg: "服务器未知错误"}
	NotFound           = JsonStruct{Code: 4000, Msg: "请求url或资源不存在"}
	NotMethod          = JsonStruct{Code: 4001, Msg: "请求方法错误"}
	NullData           = JsonStruct{Code: 4002, Msg: "未查询到数据"}
	IPWhiteList        = JsonStruct{Code: 4003, Msg: "IP地址非法"}
	SignError          = JsonStruct{Code: 4004, Msg: "签名错误，非法请求"}
)

func NewDefaultResponse() Response {
	return Response{Data: &JsonStruct{Data: make(map[string]interface{}), Code: Success.Code, Msg: Success.Msg}}
}

type Handler interface {
	Success()
	Error(errorCode JsonStruct)
	ErrorUndef()
}

func (r *Response) Success() {
	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
}

func (r *Response) SuccessMsg(msg string) {
	r.Data.Msg = msg
	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
}

func (r *Response) AppendData(key string, value interface{}) *Response {
	r.Data.Data.(map[string]interface{})[key] = value
	return r
}

func (r *Response) SetData(value interface{}) *Response {
	r.Data.Data = value
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Data.Code = code
	return r
}

func (r *Response) ErrorUndef() {
	r.Data.Msg = ServerError.Msg
	r.Data.Code = ServerError.Code
	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
	withError := r.Context.AbortWithError(200, errors.New(r.Data.Msg))
	if withError != nil {
		global.GVA_LOG.Error("api-error", zap.Error(withError))
	}
}

func (r *Response) ErrorString(err string) {
	r.Data.Msg = err
	r.Data.Code = ServerError.Code

	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
	withError := r.Context.AbortWithError(200, errors.New(r.Data.Msg))
	if withError != nil {
		global.GVA_LOG.Error("api-error", zap.Error(withError))
	}
}

func (r *Response) ErrorValidators(msg string) {
	r.Data.Msg = msg
	r.Data.Code = ParameterError.Code

	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
	withError := r.Context.AbortWithError(200, errors.New(r.Data.Msg))
	if withError != nil {
		global.GVA_LOG.Error("api-error", zap.Error(withError))
	}
}

func (r *Response) Error(err JsonStruct) {
	r.Data.Msg = err.Msg
	r.Data.Code = err.Code

	r.Context.JSON(200, r.Data)
	resp, _ := json.Marshal(r.Data)
	r.Context.Set("response", string(resp))
	r.Context.Set("response_code", r.Data.Code)
	withError := r.Context.AbortWithError(200, errors.New(r.Data.Msg))
	if withError != nil {
		global.GVA_LOG.Error("api-error", zap.Error(withError))
	}
}

// 统一生成返回体进行格式统一返回
func MakeResponse(c *gin.Context) (response *Response) {
	responseNew := NewDefaultResponse()
	responseNew.Context = c
	return &responseNew
}

// 兼容原项目

func FailWithMessage(msg string, c *gin.Context) {
	MakeResponse(c).ErrorString(msg)
}

func OkWithMessage(msg string, c *gin.Context) {
	MakeResponse(c).SuccessMsg(msg)
}

func OkWithData(data interface{}, c *gin.Context) {
	MakeResponse(c).SetData(data).Success()
}

func OkWithDetailed(obj interface{}, msg string, c *gin.Context) {
	MakeResponse(c).SetData(obj).SuccessMsg(msg)
}

func FailWithDetailed(obj interface{}, msg string, c *gin.Context) {
	MakeResponse(c).SetData(obj).ErrorString(msg)
}
