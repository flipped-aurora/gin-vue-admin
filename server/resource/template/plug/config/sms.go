package config

type AliSms struct {
	AccessKeyId  string `mapstructure:"accessKeyId" json:"accessKeyId" yaml:"accessKeyId"`    // 短信的AccessKey ID
	AccessSecret string `mapstructure:"accessSecret" json:"accessSecret" yaml:"accessSecret"` // 短信的AccessKey Secret
	SignName     string `mapstructure:"signName" json:"signName" yaml:"signName"`             // 短信的 SignName
}
