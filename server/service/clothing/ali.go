package clothing

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

type AliService struct {
}

func (a *AliService) CreateH5Order(Description string, order clothing.Order) (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", order.OrderNo)
	//bm.Set("total_amount", fmt.Sprintf("%.2f", order.Amount))
	bm.Set("total_amount", fmt.Sprintf("%.2f", 0.01))
	bm.Set("subject", Description)
	pay, err := global.GVA_ALI_PAY.TradeWapPay(
		context.Background(),
		bm,
	)
	if err != nil {
		return "", err
	}
	return pay, nil
}

func (a *AliService) CreateAppOrder(Description string, order clothing.Order) (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", order.OrderNo)
	//bm.Set("total_amount", fmt.Sprintf("%.2f", order.Amount))
	bm.Set("total_amount", fmt.Sprintf("%.2f", 0.01))
	bm.Set("subject", Description)
	pay, err := global.GVA_ALI_PAY.TradeAppPay(
		context.Background(),
		bm,
	)
	if err != nil {
		return "", err
	}
	return pay, nil
}

func (a *AliService) PayNotify(c *gin.Context) (string, string, error) {
	bodyMap, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		return "", "", err
	}
	global.GVA_LOG.Sugar().Info(bodyMap)
	if bodyMap.Get("trade_status") != "TRADE_SUCCESS" {
		return "", "", errors.New(bodyMap.Get("trade_status"))
	}
	return bodyMap.GetString("out_trade_no"), bodyMap.GetString("trade_no"), nil
}
