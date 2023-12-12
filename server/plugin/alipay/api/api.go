package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	ali "github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

const (
	//https://openapi.alipay.com/gateway.do?timestamp=2013-01-01 08:08:08&method=alipay.system.oauth.token&app_id=36340&sign_type=RSA2&sign=ERITJKEIJKJHKKKKKKKHJEREEEEEEEEEEE&version=1.0&charset=GBK&grant_type=authorization_code&code=4b203fe6c11548bcabd8da5bb087a83b&refresh_token=201208134b203fe6c11548bcabd8da5bb087a83b
	AuthCode      = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm?app_id=%s&scope=%s&redirect_uri=%s"               //正式环境使用
	DebugAuthcode = "https://openauth-sandbox.dl.alipaydev.com/oauth2/publicAppAuthorize.htm?app_id=%s&scope=%s&redirect_uri=%s" //沙箱环境使用
)

type AlipayApi struct{}

// ApiGetCode
// @Tags Alipay
// @Summary 获取code
// @Produce  application/json
// @Success 301 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/getcode [post]
func (p *AlipayApi) ApiGetCode(c *gin.Context) {
	var plug *request.ResAuthCode
	var str string
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	//把参数字符串编码一下
	s1 := fmt.Sprintf("%s/api/alipay/gettoken?attach=%s", global.GVA_CONFIG.AliPay.Load, plug.Attach)
	data := url.QueryEscape(s1)
	//判断一下是否是沙箱模式
	if ali.Client.IsProd == true {
		str = fmt.Sprintf(AuthCode, ali.Client.AppId, "auth_base", data)
	} else {
		str = fmt.Sprintf(DebugAuthcode, ali.Client.AppId, "auth_base", data)
	}
	c.Redirect(http.StatusFound, str)
	return
}

// ApiGetToken
// @Tags Alipay
// @Summary 获取token
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/gettoken [post]
func (p *AlipayApi) ApiGetToken(c *gin.Context) {
	var plug *request.ResAuthToken
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("获取参数失败", c)
		return
	}
	if plug.Code == "" {
		global.GVA_LOG.Error("失败", zap.Error(errors.New("code不能为空")))
		response.FailWithMessage("code不能为空", c)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", plug.Code)
	aliRsp, err := ali.Client.SystemOauthToken(context.Background(), bm)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	//调到html页面显示金额
	locat := fmt.Sprintf("%s/toyPre.html?attach=%s&ailopenId=%s", global.GVA_CONFIG.AliPay.Load, plug.Attach, aliRsp.Response.OpenId)
	c.Redirect(http.StatusFound, locat)
	return
}

// ApiToPayGetCode
// @Tags Alipay
// @Summary 自定义金额付款获取code
// @Produce  application/json
// @Success 301 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/getcode [post]
func (p *AlipayApi) ApiToPayGetCode(c *gin.Context) {
	var str string
	//把参数字符串编码一下
	s1 := fmt.Sprintf("%s/api/alipay/topaygettoken?", global.GVA_CONFIG.AliPay.Load)
	data := url.QueryEscape(s1)
	//判断一下是否是沙箱模式
	if ali.Client.IsProd == true {
		str = fmt.Sprintf(AuthCode, ali.Client.AppId, "auth_base", data)
	} else {
		str = fmt.Sprintf(DebugAuthcode, ali.Client.AppId, "auth_base", data)
	}
	c.Redirect(http.StatusFound, str)
	return
}

// ApiToPayGetToken
// @Tags Alipay
// @Summary 自定义金额付款获取token
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/topaygettoken [post]
func (p *AlipayApi) ApiToPayGetToken(c *gin.Context) {
	var plug *request.ResAuthToken
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("获取参数失败", c)
		return
	}
	if plug.Code == "" {
		global.GVA_LOG.Error("失败", zap.Error(errors.New("code不能为空")))
		response.FailWithMessage("code不能为空", c)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", plug.Code)
	aliRsp, err := ali.Client.SystemOauthToken(context.Background(), bm)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	//调到html页面显示金额
	locat := fmt.Sprintf("%s/topay.html?ailopenId=%s", global.GVA_CONFIG.AliPay.Load, aliRsp.Response.OpenId)
	c.Redirect(http.StatusFound, locat)
	return
}

// ApiPayJsapi
// @Tags Alipay
// @Summary 支付宝预下单
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/payjsapi [post]
func (p *AlipayApi) ApiPayJsapi(c *gin.Context) {
	var plug *request.ResAuthJsApi
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	if res, err := service.ServiceGroupApp.PlugServiceJsPay(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	} else {
		response.OkWithData(res.Response.TradeNo, c)
		return
	}
}

// ApiToPayJsapi
// @Tags Alipay
// @Summary 支付宝下单，自定义金额
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/payjsapi [post]
func (p *AlipayApi) ApiToPayJsapi(c *gin.Context) {
	var plug *request.ResAuthJsApi
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	if plug.Money <= 0 {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("金额小于等于0")))
		response.FailWithMessage("失败", c)
		return
	}
	if res, err := service.ServiceGroupApp.PlugServiceToJsPay(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	} else {
		response.OkWithData(res.Response.TradeNo, c)
		return
	}
}

// ApiNotify
// @Tags Alipay
// @Summary 支付宝回调
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/notify [post]
func (p *AlipayApi) ApiNotify(c *gin.Context) {
	bodyMap, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("fail", c)
		return
	}
	cert, err := alipay.VerifySignWithCert(global.GVA_CONFIG.AliPay.AlipayPublicContentRSA2, bodyMap)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("fail", c)
		return
	}
	if cert != true {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("验签失败")))
		response.FailWithMessage("fail", c)
		return
	}
	if err := service.ServiceGroupApp.PlugServiceNotify(&bodyMap); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		c.String(http.StatusBadRequest, "fail")
		return
	} else {
		c.String(http.StatusOK, "success")
		return
	}
}

// ApiRefunds
// @Tags Alipay
// @Summary 支付宝退款
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /alipay/refunds [post]
func (p *AlipayApi) ApiRefunds(c *gin.Context) {
	var plug *request.ReqTradeNo
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("支付宝订单退款绑定数据失败!", zap.Error(err))
		response.FailWithMessage("订单退款绑定数据失败", c)
		return
	}
	order, err := service.ServiceGroupApp.PlugServiceRefund(plug.OutTradeNo)
	if err != nil {
		global.GVA_LOG.Error("支付宝退款失败!", zap.Error(err))
		response.FailWithMessage("支付宝退款失败", c)
		return
	}
	if order.TradeState != utils.REFUND {
		global.GVA_LOG.Error("订单状态异常", zap.Error(errors.New("TradeState字段值异常")))
		response.FailWithMessage("订单状态异常", c)
		return
	}
	response.Ok(c)
	return
}
