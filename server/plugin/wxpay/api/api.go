package api

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat/v3"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type WxPay struct{}

// WxApiGetCode
// @Tags WxPay
// @Summary 微信获取code JsApi支付
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [get]
func (p *WxPay) WxApiGetCode(c *gin.Context) {
	var plug *request.JsapiGetToken
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	//-----------------
	var cfg = global.GVA_CONFIG.WxPay
	st := fmt.Sprintf("%s/api/wxpay/gettoken?attach=%s", cfg.Load, plug.Attach)
	encodedStr := url.QueryEscape(st)
	aToken := fmt.Sprintf(Urlstr, cfg.AppID, encodedStr)
	//-----
	c.Redirect(http.StatusFound, aToken)
	return
}

// WxApiGetToken
// @Tags WxPay
// @Summary 微信获取token
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [get]
func (p *WxPay) WxApiGetToken(c *gin.Context) {
	var plug *request.JsapiGetToken
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	//openid, err := service.ServiceGroupApp.PlugServiceGetToken(plug)
	cfg := global.GVA_CONFIG.WxPay
	sprintf := fmt.Sprintf(TokenUrl, cfg.AppID, cfg.Secret, plug.Code)

	str := utils.HttpGet(sprintf)
	//str := `{"access_token":"ACCESS_TOKEN","expires_in":7200,"refresh_token":"REFRESH_TOKEN","openid":"ovDTtjpt9FURG7ZUK80c6bMurQ0w","scope":"SCOPE","is_snapshotuser":1,"unionid":"UNIONID"}`
	//str := `{"access_token":"ACCESS_TOKEN","expires_in":7200,"refresh_token":"REFRESH_TOKEN","openid":"ovDTtjpt9FURG7ZUK80c6bMurQ0w","scope":"SCOPE","is_snapshotuser":1,"unionid":"UNIONID"}`
	var token AccessToken
	err = json.Unmarshal([]byte(str), &token)
	if err != nil {
		global.GVA_LOG.Error("获取openid失败", zap.Error(err))
		response.FailWithMessage("获取openid失败", c)
		return
	}

	if token.OpenID == "" {
		global.GVA_LOG.Error("获取openid失败", zap.Error(err))
		response.FailWithMessage("获取openid失败", c)
		return
	}
	//调到html页面显示金额
	locat := fmt.Sprintf("%s/toyPre.html?attach=%s&openId=%s", cfg.Load, plug.Attach, token.OpenID)
	c.Redirect(http.StatusFound, locat)
	return
}

// WxApiToPayCode
// @Tags WxPay
// @Summary 微信获取code JsApi支付 自定义金额
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [get]
func (p *WxPay) WxApiToPayCode(c *gin.Context) {
	var cfg = global.GVA_CONFIG.WxPay
	st := fmt.Sprintf("%s/api/wxpay/topaytoken?", cfg.Load)
	encodedStr := url.QueryEscape(st)
	aToken := fmt.Sprintf(Urlstr, cfg.AppID, encodedStr)
	c.Redirect(http.StatusFound, aToken)
	return
}

// WxApiToPayToken
// @Tags WxPay
// @Summary 微信获取token 自定义金额
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [get]
func (p *WxPay) WxApiToPayToken(c *gin.Context) {
	var plug *request.JsapiGetToken
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}

	cfg := global.GVA_CONFIG.WxPay
	sprintf := fmt.Sprintf(TokenUrl, cfg.AppID, cfg.Secret, plug.Code)
	str := utils.HttpGet(sprintf)
	var token AccessToken
	err = json.Unmarshal([]byte(str), &token)
	if err != nil {
		global.GVA_LOG.Error("获取openid失败", zap.Error(err))
		response.FailWithMessage("获取openid失败", c)
		return
	}

	if token.OpenID == "" {
		global.GVA_LOG.Error("获取openid失败", zap.Error(err))
		response.FailWithMessage("获取openid失败", c)
		return
	}
	//调到html页面显示金额
	locat := fmt.Sprintf("%s/topay.html?openId=%s", cfg.Load, token.OpenID)
	c.Redirect(http.StatusFound, locat)
	return
}

