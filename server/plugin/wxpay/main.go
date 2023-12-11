package wxpay

import (
	glog "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/router"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat/v3"
	"go.uber.org/zap"
	"os"
)

type WxPayPlugin struct {
}

func CreatePayPlug(mchID, appID, secret, notifyUrl, mchCertificateSerialNumber, mchAPIv3Key, pemPath, load string) *WxPayPlugin {
	global.GlobalConfig.MchID = mchID
	global.GlobalConfig.AppID = appID
	global.GlobalConfig.Secret = secret
	global.GlobalConfig.NotifyUrl = notifyUrl
	global.GlobalConfig.MchCertificateSerialNumber = mchCertificateSerialNumber
	global.GlobalConfig.MchAPIv3Key = mchAPIv3Key
	global.GlobalConfig.PemPath = pemPath
	global.GlobalConfig.Load = load
	//放这里初始化一下  不知道行不行
	PayInitialize()
	return &WxPayPlugin{}
}

func (*WxPayPlugin) Register(group *gin.RouterGroup) {
	wxpay.RouterGroupApp.InitPayRouter(group)
}

func PayInitialize() {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	str, err := os.ReadFile(global.GlobalConfig.PemPath)
	if err != nil {
		glog.GVA_LOG.Error("读取微信证书失败!", zap.Error(err))
	}
	global.WxClient, err = wechat.NewClientV3(global.GlobalConfig.MchID, global.GlobalConfig.MchCertificateSerialNumber, global.GlobalConfig.MchAPIv3Key, string(str))
	if err != nil {
		glog.GVA_LOG.Error("微信客户端初始化失败!", zap.Error(err))
		return
	}

	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = global.WxClient.AutoVerifySign()
	if err != nil {
		glog.GVA_LOG.Error("启用自动同步返回验签失败!", zap.Error(err))
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认是关闭的
	//client.DebugSwitch = gopay.DebugOn
}

func (*WxPayPlugin) RouterPath() string {
	return "api/wxpay"
}
