package service

import (
	"context"
	"errors"
	glog "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"go.uber.org/zap"
	"time"
)

type PayService struct{}

// PlugServiceJsApi JsApi下单
func (e *PayService) PlugServiceJsApi(req *request.Jsapi) (*wechat.JSAPIPayParams, error) {
	var cfg = global.GlobalConfig
	var tradeNo = req.OrderId
	//从数据库把订单信息取出来
	var orders = shop.ShopOrders{}
	err := glog.GVA_DB.Where("OutTradeNo = ?", tradeNo).First(&orders).Error
	if err != nil {
		return nil, err
	}
	if orders.TradeState == wechat.TradeStateSuccess {
		return nil, errors.New("订单已支付")
	}
	//生成微信下单body数据
	bm := WxJsapiBody(cfg.AppID, cfg.MchID, tradeNo, orders.GoodsTitle, cfg.NotifyUrl, req.Attach, req.OpenId, *orders.Total) //数据组装
	Jsapi, err := global.WxClient.V3TransactionJsapi(context.Background(), bm)
	if err != nil {
		return nil, err
	}
	if Jsapi.Code != wechat.Success {
		return nil, errors.New(Jsapi.Error)
	}
	if err != nil {
		return nil, errors.New(Jsapi.Error)
	}
	jsapi, err := global.WxClient.PaySignOfJSAPI(cfg.AppID, Jsapi.Response.PrepayId)
	if err != nil {
		return nil, err
	}
	return jsapi, nil
}

// PlugServiceToPayJsApi JsApi下单
func (e *PayService) PlugServiceToPayJsApi(req *request.ToPayJsapi) (*wechat.JSAPIPayParams, error) {
	var cfg = global.GlobalConfig
	var tradeNo = utils.GetOrderNumber()
	title := "云南串先生扫码付"
	//生成微信下单body数据
	bm := WxJsapiBody(cfg.AppID, cfg.MchID, tradeNo, title, cfg.NotifyUrl, "", req.OpenId, req.Money) //数据组装
	Jsapi, err := global.WxClient.V3TransactionJsapi(context.Background(), bm)
	if Jsapi.Code != wechat.Success {
		return nil, errors.New(Jsapi.Error)
	}
	if err != nil {
		return nil, errors.New(Jsapi.Error)
	}
	var orders = shop.ShopOrders{}
	orders.OpenId = req.OpenId
	orders.Total = &req.Money
	orders.TradeState = wechat.TradeStateNoPay
	orders.GoodsTitle = title
	orders.PayMent = utils.Wxpay
	orders.GoodsPrice = &req.Money

	orders.OutTradeNo = tradeNo
	orders.PostFee = nil

	err = glog.GVA_DB.Save(&orders).Error
	if err != nil {
		return nil, errors.New("内部错误，请检查参数")
	}

	jsapi, err := global.WxClient.PaySignOfJSAPI(cfg.AppID, Jsapi.Response.PrepayId)
	if err != nil {
		return nil, err
	}
	return jsapi, nil
}

// PlugServiceNative Native预下单
func (e *PayService) PlugServiceNative(req request.Native) (string, error) {
	var cfg = global.GlobalConfig
	var tradeNo = utils.GetOrderNumber()
	//从数据库把商品信息取出来
	var goods = shop.ShopGoods{}
	err := glog.GVA_DB.Where("id = ?", req.GoodsID).First(&goods).Error
	if err != nil {
		return "", err
	}
	//生成微信下单body数据
	Total := *goods.GoodsPrice + *goods.GoodsPostfee                                                     //这里是商品金额加上商品的邮费
	bm := WxNativeBody(cfg.AppID, cfg.MchID, tradeNo, goods.GoodsName, cfg.NotifyUrl, req.Attach, Total) //数据组装
	res, err := global.WxClient.V3TransactionNative(context.Background(), bm)

	if err != nil {
		glog.GVA_LOG.Error("下单失败!", zap.Error(err))
		return "", err
	}
	if res.Error != "" {
		glog.GVA_LOG.Error("微信下单返回错误状态："+res.Error, zap.Error(err))
		return "", errors.New(StrToJson(res.Error, "message"))
	}
	if res.Code != wechat.Success {
		glog.GVA_LOG.Error("微信下单返回错误状态：", zap.Error(err))
		return "", errors.New(StrToJson(res.Error, "message"))
	}
	order := utils.SetOrderData(cfg.AppID, cfg.MchID, req.Attach, req.CallUrl, " req.OpenId") //组装一下更新数据库需要的数据
	err = glog.GVA_DB.Create(&order).Error                                                    //写入数据库
	if err != nil {
		glog.GVA_LOG.Error("微信下单返回错误状态："+res.Error, zap.Error(err))
		return "", err
	} else {
		return res.Response.CodeUrl, nil
	}
}

