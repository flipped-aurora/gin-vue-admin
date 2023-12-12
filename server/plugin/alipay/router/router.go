package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/api"
	"github.com/gin-gonic/gin"
)

type AlipayRouter struct {
}

func (s *AlipayRouter) InitAlipayRouter(Router *gin.RouterGroup) {
	PublicGroup := Router.Group("")                                                            //fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()) //fmt.Println("鉴权插件安装==》", PrivateGroup)
	//plugRouter := Router
	plugApi := api.ApiGroupApp.AlipayApi
	{
		PublicGroup.GET("payjsapi", plugApi.ApiPayJsapi) //下单
		PublicGroup.GET("getcode", plugApi.ApiGetCode)   //获取code
		PublicGroup.GET("gettoken", plugApi.ApiGetToken) //获取token
		PublicGroup.POST("notify", plugApi.ApiNotify)    //支付宝回调
		PrivateGroup.GET("refunds", plugApi.ApiRefunds)  //订单退款
	}
	{
		PublicGroup.GET("topaycode", plugApi.ApiToPayGetCode)      //获取code 自定义金额
		PublicGroup.GET("topaygettoken", plugApi.ApiToPayGetToken) //获取token 自定义金额
		PublicGroup.GET("topayjsapi", plugApi.ApiToPayJsapi)       //下单
	}
}
