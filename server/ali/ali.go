package ali

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"os"
)

func NewAli() {
	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(global.GVA_CONFIG.Ali.AppId, global.GVA_CONFIG.Ali.PrivateKey, true)
	if err != nil {
		global.GVA_LOG.Sugar().Error("初始化支付宝失败", err)
		panic(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                         // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                        // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl(global.GVA_CONFIG.Ali.PayNotifyURL) // 设置异步通知URL

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	var appPublicCert string
	appPublicCert = os.Getenv("ALI_PUBLIC_CERT")
	if appPublicCert == "" {
		appPublicCert, _ = os.Getwd()
		appPublicCert += "/config"
	}
	appPublicCert = appPublicCert + "/appCertPublicKey.crt"
	var alipayRootCert string
	alipayRootCert = os.Getenv("ALI_PAY_ROOT_CERT")
	if alipayRootCert == "" {
		alipayRootCert, _ = os.Getwd()
		alipayRootCert += "/config"
	}
	alipayRootCert = alipayRootCert + "/alipayRootCert.crt"
	var alipayPublicCert string
	alipayPublicCert = os.Getenv("ALI_PAY_PUBLIC_CERT")
	if alipayPublicCert == "" {
		alipayPublicCert, _ = os.Getwd()
		alipayPublicCert += "/config"
	}
	alipayPublicCert = alipayPublicCert + "/alipayCertPublicKey.crt"
	bytes, err := utils.File2Bytes(alipayPublicCert)
	if err != nil {
		global.GVA_LOG.Sugar().Error("初始化支付宝失败", err)
		//panic(err)
		return
	}
	client.AutoVerifySign(bytes)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径

	if err := client.SetCertSnByPath(appPublicCert, alipayRootCert, alipayPublicCert); err != nil {
		global.GVA_LOG.Sugar().Error("初始化支付宝失败", err)
		//panic(err)
		return
	}
	global.GVA_ALI_PAY = client
	// 证书内容
	//err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
}