// PlugServiceNotify 微信回调
func (e *PayService) PlugServiceNotify(req *wechat.V3NotifyReq) (res wechat.V3NotifyRsp) {
	decode, err := wechat.V3DecryptNotifyCipherText(req.Resource.Ciphertext, req.Resource.Nonce, req.Resource.AssociatedData, global.GlobalConfig.MchAPIv3Key)
	if err != nil {
		glog.GVA_LOG.Error("微信回调失败!", zap.Error(err))
		return res
	}
	var order = shop.ShopOrders{}
	if decode.TradeState != wechat.TradeStateNoPay {
		//这里判断只要不是未支付  就修改订单状态
		utils.SetNotifyData(&order, decode)
		err = glog.GVA_DB.Where("OutTradeNo = ?", decode.OutTradeNo).Updates(&order).Error
		if err != nil {
			res.Code = "FAIL"
			res.Message = decode.TradeState
			glog.GVA_LOG.Error("微信回调修改订单状态失败!", zap.Error(err))
			return res
		} else {
			res.Code = "OK"
			res.Message = ""
			return res
		}
	}
	return res
}

// PlugServiceQueryOrder 查询订单
func (e *PayService) PlugServiceQueryOrder(req string) (res response.ResqQueryOrder, err error) {
	wxRsp, err := global.WxClient.V3TransactionQueryOrder(context.Background(), wechat.OutTradeNo, req)
	if err != nil {
		glog.GVA_LOG.Error("查询微信订单失败", zap.Error(err))
		res.Code = 2
		res.Status = "S"
		res.TradeStatus = "支付失败"
		res.Timestamp = string(time.Now().Unix())
		return res, err
	}
	if wxRsp.Code != wechat.Success {
		glog.GVA_LOG.Error("查询微信订单失败", zap.String("err", wxRsp.Error))
		return res, err

	}
	if wxRsp.Response.TradeState == wechat.TradeStateSuccess {
		return res, err
	}
	return res, err

}

// PlugServiceCloseOrder 关闭订单
func (e *PayService) PlugServiceCloseOrder(tradeNo string) (res string, err error) {
	wxRsp, err := global.WxClient.V3TransactionCloseOrder(context.Background(), tradeNo)
	if err != nil {
		glog.GVA_LOG.Error("微信关闭订单错误", zap.Error(err))
		return wxRsp.Error, err
	}
	if wxRsp.Code != wechat.Success {
		glog.GVA_LOG.Error("微信关闭订单错误", zap.String("err", wxRsp.Error))
		return wxRsp.Error, errors.New("订单状态错误")
	}
	//关闭订单更新数据库状态
	var order = &shop.ShopOrders{}
	order.TradeState = wechat.TradeStateClosed
	t := time.Now()
	order.Endtime = &t
	n := glog.GVA_DB.Where("outtradeno = ? and TradeState != ?", tradeNo, wechat.TradeStateClosed).Updates(order).RowsAffected
	if n == 0 {
		glog.GVA_LOG.Error("微信关闭订单错误", zap.String("error", "没有找到需要关闭的订单"))
		return wxRsp.Error, errors.New("找不到指定的订单")
	}
	return wxRsp.Error, nil
}

// 微信退款
func (e *PayService) PlugServerRefunds(tradeNo string) (order *shop.ShopOrders, err error) {
	err = glog.GVA_DB.Where("outtradeno = ? and TradeState = ?", tradeNo, wechat.TradeStateSuccess).Find(&order).Error
	if err != nil {
		return nil, err
	}
	if len(order.OpenId) == 0 {
		return nil, errors.New("订单不存在或参数有误")
	}

	bm := make(gopay.BodyMap)
	bm.Set("out_refund_no", order.OutTradeNo). //商户系统内部的退款单号
							Set("transaction_id", order.TransactionId).   //原支付交易对应的微信订单号
							SetBodyMap("amount", func(bm gopay.BodyMap) { //订单金额信息
			bm.Set("refund", order.Total). //退款金额
							Set("currency", "CNY").
							Set("total", order.Total) //原支付交易的订单总金额
		})
	rsp, err := global.WxClient.V3Refund(context.Background(), bm)
	if err != nil {
		return nil, err
	}
	if rsp.Code != wechat.Success {
		return nil, errors.New(rsp.Error)
	}
	t := time.Now()
	order.Endtime = &t
	order.TradeState = wechat.TradeStateRefund
	err = glog.GVA_DB.Where("outtradeno = ?", tradeNo).Updates(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
