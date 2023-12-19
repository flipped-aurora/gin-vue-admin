package shop

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat/v3"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	wxclietn  = "MicroMessenger"
	aliclient = "AlipayClient"
)

type ShopQrcodeApi struct {
}

var shopQrcodeService = service.ServiceGroupApp.ShopServiceGroup.ShopQrcodeService

// CreateShopQrcode
// @Tags Qrcode
// @Summary 支付预下单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopQrcodeRouter true "支付预下单"
// @Success 200 {string} map {"success":true,"data":{},"msg":"创建成功","serverTime":"2023-12-06 23:03:00"}
// @Router /api/mht/createShopOrders [post]
func (shopQrcodeApi *ShopQrcodeApi) CreateShopQrcode(c *gin.Context) {
	//ts=1702434294592&amp;key=539000891&amp;patternId=13&amp;macid=ZZ9KDT845Z
	body, err := c.GetRawData() //获取请求参数body原始数据
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	//替换请求里面的多余参数
	str := strings.ReplaceAll(string(body), "&amp;", "&")
	// 将字符串转换为 map 格式
	query, err := url.ParseQuery(str)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	//var plug request.RequestMhtData
	plug := request.RequestMhtData{
		Macid:     query.Get("macid"),
		PatternId: query.Get("patternId"),
	}

	if plug.PatternId == "" || plug.Macid == "" {
		global.GVA_LOG.Error("失败!", zap.String("参数", "CreateShopQrcode的macid或者patternid为空"))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if urls, err := shopQrcodeService.ServiceCreateOrders(&plug); err != nil {
		global.GVA_LOG.Error("失败", zap.Error(err))
		response.FailWithMessage("预下单失败,请检查参数", c)
	} else {
		//global.GVA_REDIS.Set(context.Background(), "", url, timer)
		ts := time.Now().Format(time.DateTime)
		response.WxQrCode(urls, wechat.TradeStateSuccess, ts, c)
	}
}

// GetOpenId
// @Tags Qrcode
// @Summary 判断用什么设备扫码，并生成url
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopQrcodeRouter true "跳转"
// @Success 302
// @Router /api/mht/openId [post]
func (shopQrcodeApi *ShopQrcodeApi) GetOpenId(c *gin.Context) {
	var plug *request.Attach
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if plug.Attach == "" {
		global.GVA_LOG.Error("失败!", zap.String("Attach为空", plug.Attach))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	agent := c.Request.UserAgent()
	if strings.Contains(agent, wxclietn) {
		locat := fmt.Sprintf("%s/api/wxpay/getcode?attach=%s", global.GVA_CONFIG.WxPay.Load, plug.Attach)
		c.Redirect(http.StatusFound, locat)
		return
	} else if strings.Contains(agent, aliclient) {
		locat := fmt.Sprintf("%s/api/alipay/getcode?attach=%s", global.GVA_CONFIG.AliPay.Load, plug.Attach)
		c.Redirect(http.StatusFound, locat)
		return
	} else {
		locat := fmt.Sprintf("/wxpay.html?%s", plug.Attach)
		c.Redirect(http.StatusFound, locat)
		return
	}
}

// QueryOrder
// @Tags Qrcode
// @Summary 棉花糖查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopQrcodeRouter true "支付预下单"
// @Success 302 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/mht/QueryOrder [post]
func (shopQrcodeApi *ShopQrcodeApi) QueryOrder(c *gin.Context) {
	//ts=1702434294592&amp;key=539000891&amp;patternId=13&amp;macid=ZZ9KDT845Z
	body, err := c.GetRawData() //获取请求参数body原始数据
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	//替换请求里面的多余参数
	str := strings.ReplaceAll(string(body), "&amp;", "&")
	// 将字符串转换为 map 格式
	query, err := url.ParseQuery(str)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	//var plug request.RequestMhtData
	plug := request.RequestMhtData{
		OutTreadNo: query.Get("outTreadNo"),
	}
	ts := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	if n, err := shopQrcodeService.ServiceQueryOrders(&plug); err != nil {
		global.GVA_LOG.Error("查询订单失败", zap.Error(err))
		response.FailMhtQueryOrder(n, "订单还未支付", "F", ts, c)
	} else {
		response.OkMhtQueryOrder(n, "支付成功", "S", ts, c)
	}
}

//------------------------------------------------------------下面是自定义金额

// ToTayGetOpenId
// @Tags Qrcode
// @Summary 判断用什么设备扫码，并生成url,自定义金额用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopQrcodeRouter true "跳转"
// @Success 302
// @Router /api/pay/topey [post]
func (shopQrcodeApi *ShopQrcodeApi) ToTayGetOpenId(c *gin.Context) {
	agent := c.Request.UserAgent()
	if strings.Contains(agent, wxclietn) {
		locat := fmt.Sprintf("%s/api/wxpay/topaycode", global.GVA_CONFIG.WxPay.Load)
		c.Redirect(http.StatusFound, locat)
		return
	} else if strings.Contains(agent, aliclient) {
		locat := fmt.Sprintf("%s/api/alipay/topaycode", global.GVA_CONFIG.AliPay.Load)
		c.Redirect(http.StatusFound, locat)
		return
	} else {
		locat := fmt.Sprintf("/wxpay.html")
		c.Redirect(http.StatusFound, locat)
		return
	}
}

// UserRefund
// @Tags Qrcode
// @Summary 棉花糖用户自定退款
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopQrcodeRouter true "支付预下单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"退款成功"}"
// @Router /api/mht/userRefund [post]
func (shopQrcodeApi *ShopQrcodeApi) UserRefund(c *gin.Context) {
	tradeNo := c.Param("OutTradeNo")
	if len(tradeNo) == 0 {
		global.GVA_LOG.Error("失败!", zap.String("参数有误", tradeNo))
		response.FailWithMessage("请求参数有误", c)
		return
	}

	if err := shopQrcodeService.ServiceUserRefund(tradeNo); err != nil {
		global.GVA_LOG.Error("退款失败", zap.Error(err))
		response.FailWithMessage("退款失败", c)
	} else {
		response.OkWithMessage("退款成功", c)
	}
}