// WxApiNativeCode
// @Tags WxPay
// @Summary 微信获取下单接口 Native支付
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [post]
func (p *WxPay) WxApiNativeCode(c *gin.Context) {
	var plug request.Native
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	if res, err := service.ServiceGroupApp.PlugServiceNative(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		//CodeUrl
		response.OkWithDetailed(res, "成功", c)
		return
	}
}

// WxApiJsApi
// @Tags WxPay
// @Summary 微信获取下单接口 JsApi支付
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [post]
func (p *WxPay) WxApiJsApi(c *gin.Context) {
	var plug *request.Jsapi
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}

	str, err := utils.ReplaceAttach(plug.Attach)
	if err != nil {
		global.GVA_LOG.Error("支付失败", zap.Error(err))
		response.Fail(c)
		return
	}
	plug.OrderId = str.Get("outTreadNo")
	plug.GoodsID = str.Get("goodsid")

	if res, err := service.ServiceGroupApp.PlugServiceJsApi(plug); err != nil {
		global.GVA_LOG.Error("支付失败", zap.Error(err))
		response.Fail(c)
		return
	} else {
		//CodeUrl
		c.JSON(0, res)
		return
	}
}

// WxApiToPayJsApi
// @Tags WxPay
// @Summary 微信获取下单接口 JsApi支付 自定义金额
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/routerName [post]
func (p *WxPay) WxApiToPayJsApi(c *gin.Context) {
	var plug *request.ToPayJsapi
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	if plug.Money <= 0 {
		global.GVA_LOG.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	if res, err := service.ServiceGroupApp.PlugServiceToPayJsApi(plug); err != nil {
		global.GVA_LOG.Error("支付失败", zap.Error(err))
		response.Fail(c)
		return
	} else {
		//CodeUrl
		c.JSON(0, res)
		return
	}
}

// WxApiNativeNotifyUrl
// @Tags WxPay
// @Summary 微信回调接口
// @Security  ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/Notifyurl [post]
func (p *WxPay) WxApiNativeNotifyUrl(c *gin.Context) {
	//这里解析微信回调到结构体  下面简单判断一下是不是无效请求
	s, err := wechat.V3ParseNotify(c.Request)
	if err != nil {
		global.GVA_LOG.Error("微信回调接口失败!", zap.Error(errors.New("参数有问题，请检查")))
		response.FailWxCall("FAIL", "失败", c)
		return
	}
	if s.Id == "" {
		global.GVA_LOG.Error("微信回调接口失败!", zap.Error(errors.New("参数有问题，请检查")))
		response.FailWxCall("FAIL", "失败", c)
		return
	}
	if s.EventType == "" {
		global.GVA_LOG.Error("微信回调接口失败!", zap.Error(errors.New("参数有问题，请检查")))
		response.FailWxCall("FAIL", "失败", c)
		return
	}
	res := service.ServiceGroupApp.PlugServiceNotify(s)
	if res.Code == "" { //这里先判断是否为nil 不然会报指针错误
		global.GVA_LOG.Error("微信回调接口失败!", zap.Error(errors.New("参数有问题，请检查")))
		response.FailWxCall("FAIL", "失败", c)
		return
	}
	if res.Code == "FAIL" {
		global.GVA_LOG.Error("微信回调接口失败!", zap.Error(errors.New(res.Message)))
		response.FailWxCall(res.Code, res.Message, c)

		return
	}
	if res.Code == "OK" {
		response.OkWxCall(c)
		return
	}
}

// WxApiNativeQueryOrder
// @Tags WxPay
// @Summary 微信订单查询接口 Native支付
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/queryorder [get]
func (p *WxPay) WxApiNativeQueryOrder(c *gin.Context) {
	var plug request.ReqTradeNo
	var err = c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("微信订单查询绑定数据失败", zap.Error(err))
		response.FailWithMessage("微信订单查询绑定数据失败", c)
		return
	}

	if res, err := service.ServiceGroupApp.PlugServiceQueryOrder(plug.OrderId); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	} else {
		response.OkWithDetailed(res, "成功", c)
		return
	}
}

