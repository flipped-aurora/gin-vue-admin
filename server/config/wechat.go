package config

type Wechat struct {
	AppID                      string `mapstructure:"appID" json:"appID" yaml:"appID"`
	AppSecret                  string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"`
	MchID                      string `mapstructure:"mchID" json:"mchID" yaml:"mchID"`
	MchCertificateSerialNumber string `mapstructure:"mchCertificateSerialNumber" json:"mchCertificateSerialNumber" yaml:"mchCertificateSerialNumber"`
	MchAPIv3Key                string `mapstructure:"mchAPIv3Key" json:"mchAPIv3Key" yaml:"mchAPIv3Key"`
	MchAPIv2Key                string `mapstructure:"mchAPIv2Key" json:"mchAPIv2Key" yaml:"mchAPIv2Key"`
	PayNotifyUrl               string `mapstructure:"payNotifyUrl" json:"payNotifyUrl" yaml:"payNotifyUrl"`
	RefundNotifyUrl            string `mapstructure:"refundNotifyUrl" json:"refundNotifyUrl" yaml:"refundNotifyUrl"`
	GetAccessTokenAccount      int    `mapstructure:"getAccessTokenAccount" json:"getAccessTokenAccount" yaml:"getAccessTokenAccount"`
	MiniprogramState           string `mapstructure:"miniprogramState" json:"miniprogramState" yaml:"miniprogramState"`
}
