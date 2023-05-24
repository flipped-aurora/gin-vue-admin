package wechat

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/config"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"os"
)

func NewWechat() {
	wc := wechat.NewWechat()
	mini(wc)
	pay()
}

// 小程序
func mini(wc *wechat.Wechat) {
	redisCfg := global.GVA_CONFIG.Redis
	redisOpts := &cache.RedisOpts{
		Host:        redisCfg.Addr,     // redis host
		Password:    redisCfg.Password, //redis password
		Database:    redisCfg.DB,       // redis db
		MaxActive:   10,                // 连接池最大活跃连接数
		MaxIdle:     10,                //连接池最大空闲连接数
		IdleTimeout: 60,                //空闲连接超时时间，单位：second
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	cfg := &config.Config{
		AppID:     global.GVA_CONFIG.Wechat.AppID,
		AppSecret: global.GVA_CONFIG.Wechat.AppSecret,
		Cache:     redisCache,
	}
	global.GVA_WECHAT_MINI = wc.GetMiniProgram(cfg)
}

// 微信支付
func pay() {
	var (
		mchID                      = global.GVA_CONFIG.Wechat.MchID                      // 商户号
		mchCertificateSerialNumber = global.GVA_CONFIG.Wechat.MchCertificateSerialNumber // 商户证书序列号
		mchAPIv3Key                = global.GVA_CONFIG.Wechat.MchAPIv3Key                // 商户APIv3密钥
	)
	var path string
	path = os.Getenv("WECHAT_PAY_KEY")
	if path == "" {
		path, _ = os.Getwd()
		path += "/config"
	}
	path = path + "/apiclient_key.pem"
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(path)
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := jsapi.JsapiApiService{Client: client}
	svc2 := app.AppApiService{Client: client}
	svc3 := h5.H5ApiService{Client: client}
	global.GVA_WECHAT_PAY = &svc
	global.GVA_WECHAT_PAY_APP = &svc2
	global.GVA_WECHAT_PAY_H5 = &svc3
	refund := refunddomestic.RefundsApiService{Client: client}
	global.GVA_WECHAT_REFUND = &refund
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mchCertificateSerialNumber, mchID, mchAPIv3Key)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	global.GVA_WECHAT_PAY_HANDLER = handler
}
