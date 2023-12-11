package shop

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat/v3"
	"go.uber.org/zap"
	"net/http"
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
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功","serverTime":"2023-12-06 23:03:00"}
// @Router /api/pay/createShopOrders [post]
func (shopQrcodeApi *ShopQrcodeApi) CreateShopQrcode(c *gin.Context) {
	//{ts=1699067732226, key=658241451, patternId=3, macid=ZZ9KDT845Z}
	body, err := c.GetRawData() //获取请求参数body原始数据
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if len(body) <= 1 {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	var strbody *request.RequestMhtData //解析棉花糖机器请求参数
	strbody = utils.ParseRequest(string(body))
	if strbody.PatternId == "" || strbody.Macid == "" {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("解析请求错误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if url, err := shopQrcodeService.ServiceCreateOrders(strbody); err != nil {
		global.GVA_LOG.Error("预下单失败", zap.Error(err))
		response.FailWithMessage("预下单失败,请检查参数", c)
	} else {
		//global.GVA_REDIS.Set(context.Background(), "", url, timer)
		ts := time.Now().Format(time.DateTime)
		response.WxQrCode(url, wechat.TradeStateSuccess, ts, c)
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
// @Router /api/pay/openId [post]
func (shopQrcodeApi *ShopQrcodeApi) GetOpenId(c *gin.Context) {
	//{ts=1699067732226, key=658241451, patternId=3, macid=ZZ9KDT845Z}
	var plug *request.Attach
	err := c.ShouldBindQuery(&plug)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if plug.Attach == "" {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
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
// @Router /api/pay/QueryOrder [post]
func (shopQrcodeApi *ShopQrcodeApi) QueryOrder(c *gin.Context) {
	//{ts=1699067732664, hlMerchantId=1369596012470, key=625031257, outTreadNo=569621434041902, macid=ZZ9KDT845Z, agencyNo=1226862}
	ts := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	body, err := c.GetRawData() //获取请求参数body原始数据
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if len(body) <= 1 {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("请求参数有误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	var strbody *request.RequestMhtData //解析棉花糖机器请求参数
	strbody = utils.ParseRequest(string(body))
	if strbody.OutTreadNo == "" {
		global.GVA_LOG.Error("失败!", zap.Error(errors.New("解析请求错误")))
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if n, err := shopQrcodeService.ServiceQueryOrders(strbody); err != nil {
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
