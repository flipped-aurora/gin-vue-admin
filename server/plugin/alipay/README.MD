## GVA 支付宝支付插件

#### 开发者：ztcjn 

### 使用步骤

#### 1. 前往GVA主程序下的 server/initialize/plugin.go 加入如下代码

    例：
    本插件可以采用gva的配置文件 
	//微信支付插件 不需要鉴权
	PluginInit(PublicGroup, wxpay.CreatePayPlug(
		global.GVA_CONFIG.WxPay.MchID,
		global.GVA_CONFIG.WxPay.AppID,
		global.GVA_CONFIG.WxPay.NotifyUrl,
		global.GVA_CONFIG.WxPay.MchCertificateSerialNumber,
		global.GVA_CONFIG.WxPay.MchAPIv3Key,
		global.GVA_CONFIG.WxPay.PemPath,
	))

#### 2.在config里面加入配置项 server/config/config.go

    //wxpay
	WxPay WxPay `mapstructure:"wxpay" json:"wxpay" yaml:"wxpay"`

#### 3.在config下新建wxpay.go 文件，路径：server/config/wxpay.go 写如下代码：

    package config
    
    type WxPay struct {
    MchID                      string `mapstructure:"mchid" json:"mchid" yaml:"mchid"`                                                                // 商户ID
    AppID                      string `mapstructure:"appid" json:"appid" yaml:"appid"`                                                                // 绑定小程序的APPID
    NotifyUrl                  string `matstructuer:"notifyurl" json:"notifyurl" yaml:"notifyurl"`                                                    // 支付回调域名
    MchCertificateSerialNumber string `matstructuer:"mchcertificateserialnumber" json:"mchcertificateserialnumber" yaml:"mchcertificateserialnumber"` // 商户证书序列号
    MchAPIv3Key                string `matstructuer:"mchapiv3key" json:"mchapiv3key" yaml:"mchapiv3key"`                                              // 商户APIv3密钥
    PemPath                    string `matstructuer:"pempath" json:"pempath" yaml:"pempath"`                                                          // 证书文件所在地址
    }


### 2. 配置说明

#### 2-1 全局配置结构体说明
    插件目录下面config/config.go
    //小程序必须和商户号绑定才可以使用微信支付

    type WxPay struct {
	MchID                      string  // 商户ID
	AppID                      string  // 绑定小程序的APPID
	NotifyUrl                  string  // 支付回调域名
	MchCertificateSerialNumber string  // 商户证书序列号
	MchAPIv3Key                string  // 商户APIv3密钥
	PemPath                    string  // 证书文件所在地址
}

#### 2-2 入参结构说明

    type Request struct {
    UserID  int    `form:"userid" json:"userid"`       // 下单用户ID
    GoodsID string `form:"goodsid" binding:"required"` // 商品ID用于查询商品价格
    Attach  string `form:"attach" json:"attach"`       // 附加参数
    CallUrl string `form:"callurl" json:"callUrl"`     //第三方回调call
    }
    
    type ReqTradeNo struct {
    OrderId string `form:"orderid" binding:"required,min=5,max=24"` // 订单ID
}

### 3. 方法API

    utils.EmailTest(邮件标题，邮件主体) 发送测试邮件
    例:utils.EmailTest("测试邮件"，"测试邮件")
    utils.ErrorToEmail(邮件标题,邮件主体) 错误监控
    例:utils.ErrorToEmail("测试邮件"，"测试邮件")
    utils.Email(目标邮箱多个的话用逗号分隔，邮件标题，邮件主体) 发送测试邮件
    例:utils.Email(”a.qq.com,b.qq.com“,"测试邮件"，"测试邮件")

### 4. 可直接调用的接口
    PublicGroup.GET("getpaycode", plugApi.WxApiGetCode)     //微信预下单
		PublicGroup.POST("notify_url", plugApi.WxapinotifyUrl)  //微信回调
		PublicGroup.GET("queryorder", plugApi.WxApiQueryOrder)  //微信订单查询
		PrivateGroup.GET("closeorder", plugApi.WxApiCloseOrder) //微信订单关闭
		PrivateGroup.POST("refunds", plugApi.WxApiRefunds)      //微信订单退款
   
