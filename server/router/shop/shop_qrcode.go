package shop

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopQrcodeRouter struct {
}

// InitShopQrcodeRecode   初始化 获取棉花糖机器支付二维码 路由信息
func (ShopQrcodeRouter *ShopQrcodeRouter) InitShopQrcodeRecode(Router *gin.RouterGroup) {
	shopGoodsPayRouter := Router.Group("mht")
	shopGoodsToPayRouter := Router.Group("pay")
	var api = v1.ApiGroupApp.ShopApiGroup.ShopQrcodeApi
	{
		shopGoodsPayRouter.POST("Qrcode", api.CreateShopQrcode)          //棉花糖机器获取授权url
		shopGoodsPayRouter.GET("openId", api.GetOpenId)                  //棉花糖机器专用判断是用什么扫码，并生成对应url
		shopGoodsPayRouter.POST("QueryOrder", api.QueryOrder)            //查询订单
		shopGoodsPayRouter.GET("userRefund/:OutTradeNo", api.UserRefund) //棉花糖机器自动退款

	}
	{
		shopGoodsToPayRouter.GET("openId", api.ToTayGetOpenId) //自定义金额支付，并生成对应授权url 跳过去授权
	}

}