// WxApiNativeCloseOrder
// @Tags WxPay
// @Summary 微信关闭订单 Native支付
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/CloseOrder [get]
func (p *WxPay) WxApiNativeCloseOrder(c *gin.Context) {
	var plug request.ReqTradeNo
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("关闭订单绑定数据失败", zap.Error(err))
		response.FailWithMessage("关闭订单绑定数据失败", c)
		return
	}
	var str string
	if str, err = service.ServiceGroupApp.PlugServiceCloseOrder(plug.OrderId); err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	if str != "" {
		global.GVA_LOG.Error(str, zap.Error(err))
		response.FailWithMessage(str, c)
		return
	} else {
		response.Ok(c)
		return
	}
}

// WxApiNativeRefunds
// @Tags WxPay
// @Summary 微信订单退款
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/CloseOrder [get]
func (p *WxPay) WxApiNativeRefunds(c *gin.Context) {
	var plug request.ReqTradeNo
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("订单退款绑定数据失败", zap.Error(err))
		response.FailWithMessage("订单退款绑定数据失败", c)
		return
	}
	n, err := service.ServiceGroupApp.PlugServerRefunds(plug.OrderId)
	if err != nil {
		global.GVA_LOG.Error("微信退款失败", zap.Error(err))
		response.FailWithMessage("微信退款失败", c)
		return
	}
	if n != 0 {
		global.GVA_LOG.Error("账户更新数据条数为0", zap.Error(errors.New("账户更新数据条数为0")))
		response.FailWithMessage("微信退款失败", c)
		return
	}
	response.Ok(c)
	return

}

// WxTest 获取版本
func (p *WxPay) WxTest(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var m map[string]interface{}
	json.Unmarshal(jsonData, &m)
	fmt.Println(m)

	//log.Print("handle log")
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	fmt.Println("header \r\n", c.Request.Header)

	a := make(map[string]string)
	a["code"] = strconv.Itoa(0)
	a["msg"] = "success"
	a["data"] = ""
	a["ServerTime"] = time.Now().Format(time.DateTime)

	c.JSON(http.StatusOK, a)
}

// WxgetVolume 获取音量
func (p *WxPay) WxgetVolume(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var m map[string]interface{}
	json.Unmarshal(jsonData, &m)
	fmt.Println(m)
	//{"code":0,"msg":"success","data":"修改成功","ServerTime":"2023-11-04 10:02:39"}

	a := make(map[string]string)
	a["code"] = strconv.Itoa(0)
	a["msg"] = "success"
	a["data"] = "修改成功"
	a["ServerTime"] = time.Now().Format(time.DateTime)

	c.JSON(http.StatusOK, a)

}

// WxcandyInfo 返回棉花糖数据
func (p *WxPay) WxcandyInfo(c *gin.Context) {
	f, err := os.OpenFile("./candy_info", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	type_b, _ := buf.Peek(512)

	c.Writer.Header().Set("Content-type", http.DetectContentType(type_b)) //设置文件格式

	fileinfo, _ := f.Stat()
	//"application/octet-stream"
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(fileinfo.Size(), 10))

	//if ok, _ := strconv.ParseBool(show); ok {
	//	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+filename) //下载文件名,不设置网页直接打开PDF.jpg.png 等格式
	//}

	buf_b := make([]byte, 1024*1024) //发送大小
	for {
		n, err := buf.Read(buf_b)
		if err == io.EOF || n == 0 {
			break
		}
		c.Writer.Write(buf_b[:n])
	}
	//fmt.Println(filename, "结束")
	//c.JSON(http.StatusOK, bm)

}

func (p *WxPay) WxsugarFlow(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var m map[string]interface{}
	json.Unmarshal(jsonData, &m)
	fmt.Println(m)

}
