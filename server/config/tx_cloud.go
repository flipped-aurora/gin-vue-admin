package config

type TXCloud struct {
	SecretId    string `mapstructure:"secretId" json:"secretId" yaml:"secretId"`
	SecretKey   string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	SmsSdkAppId string `mapstructure:"smsSdkAppId" json:"smsSdkAppId" yaml:"smsSdkAppId"`
	SmsSign     string `mapstructure:"smsSign" json:"smsSign" yaml:"smsSign"`
	SmsTemplate string `mapstructure:"smsTemplate" json:"smsTemplate" yaml:"smsTemplate"`
}
