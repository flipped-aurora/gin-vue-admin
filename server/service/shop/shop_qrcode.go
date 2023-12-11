package shop

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay"
)

type ShopQrcodeService struct {
}

// ServiceCreateOrders 棉花糖机器下单，返回二维码url
func (ShopQrcodeService *ShopQrcodeService) ServiceCreateOrders(req *request.RequestMhtData) (gopay.BodyMap, error) {
	//从数据库把商品信息取出来
	var goods = shop.ShopGoods{}
	err := global.GVA_DB.Where("id = ? and goods_Macid = ?", req.PatternId, req.Macid).First(&goods).Error
	if err != nil {
		return nil, errors.New("商品信息错误")
	}
	var tradeNo = utils.GetOrderNumber()
	var order = shop.ShopOrders{}
	total := *goods.GoodsPrice + *goods.GoodsPostfee
	order.DeviceId = req.Macid
	order.OutTradeNo = tradeNo
	order.GoodsTitle = goods.GoodsName
	order.GoodsId = int(goods.ID)
	order.Total = &total
	order.PostFee = goods.GoodsPostfee
	order.GoodsPrice = goods.GoodsPrice
	order.ShopID = goods.GoodsShopid
	order.TradeState = utils.NOTPAY //"NOTPAY"
	err = global.GVA_DB.Create(&order).Error
	if err != nil {
		return nil, err
	}
	aToken := fmt.Sprintf("%s/api/mht/openId?attach=%s,%d,%s,%d,1226862,1369596012470", global.GVA_CONFIG.AliPay.Load, goods.GoodsMacid, goods.ID, tradeNo, total)
	bm := make(gopay.BodyMap).
		Set("payQr", aToken).
		Set("agencyNo", "1226862").
		Set("money", goods.GoodsPrice).
		Set("hlMerchantId", goods.GoodsShopid).
		Set("orderId", tradeNo)
	return bm, nil
}

// ServiceQueryOrders 棉花糖机器查询订单
func (ShopQrcodeService *ShopQrcodeService) ServiceQueryOrders(req *request.RequestMhtData) (int, error) {
	//从数据库把商品信息取出来
	var order = shop.ShopOrders{}
	err := global.GVA_DB.Where("OutTradeNo = ? ", req.OutTreadNo).First(&order).Error
	if err != nil {
		return 2, errors.New("商品信息错误")
	}
	if order.TradeState == utils.SUCCESS || order.TradeState == utils.TRADE_SUCCESS {
		return 0, nil
	} else {
		return 9, errors.New("支付状态异常")
	}
}
