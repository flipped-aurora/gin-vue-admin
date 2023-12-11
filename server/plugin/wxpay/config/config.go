package config

//	type WxPay struct {
//		MchID                      string // 商户ID
//		AppID                      string // 绑定小程序的APPID
//		Secret                     string // AppSecret 调用授权接口需要用
//		NotifyUrl                  string // 支付回调域名
//		MchCertificateSerialNumber string // 商户证书序列号
//		MchAPIv3Key                string // 商户APIv3密钥
//		PemPath                    string // 证书文件所在地址
//	}
type WxPay struct {
	MchID                      string `mapstructure:"mchid" json:"mchid" yaml:"mchid"`                                                                // 商户ID
	AppID                      string `mapstructure:"appid" json:"appid" yaml:"appid"`                                                                // 绑定小程序的APPID
	Secret                     string `mapstructure:"secret" json:"secret" yaml:"secret"`                                                             // 授权登录用到AppSecret
	NotifyUrl                  string `matstructuer:"notifyurl" json:"notifyurl" yaml:"notifyurl"`                                                    // 支付回调域名
	MchCertificateSerialNumber string `matstructuer:"mchcertificateserialnumber" json:"mchcertificateserialnumber" yaml:"mchcertificateserialnumber"` // 商户证书序列号
	MchAPIv3Key                string `matstructuer:"mchapiv3key" json:"mchapiv3key" yaml:"mchapiv3key"`                                              // 商户APIv3密钥
	PemPath                    string `matstructuer:"pempath" json:"pempath" yaml:"pempath"`                                                          // 证书文件所在地址
	Load                       string `matstructuer:"load" json:"load" yaml:"load"`                                                                   // 微信接口前缀路径
}
