package wxpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	wxapi "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/api"
	"github.com/gin-gonic/gin"
)

type PayRouter struct {
}

func (s *PayRouter) InitPayRouter(Router *gin.RouterGroup) {
	PublicGroup := Router.Group("")  //fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("") //fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	plugApi := wxapi.WxApiGroupApp.WxPay
	{
		PublicGroup.GET("native", plugApi.WxApiNativeCode)            //微信Native预下单
		PublicGroup.POST("notify", plugApi.WxApiNativeNotifyUrl)      //微信回调
		PrivateGroup.GET("queryorder", plugApi.WxApiNativeQueryOrder) //微信订单查询
		PrivateGroup.GET("closeorder", plugApi.WxApiNativeCloseOrder) //微信订单关闭
		PrivateGroup.POST("refunds", plugApi.WxApiNativeRefunds)      //微信订单退款

		PublicGroup.GET("getcode", plugApi.WxApiGetCode)   //微信JsApi 获取code
		PublicGroup.GET("gettoken", plugApi.WxApiGetToken) //微信JsApi 获取token
		PublicGroup.GET("payjsapi", plugApi.WxApiJsApi)    //微信JsApi下单

		PublicGroup.POST("set_candy", plugApi.WxTest)       //测试获取版本
		PublicGroup.POST("getVolume", plugApi.WxgetVolume)  //测试获取音量
		PublicGroup.POST("candy_info", plugApi.WxcandyInfo) //测试发送糖型号
		PublicGroup.POST("sugarFlow", plugApi.WxsugarFlow)  //测试流量检测
	}
	{
		PublicGroup.GET("topaycode", plugApi.WxApiToPayCode)   //微信JsApi 获取code 自定义金额
		PublicGroup.GET("topaytoken", plugApi.WxApiToPayToken) //微信JsApi 获取token 自定义金额
		PublicGroup.GET("topayjsapi", plugApi.WxApiToPayJsApi) //微信JsApi 获取token 自定义金额
	}
}
