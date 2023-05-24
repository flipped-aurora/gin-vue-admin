package ali

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	appid := "2021003187675861"
	key := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCSwVikI1LZ0iZDsjw4HRWHqKHdW/bX/fu15B2+NF3df6935iUfm91rKTSBJZafZyIGosze5Ozg3nXNLTVBr8mx/CWA53ykFj5uPfkvP0VR+1yjZdlRWuVuURebrKNqT/UDoO0Hes/K/d6zPPK70D14JIZHuif7WJF3SHq8FTcQ56AskJnEYXShRfGflOBGkB6xHslXaSAlEM2idfL9hr6okICaryvd71smqGH9KuNBKhMARJJv/Q3S7VMKNiPLdm34VlAcg/zKbEAFlmEeCMmfLdThEYeYCnxLN90XaNJUsMy7tB944syj9l1KgevGfoXf54Hn18rI7CcEHy/jKe4FAgMBAAECggEBAIadvIWzsUgFMPxf/kUqyJkL4o5+0UsHu2D5wx66Qm/vtDEzd1IrU0v9HCBrLmNB40BbDbpYG8GZzk1K7tM2QLl3JZ+8t+R8UwAk8PKaAyhUfoGT52htDI/I6vLtxWi7t3CAqWx3ByyY+zNG8RG+zuPp/RLGe+5fK/BTF+J3Q2/I17bUIzvRfp5bTmtuawecpV9eua9xCiibVu/G+wLhcvIeP8JiutyqfVXZIQ8SPcDmb5VfQeAEsQOStgls78l9Wb7Q0KNZZ4ziuBf3Kyy5bMAq7NzKFrZUwSisR/mf+ASIlJWN9ihq3S3fjzP75UzaAqR21xTI6WsA3cdTBWMpFQECgYEA/iPZZ7XK44ZXOaLwsxcwJRzpk+jIHzv52aqtWU6fghhAMCnTtwqwVveuLb+R9TEolOPKAzxbLnJiIQOXKicMZ2kWpRVjf/9gEZ+sfeM5MvWe4T3iL3koe0DhMPRFkOOYAnj+A1/js3Sjz1St0BVAeImyTMIW1xlVWoo/YIAHDHUCgYEAk9RNrgdYgKa4XFnLlQOZfl+1nx4DgLmCECU9DE/0jJcN0tTVrTWcN2cKuz/H2rdCXz83KPw+pYvlUGV8egg+eT86+BZ3MaFsrFWHword7LXBhT42owQkYv7M87AXlJ2JCaJPlb4Q6KttWBFGE07+p5nKmGvzJBQ/qvhLlFsnaVECgYEAlxFeUWzZ/ufDf8pgxER7ZQRLwJV2KELxoTb4toarAJ1Dwp2KcVsUOY4fZqhirdqb4e2KDuRY20I5HKDSwcbWc74fLOCK9fIq/C7rTSW9wZIJ8KBljIy7yJBnjD8X/ymbGZUsgXsFko46NauDapbHbR+E8mXJR/GSMDDPvaA9EGUCgYBjoDWFri+n7DPdESsRTu636zqUdH651stfpCkDTjMqOHa5enRSmgmgkNH7pZ1DujOByzGabnytxp0RVIPSHhI0VjegkEtd6EM71Rx8MGglqcw0xtu9ewImsT9gj3Ek2BJNqtRlv0TX+QfkzSdjxd37B7swg7fFvdu2lkUhQ1lJMQKBgDJZswoF8j63y8bCYTvY4SBjoJVU9yMT7aSYxmOJP2lvHGEvpD0vAUj4GiWzPIK6oXcD0ZRBpFdMnnlofhDINtsCX5w0qfbWZ77zhn4v2B/JtSsC0fdemJ3DJ9Tgfg6PKIWtAXdLq6RhKuMgwYnK2K0Rmf5CqvfmRSF9Riwd/yxj"
	client, err := alipay.NewClient(appid, key, true)
	if err != nil {
		global.GVA_LOG.Sugar().Error("初始化支付宝失败", err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                                     // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                                    // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl("https://app.soonai.cn/apihash1/ali/aliNotify") // 设置异步通知URL

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	var appPublicCert string
	appPublicCert = os.Getenv("ALI_PUBLIC_CERT")
	if appPublicCert == "" {
		appPublicCert, _ = os.Getwd()
	}
	appPublicCert = appPublicCert + "/appCertPublicKey.crt"
	var alipayRootCert string
	alipayRootCert = os.Getenv("ALI_PAY_ROOT_CERT")
	if alipayRootCert == "" {
		alipayRootCert, _ = os.Getwd()
	}
	alipayRootCert = alipayRootCert + "/alipayRootCert.crt"
	var alipayPublicCert string
	alipayPublicCert = os.Getenv("ALI_PAY_PUBLIC_CERT")
	if alipayPublicCert == "" {
		alipayPublicCert, _ = os.Getwd()
	}
	alipayPublicCert = alipayPublicCert + "/alipayCertPublicKey.crt"
	bytes, err := utils.File2Bytes(alipayPublicCert)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	client.AutoVerifySign(bytes)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径

	if err := client.SetCertSnByPath(appPublicCert, alipayRootCert, alipayPublicCert); err != nil {
		global.GVA_LOG.Sugar().Error("初始化支付宝失败", err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "123456")
	bm.Set("total_amount", fmt.Sprintf("%.2f", 0.01))
	bm.Set("subject", "test")
	pay, err := client.TradeWapPay(
		context.Background(),
		bm,
	)
	fmt.Println(pay)
}
