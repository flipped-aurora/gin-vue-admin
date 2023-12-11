package service

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	ali "github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/global"
	alireq "github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

type AliPayService struct{}

// PlugServiceJsPay  创建订单
func (A *AliPayService) PlugServiceJsPay(req *alireq.ResAuthJsApi) (aliRsp *alipay.TradeCreateResponse, err error) {
	str, err := utils.ReplaceAttach(req.Attach)
	if err != nil {
		return nil, err
	}
	//把请求参数里面的单号取出来
	outTreadNo := str.Get("outTreadNo")
	if outTreadNo == "" {
		err = errors.New("参数错误")
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	//从数据库把订单信息取出来
	var orders = shop.ShopOrders{}
	err = global.GVA_DB.Where("OutTradeNo = ?", outTreadNo).First(&orders).Error
	if err != nil {
		return nil, err
	}
	if orders.TradeState == utils.SUCCESS || orders.TradeState == utils.TRADE_SUCCESS {
		return nil, errors.New("订单已支付")
	}

	// 请求参数
	//计算一下金额，支付宝需要用双进度数值表示 例如：0.01,1.21
	money := float64(*orders.Total)
	total := decimal.NewFromFloat(money).Div(decimal.NewFromFloat(100))
	bm := make(gopay.BodyMap)
	bm.Set("subject", orders.GoodsTitle).
		Set("out_trade_no", outTreadNo).
		Set("total_amount", total).
		Set("passback_params", req.Attach).
		Set("buyer_open_id", req.OpenId)
	// 创建订单
	paystr, err := ali.Client.TradeCreate(context.Background(), bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			global.GVA_LOG.Error("失败!", zap.Error(bizErr))
			return nil, err
		}
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	if paystr.Response.TradeNo == "" {
		err = errors.New("返回交易号为空")
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	return paystr, nil
}

// PlugServiceJsPay  创建订单 ，自定义金额付款
func (A *AliPayService) PlugServiceToJsPay(req *alireq.ResAuthJsApi) (aliRsp *alipay.TradeCreateResponse, err error) {
	outTreadNo := utils.GetOrderNumber()
	// 请求参数
	//计算一下金额，支付宝需要用双进度数值表示 例如：0.01,1.21
	n := decimal.NewFromFloat(float64(req.Money)).Div(decimal.NewFromFloat(100))
	//total := n.IntPart()
	//fmt.Println(total)
	bm := make(gopay.BodyMap)
	title := "云南串先生扫码付"
	bm.Set("subject", title).
		Set("out_trade_no", outTreadNo).
		Set("total_amount", n).
		Set("passback_params", req.Attach).
		Set("buyer_open_id", req.OpenId)
	// 创建订单
	paystr, err := ali.Client.TradeCreate(context.Background(), bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			global.GVA_LOG.Error("失败!", zap.Error(bizErr))
			return nil, err
		}
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	if paystr.Response.TradeNo == "" {
		err = errors.New("返回交易号为空")
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	var orders = shop.ShopOrders{}
	orders.OpenId = req.OpenId
	orders.Total = &req.Money
	orders.TradeState = utils.NOTPAY
	orders.GoodsTitle = title
	orders.PayMent = utils.Alipay
	orders.GoodsPrice = &req.Money
	//orders.TradeType = "NATIVE"
	orders.OutTradeNo = outTreadNo
	orders.PostFee = nil

	err = global.GVA_DB.Save(&orders).Error
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		return nil, err
	}
	return paystr, nil
}

// 回调
func (A *AliPayService) PlugServiceNotify(req *gopay.BodyMap) error {
	var order = shop.ShopOrders{}
	order.AppId = req.Get("app_id")
	order.MchId = req.Get("seller_id")        //卖家支付宝用户号
	order.Attach = req.Get("passback_params") //公共回传参数，如果请求时传递了该参数，则返回给商家时会在异步通知时将该参数原样返回
	order.GoodsTitle = req.Get("subject")     //订单标题
	order.TransactionId = req.Get("trade_no")
	s, err := StrToInt64(req.Get("receipt_amount"))
	if err != nil {
		return errors.New("获取订单金额错误")
	}
	order.PayerTotal = s //实际支付金额
	m, err := StrToInt64(req.Get("total_amount"))
	if err != nil {
		return errors.New("获取支付金额错误")
	}
	order.Total = m //订单金额
	order.TradeState = req.Get("trade_status")
	order.TradeType = "JSAPI"
	order.OpenId = req.Get("buyer_open_id")
	t1 := time.Now()
	order.SuccessTime = &t1
	order.PayMent = utils.Alipay       //支付方式，1 微信 2 支付宝
	tradeno := req.Get("out_trade_no") //商户订单号
	if tradeno == "" {
		return errors.New("订单号为空")
	}
	err = global.GVA_DB.Where("OutTradeNo=?", tradeno).Updates(&order).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// PlugServiceRefund SUCCESS:支付成功, REFUND:转入退款, NOTPAY:未支付, CLOSED:已关闭, REVOKED:已撤销, USERPAYING:支付中, PAYERROR:支付失败,WAIT_BUYER_PAY:交易创建,TRADE_CLOSED:交易超时,TRADE_SUCCESS:支付成功,TRADE_FINISHED:交易结束
// PlugServiceRefund 退款
func (A *AliPayService) PlugServiceRefund(tradeNo *alireq.ReqTradeNo) (aliRsp *alipay.TradeRefundResponse, err error) {
	var order *shop.ShopOrders
	err = global.GVA_DB.Where("OutTradeNo = ? and TradeState = ?", tradeNo.OutTradeNo, utils.TRADE_SUCCESS).First(&order).Error
	if err != nil {
		return nil, err
	}
	if len(order.OpenId) == 0 {
		return nil, errors.New("订单不存在")
	}
	// 请求参数
	//数据库以分为单位，需要除以100
	money := float64(*order.PayerTotal)
	amount := decimal.NewFromFloat(money).Div(decimal.NewFromFloat(100))
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", tradeNo.OutTradeNo).
		Set("refund_amount", amount)

	// 发起退款请求
	aliRsp, err = ali.Client.TradeRefund(context.Background(), bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			return nil, bizErr
		}
		return nil, err
	}
	order.TradeState = utils.REFUND //退款标记
	t := time.Now()
	order.Endtime = &t
	err = global.GVA_DB.Save(&order).Error
	if err != nil {
		return nil, err
	}
	return aliRsp, nil
}
