package config

type Alipay struct {
	Appid                   string `mapstructure:"appid" json:"appid" yaml:"appid"`                                                       // appid
	EncryptKey              string `mapstructure:"encryptkey" json:"encryptkey" yaml:"encryptkey"`                                        //接口解密密钥
	PrivateKey              string `mapstructure:"privatekey" json:"privatekey" yaml:"privatekey"`                                        // 支付宝私钥
	AlipayPublicContentRSA2 string `mapstructure:"alipaypublicContentrsa2" json:"alipaypublicContentrsa2" yaml:"alipaypublicContentrsa2"` // 支付宝公钥证书
	AlipayRootContent       string `mapstructure:"alipayrootcontent" json:"alipayrootcontent" yaml:"alipayrootcontent"`                   // 支付宝根证书
	AppPublicContent        string `mapstructure:"apppubliccontent" json:"apppubliccontent" yaml:"apppubliccontent"`                      // 应用公钥证书
	NotifyUrl               string `mapstructure:"alinotifyurl" json:"alinotifyurl" yaml:"alinotifyurl"`                                  // 回调接口
	ReturnURL               string `mapstructure:"returnurl" json:"returnurl" yaml:"returnurl"`                                           // 返回地址
	Load                    string `mapstructure:"load" json:"load" yaml:"load"`                                                          // 支付宝支付前缀路径
}
