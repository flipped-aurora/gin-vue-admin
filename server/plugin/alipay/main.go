package alipay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/router"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

type AlipayPlugin struct {
}

func CreateAlipayPlug(appid, encryptKey, privateKey, alipayPublicContentRSA2, alipayRootContent, appPublicContent, notifyurl, returnurl, load string) *AlipayPlugin {
	global.GlobalConfig.Appid = appid
	global.GlobalConfig.PrivateKey = privateKey
	global.GlobalConfig.AlipayPublicContentRSA2 = alipayPublicContentRSA2
	global.GlobalConfig.AppPublicContent = appPublicContent
	global.GlobalConfig.AlipayRootContent = alipayRootContent
	global.GlobalConfig.EncryptKey = encryptKey
	global.GlobalConfig.NotifyUrl = notifyurl
	global.GlobalConfig.ReturnURL = returnurl
	global.GlobalConfig.Load = load
	AliPayInitialize()
	return &AlipayPlugin{}
}

func (*AlipayPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitAlipayRouter(group)
}

func (*AlipayPlugin) RouterPath() string {
	return "alipay"
}

func AliPayInitialize() {

	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	var err error
	var cfg = global.GlobalConfig
	global.Client, err = alipay.NewClient(cfg.Appid, cfg.PrivateKey, true)
	if err != nil {
		xlog.Error(err)
		return
	}
	// Debug开关，输出/关闭日志
	//global.Client.DebugSwitch = gopay.DebugOff

	// 配置公共参数
	global.Client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		// SetAppAuthToken("")
		//SetReturnUrl("https://www.fmm.ink").
		SetNotifyUrl(global.GlobalConfig.NotifyUrl)

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前未测试成功）
	//client.SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==")

	// 自动同步验签（只支持证书模式）
	// 传入 支付宝公钥证书 alipayPublicCert.crt 内容
	//str, err := os.ReadFile(global.GlobalConfig.AlipayPublicContentRSA2)
	//if err != nil {
	//	glog.GVA_LOG.Error("读取支付宝证书失败!", zap.Error(err))
	//}
	//global.Client.AutoVerifySign(str)

	// 传入证书内容
	//err = global.Client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	// 传入证书文件路径
	err = global.Client.SetCertSnByPath(cfg.AppPublicContent, cfg.AlipayRootContent, cfg.AlipayPublicContentRSA2)
	if err != nil {
		xlog.Debug("SetCertSn:", err)
		return
	}
}
